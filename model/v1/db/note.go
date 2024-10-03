package model

import (
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	NoteID  string `json:"note_id" gorm:"uniqueIndex"`
	UserID  string `json:"user_id" gorm:"uniqueIndex"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Tags    string `json:"tags"`
}

func NewNote() *Note {
	return &Note{}
}

type NoteResponse struct {
	Date    string `json:"date"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Tags    string `json:"tags"`
}

func NewNoteResponse() *NoteResponse {
	return &NoteResponse{}
}
