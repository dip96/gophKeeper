// паттерн Unit of Work -  позволяет управлять транзакциями для всех репозиториев в одном месте
package uow

import (
	"context"
	"database/sql"
	models "gophKeeper/internal/models/entities"
	"gophKeeper/internal/repositories/entities"
	"gophKeeper/internal/repositories/entities/binary_data"
	"gophKeeper/internal/repositories/entities/login_data"
	"gophKeeper/internal/repositories/entities/text_data"
	"gophKeeper/internal/repositories/user"
	"gophKeeper/internal/storage"
)

type unitOfWork struct {
	db             storage.Storage
	userRepo       user.UserRepository
	binaryDataRepo entities.DataRepository[models.BinaryData]
	loginDataRepo  entities.DataRepository[models.LoginData]
	textDataRepo   entities.DataRepository[models.TextData]
}

func NewUnitOfWork(db storage.Storage) UnitOfWork {
	return &unitOfWork{
		db:             db,
		userRepo:       user.NewUserRepository(db),
		binaryDataRepo: binary_data.NewBinaryDataRepository(db),
		loginDataRepo:  login_data.NewLoginDataRepository(db),
		textDataRepo:   text_data.NewTextDataRepository(db),
	}
}

func (u *unitOfWork) UserRepository() user.UserRepository {
	return u.userRepo
}

func (u *unitOfWork) BinaryDataRepository() entities.DataRepository[models.BinaryData] {
	return u.binaryDataRepo
}

func (u *unitOfWork) LoginDataRepository() entities.DataRepository[models.LoginData] {
	return u.loginDataRepo
}

func (u *unitOfWork) TextDataRepository() entities.DataRepository[models.TextData] {
	return u.textDataRepo
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
