package bootstrap

import (
	"context"
	"log/slog"
	"os"
)

func Run() {
	ctx := context.Background()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := InitConfig(ctx, logger)
	if err != nil {
		logger.Error("Failed to initialize application", "error", err)
		os.Exit(1)
		return
	}
	defer db.Close()
}
