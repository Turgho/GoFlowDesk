package bootstrap

import (
	"fmt"
	"os"
)

type Config struct {
	Environment string
	DatabaseURL string
	ServerPort  string
}

func Load() (*Config, error) {
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "development"
	}

	db := os.Getenv("DATABASE_URL")
	if db == "" {
		return nil, fmt.Errorf("DATABASE_URL is required")
	}

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		Environment: env,
		DatabaseURL: db,
		ServerPort:  port,
	}, nil
}
