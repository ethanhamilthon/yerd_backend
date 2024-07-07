package sqlite

func (repo *Repository) runMigration() error {
	_, err := repo.db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			full_name TEXT,
			email TEXT NOT NULL UNIQUE,
			avatar TEXT,
			language TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			role TEXT
		)
	`)
	if err != nil {
		return err
	}

	// Создание таблицы languages
	_, err = repo.db.Exec(`
		CREATE TABLE IF NOT EXISTS languages (
			id TEXT PRIMARY KEY,
			user_id TEXT,
			name TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY(user_id) REFERENCES users(id)
		)
	`)
	if err != nil {
		return err
	}

	// Создание таблицы words
	_, err = repo.db.Exec(`
		CREATE TABLE IF NOT EXISTS words (
			id TEXT PRIMARY KEY,
			title TEXT NOT NULL,
			description TEXT,
			from_language TEXT,
			to_language TEXT,
			type TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			user_id TEXT,
			is_deleted INTEGER DEFAULT 0,
			FOREIGN KEY(user_id) REFERENCES users(id)
		)
	`)
	if err != nil {
		return err
	}

	// Создание таблицы logs
	_, err = repo.logdb.Exec(`
		CREATE TABLE IF NOT EXISTS logs (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			type TEXT,
			data TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return err
	}

	return nil
}
