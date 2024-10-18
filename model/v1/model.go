package model

import (
	db "github.com/yashikota/chronotes/model/v1/db"
)

type (
	User         = db.User
	Login        = db.Login
	Role         = db.Role
	Identity     = db.Identity
	Note         = db.Note
	NoteResponse = db.NoteResponse
	Summary      = db.Summary
	Password     = db.Password
	Accounts     = db.Accounts
)

const (
	Normal = db.Normal
	Admin  = db.Admin

	UserID = db.UserID
	Email  = db.Email
)

var (
	NewUser         = db.NewUser
	NewLogin        = db.NewLogin
	NewNote         = db.NewNote
	NewNoteResponse = db.NewNoteResponse
	NewSummary      = db.NewSummary
	NewPassword     = db.NewPassword
	NewAccounts     = db.NewAccounts
)
