package users

import (
	"errors"
	"log/slog"

	"github.com/yashikota/chronotes/model/v1"
	"github.com/yashikota/chronotes/pkg/db"
)

func PromoteUser(u *model.User) error {
	if db.DB == nil {
		return errors.New("database connection is not initialized")
	}

	// Find the user by UserID
	r := model.NewUser()
	result := db.DB.Where("user_id = ?", u.UserID).First(&r)
	if result.Error != nil {
		return result.Error
	}

	slog.Info("User found")

	// Promote the user
	r.Role = model.Admin
	result = db.DB.Save(&r)
	if result.Error != nil {
		return result.Error
	}

	slog.Info("User promoted")

	return nil
}
