package database

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type wrappedRow struct {
	rows pgx.Rows
}

func (r wrappedRow) Scan(dest ...interface{}) error {
	if r.rows.Err() != nil {
		return r.rows.Err()
	}

	if !r.rows.Next() {
		if r.rows.Err() == nil {
			return pgx.ErrNoRows
		}
		return r.rows.Err()
	}

	r.rows.Scan(dest...)
	r.rows.Close()
	return r.rows.Err()
}

type wrappedRowError struct {
	err error
}

func (r wrappedRowError) Scan(dest ...interface{}) error {
	return r.err
}

type wrappedPool struct {
	pool *pgxpool.Pool
}

func (w wrappedPool) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	var ct pgconn.CommandTag
	err := retry(ctx, func() (bool, error) {
		var err error
		ct, err = w.pool.Exec(ctx, query, args...)
		return true, wrapError(err)
	}, false)
	return ct, err
}

func (w wrappedPool) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	var rows pgx.Rows
	err := retry(ctx, func() (bool, error) {
		var err error
		rows, err = w.pool.Query(ctx, query, args...)
		return true, wrapError(err)
	}, isReadOnlyQuery(query))
	return rows, err
}

func (w wrappedPool) CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
	var n int64
	err := retry(ctx, func() (bool, error) {
		var err error
		n, err = w.pool.CopyFrom(ctx, tableName, columnNames, rowSrc)
		return true, wrapError(err)
	}, false)
	return n, err
}

func (w wrappedPool) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	rows, err := w.Query(ctx, query, args...)
	if err != nil {
		return wrappedRowError{err}
	}
	return wrappedRow{rows}
}

type wrappedTransaction struct {
	tx pgx.Tx
}

func (t wrappedTransaction) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	ct, err := t.tx.Exec(ctx, query, args...)
	return ct, wrapError(err)
}

func (t wrappedTransaction) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	rows, err := t.tx.Query(ctx, query, args...)
	return rows, wrapError(err)
}

func (t wrappedTransaction) CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
	n, err := t.tx.CopyFrom(ctx, tableName, columnNames, rowSrc)
	return n, wrapError(err)
}

func (t wrappedTransaction) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	rows, err := t.Query(ctx, query, args...)
	if err != nil {
		return wrappedRowError{err}
	}
	return wrappedRow{rows}
}
