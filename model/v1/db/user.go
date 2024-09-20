package db

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       string `json:"id" gorm:"primaryKey"`
	Name     string `json:"username"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}
