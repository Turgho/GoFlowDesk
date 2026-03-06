package infrastructure

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

type Server struct {
	httpServer *http.Server
	Log        *zap.Logger
}

func NewServer(addr string, handler http.Handler, log *zap.Logger) *Server {
	// Configura o servidor HTTP com timeouts
	return &Server{
		httpServer: &http.Server{
			Addr:         addr,
			Handler:      handler,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  60 * time.Second,
		},
		Log: log,
	}
}

func (s *Server) Start() error {
	// Canal para capturar CTRL+C ou SIGTERM
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Rodar servidor em goroutine
	go func() {
		s.Log.Info("Servidor iniciado em", zap.String("endereço", s.httpServer.Addr))
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.Log.Fatal("Erro ao iniciar servidor", zap.Error(err))
		}
	}()

	// Esperar sinal
	<-stop
	s.Log.Info("Desligando servidor...")

	// Timeout para shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return s.httpServer.Shutdown(ctx)
}
