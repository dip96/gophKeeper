package postgres

import (
	"context"
	dSql "database/sql"
	"errors"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/stdlib"
	"gophKeeper/internal/config"
	postgresError "gophKeeper/internal/errors/postgres"
	storageInterface "gophKeeper/internal/storage"
)

type storage struct {
	db *dSql.DB
}

func NewDB(cnf config.ConfigServer) (storageInterface.Storage, error) {
	cnfPgx, err := pgx.ParseConfig(cnf.GetDatabaseURI())
	if err != nil {
		return nil, postgresError.New("error parsing database URI", err)
	}

	db := stdlib.OpenDB(*cnfPgx)
	return &storage{db: db}, nil
}

func (s *storage) Exec(ctx context.Context, sql string, args ...any) (dSql.Result, error) {
	retryDelays := []time.Duration{1 * time.Second, 1 * time.Second, 1 * time.Second}
	var err error

	for attempt, delay := range retryDelays {
		var result dSql.Result
		result, err = s.db.ExecContext(ctx, sql, args...)

		if err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				if pgErr.Code == "23505" {
					return nil, postgresError.New("duplicate login", pgErr)
				}
			}

			log.Printf("Err (attempt %d/%d): %v", attempt+1, len(retryDelays), err)
			time.Sleep(delay)
			continue
		}

		return result, nil
	}

	return nil, err
}

func (s *storage) Query(ctx context.Context, sql string, args ...any) (*dSql.Rows, error) {
	return s.db.QueryContext(ctx, sql, args...)
}

func (s *storage) QueryRow(ctx context.Context, sql string, args ...any) *dSql.Row {
	return s.db.QueryRowContext(ctx, sql, args...)
}

func (s *storage) Begin(ctx context.Context) (*dSql.Tx, error) {
	return s.db.BeginTx(ctx, nil)
}

func (s *storage) Close() error {
	return s.db.Close()
}
