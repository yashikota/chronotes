package users

import (
	"errors"
	"log/slog"

	model "github.com/yashikota/chronotes/model/v1/db"
	"github.com/yashikota/chronotes/pkg/db"
)

func DeleteUser(u *model.User) error {
	if db.DB == nil {
		return errors.New("database connection is not initialized")
	}

	// Find the user by ID
	user := model.User{}
	result := db.DB.Where("user_id = ?", u.UserID).First(&user)
	if result.Error != nil {
		return result.Error
	}

	slog.Info("User found")

	// Delete the user
	result = db.DB.Delete(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
