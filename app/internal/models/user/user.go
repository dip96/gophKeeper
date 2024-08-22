package user

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID        int64
	Login     string
	Password  []byte
	Salt      string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}
