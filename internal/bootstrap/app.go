package bootstrap

import (
	httpadapter "github.com/Turgho/GoFlowDesk/internal/adapters/http"
	"github.com/Turgho/GoFlowDesk/internal/infrastructure"
)

type App struct {
	server *infrastructure.Server
}

func NewApp() *App {
	router := httpadapter.NewRouter()

	server := infrastructure.NewServer(":8080", router)

	return &App{
		server: server,
	}
}

func (a *App) Run() error {
	return a.server.Start()
}
