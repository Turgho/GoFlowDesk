package database

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type PostgresDB struct {
	Pool *pgxpool.Pool
}

func NewPostgresDB(ctx context.Context, databaseURL string, log *zap.Logger) (*PostgresDB, error) {
	if log == nil {
		return nil, errors.New("logger não inicializado")
	}

	pool, err := databaseConnection(ctx, databaseURL, log)
	if err != nil {
		return nil, err
	}

	return &PostgresDB{
		Pool: pool,
	}, nil
}

// conexões e migrações
func databaseConnection(ctx context.Context, databaseURL string, log *zap.Logger) (*pgxpool.Pool, error) {
	log.Info("Tentando conectar ao banco de dados")

	cfg, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		return nil, err
	}

	// Configurações de pool de conexões
	cfg.MaxConns = 10
	cfg.MinConns = 2
	cfg.MaxConnLifetime = 30 * 60 // 30 minutes
	cfg.MaxConnIdleTime = 5 * 60  // 5 minutes

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		log.Error("Erro ao conectar ao banco de dados", zap.String("error:", err.Error()))
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		log.Error("Falha ao testar conexão com o banco de dados", zap.String("error:", err.Error()))
		pool.Close()
		return nil, err
	}

	log.Info("Conexão com o banco de dados estabelecida com sucesso")
	return pool, nil
}
