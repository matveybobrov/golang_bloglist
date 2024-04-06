package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func CreateDatabase() (*sql.DB, error) {
	dbURL := os.Getenv("DATABASE_URL")
	// err returned only if dbURL is not well-formed
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}
	// check if the connection established
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
