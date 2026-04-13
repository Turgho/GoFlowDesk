package health

import (
	"net/http"

	handlerpkg "github.com/Turgho/GoFlowDesk/internal/handler"
)

type HealthHandler interface {
	Check(w http.ResponseWriter, r *http.Request)
}

type healthHandler struct{}

func NewHealthHandler() HealthHandler {
	return &healthHandler{}
}

func (h *healthHandler) Check(w http.ResponseWriter, r *http.Request) {
	response := handlerpkg.APIResponse{
		Data: map[string]any{
			"status":  "ok",
			"service": "goflowdesk-api",
			"version": "1.0.0",
		},
	}

	_ = response.Send(w, http.StatusOK)
}
