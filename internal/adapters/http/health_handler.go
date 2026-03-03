package httpadapter

import (
	"net/http"
	"time"
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

	_ = WriteJSON(w, http.StatusOK, response, nil)
}
