package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func RunMigrations() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	dsn := os.Getenv("DATABASE_URL")

	m, err := migrate.New("file://migrations", dsn)
	if err != nil {
		log.Fatal("Migrations failed: ", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Migrate failed: ", err)
	}

	log.Println("Migrations done")
}

func Connect() (*pgxpool.Pool, error) {
	dsn := os.Getenv("DATABASE_URL")

	dbpool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		return nil, err
	}

	if err := dbpool.Ping(context.Background()); err != nil {
		log.Printf("Error pinging db")
		return nil, err
	}

	return dbpool, nil
}
