package sqlite

import (
	"time"
	"word/internal/entities"

	sq "github.com/Masterminds/squirrel"
)

func (repo *Repository) AddLog(Type, Data string, Time time.Time) error {
	sql, agrs, err := sq.Insert("logs").Columns("type", "data", "created_at").Values(Type, Data, Time).ToSql()
	if err != nil {
		return err
	}
	_, err = repo.logdb.Exec(sql, agrs...)
	return err
}

func (repo *Repository) GetLogs(Type string) ([]entities.Log, error) {
	query, args, err := sq.Select("id", "type", "data", "created_at").
		From("logs").Where(sq.Eq{"type": Type}).OrderBy("created_at DESC").ToSql()
	if err != nil {
		return []entities.Log{}, nil
	}
	logs := make([]entities.Log, 0, 0)
	err = repo.logdb.Select(&logs, query, args...)
	if err != nil {
		return []entities.Log{}, err
	}
	return logs, nil

}
