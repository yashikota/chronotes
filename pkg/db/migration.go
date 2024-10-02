package db

import (
	"log/slog"

	"gorm.io/gorm"

	model "github.com/yashikota/chronotes/model/v1/db"
)

func Migration(db *gorm.DB) {
	err := DB.AutoMigrate(
		&model.User{},
		&model.Note{},
		&model.Account{},
	)

	if err != nil {
		slog.Error("Failed to migrate database:" + err.Error())
	}
}
