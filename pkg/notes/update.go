package notes

import (
	"errors"
	"log/slog"

	"github.com/yashikota/chronotes/model/v1"
	"github.com/yashikota/chronotes/pkg/db"
)

func UpdateNote(note model.Note, user *model.User) error {
	slog.Info("Update note")

	// Check ownership
	if note.UserID != user.UserID && user.Role != model.Admin {
		return errors.New("note does not belong to user")
	}

	result := db.DB.Model(&model.Note{}).Where("note_id = ?", note.NoteID).Updates(&note)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
