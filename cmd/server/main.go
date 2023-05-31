package main

import (
	"database/sql"
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/mazlon/gobeyond/internal/router"
)

func main() {
	
	db, err := sql.Open("postgres", "postgres://maz:test_password@172.17.0.2:5432/maz?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	err = migrateDB(db)
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	defer db.Close()
	router.PrepareRoutes(r, *db)
	// Run the server
	r.Run(":8080")
}

func migrateDB(db *sql.DB) error {

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
