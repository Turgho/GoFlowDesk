package main

import (
	"os"

	"github.com/Turgho/GoFlowDesk/internal/app"
	"github.com/Turgho/GoFlowDesk/internal/infrastructure/logging"
	"go.uber.org/zap"
)

func main() {
	debug := os.Getenv("ENVIRONMENT") != "production"

	// Cria logger primeiro
	appLogger, err := logging.NewLogger(debug)
	if err != nil {
		panic(err)
	}
	defer appLogger.Sync()

	appLogger.Info("Iniciando aplicação...")

	// Cria app passando o logger
	app := app.NewApp(appLogger)

	// Roda servidor
	if err := app.Run(); err != nil {
		appLogger.Error("Erro ao iniciar app", zap.Error(err))
		os.Exit(1)
	}
}
