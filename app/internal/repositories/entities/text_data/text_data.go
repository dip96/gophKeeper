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
	sqlQuery := "SELECT id, entry_id, text, created_at, updated_at FROM text_data WHERE entry_id = $1"

	if tx != nil {
		err := tx.QueryRowContext(ctx, sqlQuery, entityId).Scan(&textData.ID, &textData.EntryID, &textData.Text, &textData.CreatedAt, &textData.UpdatedAt)
		if err != nil {
			return models.TextData{}, err
		}
		return textData, nil
	}

	err := r.db.QueryRow(ctx, sqlQuery, entityId).Scan(&textData.ID, &textData.EntryID, &textData.Text, &textData.CreatedAt, &textData.UpdatedAt)
	if err != nil {
		return models.TextData{}, err
	}
	return textData, nil
}

func (r *textDataRepository) SaveData(ctx context.Context, tx *sql.Tx, data models.TextData) error {
	sqlQuery := "INSERT INTO text_data (entry_id, text, created_at, updated_at) VALUES ($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)"
	if tx != nil {
		_, err := tx.ExecContext(ctx, sqlQuery, data.EntryID, data.Text)
		return err
	}

	_, err := r.db.Exec(ctx, sqlQuery, data.EntryID, data.Text)
	return err
}

func (r *textDataRepository) EditData(ctx context.Context, tx *sql.Tx, entityId int, data models.TextData) error {
	sqlQuery := "UPDATE text_data SET text = $1, updated_at = CURRENT_TIMESTAMP WHERE entry_id = $2"
	if tx != nil {
		_, err := tx.ExecContext(ctx, sqlQuery, data.Text, entityId)
		return err
	}

	_, err := r.db.Exec(ctx, sqlQuery, data.Text, entityId)
	return err
}

func (r *textDataRepository) DeleteData(ctx context.Context, tx *sql.Tx, entityId int) error {
	sqlQuery := "DELETE FROM text_data WHERE entry_id = $1"
	if tx != nil {
		_, err := tx.ExecContext(ctx, sqlQuery, entityId)
		return err
	}

	_, err := r.db.Exec(ctx, sqlQuery, entityId)
	return err
}
