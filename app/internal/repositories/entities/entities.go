package entities

import (
	"context"
	"database/sql"
)

// Общий интерфейс с generic
type DataRepository[T any] interface {
	GetDataById(ctx context.Context, tx *sql.Tx, entityId int) (T, error)
	SaveData(ctx context.Context, tx *sql.Tx, data T) error
	EditData(ctx context.Context, tx *sql.Tx, entityId int, data T) error
	DeleteData(ctx context.Context, tx *sql.Tx, entityId int) error
}
