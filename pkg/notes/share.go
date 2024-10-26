package notes

import (
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

	return note.ShareURL, nil
}
