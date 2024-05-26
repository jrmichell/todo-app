package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// DB is the database connection pool
var DB *sql.DB

// InitDB initializes the database connection
func InitDB(connectionString string) error {
	var err error
	DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		return fmt.Errorf("error opening database: %w", err)
	}

	if err = DB.Ping(); err != nil {
		return fmt.Errorf("error connecting to the database: %w", err)
	}

	fmt.Println("Successfully connected to the database")
	return nil
}
