package db

import (
	"time"
)

type User struct {
	ID       string `json:"user_id" gorm:"column:user_id;primaryKey"`
	Name     string `json:"username"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type Note struct {
	ID      string `json:"note_id" gorm:"column:id;primaryKey"`
	UserID  string `json:"user_id" gorm:"column:user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Tags    string `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type NoteResponse struct {
	Date	string `json:"date"`
	Title	string `json:"title"`
	Content	string `json:"content"`
	Tags	string `json:"tags"`
}

type Account struct {
	ID       string `json:"id" gorm:"primaryKey"`
	Provider string `json:"provider"`
}
