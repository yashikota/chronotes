package notes

import (
	"github.com/yashikota/chronotes/model/v1"
	"github.com/yashikota/chronotes/pkg/db"
)

func GetAccounts(userID string) (*model.Accounts, error) {
	// Get accounts from database
	accounts := model.NewAccounts()
	result := db.DB.Where("user_id = ?", userID).Find(&accounts)
	if result.Error != nil {
		return nil, result.Error
	}

	return accounts, nil
}
