package admin

import (
	"errors"
	"log/slog"

	"github.com/yashikota/chronotes/model/v1"
	"github.com/yashikota/chronotes/pkg/db"
)

func IsAdmin(userID string) (bool, error) {
	if db.DB == nil {
		return false, errors.New("database connection is not initialized")
	}

	r := model.NewUser()
	result := db.DB.Where("user_id = ?", userID).First(&r)
	if result.Error != nil {
		return false, result.Error
	}

	slog.Info("User found")

	if r.Role != model.Admin {
		return false, errors.New("user is not an admin")
	} else {
		return true, nil
	}
}
