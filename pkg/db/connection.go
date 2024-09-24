package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load(fmt.Sprintf(".env.%s", os.Getenv("GO_ENV")))
	if err != nil && !os.IsNotExist(err) {
		log.Fatal("Failed to load .env file:", err)
	}

	// Connect to the database
	pgpw := os.Getenv("POSTGRES_PASSWORD")
	dsn := fmt.Sprintf("postgres://postgres:%s@db:5432/chronotes?sslmode=disable", pgpw)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Connected to database")

	// Migrate the database
	Migration(DB)

	log.Println("Migrated the database")

	// Seed the database
	// Seed(DB)

	log.Println("Database connection initialized")
}
