package db

import (
	"time"

	model "github.com/yashikota/chronotes/model/v1/provider"
)

type User struct {
	ID        string    `json:"user_id" gorm:"column:user_id;primaryKey"`
	Name      string    `json:"username"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type Note struct {
	ID        string    `json:"note_id" gorm:"column:id;primaryKey"`
	UserID    string    `json:"user_id" gorm:"column:user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Tags      string    `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type NoteResponse struct {
	Date    string `json:"date"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Tags    string `json:"tags"`
}

type Account struct {
	UserID string `json:"user_id" gorm:"column:user_id"`
	model.Gemini
}

type Summary struct {
        ID        string    `json:"summary_id" gorm:"column:id;primaryKey"`
        UserID    string    `json:"user_id" gorm:"column:user_id"`
        Content   string    `json:"content"`
        StartDate time.Time `json:"start_date"`
        EndDate   time.Time `json:"end_date"`
        DaysCount int       `json:"days_count"`
        CreatedAt time.Time `json:"created_at"`
}