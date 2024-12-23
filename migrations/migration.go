package migrations

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Could not create postgres driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations/sql",
		"postgres", driver)
	if err != nil {
		log.Fatalf("Could not create migrate instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Could not run up migrations: %v", err)
	}

	log.Println("Migrations ran successfully")
}
