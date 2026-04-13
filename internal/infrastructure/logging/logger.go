package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger cria e retorna um logger configurado.
// debugMode: se true, usa formato legível (desenvolvimento), senão usa JSON (produção)
func NewLogger(debugMode bool) (*zap.Logger, error) {
	if debugMode {
		// Logger de desenvolvimento
		return zap.NewDevelopment()
	}

	// Logger de produção (JSON, mais rápido)
	cfg := zap.NewProductionConfig()

	// Exemplo de customização de nível
	cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)

	// Exemplo: customizar timestamp
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	return cfg.Build()
}
