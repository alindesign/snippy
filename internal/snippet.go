package internal

import "time"

type Snippet struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	ID        string
	Filename  string
	Extension string
	Contents  string
}
