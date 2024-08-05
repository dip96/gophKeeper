package binary_data

import (
	"context"
	"database/sql"
	models "gophKeeper/internal/models/entities"
	"gophKeeper/internal/repositories/entities"
	"gophKeeper/internal/storage"
)

type binaryDataRepository struct {
	db storage.Storage
}

func NewBinaryDataRepository(storage storage.Storage) entities.DataRepository[models.BinaryData] {
	return &binaryDataRepository{db: storage}
}

func (r *binaryDataRepository) GetDataById(ctx context.Context, tx *sql.Tx, entityId int) (models.BinaryData, error) {
	var binaryData models.BinaryData
	sqlQuery := "SELECT id, entry_id, path, notes, created_at, updated_at FROM binary_data WHERE entry_id = $1"

	if tx != nil {
		err := tx.QueryRowContext(ctx, sqlQuery, entityId).Scan(&binaryData.ID, &binaryData.EntryID, &binaryData.Path, &binaryData.Notes, &binaryData.CreatedAt, &binaryData.UpdatedAt)
		if err != nil {
			return models.BinaryData{}, err
		}
		return binaryData, nil
	}

	err := r.db.QueryRow(ctx, sqlQuery, entityId).Scan(&binaryData.ID, &binaryData.EntryID, &binaryData.Path, &binaryData.Notes, &binaryData.CreatedAt, &binaryData.UpdatedAt)
	if err != nil {
		return models.BinaryData{}, err
	}
	return binaryData, nil
}

func (r *binaryDataRepository) GetAllData(ctx context.Context, tx *sql.Tx, page, limit int) ([]models.BinaryData, error) {
	var binaryData []models.BinaryData
	return binaryData, nil
}

func (r *binaryDataRepository) SaveData(ctx context.Context, tx *sql.Tx, data models.BinaryData) error {
	sqlQuery := "INSERT INTO binary_data (entry_id, path, notes, created_at, updated_at) VALUES ($1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)"
	if tx != nil {
		_, err := tx.ExecContext(ctx, sqlQuery, data.EntryID, data.Path, data.Notes)
		return err
	}

	_, err := r.db.Exec(ctx, sqlQuery, data.EntryID, data.Path, data.Notes)
	return err
}

func (r *binaryDataRepository) EditData(ctx context.Context, tx *sql.Tx, entityId int, data models.BinaryData) error {
	sqlQuery := "UPDATE binary_data SET path = $1, notes = $2, updated_at = CURRENT_TIMESTAMP WHERE entry_id = $3"
	if tx != nil {
		_, err := tx.ExecContext(ctx, sqlQuery, data.Path, data.Notes, entityId)
		return err
	}

	_, err := r.db.Exec(ctx, sqlQuery, data.Path, data.Notes, entityId)
	return err
}

func (r *binaryDataRepository) DeleteData(ctx context.Context, tx *sql.Tx, entityId int) error {
	sqlQuery := "DELETE FROM binary_data WHERE entry_id = $1"
	if tx != nil {
		_, err := tx.ExecContext(ctx, sqlQuery, entityId)
		return err
	}

	_, err := r.db.Exec(ctx, sqlQuery, entityId)
	return err
}
