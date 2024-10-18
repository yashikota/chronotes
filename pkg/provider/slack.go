package provider

import (
	"log/slog"
	"os"
	"strconv"
	"time"

	model "github.com/yashikota/chronotes/model/v1/provider"
	"github.com/yashikota/chronotes/pkg/gemini"

	"github.com/slack-go/slack"
)

func SlackProvider(channelID string) ([]string, error) {
	token := os.Getenv("SLACK_TOKEN")

	if token == "" {
		slog.Error("SLACK_TOKEN environment variable is not set")
		return []string{}, nil
	}

	api := slack.New(token)

	history, err := api.GetConversationHistory(&slack.GetConversationHistoryParameters{
		ChannelID: channelID,
		Limit:     100,
	})

	if err != nil {
		slog.Error("Slack : error fetching channel history" + err.Error())
		return []string{}, nil
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
	contens := extractContentSlack(todayMessages)

	if contens == nil {
		slog.Error("Slack : could not fetch commits")
		return []string{}, nil
	}

	// fmt.Println("Contents:", contens)
	summaries, err := gemini.SummarizeText(contens)
	if err != nil {
		slog.Error("Slack : error summarizing text" + err.Error())
		return []string{}, nil
	}
	// fmt.Println("Summarized texts:", summaries)
	return summaries, nil
}

func extractContentSlack(messages []model.SlackMessage) []string {
	var contents []string
	for _, msg := range messages {
		// Text フィールドが存在する場合
		contents = append(contents, msg.Text)
	}
	return contents
}
