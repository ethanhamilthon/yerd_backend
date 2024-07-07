package sqlite

import (
	"word/internal/entities"

	sq "github.com/Masterminds/squirrel"
)

func (repo *Repository) CreateUser(User entities.User) error {
	query, args, err := sq.Insert("users").
		Columns("id", "name", "full_name", "email", "avatar", "language").
		Values(User.ID, User.Name, User.FullName, User.Email, User.Avatar, User.Language).
		ToSql()
	if err != nil {
		return err
	}
	_, err = repo.db.Exec(query, args...)
	return err
}

func (repo *Repository) UserByEmail(email string) (entities.User, error) {
	query, args, err := sq.Select("id", "name", "full_name", "email", "avatar", "language", "created_at").
		From("users").
		Where(sq.Eq{"email": email}).
		ToSql()
	if err != nil {
		return entities.User{}, err
	}

	var user entities.User
	err = repo.db.Get(&user, query, args...)
	return user, err
}

func (repo *Repository) Languages(userID string) ([]entities.Language, error) {
	query, args, err := sq.Select("id", "user_id", "name", "created_at").
		From("languages").
		Where(sq.Eq{"user_id": userID}).
		ToSql()
	if err != nil {
		return nil, err
	}

	languages := make([]entities.Language, 0)
	err = repo.db.Select(&languages, query, args...)
	return languages, err
}

func (repo *Repository) UpdateUserLanguage(userLanguage string, userID string) error {
	query, args, err := sq.Update("users").
		Set("language", userLanguage).
		Where(sq.Eq{"id": userID}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = repo.db.Exec(query, args...)
	return err
}

func (repo *Repository) CreateLanguages(languages []entities.Language) error {
	insertQuery := sq.Insert("languages").
		Columns("id", "user_id", "name", "created_at")
	for _, language := range languages {
		insertQuery = insertQuery.Values(language.ID, language.UserID, language.LanguageName, language.CreatedAt)
	}
	query, args, err := insertQuery.ToSql()
	if err != nil {
		return err
	}
	_, err = repo.db.Exec(query, args...)
	return err
}
