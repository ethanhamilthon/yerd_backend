package storage

import (
	"word/internal/storage/sqlite"
)

type Storage struct {
	DB *sqlite.Repository
}

func New() (*Storage, error) {
	db, err := sqlite.New()
	if err != nil {
		return &Storage{}, err
	}
	return &Storage{
		DB: db,
	}, nil
}

func (repo *Storage) CloseConnections() {
	repo.DB.Close()
}
