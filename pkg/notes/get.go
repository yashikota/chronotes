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

func GetNoteIgnoreContent(userID string, dateTime time.Time) (model.Note, error) {
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
	note.Content = ""

	return note, nil
}

func GetSummary(userID string, startDate time.Time, endDate time.Time, daysCount int) (model.Summary, error) {
	result := db.DB.Where("user_id = ? AND start_date = ? AND end_date = ? AND days_count = ?",
		userID, startDate, endDate, daysCount).First(&model.Summary{})

	var summary model.Summary
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return model.Summary{}, nil
	} else if result.Error != nil {
		return model.Summary{}, result.Error
	}

	return summary, nil
}

func GetNoteContents(userID string, startDate time.Time, endDate time.Time) ([]string, error) {
	query := db.DB.Where("user_id = ? AND created_at::date BETWEEN ? AND ?",
		userID, startDate, endDate)

	var notes []model.Note
	result := query.Select("content").Find(&notes)

	if result.Error != nil {
		return nil, result.Error
	}

	contents := make([]string, len(notes))
	for i, note := range notes {
		contents[i] = note.Content
	}

	return contents, nil
}
