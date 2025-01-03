package db

import (
	"database/sql"
	"log"

	"github.com/QDEX-Core/oneart-identity-service/internal/config"
	_ "github.com/lib/pq" // Postgres driver
)

func NewDB(cfg *config.Config) (*sql.DB, error) {
	// Log the DSN being used
	log.Printf("Connecting to DB with DSN: %s", cfg.DBDSN)
	db, err := sql.Open("postgres", cfg.DBDSN)
	if err != nil {
		return nil, err
	}

	// Test the DB connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Successfully connected to the database!")
	return db, nil
}
