package model

import (
	"time"

	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	NoteID           string `json:"note_id" gorm:"uniqueIndex"`
	UserID           string `json:"user_id"`
	Title            string `json:"title"`
	Content          string `json:"content"`
	TokenizedContent string `json:"tokenized_content"`
	Length           int    `json:"length"`
	Tags             string `json:"tags"`
	ShareURL         string `json:"share_url"`
}

func NewNote() *Note {
	return &Note{}
}

type NoteResponse struct {
	Date    time.Time `json:"date"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Tags    string    `json:"tags"`
}

func NewNoteResponse() *NoteResponse {
	return &NoteResponse{}
}
