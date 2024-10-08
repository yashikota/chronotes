package users

import (
	"github.com/yashikota/chronotes/model/v1"
	"github.com/yashikota/chronotes/pkg/db"
	"github.com/yashikota/chronotes/pkg/utils"
)

func CreateUser(user *model.User) error {
	err := utils.GeneratePassword(user)
	if err != nil {
		return err
	}

	// Create the user in the database using the existing connection
	result := db.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func IsEmailTaken(email string) (bool, error) {
	var count int64
	user := model.NewUser()
	if err := db.DB.Model(user).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
