package http

import (
	"database/sql"
	"log/slog"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB, log *slog.Logger) *gin.Engine {
	router := gin.Default()

	healthHandler := NewHealthHandler(db, log)

	v1 := router.Group("/api/v1")
	{
		v1.GET("/live", healthHandler.Liveness)
		v1.GET("/ready", healthHandler.Readiness)
	}

	return router
}
