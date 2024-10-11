package notes

import (
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/yashikota/chronotes/model/v1"
	"github.com/yashikota/chronotes/pkg/db"
)

func GetNote(user model.User, dateTime time.Time) (*model.Note, error) {
	// Get note from database
	note := model.NewNote()
	result := db.DB.Where("user_id = ? AND created_at::date = ?::date", user.UserID, dateTime).First(&note)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if result.Error != nil {
		return nil, result.Error
	}

	return note, nil
}

func GetSummary(userID string, startDate time.Time, endDate time.Time, daysCount int) (*model.Summary, error) {
	summary := model.NewSummary()
	result := db.DB.Where("user_id = ? AND start_date = ? AND end_date = ? AND days_count = ?",
		userID, startDate, endDate, daysCount).First(summary)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if result.Error != nil {
		return nil, result.Error
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

func GetNoteList(userID string, startDate time.Time, endDate time.Time, fields []string) ([]map[string]string, error) {
	query := db.DB.Where("user_id = ? AND created_at::date BETWEEN ? AND ?",
		userID, startDate, endDate)

	var noteList []map[string]string
	note := model.NewNote()
	rows, err := query.Model(note).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var note model.Note
		db.DB.ScanRows(rows, &note)

		noteMap := make(map[string]string)
		for _, field := range fields {
			switch field {
			case "note_id":
				noteMap[field] = note.NoteID
			case "user_id":
				noteMap[field] = note.UserID
			case "title":
				noteMap[field] = note.Title
			case "content":
				noteMap[field] = note.Content
			case "tags":
				noteMap[field] = note.Tags
			case "created_at":
				noteMap[field] = note.CreatedAt.String()
			case "updated_at":
				noteMap[field] = note.UpdatedAt.String()
			case "deleted_at":
				noteMap[field] = note.DeletedAt.Time.String()
			}
		}
		noteList = append(noteList, noteMap)
	}
	return noteList, nil
}
