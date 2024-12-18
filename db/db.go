package db

import (
	"database/sql"
	//"github.com/lib/pq" // PostgreSQL driver
	"log"
)

var DB *sql.DB

func InitDB() {
	var err error
	connStr := "postgres://user:password@localhost:5432/services_db?sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Connected to database successfully!")
}
