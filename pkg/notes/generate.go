package notes

import (
	"log/slog"

	"github.com/yashikota/chronotes/model/v1"
	"github.com/yashikota/chronotes/pkg/db"
	"github.com/yashikota/chronotes/pkg/provider"
	"github.com/yashikota/chronotes/pkg/utils"
)

func GenerateNote(userID string, date string, accounts model.Accounts) (*model.Note, error) {
	response, err := provider.Gemini(accounts)
	if err != nil {
		return nil, err
	}

	slog.Info("Gemini response:" + response.Result)

	contentHTML := utils.Md2HTML(response.Result)
	slog.Info("Gemini contentHTML:" + contentHTML)
	content, err := utils.CustomJSONEncoder(contentHTML)
	slog.Info("Gemini content:" + content)
	if err != nil {
		return nil, err
	}

	slog.Info("Gemini content:" + content)

	// Generate note
	note := model.Note{
		NoteID:  utils.GenerateULID(),
		UserID:  userID,
		Title:   response.Title,
		Content: content,
		Length:  utils.GetCharacterLength(response.Result),
		Tags:    response.Tag,
	}

	slog.Info("Note:" + note.Title)

	// Save note to database
	result := db.DB.Create(&note)
	if result.Error != nil {
		return nil, result.Error
	}

	slog.Info("Save note to database passed")

	return &note, nil
}
