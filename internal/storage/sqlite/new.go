package sqlite

import (
	"word/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Repository struct {
	db    *sqlx.DB
	logdb *sqlx.DB
}

func New() (*Repository, error) {
	db, err := sqlx.Open("sqlite3", config.SQLiteMainPath)
	if err != nil {
		return &Repository{}, err
	}

	err = db.Ping()
	if err != nil {
		return &Repository{}, err
	}
	logdb, err := sqlx.Open("sqlite3", config.SQLiteLogPath)
	if err != nil {
		return &Repository{}, err
	}

	err = logdb.Ping()
	if err != nil {
		return &Repository{}, err
	}

	newDB := &Repository{
		db:    db,
		logdb: logdb,
	}
	err = newDB.runMigration()
	if err != nil {
		return &Repository{}, err
	}
	return newDB, nil
}

func (repo *Repository) Close() {
	repo.db.Close()
	repo.logdb.Close()
}
