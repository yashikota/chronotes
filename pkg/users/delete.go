package users

import (
	"errors"
	"log"

	model "github.com/yashikota/chronotes/model/v1/db"
	"github.com/yashikota/chronotes/pkg/db"
)

func DeleteUser(deleteUser *model.User) error {
	if db.DB == nil {
		return errors.New("database connection is not initialized")
	}

	// Find the user by ID
	user := model.User{}
	result := db.DB.Where("id = ?", deleteUser.ID).First(&user)
	if result.Error != nil {
		return result.Error
	}

	log.Println("User found")

	// Delete the user
	result = db.DB.Delete(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
