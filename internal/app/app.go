package app

import (
	"context"
	"os"

	"github.com/Turgho/GoFlowDesk/internal/handler"
	"github.com/Turgho/GoFlowDesk/internal/infrastructure"
	"github.com/Turgho/GoFlowDesk/internal/infrastructure/database"
	"go.uber.org/zap"
)

type App struct {
	server   *infrastructure.Server
	database *database.PostgresDB
	Log      *zap.Logger
}

// NewApp inicializa a aplicação, configurando o banco de dados, os handlers e o servidor HTTP.
func NewApp(log *zap.Logger) *App {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL environment variable not set")
	}

	// Database
	postgresDB, err := database.NewPostgresDB(context.Background(), databaseURL, log)
	if err != nil {
		log.Fatal("Erro ao conectar no banco", zap.Error(err))
	}

	// Router
	router := handler.NewRouter()

	// Server
	server := infrastructure.NewServer(":8080", router, log)

	log.Info("Aplicação inicializada com sucesso")

	return &App{
		server:   server,
		database: postgresDB,
		Log:      log,
	}
}

func (a *App) Run() error {
	a.Log.Info("Iniciando servidor HTTP na porta 8080")
	return a.server.Start()
}
