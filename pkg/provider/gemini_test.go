package provider_test

import (
	"os"
	"testing"

	"github.com/yashikota/chronotes/model/v1"
	"github.com/yashikota/chronotes/pkg/provider"
)

func TestGeminiHandler(t *testing.T) {
	// 環境変数からのデータ取得
	input := model.Accounts{
		GitHubUserID:     os.Getenv("GITHUB_USER_ID"),
		SlackChannelID:   os.Getenv("SLACK_CHANNEL_ID"),
		DiscordChannelID: os.Getenv("DISCORD_CHANNEL_ID"),
		QiitaUserID:      os.Getenv("QIITA_USER_ID"),
	}

	// 環境変数が設定されていない場合のデフォルト値
	if input.GitHubUserID == "" {
		input.GitHubUserID = "GITHUB_USER_ID"
	}
	if input.SlackChannelID == "" {
		input.SlackChannelID = "SLACK_CHANNEL_ID"
	}
	if input.DiscordChannelID == "" {
		input.DiscordChannelID = "DISCORD_CHANNEL_ID"
	}

	if input.QiitaUserID == "" {
		input.QiitaUserID = "QIITA_USER_ID"
	}

	response, err := provider.Gemini(input)

	if err != nil {
		t.Error(err)
	}

	if response.Result == "" {
		t.Error("could not fetch commits")
	}

	if response.Day == "" {
		t.Error("could not fetch day")
	}

}
