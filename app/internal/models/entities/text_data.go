package entities

import "time"

type TextData struct {
	ID        int
	EntryID   int
	Text      string
	Notes     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
