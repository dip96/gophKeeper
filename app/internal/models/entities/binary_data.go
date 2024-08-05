package entities

import "time"

type BinaryData struct {
	ID        int
	EntryID   int
	Path      string
	Notes     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
