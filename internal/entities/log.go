package entities

import "time"

type Log struct {
	ID        int       `json:"id" db:"id"`
	Type      string    `json:"type" db:"type"`
	Data      string    `json:"data" db:"data"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
