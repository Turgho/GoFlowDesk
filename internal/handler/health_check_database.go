package handler

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
	DB *sql.DB
}

type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func NewHealthHandler(db *sql.DB) *HealthHandler {
	return &HealthHandler{
		DB: db,
	}
}

func (h *HealthHandler) HealthCheckDatabase(c *gin.Context) {
	err := h.DB.Ping()
	if err != nil {
		c.JSON(http.StatusInternalServerError, HealthResponse{
			Status:  "error",
			Message: "Database connection failed",
		})
		return
	}

	c.JSON(http.StatusOK, HealthResponse{
		Status:  "success",
		Message: "Database connection is healthy",
	})
}
