package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/mazlon/gobeyond/internal/models"
	"github.com/mazlon/gobeyond/internal/router"
)

func main() {

	db, err := sql.Open("postgres", "postgres://maz:test_password@172.17.0.2:5432/maz?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	err = models.MigrateDB(db)
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	defer db.Close()
	router.NewGbtServer(db, r)

	// Run the server
	r.Run(":8080")
}
