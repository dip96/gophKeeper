package entities

import "time"

type LoginData struct {
	ID        int64     `json:"id"`
	EntryID   int       `json:"entry_id"`
	UserID    int       `json:"user_id"`
	Username  string    `json:"username"`
	Password  []byte    `json:"password"`
	URL       string    `json:"url"`
	Notes     string    `json:"notes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
