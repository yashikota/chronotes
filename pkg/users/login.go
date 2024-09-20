package users

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"

	model "github.com/yashikota/chronotes/model/v1/db"
	"github.com/yashikota/chronotes/pkg/db"
)

func LoginUser(loginUser *model.User) error {
	if db.DB == nil {
		return errors.New("database connection is not initialized")
	}

	// Find the user by email
	registeredUser := model.User{}
	result := db.DB.Where("email = ?", loginUser.Email).First(&registeredUser)
	if result.Error != nil {
		return result.Error
	}

	log.Println("User found")

	// Compare the hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(registeredUser.Password), []byte(loginUser.Password)); err != nil {
		return errors.New("password does not match")
	}

	log.Println("Password matched")

	loginUser.UserID = registeredUser.UserID
	loginUser.Name = registeredUser.Name

	return nil
}
