package entities

import "time"

type CreditCardData struct {
	ID             int       `json:"id"`
	UserID         int       `json:"user_id"`
	CardNumber     string    `json:"card_number"`
	CardholderName string    `json:"cardholder_name"`
	ExpirationDate time.Time `json:"expiration_date"`
	CVV            string    `json:"cvv"`
	Notes          string    `json:"notes"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
