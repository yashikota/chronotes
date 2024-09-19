package users

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	model "github.com/yashikota/chronotes/model/v1/db"
	"github.com/yashikota/chronotes/pkg/db"
)

func LoginUser(user *model.User) (string, error) {
	if db.DB == nil {
		return "", errors.New("database connection is not initialized")
	}

	// Find the user by email
	result := db.DB.Where("email = ?", user.Email).First(&user)
	if result.Error != nil {
		return "", result.Error
	}

	// Compare the hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(user.Password)); err != nil {
		return "", errors.New("invalid password")
	}

	// Generate token
	// TODO: Implement token generation

	return "token", nil
}
