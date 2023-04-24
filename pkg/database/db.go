package database

import (
	"context"
	"flag"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/Jille/trxwrap"
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
	db trxwrap.TrxWrap[gendb.Queries]
)

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

	pgx, err := pgxpool.ConnectConfig(context.Background(), cc)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	db = trxwrap.New(pgx, func(tx trxwrap.PGDBTX) *gendb.Queries {
		return gendb.New(tx)
	})
	return pgx
}

func RunRWTransaction(ctx context.Context, isolationLevel pgx.TxIsoLevel, runner func(*gendb.Queries) error) error {
	return db.RunRWTransaction(ctx, isolationLevel, runner)
}

func RunROTransaction(ctx context.Context, isolationLevel pgx.TxIsoLevel, runner func(*gendb.Queries) error) error {
	return db.RunROTransaction(ctx, isolationLevel, runner)
}

func binaryName() string {
	p, err := os.Executable()
	if err != nil {
		p = os.Args[0]
	}
	return filepath.Base(p)
}
