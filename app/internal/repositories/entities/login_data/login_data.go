package login_data

import (
	"context"
	"database/sql"
	models "gophKeeper/internal/models/entities"
	"gophKeeper/internal/repositories/entities"
	"gophKeeper/internal/storage"
)

type loginDataRepository struct {
	db storage.Storage
}

func NewLoginDataRepository(storage storage.Storage) entities.DataRepository[models.LoginData] {
	return &loginDataRepository{db: storage}
}

func (r *loginDataRepository) GetDataById(ctx context.Context, tx *sql.Tx, entityId int) (models.LoginData, error) {
	var loginData models.LoginData
	sqlQuery := "SELECT id, entry_id, username, password, url, notes, created_at, updated_at FROM login_data WHERE entry_id = $1"

	if tx != nil {
		err := tx.QueryRowContext(ctx, sqlQuery, entityId).Scan(&loginData.ID, &loginData.EntryID, &loginData.Username, &loginData.Password, &loginData.URL, &loginData.Notes, &loginData.CreatedAt, &loginData.UpdatedAt)
		if err != nil {
			return models.LoginData{}, err
		}
		return loginData, nil
	}

	err := r.db.QueryRow(ctx, sqlQuery, entityId).Scan(&loginData.ID, &loginData.EntryID, &loginData.Username, &loginData.Password, &loginData.URL, &loginData.Notes, &loginData.CreatedAt, &loginData.UpdatedAt)
	if err != nil {
		return models.LoginData{}, err
	}
	return loginData, nil
}

func (r *loginDataRepository) SaveData(ctx context.Context, tx *sql.Tx, data models.LoginData) error {
	sqlQuery := "INSERT INTO login_data (entry_id, username, password, url, notes, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)"
	if tx != nil {
		_, err := tx.ExecContext(ctx, sqlQuery, data.EntryID, data.Username, data.Password, data.URL, data.Notes)
		return err
	}

	_, err := r.db.Exec(ctx, sqlQuery, data.EntryID, data.Username, data.Password, data.URL, data.Notes)
	return err
}

func (r *loginDataRepository) EditData(ctx context.Context, tx *sql.Tx, entityId int, data models.LoginData) error {
	sqlQuery := "UPDATE login_data SET username = $1, password = $2, url = $3, notes = $4, updated_at = CURRENT_TIMESTAMP WHERE entry_id = $5"
	if tx != nil {
		_, err := tx.ExecContext(ctx, sqlQuery, data.Username, data.Password, data.URL, data.Notes, entityId)
		return err
	}

	_, err := r.db.Exec(ctx, sqlQuery, data.Username, data.Password, data.URL, data.Notes, entityId)
	return err
}

func (r *loginDataRepository) DeleteData(ctx context.Context, tx *sql.Tx, entityId int) error {
	sqlQuery := "DELETE FROM login_data WHERE entry_id = $1"
	if tx != nil {
		_, err := tx.ExecContext(ctx, sqlQuery, entityId)
		return err
	}

	_, err := r.db.Exec(ctx, sqlQuery, entityId)
	return err
}
