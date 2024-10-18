package db

import (
	"errors"
	"fmt"
	"log/slog"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	pgpw := os.Getenv("POSTGRES_PASSWORD")
	dsn := fmt.Sprintf("postgres://postgres:%s@db:5432/chronotes?sslmode=disable", pgpw)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error("Failed to connect to database:" + err.Error())
	}

	if DB == nil {
		slog.Error("Failed to connect to database")
		panic(errors.New("failed to connect to database"))
	}

	slog.Info("Connected to database")

	if os.Getenv("IS_MIGRATE") == "true" {
		Migration(DB)
		slog.Info("Migrated the database")
	}
}
