package users

import (
	"errors"
	"log/slog"

	"gorm.io/gorm"

	"github.com/yashikota/chronotes/model/v1"
	"github.com/yashikota/chronotes/pkg/db"
)

func GetAccounts(userID string) (*model.Accounts, error) {
	slog.Info("Updating accounts: " + userID)

	accounts := model.NewAccounts()
	result := db.DB.Model(&model.Accounts{}).Where("user_id = ?", userID).First(&accounts)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return accounts, nil
		}
		return nil, result.Error
	}
	return accounts, nil
}

func UpdateAccounts(accounts *model.Accounts) error {
	slog.Info("Updating accounts: " + accounts.UserID)

	result := db.DB.Model(&model.Accounts{}).Where("user_id = ?", accounts.UserID).First(&model.Accounts{})
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			result = db.DB.Create(accounts)
			if result.Error != nil {
				return result.Error
			}
		} else {
			return result.Error
		}
	} else {
		result = db.DB.Model(&model.Accounts{}).Where("user_id = ?", accounts.UserID).Updates(accounts)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}
