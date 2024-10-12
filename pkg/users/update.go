package users

import (
	"log/slog"

	"github.com/yashikota/chronotes/model/v1"
	"github.com/yashikota/chronotes/pkg/db"
)

func UpdateUser(req *model.User) error {
	result := db.DB.Model(&model.User{}).Where("user_id = ?", req.UserID).Updates(&req)
	if result.Error != nil {
		return result.Error
	}
	slog.Info("Update user successful")

	return nil
}
