package database

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

// conexões e migrações
func DatabaseConnection(ctx context.Context, databaseURL string, logger *slog.Logger) (*pgxpool.Pool, error) {
	logger.With("Database", "Connection").Info("Connecting to database")

	cfg, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		return nil, err
	}

	// Configurações de pool de conexões
	cfg.MaxConns = 10
	cfg.MinConns = 2
	cfg.MaxConnLifetime = 30 * 60 // 30 minutes
	cfg.MaxConnIdleTime = 5 * 60  // 5 minutes

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		logger.Error("Failed to connect to database", "error", err)
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		logger.Error("Failed to ping database", "error", err)
		pool.Close()
		return nil, err
	}

	logger.Info("Database connection established successfully")

	return pool, nil
}
