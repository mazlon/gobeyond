package models

import (
	"database/sql"
	"errors"
	"log"

	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func MigrateDB(db *sql.DB) error {
	log.Print("Starting to migrate DB")
	driver, err := postgres.WithInstance(db, &postgres.Config{
		DatabaseName: "maz",
	})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		return err
	}
	// m.Log = new(migrateLogger)

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
