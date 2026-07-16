package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var dbPool *pgxpool.Pool

func ConnectToPostgres() {
	dbUrl := os.Getenv("POSTGRES_URL")
	if dbUrl == "" {
		log.Fatalf("POSTGRES_URL environment variable is not set")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	var err error
	dbPool, err = pgxpool.New(ctx, dbUrl)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	if err = dbPool.Ping(ctx); err != nil {
		log.Fatalf("Unable to ping database: %v", err)
	}

	log.Println("Connected to PostgreSQL database successfully")
}

func GetDB() *pgxpool.Pool {
	if dbPool == nil {
		log.Fatal("Database connection pool is not initialized")
	}
	return dbPool
}

func ClosePostgres() {
	if dbPool != nil {
		dbPool.Close()
	}

}
