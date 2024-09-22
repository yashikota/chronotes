package notes

import (
	"errors"
	"time"

	"gorm.io/gorm"

	model "github.com/yashikota/chronotes/model/v1/db"
	"github.com/yashikota/chronotes/pkg/db"
)

func GetNote(userID string, dateTime time.Time) (model.Note, error) {
	if db.DB == nil {
		return model.Note{}, errors.New("database connection is not initialized")
	}

	// Get note from database
	note := model.Note{}
	result := db.DB.Where("user_id = ? AND created_at::date = ?::date", userID, dateTime).First(&note)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return model.Note{}, nil
	} else if result.Error != nil {
		return model.Note{}, result.Error
	}

	return note, nil
}

func GetNoteIgnoreContent(userID string, date string) (model.Note, error) {
	if db.DB == nil {
		return model.Note{}, errors.New("database connection is not initialized")
	}

	// Get note from database
	note := model.Note{}
	result := db.DB.Where("user_id = ? AND created_at::date = ?::date", userID, dateTime).First(&note)
	if result.Error != nil {
		return model.Note{}, result.Error
	}
	note.Content = ""

	return note, nil
}
