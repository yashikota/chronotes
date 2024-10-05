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
	UserName string `json:"user_name"`
	Email    string `json:"email" gorm:"uniqueIndex"`
	Password string `json:"password"`
	Role     Role   `json:"role"`
}

func NewUser() *User {
	return &User{
		Role: Normal, // default role
	}
}

type Identity int

const (
	UserID Identity = iota
	Email
)

type Login struct {
	UserID	   string `json:"user_id"`
	Email 	   string `json:"email"`
	Password   string `json:"password"`
}

func NewLogin() *Login {
	return &Login{}
}

type Password struct {
	Password string `json:"password"`
}

func NewPassword() *Password {
	return &Password{}
}
