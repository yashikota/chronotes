package db

import (
	"log/slog"

	"gorm.io/gorm"

	"github.com/yashikota/chronotes/model/v1"
)

func Migration(db *gorm.DB) {
	err := DB.AutoMigrate(
		&model.User{},
		&model.Note{},
		&model.Accounts{},
	)

	if err != nil {
		slog.Error("Failed to migrate database:" + err.Error())
	}
}
