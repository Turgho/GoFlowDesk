package bootstrap

import (
	"fmt"

	"github.com/Turgho/GoFlowDesk/internal/infrastructure/database"
	"github.com/Turgho/GoFlowDesk/internal/infrastructure/logger"
)

// Run inicializa a aplicação inteira
func Run() error {
	// Carrega config
	cfg, err := Load()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	// Logger
	log := logger.InitLogger(cfg.Environment).
		With("service", "api").
		With("env", cfg.Environment)
	log.Info("starting GoFlowDesk API server")

	// Database
	dbConn, err := database.SetupDatabase(cfg.DatabaseURL)
	if err != nil {
		log.Error("failed to set up database", "error", err)
		return err
	}
	defer dbConn.Close()

	gormDB := database.SetupGorm(dbConn)
	if err := database.AutoMigrate(gormDB); err != nil {
		log.Error("failed to auto-migrate DB", "error", err)
		return err
	}

	// HTTP Server
	srv := NewServer(dbConn, log, cfg)
	return srv.Start()
}
