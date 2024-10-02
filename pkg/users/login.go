package users

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"

	model "github.com/yashikota/chronotes/model/v1/db"
	"github.com/yashikota/chronotes/pkg/db"
)

func LoginUser(u *model.User) error {
	if db.DB == nil {
		return errors.New("database connection is not initialized")
	}

	// Find the user by email
	r := model.User{}
	result := db.DB.Where("email = ?", u.Email).First(&r)
	if result.Error != nil {
		return result.Error
	}

	log.Println("User found")

	// Compare the hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(r.Password), []byte(u.Password)); err != nil {
		return errors.New("password does not match")
	}

	log.Println("Password matched")

	u.UserID = r.UserID
	u.UserName = r.UserName

	return nil
}
