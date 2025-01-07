package main

import (
	"log"
	"net/http"

	"github.com/QDEX-Core/oneart-identity-service/internal/config"
	"github.com/QDEX-Core/oneart-identity-service/internal/db"
	handler "github.com/QDEX-Core/oneart-identity-service/internal/handlers"
	"github.com/QDEX-Core/oneart-identity-service/internal/repository"
	service "github.com/QDEX-Core/oneart-identity-service/internal/services"

	"github.com/gorilla/mux"
	"github.com/rs/cors" // Importing CORS package
)

func main() {
	cfg := config.NewConfig()

	log.Println("Initializing database connection...")
	database, err := db.NewDB(cfg)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer database.Close()

	log.Println("Successfully connected to the database!")

	// Initialize repository, services, and handlers
	userRepo := repository.NewUserRepository(database)
	userService := service.NewUserService(userRepo, cfg.JWTSecret)
	userHandler := handler.NewUserHandler(userService)

	// Set up router
	log.Println("Setting up routes...")
	r := mux.NewRouter()
	r.HandleFunc("/auth/register", userHandler.Register).Methods("POST")
	r.HandleFunc("/auth/login", userHandler.Login).Methods("POST")

	// Add CORS middleware to allow frontend connections
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8081", "http://127.0.0.1:8081"}, // Add allowed frontend origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}).Handler(r)

	// Start the server
	port := "8080"
	log.Printf("Starting OneArt Identity Service on port %s...", port)
	if err := http.ListenAndServe(":"+port, corsMiddleware); err != nil {
		log.Fatal("Failed to start the server:", err)
	}
}
