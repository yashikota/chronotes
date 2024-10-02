package db

import (
	"time"

	"gorm.io/gorm"

	model "github.com/yashikota/chronotes/model/v1/provider"
)

type User struct {
	gorm.Model
	UserID   string `json:"user_id" gorm:"uniqueIndex"`
	UserName string `json:"username" gorm:"uniqueIndex"`
	Email    string `json:"email" gorm:"uniqueIndex"`
	Password string `json:"password"`
}

type Note struct {
	gorm.Model
	NoteID  string `json:"note_id" gorm:"uniqueIndex"`
	UserID  string `json:"user_id" gorm:"uniqueIndex"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Tags    string `json:"tags"`
}

type NoteResponse struct {
	Date    string `json:"date"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Tags    string `json:"tags"`
}

type Account struct {
	UserID string `json:"user_id"`
	model.Gemini
}

type Summary struct {
	gorm.Model
	ID        string    `json:"summary_id"`
	UserID    string    `json:"user_id"`
	Content   string    `json:"content"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	DaysCount int       `json:"days_count"`
}
