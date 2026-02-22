package database

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func SetupDatabase(databaseURL string) (*sql.DB, error) {
	// Connect to the database
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		panic(err)
	}

	// Verify the connection is successful
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	return db, nil
}
