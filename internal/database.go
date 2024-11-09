package internal

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type (
	DatabaseConnection = *sql.DB
)

func NewDatabase(config Config) (DatabaseConnection, error) {
	log.Printf("connecting to database: %s", config.DatabaseConnection)
	conn, err := sql.Open("sqlite3", config.DatabaseConnection)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
