package pg

import "word/internal/entities"

func (repo *Repository) CreateWord(Word entities.Word) error {
	query := `
    INSERT INTO words (id, title, description, from_language, to_language, type, created_at, updated_at, user_id)
    VALUES (:id, :title, :description, :from_language, :to_language, :type, :created_at, :updated_at, :user_id)
    `
	_, err := repo.NamedExec(query, Word)
	return err
}
