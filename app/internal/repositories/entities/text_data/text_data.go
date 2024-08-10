package text_data

import (
	"context"
	"database/sql"
	models "gophKeeper/internal/models/entities"
	"gophKeeper/internal/repositories/entities"
	"gophKeeper/internal/storage"
)

type textDataRepository struct {
	db storage.Storage
}

func NewTextDataRepository(storage storage.Storage) entities.DataRepository[models.TextData] {
	return &textDataRepository{db: storage}
}

func (r *textDataRepository) GetDataById(ctx context.Context, tx *sql.Tx, entityId int) (models.TextData, error) {
	var textData models.TextData
	sqlQuery := "SELECT id, text, created_at, updated_at FROM text_data WHERE id = $1"

	if tx != nil {
		err := tx.QueryRowContext(ctx, sqlQuery, entityId).Scan(&textData.ID, &textData.Text, &textData.CreatedAt, &textData.UpdatedAt)
		if err != nil {
			return models.TextData{}, err
		}
		return textData, nil
	}

	err := r.db.QueryRow(ctx, sqlQuery, entityId).Scan(&textData.ID, &textData.Text, &textData.CreatedAt, &textData.UpdatedAt)
	if err != nil {
		return models.TextData{}, err
	}
	return textData, nil
}

func (r *textDataRepository) GetAllData(ctx context.Context, tx *sql.Tx, page, limit int) ([]models.TextData, error) {
	offset := (page - 1) * limit
	sqlQuery := `
        SELECT id, user_id, text 
        FROM text_data 
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

	var textData []models.TextData

	for rows.Next() {
		var data models.TextData
		err := rows.Scan(
			&data.ID,
			&data.UserID,
			&data.Text,
		)
		if err != nil {
			return nil, err
		}
		textData = append(textData, data)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return textData, nil
}

func (r *textDataRepository) SaveData(ctx context.Context, tx *sql.Tx, data models.TextData) error {
	sqlQuery := "INSERT INTO text_data (text, user_id, created_at, updated_at) VALUES ($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)"
	if tx != nil {
		_, err := tx.ExecContext(ctx, sqlQuery, data.Text, data.UserID)
		return err
	}

	_, err := r.db.Exec(ctx, sqlQuery, data.Text, data.UserID)
	return err
}

func (r *textDataRepository) EditData(ctx context.Context, tx *sql.Tx, entityId int, data models.TextData) error {
	sqlQuery := "UPDATE text_data SET text = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2"
	if tx != nil {
		_, err := tx.ExecContext(ctx, sqlQuery, data.Text, entityId)
		return err
	}

	_, err := r.db.Exec(ctx, sqlQuery, data.Text, entityId)
	return err
}

func (r *textDataRepository) DeleteData(ctx context.Context, tx *sql.Tx, entityId int) error {
	sqlQuery := "DELETE FROM text_data WHERE id = $1"
	if tx != nil {
		_, err := tx.ExecContext(ctx, sqlQuery, entityId)
		return err
	}

	_, err := r.db.Exec(ctx, sqlQuery, entityId)
	return err
}
