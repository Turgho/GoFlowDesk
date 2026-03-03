package httpadapter

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()

	// ===== Middlewares Globais =====
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// ===== Health (fora da versão) =====
	healthHandler := NewHealthHandler()
	r.Get("/health", healthHandler.Check)

	return r
}
