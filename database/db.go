package database

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/tushargupta7/kong/errors"
	"log"
)

var DB *sql.DB

func InitDB(connStr ...string) *errors.AppError {
	// Default connection string
	defaultConnStr := "postgres://postgres:password@localhost:5432/kong?sslmode=disable"

	// Use the provided connection string if available, otherwise default
	var finalConnStr string
	if len(connStr) > 0 && connStr[0] != "" {
		finalConnStr = connStr[0]
	} else {
		finalConnStr = defaultConnStr
	}

	if DB == nil {
		var err error
		DB, err = sql.Open("postgres", finalConnStr)
		if err != nil {
			return errors.New(fiber.StatusInternalServerError, errors.ErrConnectDatabase, err, map[string]interface{}{})
		}
	}

	if err := DB.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	return nil
}
