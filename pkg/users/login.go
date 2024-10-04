package users

import (
	"errors"
	"log/slog"

	"golang.org/x/crypto/bcrypt"

	"github.com/yashikota/chronotes/model/v1"
	"github.com/yashikota/chronotes/pkg/db"
)

func LoginUser(u *model.User) error {
	if db.DB == nil {
		return errors.New("database connection is not initialized")
	}

	// Find the user by email
	r := model.NewUser()
	result := db.DB.Where("email = ?", u.Email).First(&r)
	if result.Error != nil {
		return result.Error
	}

	slog.Info("User found")

	// Compare the hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(r.Password), []byte(u.Password)); err != nil {
		return errors.New("password does not match")
	}

	slog.Info("Password matched")

	u.UserID = r.UserID
	u.UserName = r.UserName

	return nil
}
