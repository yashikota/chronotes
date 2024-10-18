package users

import (
	"log/slog"

	"github.com/yashikota/chronotes/model/v1"
	"github.com/yashikota/chronotes/pkg/db"
)

func GetUser(u *model.User) (*model.User, error) {
	user := model.NewUser()
	result := db.DB.Where("user_id = ?", u.UserID).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	slog.Info("User found: " + user.UserID)

	return user, nil
}
