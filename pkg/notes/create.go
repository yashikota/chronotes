package notes

import (
	"log/slog"

	"github.com/yashikota/chronotes/model/v1"
	"github.com/yashikota/chronotes/pkg/db"
)

func CreateNote(note model.Note) error {
	slog.Info("Create note")

	result := db.DB.Create(&note)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
