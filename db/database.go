package db

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
)

var DB *sql.DB

func CreateDatabase() (*sql.DB, error) {
	godotenv.Load()

	dbURL := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
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
