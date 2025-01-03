package config

import (
	"log"
	"os"
)

type Config struct {
	DBDSN     string
	JWTSecret string
}

func NewConfig() *Config {
	dbDSN := os.Getenv("DB_DSN")
	if dbDSN == "" {
		//dbDSN = "postgres://postgres:oneart-secret@35.244.41.139:5432/postgres?sslmode=disable"
		dbDSN = "postgres://postgres:oneart-secret@/postgres?host=/cloudsql/qdex-401002:asia-south1:oneart-postgres"
		log.Printf("DB_DSN not set, defaulting to: %s", dbDSN)
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "supersecret"
		log.Println("JWT_SECRET not set, defaulting to 'supersecret'")
	}

	return &Config{
		DBDSN:     dbDSN,
		JWTSecret: jwtSecret,
	}
}
