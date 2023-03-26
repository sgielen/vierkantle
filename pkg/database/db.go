package database

import (
	"context"
	"errors"
	"flag"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/logrusadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sgielen/vierkantle/pkg/database/gendb"
	"github.com/sirupsen/logrus"
)

var (
	queryLogging = flag.Bool("query_logging", true, "Whether to enable logging for the database")
)

var (
	db *pgxpool.Pool
)

type TransactionRunner func(*gendb.Queries) error

func Init() {
	dsn := os.Getenv("VIERKANTLE_DSN")
	if dsn == "" {
		log.Fatalf("Unset $VIERKANTLE_DSN")
	}
	InitWithDSN(dsn)
}

// This function should *only* be used in the tests and by Init above
func InitWithDSN(dsn string) *pgxpool.Pool {
	cc, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	if cc.ConnConfig.RuntimeParams["application_name"] == "" {
		cc.ConnConfig.RuntimeParams["application_name"] = binaryName()
	}
	if *queryLogging {
		cc.ConnConfig.Logger = logrusadapter.NewLogger(&logrus.Logger{
			Out:          os.Stderr,
			Formatter:    new(logrus.JSONFormatter),
			Hooks:        make(logrus.LevelHooks),
			Level:        logrus.InfoLevel,
			ExitFunc:     os.Exit,
			ReportCaller: false,
		})
	}
	cc.MaxConnIdleTime = 30 * time.Second
	cc.MinConns = 2
	cc.MaxConns = 10

	db, err = pgxpool.ConnectConfig(context.Background(), cc)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}

func RunRWTransaction(ctx context.Context, isolationLevel pgx.TxIsoLevel, runner TransactionRunner) error {
	txo := pgx.TxOptions{
		IsoLevel:   isolationLevel,
		AccessMode: pgx.ReadWrite,
	}
	return RunTransaction(ctx, txo, false, runner)
}

func RunROTransaction(ctx context.Context, isolationLevel pgx.TxIsoLevel, runner TransactionRunner) error {
	txo := pgx.TxOptions{
		IsoLevel:   isolationLevel,
		AccessMode: pgx.ReadOnly,
	}
	return RunTransaction(ctx, txo, true, runner)
}

func RunTransaction(ctx context.Context, txo pgx.TxOptions, idempotent bool, runner TransactionRunner) error {
	return retry(ctx, func() (bool, error) {
		return runTransactionOnce(ctx, txo, runner)
	}, idempotent || txo.AccessMode == pgx.ReadOnly)
}

func retry(ctx context.Context, f func() (bool, error), idempotent bool) error {
	for attempt := 0; ; attempt++ {
		commitAttempted, err := f()
		if attempt >= 3 {
			return err
		}
		retry := pgconn.SafeToRetry(err)
		switch ToSQLState(err) {
		case "40001", "40P01":
			retry = true
		case "08Q99", // io.EOF or io.UnexpectedEOF
			"57P01", // admin_shutdown
			"57P02", // crash_shutdown
			"57P03", // cannot_connect_now
			"08000", // connection_exception
			"08003", // connection_does_not_exist
			"08006", // connection_failure
			"08001", // sqlclient_unable_to_establish_sqlconnection
			"08004": // sqlserver_rejected_establishment_of_sqlconnection
			retry = !commitAttempted || idempotent
		}
		if retry {
			// TODO: Exponential backoff?
			time.Sleep(500 * time.Millisecond)
			continue
		}
		return err
	}
}

func runTransactionOnce(ctx context.Context, txo pgx.TxOptions, runner TransactionRunner) (commitAttempted bool, _ error) {
	tx, err := db.BeginTx(ctx, txo)
	if err != nil {
		return false, wrapError(err)
	}
	q := gendb.New(wrappedTransaction{tx})
	if err := runner(q); err != nil {
		// TODO: Use multi-error to combine a possible error from rollback with err.
		tx.Rollback(ctx)
		return false, err
	}
	return true, wrapError(tx.Commit(ctx))
}

func ToSQLState(err error) string {
	if e, ok := err.(Error); ok {
		if e.parent == io.EOF || e.parent == io.ErrUnexpectedEOF {
			return "08Q99"
		}
	}
	var pge *pgconn.PgError
	if errors.As(err, &pge) {
		return pge.SQLState()
	}
	return ""
}

// wrapError wraps an error that comes from the database.
// If the given error is nil, nil wil be returned.
// The returned error will hide the actual error when returned through gRPC.
func wrapError(err error) error {
	if err == nil {
		return nil
	}
	if err == pgx.ErrNoRows || err == pgx.ErrTxClosed || err == pgx.ErrTxCommitRollback {
		return err
	}
	return Error{err}
}

type Error struct {
	parent error
}

func (e Error) Error() string {
	return e.parent.Error()
}

/* Enable this when going live
func (e Error) GRPCStatus() *status.Status {
	return status.New(codes.Internal, "database error")
}
*/

func (e Error) Unwrap() error {
	return e.parent
}

func binaryName() string {
	p, err := os.Executable()
	if err != nil {
		p = os.Args[0]
	}
	return filepath.Base(p)
}

func isReadOnlyQuery(sql string) bool {
	for strings.HasPrefix(sql, "--") {
		sql = sql[strings.Index(sql, "\n")+1:]
	}
	return strings.HasPrefix(sql, "SELECT")
}

func MaxPoolSize() int {
	return int(db.Config().MaxConns)
}
