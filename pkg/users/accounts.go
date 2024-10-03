package users

import (
	"errors"
	"log/slog"

	"gorm.io/gorm"

	dbmodel "github.com/yashikota/chronotes/model/v1/db"
	"github.com/yashikota/chronotes/pkg/db"
)

func UpdateAccounts(newAccounts *model.Account) error {
	if db.DB == nil {
		return errors.New("database connection is not initialized")
	}

	slog.Info("Updating accounts: " + newAccounts.UserID)

	oldAccounts := model.Account{}
	result := db.DB.Where("user_id = ?", newAccounts.UserID).First(&oldAccounts)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// Create new account
		if err := db.DB.Create(&newAccounts).Error; err != nil {
			return err
		}
	} else if result.Error != nil {
		return result.Error
	}

	updates := map[string]interface{}{}
	if newAccounts.SlackChannelID != "" {
		updates["slack_channel_id"] = newAccounts.SlackChannelID
	}
	if newAccounts.GitHubUserID != "" {
		updates["git_hub_user_id"] = newAccounts.GitHubUserID
	}
	if newAccounts.DiscordChannelID != "" {
		updates["discord_channel_id"] = newAccounts.DiscordChannelID
	}
	if newAccounts.QiitaUserID != "" {
		updates["qiita_user_id"] = newAccounts.QiitaUserID
	}

	if len(updates) > 0 {
		if err := db.DB.Model(&oldAccounts).Where("user_id = ?", newAccounts.UserID).Updates(updates).Error; err != nil {
			return err
		}
	}

	return nil
}
