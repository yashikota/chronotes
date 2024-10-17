package users

import (
	"log/slog"

	"github.com/yashikota/chronotes/model/v1"
	"github.com/yashikota/chronotes/pkg/db"
)

func GetUsersList() ([]string, error) {
	var userIDs []string
	err := db.DB.Model(&model.User{}).Pluck("user_id", &userIDs).Error
	if err != nil {
		slog.Error("Error getting user list")
		return nil, err
	}
	return userIDs, nil
}
