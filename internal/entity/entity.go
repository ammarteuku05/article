package entity

import (
	"time"
)

type Author struct {
	ID        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type Article struct {
	ID        string    `json:"id" db:"id"`
	AuthorID  string    `json:"author_id" db:"author_id"`
	Title     string    `json:"title" db:"title"`
	Body      string    `json:"body" db:"body"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
