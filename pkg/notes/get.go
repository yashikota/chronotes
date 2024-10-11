package notes

import (
	"strings"
	"time"

	"github.com/yashikota/chronotes/model/v1"
	"github.com/yashikota/chronotes/pkg/db"
)

func GetNotes(userID string, startDate time.Time, endDate time.Time, fields []string) ([]map[string]string, error) {
	query := db.DB.Where("user_id = ? AND created_at::date BETWEEN ?::date AND ?::date",
		userID, startDate, endDate)

	var noteList []map[string]string
	note := model.NewNote()
	filed := strings.Join(fields, ",")
	rows, err := query.Model(note).Select(filed).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var note model.Note
		db.DB.ScanRows(rows, &note)
		noteList = append(noteList, toMap(&note))
	}
	return noteList, nil
}

func toMap(note *model.Note) map[string]string {
	result := make(map[string]string)

	if note.NoteID != "" {
		result["note_id"] = note.NoteID
	}
	if note.UserID != "" {
		result["user_id"] = note.UserID
	}
	if note.Title != "" {
		result["title"] = note.Title
	}
	if note.Content != "" {
		result["content"] = note.Content
	}
	if note.Tags != "" {
		result["tags"] = note.Tags
	}
	if !note.CreatedAt.IsZero() {
		result["created_at"] = note.CreatedAt.String()
	}
	if !note.UpdatedAt.IsZero() {
		result["updated_at"] = note.UpdatedAt.String()
	}
	if !note.DeletedAt.Time.IsZero() {
		result["deleted_at"] = note.DeletedAt.Time.String()
	}

	return result
}
