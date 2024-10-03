package v1

import (
	db "github.com/yashikota/chronotes/model/v1/db"
	provider "github.com/yashikota/chronotes/model/v1/provider"
)

type (
	User = db.User
	Note = db.Note
)

var (
	CreateUser = db.CreateUser
	GetUser    = db.GetUser
)

type (
	AuthProvider = provider.AuthProvider
)

var (
	NewAuthProvider = provider.NewAuthProvider
)
