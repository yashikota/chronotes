package utils

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if os.Getenv("GO_ENV") == "dev" {
		err := godotenv.Load(".env")
		if err != nil {
			panic("Error loading .env file")
		}
		slog.Info("Loaded .env file")
	}
	slog.Info("Loaded environment variables")
}
