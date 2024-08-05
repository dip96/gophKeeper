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
	sqlQuery := "SELECT * FROM login_data WHERE id = $1"

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

func (r *loginDataRepository) GetAllData(ctx context.Context, tx *sql.Tx, page, limit int) ([]models.LoginData, error) {
	offset := (page - 1) * limit
	sqlQuery := `
        SELECT id, user_id, username, password, url, notes 
        FROM login_data 
        ORDER BY id 
        LIMIT $1 OFFSET $2
    `

	var rows *sql.Rows
	var err error

	if tx != nil {
		rows, err = tx.QueryContext(ctx, sqlQuery, limit, offset)
	} else {
		rows, err = r.db.Query(ctx, sqlQuery, limit, offset)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var loginData []models.LoginData

	for rows.Next() {
		var data models.LoginData
		err := rows.Scan(
			&data.ID,
			&data.EntryID,
			&data.Username,
			&data.Password,
			&data.URL,
			&data.Notes,
		)
		if err != nil {
			return nil, err
		}
		loginData = append(loginData, data)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return loginData, nil
}

func (r *loginDataRepository) SaveData(ctx context.Context, tx *sql.Tx, data models.LoginData) error {
	sqlQuery := "INSERT INTO login_data (user_id, username, password, url, notes, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)"
	if tx != nil {
		_, err := tx.ExecContext(ctx, sqlQuery, data.UserID, data.Username, data.Password, data.URL, data.Notes)
		return err
	}

	_, err := r.db.Exec(ctx, sqlQuery, data.EntryID, data.Username, data.Password, data.URL, data.Notes)
	return err
}

func (r *loginDataRepository) EditData(ctx context.Context, tx *sql.Tx, entityId int, data models.LoginData) error {
	sqlQuery := "UPDATE login_data SET username = $1, password = $2, url = $3, notes = $4, updated_at = CURRENT_TIMESTAMP WHERE id = $5"
	if tx != nil {
		_, err := tx.ExecContext(ctx, sqlQuery, data.Username, data.Password, data.URL, data.Notes, entityId)
		return err
	}

	_, err := r.db.Exec(ctx, sqlQuery, data.Username, data.Password, data.URL, data.Notes, entityId)
	return err
}

func (r *loginDataRepository) DeleteData(ctx context.Context, tx *sql.Tx, entityId int) error {
	sqlQuery := "DELETE FROM login_data WHERE id = $1"
	if tx != nil {
		_, err := tx.ExecContext(ctx, sqlQuery, entityId)
		return err
	}

	_, err := r.db.Exec(ctx, sqlQuery, entityId)
	return err
}
