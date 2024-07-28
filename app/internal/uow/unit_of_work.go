// паттерн Unit of Work -  позволяет управлять транзакциями для всех репозиториев в одном месте
package uow

import (
	"context"
	"database/sql"
	"gophKeeper/internal/repositories/user"
	"gophKeeper/internal/storage"
)

type unitOfWork struct {
	db       storage.Storage
	userRepo user.UserRepository
}

func NewUnitOfWork(db storage.Storage) UnitOfWork {
	return &unitOfWork{
		db:       db,
		userRepo: user.NewUserRepository(db),
	}
}

func (u *unitOfWork) UserRepository() user.UserRepository {
	return u.userRepo
}

func (u *unitOfWork) BeginTx(ctx context.Context) (*sql.Tx, error) {
	return u.db.BeginTx(ctx)
}

func (u *unitOfWork) Commit(tx *sql.Tx) error {
	return tx.Commit()
}

func (u *unitOfWork) Rollback(tx *sql.Tx) error {
	return tx.Rollback()
}
