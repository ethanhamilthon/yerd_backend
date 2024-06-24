package pg

import (
	"word/config"
	"word/internal/entities"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repository struct {
	*sqlx.DB
}

func New() (*Repository, error) {
	db, err := sqlx.Open("postgres", config.PgConnStr)
	if err != nil {
		return &Repository{}, err
	}

	err = db.Ping()
	if err != nil {
		return &Repository{}, err
	}
	return &Repository{
		DB: db,
	}, nil
}

func (repo *Repository) GetWordByID(ID string) (entities.Word, error) {
	return entities.Word{}, nil
}
