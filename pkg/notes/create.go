package notes

import (
	"errors"
	"log/slog"

	"github.com/yashikota/chronotes/model/v1"
	"github.com/yashikota/chronotes/pkg/db"
)

func CreateNote(note model.Note) error {
	if db.DB == nil {
		return errors.New("database connection is not initialized")
	}

	slog.Info("Create note")

	result := db.DB.Create(&note)
	if result.Error != nil {
		return result.Error
	}

	slog.Info("Save note to database passed")

	return nil
}
