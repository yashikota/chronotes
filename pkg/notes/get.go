package notes

import (
	"errors"

	model "github.com/yashikota/chronotes/model/v1/db"
	"github.com/yashikota/chronotes/pkg/db"
)

func GetNote(userID string, date string) (model.Note, error) {
	if db.DB == nil {
		return model.Note{}, errors.New("database connection is not initialized")
	}

	// Get note from database
	note := model.Note{}
	result := db.DB.Where("user_id = ? AND date = ?", userID, date).First(&note)
	if result.Error != nil {
		return model.Note{}, result.Error
	}

	return note, nil
}
