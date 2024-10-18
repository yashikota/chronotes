package users

import (
	"log/slog"

	"github.com/yashikota/chronotes/model/v1"
	"github.com/yashikota/chronotes/pkg/db"
)

func DeleteUser(u *model.User) error {
	// Find the user by ID
	user := model.NewUser()
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
