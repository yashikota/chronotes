package notes

import (
	"errors"

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

	contentHTML := utils.Md2HTML(response.Result[0])
	content, err := utils.CustomJSONEncoder(contentHTML)
	if err != nil {
		return dbModel.Note{}, err
	}

	// Generate note
	note := dbModel.Note{
		ID:      utils.GenerateULID(),
		Title:   "Gemini",
		Content: content,
		Tags:    "gemini,google,go",
	}

	// Save note to database
	result := db.DB.Create(&note)
	if result.Error != nil {
		return dbModel.Note{}, result.Error
	}

	return note, nil
}
