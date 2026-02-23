package bootstrap

import (
	"database/sql"
	"log/slog"

	"github.com/Turgho/GoFlowDesk/internal/interfaces/http"
)

type Server struct {
	db  *sql.DB
	log *slog.Logger
	cfg *Config
}

func NewServer(db *sql.DB, log *slog.Logger, cfg *Config) *Server {
	return &Server{db: db, log: log, cfg: cfg}
}

func (s *Server) Start() error {
	r := http.SetupRouter(s.db, s.log)
	s.log.Info("server running", "port", s.cfg.ServerPort)
	return r.Run(":" + s.cfg.ServerPort)
}
