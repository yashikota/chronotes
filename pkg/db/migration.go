package db

import (
	"log"

	"gorm.io/gorm"

	model "github.com/yashikota/chronotes/model/v1/db"
)

func Migration(db *gorm.DB) {
	err := DB.AutoMigrate(
		&model.User{},
		&model.Note{},
	)

	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
