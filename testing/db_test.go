package testing

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"

	dbconnection "loggingserver/db"
)

func TestInitDB(t *testing.T) {
	// Create a new mock DB connection and a SQL mock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' occurred when opening a stub database connection", err)
	}
	defer db.Close()

	// Set expectations on the mock for Ping and connection settings
	mock.ExpectPing()

	// Call the InitDB function with the mock DB
	dbpool, err := dbconnection.InitDB()
	if err != nil {
		t.Fatalf("error initializing DB: %v", err)
	}

	// Check if the returned dbpool is not nil
	if dbpool == nil {
		t.Error("expected non-nil dbpool, got nil")
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
