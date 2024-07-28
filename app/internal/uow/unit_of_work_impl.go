// паттерн Unit of Work -  позволяет управлять транзакциями для всех репозиториев в одном месте
package uow

import (
	"context"
	"database/sql"
	"gophKeeper/internal/repositories/user"
)

type UnitOfWork interface {
	UserRepository() user.UserRepository

	BeginTx(ctx context.Context) (*sql.Tx, error)
	Commit(tx *sql.Tx) error
	Rollback(tx *sql.Tx) error
}
