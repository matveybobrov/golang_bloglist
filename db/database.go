package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func CreateDatabase() (*sql.DB, error) {
	dbURL := os.Getenv("DATABASE_URL")
	// somehow doesn't return an error if url is incorrect
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}
	// will return an error if url is incorrect
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func Init() error {
	db, err := CreateDatabase()
	if err != nil {
		return err
	}

	DB = db
	return nil
}
