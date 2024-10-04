package model

import (
	db "github.com/yashikota/chronotes/model/v1/db"
	provider "github.com/yashikota/chronotes/model/v1/provider"
)

// db package
type (
	User = db.User
	Note = db.Note
	NoteResponse = db.NoteResponse
	Summary = db.Summary
)

var (
	NewUser = db.NewUser
	NewNote = db.NewNote
	NewNoteResponse = db.NewNoteResponse
	NewSummary = db.NewSummary
)

// provider package
type (
	Accounts = provider.Accounts
)

var (
	NewAccounts = provider.NewAccounts
)
