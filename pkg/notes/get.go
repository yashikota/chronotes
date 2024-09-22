package notes

import (
	"errors"
	"time"

	"gorm.io/gorm"

	model "github.com/yashikota/chronotes/model/v1/db"
	"github.com/yashikota/chronotes/pkg/db"
)

func commonQuery(query interface{}, args ...interface{}) (*gorm.DB, error) {
	if db.DB == nil {
		return nil, errors.New("database connection is not initialized")
	}
	return db.DB.Where(query, args...), nil
}

func GetNote(userID string, dateTime time.Time) (model.Note, error) {
	return getNote(userID, dateTime, false)
}

func GetNoteIgnoreContent(userID string, dateTime time.Time) (model.Note, error) {
	return getNote(userID, dateTime, true)
}

func getNote(userID string, dateTime time.Time, ignoreContent bool) (model.Note, error) {
	query, err := commonQuery("user_id = ? AND created_at::date = ?::date", userID, dateTime)
	if err != nil {
		return model.Note{}, err
	}

	var note model.Note
	result := query.First(&note)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return model.Note{}, nil
	} else if result.Error != nil {
		return model.Note{}, result.Error
	}

	if ignoreContent {
		note.Content = ""
	}

	return note, nil
}

func GetSummary(userID string, startDate time.Time, endDate time.Time, daysCount int) (model.Summary, error) {
	query, err := commonQuery("user_id = ? AND start_date = ? AND end_date = ? AND days_count = ?",
		userID, startDate, endDate, daysCount)
	if err != nil {
		return model.Summary{}, err
	}

	var summary model.Summary
	result := query.First(&summary)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return model.Summary{}, nil
	} else if result.Error != nil {
		return model.Summary{}, result.Error
	}

	return summary, nil
}

func GetNoteContents(userID string, startDate time.Time, endDate time.Time) ([]string, error) {
	query, err := commonQuery("user_id = ? AND created_at::date BETWEEN ? AND ?",
		userID, startDate, endDate)
	if err != nil {
		return nil, err
	}

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
