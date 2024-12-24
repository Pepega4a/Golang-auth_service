package database

import (
	"database/sql"
	"os"
)

func InitDatabase() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	createTableSQL := `
       CREATE TABLE IF NOT EXISTS refresh_tokens (
           user_id VARCHAR(255) PRIMARY KEY,
           token_hash VARCHAR(255) NOT NULL,
           ip VARCHAR(45) NOT NULL
       );`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		return nil, err
	}
	return db, nil
}
