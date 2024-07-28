package user

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
	userModel "gophKeeper/internal/models/user"
	"gophKeeper/internal/storage"
)

type UserRepository interface {
	CreateUser(ctx context.Context, tx *sql.Tx, login string, password string) error
	GetUserByLogin(ctx context.Context, tx *sql.Tx, login string) (*userModel.User, error)
	Exists(ctx context.Context, tx *sql.Tx, login string) (bool, error)
}

type userRepository struct {
	db storage.Storage
}

func NewUserRepository(storage storage.Storage) UserRepository {
	return &userRepository{db: storage}
}

func generateSalt() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

func (r *userRepository) CreateUser(ctx context.Context, tx *sql.Tx, login string, password string) error {
	salt, err := generateSalt()
	if err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	sqlQuery := "INSERT INTO users (login, password, salt) VALUES ($1, $2, $3)"

	if tx == nil {
		_, err := tx.ExecContext(ctx, sqlQuery, login, hashedPassword, salt)
		return err
	}

	_, err = r.db.Exec(ctx, sqlQuery, login, hashedPassword, salt)
	return err
}

func (r *userRepository) GetUserByLogin(ctx context.Context, tx *sql.Tx, login string) (*userModel.User, error) {
	sqlSelect := "SELECT id, login, password, salt, created_at FROM users WHERE login = $1"
	var user userModel.User

	if tx == nil {
		err := tx.QueryRowContext(ctx, sqlSelect, login).Scan(&user.ID, &user.Login, &user.Password, &user.Salt, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		return &user, nil
	}

	err := r.db.QueryRow(ctx, sqlSelect, login).Scan(&user.ID, &user.Login, &user.Password, &user.Salt, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Exists(ctx context.Context, tx *sql.Tx, login string) (bool, error) {
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE login = $1)"

	var exists bool
	var err error

	if tx != nil {
		err = tx.QueryRowContext(ctx, query, login).Scan(&exists)
	} else {
		err = r.db.QueryRow(ctx, query, login).Scan(&exists)
	}

	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *userRepository) GetDB() storage.Storage {
	return r.db
}
