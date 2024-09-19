package users

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	model "github.com/yashikota/chronotes/model/v1/db"
	"github.com/yashikota/chronotes/pkg/db"
)

func CreateUser(user *model.User) error {
	if db.DB == nil {
		return errors.New("database connection is not initialized")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Create the user in the database using the existing connection
	result := db.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func IsEmailTaken(email string) (bool, error) {
	var count int64
	if err := db.DB.Model(&model.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
