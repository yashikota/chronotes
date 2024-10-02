package notes

import (
	"errors"
	"log/slog"

	dbModel "github.com/yashikota/chronotes/model/v1/db"
	noteModel "github.com/yashikota/chronotes/model/v1/provider"
	"github.com/yashikota/chronotes/pkg/db"
	"github.com/yashikota/chronotes/pkg/provider"
	"github.com/yashikota/chronotes/pkg/utils"
)

func GenerateNote(userID string, date string, accounts noteModel.Gemini) (dbModel.Note, error) {
	if db.DB == nil {
		return dbModel.Note{}, errors.New("database connection is not initialized")
	}

	response, err := provider.Gemini(accounts)
	if err != nil {
		return dbModel.Note{}, err
	}

	slog.Info("Gemini response:", response)

	contentHTML := utils.Md2HTML(response.Result)
	slog.Info("Gemini contentHTML:", contentHTML)
	content, err := utils.CustomJSONEncoder(contentHTML)
	slog.Info("Gemini content:", content)
	if err != nil {
		return dbModel.Note{}, err
	}

	slog.Info("Gemini content:", content)

	// Generate note
	note := dbModel.Note{
		NoteID:  utils.GenerateULID(),
		UserID:  userID,
		Title:   response.Title,
		Content: content,
		Tags:    response.Tag,
	}

	slog.Info("Note:", note)

	// Save note to database
	result := db.DB.Create(&note)
	if result.Error != nil {
		return dbModel.Note{}, result.Error
	}

	slog.Info("Save note to database passed")

	return note, nil
}
