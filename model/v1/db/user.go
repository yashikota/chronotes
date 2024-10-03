package model

import (
	"gorm.io/gorm"
)

type Role int

const (
	Normal Role = iota
	Admin
)

type User struct {
	gorm.Model
	UserID   string `json:"user_id" gorm:"uniqueIndex"`
	UserName string `json:"username" gorm:"uniqueIndex"`
	Email    string `json:"email" gorm:"uniqueIndex"`
	Password string `json:"-"`
	Role     Role   `json:"role"`
}

func NewUser() *User {
	return &User{
		Role: Normal,
	}
}
