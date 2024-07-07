package pg

func (repo *Repository) AddLog(Type, Data, Time string) error {
	query := `
    INSERT INTO logs (type, data, created_at)
    VALUES ($1, $2, $3)
     `
	_, err := repo.Exec(query, Type, Data, Time)
	return err
}
