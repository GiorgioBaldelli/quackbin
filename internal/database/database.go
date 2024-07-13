package database

import (
	"database/sql"

	_ "github.com/marcboeker/go-duckdb"
)

func InitDB(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("duckdb", dbPath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS pastes (
		id VARCHAR PRIMARY KEY,
		content TEXT,
		is_private BOOLEAN,
		password_hash VARCHAR
	)`)
	if err != nil {
		return nil, err
	}

	return db, nil
}
