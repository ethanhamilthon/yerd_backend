package entities

import (
	"time"
)

type WordBasic struct {
	ID           string `db:"id" json:"id"`
	Title        string `db:"title" json:"title"`
	Description  string `db:"description" json:"description"`
	FromLanguage string `db:"from_language" json:"from_language"`
	ToLanguage   string `db:"to_language" json:"to_language"`
	Type         string `db:"type" json:"type"`
}

type Word struct {
	WordBasic
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	UserID    string    `db:"user_id" json:"user_id"`
}

type WordAllResponse struct {
	Language string `json:"language"`
	Words    []Word `json:"words"`
}
