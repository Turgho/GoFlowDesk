package logger

import (
	"log/slog"
	"os"
)

// InitLogger initializes a new logger based on the provided environment.
func InitLogger(env string) *slog.Logger {
	if env == "development" {
		return slog.New(slog.NewTextHandler(os.Stdout, nil))
	}
	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}
