package main

import (
	"log/slog"
	"os"

	"github.com/Turgho/GoFlowDesk/internal/bootstrap"
)

func main() {
	// Configura o logger para JSON e define como padrão
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	app := bootstrap.NewApp()
	if err := app.Run(); err != nil {
		logger.Error("Failed to run the application", "error", err)
		os.Exit(1)
	}
}
