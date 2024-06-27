package pg

import (
	"database/sql"
	"time"
	"word/internal/entities"
)

func (repo *Repository) CreateWord(Word entities.Word) error {
	query := `
    INSERT INTO words (id, title, description, from_language, to_language, type, created_at, updated_at, user_id)
    VALUES (:id, :title, :description, :from_language, :to_language, :type, :created_at, :updated_at, :user_id)
    `
	_, err := repo.NamedExec(query, Word)
	return err
}

func (repo *Repository) Word(ID string) (entities.Word, error) {
	query := `
  SELECT id, title, description, from_language, to_language, type, created_at, updated_at, user_id
  FROM words
  WHERE id = $1
  `
	var word entities.Word
	err := repo.Get(&word, query, ID)
	if err != nil {
		return word, err
	}

	return word, nil
}

func (repo *Repository) Words(UserID string) ([]entities.Word, error) {
	query := `
		SELECT id, title, description, from_language, to_language, created_at, updated_at, user_id
		FROM words
		WHERE user_id = $1
    ORDER BY created_at DESC
	`

	var words []entities.Word
	err := repo.Select(&words, query, UserID)
	if err != nil {
		return nil, err
	}

	return words, nil
}

func (repo *Repository) UpdateWord(ID, Title, Description, UserID string, UpdatedAt time.Time) error {
	query := `
    UPDATE words
    SET title = :title,
        description = :description,
        updated_at = :updated_at
    WHERE id = :id AND user_id = :user_id
    `

	params := map[string]interface{}{
		"id":          ID,
		"title":       Title,
		"description": Description,
		"user_id":     UserID,
		"updated_at":  UpdatedAt,
	}

	_, err := repo.NamedExec(query, params)
	return err
}

func (repo *Repository) DeleteWord(ID, UserID string) error {
	query := `
        DELETE FROM words
        WHERE id = $1 AND user_id = $2
    `

	result, err := repo.Exec(query, ID, UserID)
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
