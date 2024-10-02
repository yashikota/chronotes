package db

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load(fmt.Sprintf(".env.%s", os.Getenv("GO_ENV")))
	if err != nil && !os.IsNotExist(err) {
		slog.Error("Failed to load .env file:" + err.Error())
	}

	// Connect to the database
	pgpw := os.Getenv("POSTGRES_PASSWORD")
	dsn := fmt.Sprintf("postgres://postgres:%s@db:5432/chronotes?sslmode=disable", pgpw)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error("Failed to connect to database:" + err.Error())
	}

	slog.Info("Connected to database")

	// Migrate the database
	Migration(DB)

	slog.Info("Migrated the database")

	// Seed the database
	// Seed(DB)

	slog.Info("Database connection initialized")
}
