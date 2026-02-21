package router

import (
	"database/sql"

	"github.com/Turgho/GoFlowDesk/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	h := handler.NewHealthHandler(db)

	v1 := r.Group("/api/v1")
	{
		v1.GET("/health-check/database", h.HealthCheckDatabase)
	}

	return r
}
