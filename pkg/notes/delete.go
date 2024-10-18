package notes

import (
	"errors"
	"log/slog"

	"github.com/yashikota/chronotes/model/v1"
	"github.com/yashikota/chronotes/pkg/db"
)

func DeleteNote(note model.Note, user *model.User) error {
	slog.Info("Delete note")

	// Check note ownership
	if note.UserID != user.UserID && user.Role != model.Admin {
		return errors.New("note does not belong to user")
	}

	result := db.DB.Where("note_id = ?", note.NoteID).Delete(&model.Note{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
