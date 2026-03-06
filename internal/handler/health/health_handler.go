package health

import (
	"net/http"
	"time"

	"github.com/Turgho/GoFlowDesk/internal/handler/render"
)

type HealthHandler interface {
	Check(w http.ResponseWriter, r *http.Request)
}

type healthHandler struct{}

func NewHealthHandler() HealthHandler {
	return &healthHandler{}
}

func (h *healthHandler) Check(w http.ResponseWriter, r *http.Request) {
	response := map[string]any{
		"status":  "ok",
		"service": "goflowdesk-api",
		"version": "1.0.0",
		"time":    time.Now().UTC(),
	}

	_ = render.WriteJSON(w, http.StatusOK, response, nil)
}
