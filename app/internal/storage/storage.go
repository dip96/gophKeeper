package storage

import (
	"context"
	"database/sql"
)

type Storage interface {
	Exec(ctx context.Context, sql string, args ...any) (sql.Result, error)
	Query(ctx context.Context, sql string, args ...any) (*sql.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) *sql.Row
	Begin(ctx context.Context) (*sql.Tx, error)
	Close() error
}
