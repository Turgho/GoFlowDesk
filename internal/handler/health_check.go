package handler

import (
	"context"
	"database/sql"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
	DB     *sql.DB
	Logger *slog.Logger
}

type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

// NewHealthHandler creates a new instance of HealthHandler with the given database connection and logger.
func NewHealthHandler(db *sql.DB, log *slog.Logger) *HealthHandler {
	return &HealthHandler{
		DB:     db,
		Logger: log,
	}
}

// Liveness checks if the application is running and can respond to requests.
func (h *HealthHandler) Liveness(c *gin.Context) {
	c.JSON(http.StatusOK, HealthResponse{
		Status: "alive",
	})
}

// Readiness checks if the database connection is healthy.
func (h *HealthHandler) Readiness(c *gin.Context) {
	// Use a context with timeout to avoid hanging if the database is unresponsive
	ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()

	// Check database connectivity
	if err := h.DB.PingContext(ctx); err != nil {
		c.JSON(http.StatusInternalServerError, HealthResponse{
			Status:  "error",
			Message: "database unreachable",
		})
		h.Logger.Error("Database ping failed", "error", err)
		return
	}

	c.JSON(http.StatusOK, HealthResponse{
		Status: "ready",
	})
}
