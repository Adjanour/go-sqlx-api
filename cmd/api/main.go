package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Adjanour/go-sqlx-api/internal/database"
	"github.com/Adjanour/go-sqlx-api/internal/handlers"
	"github.com/Adjanour/go-sqlx-api/internal/middleware"
)

func main() {
	// Load configuration from environment
	cfg := loadConfig()

	// Initialize database
	db, err := database.NewConnection(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Run migrations
	if err := database.RunMigrations(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Initialize handlers
	h := handlers.NewHandlers(db)

	// Setup routes
	mux := http.NewServeMux()

	// Health check endpoint
	mux.HandleFunc("/health", h.HealthCheck)

	// API routes
	mux.HandleFunc("/api/v1/users", h.HandleUsers)
	mux.HandleFunc("/api/v1/users/", h.HandleUser)

	// Apply middleware
	handler := middleware.Logger(
		middleware.Recovery(
			middleware.CORS(mux),
		),
	)

	// Create server
	srv := &http.Server{
		Addr:         cfg.ServerAddress,
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Server starting on %s", cfg.ServerAddress)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Server is shutting down...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited properly")
}

// Config holds application configuration
type Config struct {
	ServerAddress string
	DatabaseURL   string
}

// loadConfig loads configuration from environment variables
func loadConfig() *Config {
	serverAddr := os.Getenv("SERVER_ADDRESS")
	if serverAddr == "" {
		serverAddr = ":8080"
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		// Default to PostgreSQL, but can be changed to any database sqlx supports
		dbURL = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			getEnvOrDefault("DB_HOST", "localhost"),
			getEnvOrDefault("DB_PORT", "5432"),
			getEnvOrDefault("DB_USER", "postgres"),
			getEnvOrDefault("DB_PASSWORD", "postgres"),
			getEnvOrDefault("DB_NAME", "api_db"),
		)
	}

	return &Config{
		ServerAddress: serverAddr,
		DatabaseURL:   dbURL,
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
