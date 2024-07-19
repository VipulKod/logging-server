package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	user := os.Getenv("DB_USER")
	if user == "" {
		user = "postgres" // Default user
	}

	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		dbname = "postgres" // Default database name
	}

	sslmode := os.Getenv("DB_SSLMODE")
	if sslmode == "" {
		sslmode = "disable" // Default SSL mode
	}

	connStr := fmt.Sprintf("user=%s dbname=%s sslmode=%s", user, dbname, sslmode)

	// Open a database connection
	dbpool, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Set maximum number of open connections to 25
	dbpool.SetMaxOpenConns(25)
	dbpool.SetMaxIdleConns(10)

	return dbpool, nil
}
