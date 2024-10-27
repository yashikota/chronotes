package notes

import (
	"fmt"
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

func GetByNoteID(noteIDs []string, fields []string) ([]map[string]string, error) {
	query := db.DB.Where("id IN (?)", noteIDs)

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

func GetUSerAllNotes(userID string, fields []string) ([]map[string]string, error) {
	query := db.DB.Where("user_id = ?", userID)

	note := model.NewNote()
	filed := strings.Join(fields, ",")
	rows, err := query.Model(note).Select(filed).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var noteList []map[string]string
	for rows.Next() {
		var note model.Note
		if err := db.DB.ScanRows(rows, &note); err != nil {
			return nil, err
		}
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
	if note.Length != 0 {
		result["length"] = fmt.Sprint(note.Length)
	}
	if !note.CreatedAt.IsZero() {
		result["created_at"] = note.CreatedAt.Format(time.RFC3339)
	}
	if !note.UpdatedAt.IsZero() {
		result["updated_at"] = note.UpdatedAt.Format(time.RFC3339)
	}
	if !note.DeletedAt.Time.IsZero() {
		result["deleted_at"] = note.DeletedAt.Time.Format(time.RFC3339)
	}

	return result
}

func GetNoteByNoteID(noteID string) (*model.Note, error) {
	query := db.DB.Where("note_id = ?", noteID)
	note := model.NewNote()
	err := query.Model(note).Take(note).Error
	if err != nil {
		return nil, err
	}
	return note, nil
}

func GetNoteByNoteShareURL(shareURL string) (*model.Note, error) {
	query := db.DB.Where("share_url = ?", shareURL)
	note := model.NewNote()
	err := query.Model(note).Take(note).Error
	if err != nil {
		return nil, err
	}
	return note, nil
}
