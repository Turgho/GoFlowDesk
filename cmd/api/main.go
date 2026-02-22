package main

import (
	"os"

	"github.com/Turgho/GoFlowDesk/internal/database"
	"github.com/Turgho/GoFlowDesk/internal/logger"
	"github.com/Turgho/GoFlowDesk/internal/router"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		// não é erro fatal
	}

	// Determine environment
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "development"
	}

	// Initialize logger AFTER env is known
	log := logger.InitLogger(env).
		With("service", "api").
		With("env", env)

	log.Info("starting GoFlowDesk API server")

	// Validate required config
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Error("missing required environment variable", "var", "DATABASE_URL")
		os.Exit(1)
	}

	// Set up database connection
	dbConnection, err := database.SetupDatabase(databaseURL)
	if err != nil {
		log.Error("failed to set up database connection", "error", err)
		os.Exit(1)
	}
	defer dbConnection.Close()

	// Set up GORM
	gormDB := database.SetupGorm(dbConnection)
	if gormDB == nil {
		log.Error("failed to set up GORM")
		os.Exit(1)
	}

	// Auto-migrate database schema
	if err := database.AutoMigrate(gormDB); err != nil {
		log.Error("failed to auto-migrate database schema", "error", err)
		os.Exit(1)
	}

	// Init router with database and logger
	r := router.SetupRouter(dbConnection, log)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	log.Info("server running", "port", port)

	// Start the server
	if err := r.Run(":" + port); err != nil {
		log.Error("server failed", "error", err)
		os.Exit(1)
	}
}
