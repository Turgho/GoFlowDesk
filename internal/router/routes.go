package router

import (
	"github.com/go-chi/chi/v5"

	"github.com/Turgho/GoFlowDesk/internal/handler/health"
	"github.com/Turgho/GoFlowDesk/internal/handler/user"
)

// RegisterRoutes attaches application routes to the provided router.
func RegisterRoutes(r chi.Router, uh *user.UserHandler) {
	r.Get("/health", health.NewHealthHandler().Check)
	r.Route("/users", func(r chi.Router) {
		r.Post("/", uh.Create)
	})
}
