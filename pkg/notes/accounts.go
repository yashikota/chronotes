package notes

import (
	"errors"

	model "github.com/yashikota/chronotes/model/v1/provider"
	"github.com/yashikota/chronotes/pkg/db"
)

func GetAccounts(userID string) (model.Gemini, error) {
	if db.DB == nil {
		return model.Gemini{}, errors.New("database connection is not initialized")
	}

	// Get accounts from database
	accounts := model.Gemini{}
	result := db.DB.Where("user_id = ?", userID).Find(&accounts)
	if result.Error != nil {
		return model.Gemini{}, result.Error
	}

	return accounts, nil
}
