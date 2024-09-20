package db

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID   string `json:"id" gorm:"primaryKey"`
	Name     string `json:"username"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

type Note struct {
	gorm.Model
	NoteID  string `json:"note_id" gorm:"primaryKey"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Tags    string `json:"tags"`
	Images  string `json:"images"`
}
