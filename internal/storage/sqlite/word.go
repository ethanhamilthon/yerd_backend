package sqlite

import (
	"database/sql"
	"time"
	"word/internal/entities"

	sq "github.com/Masterminds/squirrel"
)

func (repo *Repository) CreateWord(word entities.Word) error {
	query, args, err := sq.Insert("words").
		Columns("id", "title", "description", "from_language", "to_language", "type", "created_at", "updated_at", "user_id").
		Values(word.ID, word.Title, word.Description, word.FromLanguage, word.ToLanguage, word.Type, word.CreatedAt, word.UpdatedAt, word.UserID).
		ToSql()
	if err != nil {
		return err
	}

	_, err = repo.db.Exec(query, args...)
	return err
}

func (repo *Repository) Word(ID string) (entities.Word, error) {
	query, args, err := sq.Select("id", "title", "description", "from_language", "to_language", "type", "created_at", "updated_at", "user_id").
		From("words").
		Where(sq.Eq{"id": ID}).
		ToSql()
	if err != nil {
		return entities.Word{}, err
	}

	var word entities.Word
	err = repo.db.Get(&word, query, args...)
	if err != nil {
		return word, err
	}

	return word, nil
}

func (repo *Repository) Words(UserID string) ([]entities.Word, error) {
	query, args, err := sq.Select("id", "title", "description", "from_language", "to_language", "created_at", "updated_at", "user_id").
		From("words").
		Where(sq.Eq{"user_id": UserID}).
		OrderBy("created_at DESC").
		ToSql()
	if err != nil {
		return nil, err
	}

	var words []entities.Word
	err = repo.db.Select(&words, query, args...)
	if err != nil {
		return nil, err
	}

	return words, nil
}

func (repo *Repository) UpdateWord(ID, Title, Description, UserID string, UpdatedAt time.Time) error {
	query, args, err := sq.Update("words").
		SetMap(map[string]interface{}{
			"title":       Title,
			"description": Description,
			"updated_at":  UpdatedAt,
		}).
		Where(sq.Eq{"id": ID, "user_id": UserID}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = repo.db.Exec(query, args...)
	return err
}

func (repo *Repository) DeleteWord(ID, UserID string) error {
	query, args, err := sq.Delete("words").
		Where(sq.Eq{"id": ID, "user_id": UserID}).
		ToSql()
	if err != nil {
		return err
	}

	result, err := repo.db.Exec(query, args...)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
