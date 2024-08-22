package credit_card_data

import (
	"context"
	"database/sql"
	models "gophKeeper/internal/models/entities"
	"gophKeeper/internal/repositories/entities"
	"gophKeeper/internal/storage"
)

type creditCardDataRepository struct {
	db storage.Storage
}

// NewCreditCardDataRepository создаёт новый экземпляр репозитория для данных банковских карт
func NewCreditCardDataRepository(storage storage.Storage) entities.DataRepository[models.CreditCardData] {
	return &creditCardDataRepository{db: storage}
}

// GetDataById получает данные банковской карты по ID
func (r *creditCardDataRepository) GetDataById(ctx context.Context, tx *sql.Tx, entityId int) (models.CreditCardData, error) {
	var creditCardData models.CreditCardData
	sqlQuery := "SELECT id, card_number, cardholder_name, expiration_date, cvv, notes, created_at, updated_at FROM credit_card_data WHERE id = $1"

	if tx != nil {
		err := tx.QueryRowContext(ctx, sqlQuery, entityId).Scan(
			&creditCardData.ID,
			&creditCardData.CardNumber,
			&creditCardData.CardholderName,
			&creditCardData.ExpirationDate,
			&creditCardData.CVV,
			&creditCardData.Notes,
			&creditCardData.CreatedAt,
			&creditCardData.UpdatedAt,
		)
		if err != nil {
			return models.CreditCardData{}, err
		}
		return creditCardData, nil
	}

	err := r.db.QueryRow(ctx, sqlQuery, entityId).Scan(
		&creditCardData.ID,
		&creditCardData.CardNumber,
		&creditCardData.CardholderName,
		&creditCardData.ExpirationDate,
		&creditCardData.CVV,
		&creditCardData.Notes,
		&creditCardData.CreatedAt,
		&creditCardData.UpdatedAt,
	)
	if err != nil {
		return models.CreditCardData{}, err
	}
	return creditCardData, nil
}

// GetAllData получает все данные банковских карт с пагинацией
func (r *creditCardDataRepository) GetAllData(ctx context.Context, tx *sql.Tx, page, limit int) ([]models.CreditCardData, error) {
	offset := (page - 1) * limit
	sqlQuery := `
        SELECT id, card_number, cardholder_name, expiration_date, cvv, notes
        FROM credit_card_data 
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

	var creditCardData []models.CreditCardData

	for rows.Next() {
		var data models.CreditCardData
		err := rows.Scan(
			&data.ID,
			&data.CardNumber,
			&data.CardholderName,
			&data.ExpirationDate,
			&data.CVV,
			&data.Notes,
		)
		if err != nil {
			return nil, err
		}
		creditCardData = append(creditCardData, data)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return creditCardData, nil
}

// SaveData сохраняет данные банковской карты в базу данных
func (r *creditCardDataRepository) SaveData(ctx context.Context, tx *sql.Tx, data models.CreditCardData) error {
	sqlQuery := "INSERT INTO credit_card_data (card_number, cardholder_name, expiration_date, cvv, notes, created_at, updated_at, user_id) VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, $6)"
	if tx != nil {
		_, err := tx.ExecContext(ctx, sqlQuery, data.CardNumber, data.CardholderName, data.ExpirationDate, data.CVV, data.Notes, data.UserID)
		return err
	}

	_, err := r.db.Exec(ctx, sqlQuery, data.CardNumber, data.CardholderName, data.ExpirationDate, data.CVV, data.Notes)
	return err
}

// EditData редактирует данные банковской карты в базе данных
func (r *creditCardDataRepository) EditData(ctx context.Context, tx *sql.Tx, entityId int, data models.CreditCardData) error {
	sqlQuery := "UPDATE credit_card_data SET card_number = $1, cardholder_name = $2, expiration_date = $3, cvv = $4, notes = $5, updated_at = CURRENT_TIMESTAMP WHERE id = $6"
	if tx != nil {
		_, err := tx.ExecContext(ctx, sqlQuery, data.CardNumber, data.CardholderName, data.ExpirationDate, data.CVV, data.Notes, entityId)
		return err
	}

	_, err := r.db.Exec(ctx, sqlQuery, data.CardNumber, data.CardholderName, data.ExpirationDate, data.CVV, data.Notes, entityId)
	return err
}

// DeleteData удаляет данные банковской карты из базы данных
func (r *creditCardDataRepository) DeleteData(ctx context.Context, tx *sql.Tx, entityId int) error {
	sqlQuery := "DELETE FROM credit_card_data WHERE id = $1"
	if tx != nil {
		_, err := tx.ExecContext(ctx, sqlQuery, entityId)
		return err
	}

	_, err := r.db.Exec(ctx, sqlQuery, entityId)
	return err
}
