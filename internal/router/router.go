package router

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	// Middlewares Globais
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Middleware de timeout para todas as rotas
	r.Use(middleware.Timeout(60 * time.Second))

	// Rate Limit para IPs
	r.Use(httprate.LimitByIP(100, time.Minute))

	return r
}
