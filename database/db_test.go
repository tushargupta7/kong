package database

import (
	"bytes"
	"errors"
	"github.com/gofiber/fiber/v2/log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock" // Library for mocking sql.DB
)

func TestInitDB_DefaultConnectionString(t *testing.T) {
	// Mock the database connection
	db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}
	defer db.Close()

	// Mock DB Ping behavior to succeed
	mock.ExpectPing().WillReturnError(nil)

	// Override the global DB variable for testing
	DB = db

	// Call InitDB with no arguments to use the default connection string
	InitDB()

	// Ensure no unmet expectations
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unmet expectations: %v", err)
	}
}

func TestInitDB_CustomConnectionString(t *testing.T) {
	// Mock the database connection
	db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}
	defer db.Close()

	// Mock DB Ping behavior to succeed
	mock.ExpectPing().WillReturnError(nil)

	// Override the global DB variable for testing
	DB = db

	// Call InitDB with a custom connection string
	customConnStr := "postgres://postres:password@localhost:5432/postgres?sslmode=disable"
	InitDB(customConnStr)

	// Ensure no unmet expectations
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unmet expectations: %v", err)
	}
}

func TestInitDB_ConnectionError(t *testing.T) {
	// Mock the database connection with MonitorPingsOption enabled
	db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}
	defer db.Close()

	// Mock DB Ping behavior to return an error
	mock.ExpectPing().WillReturnError(errors.New("mock ping error"))

	// Override the global DB variable for testing
	DB = db

	// Capture log output
	var buf bytes.Buffer
	log.SetOutput(&buf)

	// Call InitDB with no arguments, expecting it to fail
	defer func() {
		// Check if log.Fatalf was called and verify the error message
		if r := recover(); r == nil {
			t.Errorf("expected InitDB to call log.Fatalf but it did not")
		} else {
			// Check if the captured log output contains the expected error message
			if !bytes.Contains(buf.Bytes(), []byte("Failed to ping database: mock ping error")) {
				t.Errorf("expected log output to contain 'Failed to ping database: mock ping error', but got %s", buf.String())
			}
		}
	}()

	// Call InitDB, which should cause a panic due to the failed ping
	InitDB()
}
