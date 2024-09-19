package provider

import (
	"fmt"
	"os"
	"strconv"
	"time"

	model "github.com/yashikota/chronotes/model/v1/provider"

	"github.com/slack-go/slack"
)

func SlackProvider(channelID string) (map[string][]model.SlackMessage, error) {
	token := os.Getenv("SLACK_TOKEN")

	if token == "" {
		return nil, fmt.Errorf("SLACK_TOKEN environment variable is not set")
	}

	api := slack.New(token)
	history, err := api.GetConversationHistory(&slack.GetConversationHistoryParameters{
		ChannelID: channelID,
		Limit:     100,
	})

	if err != nil {
		return nil, fmt.Errorf("error fetching channel history: %v", err)
	}

	// カテゴリごとのメッセージを格納するためのマップ
	categorizedMessages := map[string][]model.SlackMessage{}

	now := time.Now()
	today := now.Format("2006-01-02")

	var todayMessages []model.SlackMessage

	for _, message := range history.Messages {
		ts, err := strconv.ParseFloat(message.Timestamp, 64)
		if err != nil {
			continue
		}

		date := time.Unix(int64(ts), 0)

		if date.Format("2006-01-02") == today {
			slackMessage := model.SlackMessage{
				ID:        message.Timestamp,
				User:      message.User,
				Text:      message.Text,
				Timestamp: ts,
				Channel:   channelID,
				CreatedAt: date,
			}

			todayMessages = append(todayMessages, slackMessage)
		}
	}

	// 今日のメッセージが存在する場合は "Today" カテゴリに追加
	if len(todayMessages) > 0 {
		categorizedMessages["Today"] = todayMessages
	}

	return categorizedMessages, nil
}
