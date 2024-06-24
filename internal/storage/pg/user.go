package pg

import "word/internal/entities"

func (repo *Repository) CreateUser(User entities.User) error {
	query := `
    INSERT INTO users (id, name, full_name, email, avatar, language)
    VALUES ($1, $2, $3, $4, $5, $6)
    `
	_, err := repo.Exec(query, User.ID, User.Name, User.FullName, User.Email, User.Avatar, User.Language)
	return err
}

func (repo *Repository) UserByEmail(Email string) (entities.User, error) {
	var user entities.User
	err := repo.Get(&user, "SELECT id, name, full_name, email, avatar, language, created_at FROM users WHERE email = $1", Email)
	return user, err
}

func (repo *Repository) Languages(UserID string) ([]entities.Language, error) {
	languages := make([]entities.Language, 0)
	err := repo.Select(&languages, "SELECT id, user_id, name, created_at FROM languages WHERE user_id = $1", UserID)
	return languages, err
}

func (repo *Repository) UpdateUserLanguage(UserLanguage string, UserID string) error {
	query := `UPDATE users SET language = $2 WHERE id = $1`
	_, err := repo.Exec(query, UserID, UserLanguage)
	return err
}

func (repo *Repository) CreateLanguages(Languages []entities.Language) error {
	query := `INSERT INTO languages (id, user_id, name, created_at)
						VALUES (:id, :user_id, :name, :created_at)`
	_, err := repo.NamedExec(query, Languages)
	return err
}
