package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/mazlon/gobeyond/internal/messaging"
	"github.com/mazlon/gobeyond/internal/models"
	"github.com/mazlon/gobeyond/internal/router"
)

func main() {
	dbConnection := os.Getenv("DATABASE_URL")
	log.Println(dbConnection)
	db, err := sql.Open("postgres", dbConnection)
	if err != nil {
		log.Fatal(err)
	}
	err = models.MigrateDB(db)
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()
	defer db.Close()
	ctx, shutdown := context.WithCancel(context.Background())
	defer shutdown()
	pgxCfg, err := pgxpool.ParseConfig(dbConnection)
	if err != nil {
		log.Println(err)
	}
	connectionPool, err := pgxpool.NewWithConfig(ctx, pgxCfg)
	if err != nil {
		log.Fatal(err)
	}
	defer connectionPool.Close()
	gue, err := messaging.Queue(ctx, connectionPool)
	if err != nil {
		log.Println("Error in Queue system", err)
	}
	router.NewGbtServer(db, r, gue)

	// Run the server
	r.Run(":8080")
}
