package notes

import (
	"github.com/yashikota/chronotes/model/v1"
	"github.com/yashikota/chronotes/pkg/db"
	"github.com/yashikota/chronotes/pkg/utils"
)

func ShareNote(noteID string) (string, error) {
	// Get note
	note, err := GetNoteByNoteID(noteID)
	if err != nil {
		return "", err
	}

	shareID := utils.GenerateULID()
	note.ShareURL = shareID

	// update note
	result := db.DB.Model(&model.Note{}).Where("note_id = ?", note.NoteID).Updates(&note)
	if result.Error != nil {
		return "", result.Error
	}

	return note.ShareURL, nil
}

func UnShareNote(noteID string) error {
	// Get note
	note, err := GetNoteByNoteID(noteID)
	if err != nil {
		return err
	}

	// update note
	result := db.DB.Model(&model.Note{}).Where("note_id = ?", note.NoteID).UpdateColumn("share_url", nil)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
