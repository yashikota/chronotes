package model

import (
	"time"

	"gorm.io/gorm"
)

type Summary struct {
	gorm.Model
	ID        string    `json:"summary_id"`
	UserID    string    `json:"user_id"`
	Content   string    `json:"content"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	DaysCount int       `json:"days_count"`
}

func NewSummary() *Summary {
	return &Summary{}
}
