package bootstrap

import (
	"context"
	"log/slog"
	"os"

	"github.com/Turgho/GoFlowDesk/internal/infrastructure/database"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func InitConfig(ctx context.Context, logger *slog.Logger) (*pgxpool.Pool, error) {
	logger.With("Bootstrap", "Initialize").Info("Initializing application")

	err := godotenv.Load()
	if err != nil {
		logger.Error("Failed to load .env file", "error", err)
		return nil, err
	}
	databaseURL := os.Getenv("DATABASE_URL")

	db, err := database.DatabaseConnection(ctx, databaseURL, logger)
	if err != nil {
		return nil, err

	}
	return db, nil
}
