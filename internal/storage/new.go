package storage

import "word/internal/storage/pg"

type Storage struct {
	DB *pg.Repository
}

func New() (*Storage, error) {
	db, err := pg.New()
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
