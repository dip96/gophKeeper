// паттерн Unit of Work -  позволяет управлять транзакциями для всех репозиториев в одном месте
package uow

import (
	"context"
	"database/sql"
	models "gophKeeper/internal/models/entities"
	"gophKeeper/internal/repositories/entities"
	"gophKeeper/internal/repositories/user"
)

type UnitOfWork interface {
	UserRepository() user.UserRepository
	BinaryDataRepository() entities.DataRepository[models.BinaryData]
	LoginDataRepository() entities.DataRepository[models.LoginData]
	TextDataRepository() entities.DataRepository[models.TextData]

	BeginTx(ctx context.Context) (*sql.Tx, error)
	Commit(tx *sql.Tx) error
	Rollback(tx *sql.Tx) error
}
