package provider

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"

	model "github.com/yashikota/chronotes/model/v1/provider"
	"github.com/yashikota/chronotes/pkg/utils"
)

func DiscordProvider(channelID string) ([]string, error) {
	if channelID == "" {
		log.Printf("Discord : channelID is not set")
		return []string{}, nil
	}
	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		log.Printf("Discord : DISCORD_TOKEN environment variable is not set")
		return []string{}, nil
	}

	messages, err := runBot(channelID, token)
	if err != nil {
		log.Printf("Discord : Error running bot: %v", err)
		return []string{}, nil
	}

	categorizedMessages := categorizeMessages(messages)

	// 「Today」カテゴリのメッセージのみを取り出す
	todayMessages := categorizedMessages["Today"]
	contents := extractContentsDiscord(todayMessages)
	summaries, err := utils.SummarizeText(contents)
	if err != nil {
		log.Printf("Discord : Error summarizing text: %v", err)
		return []string{}, nil
	}
	return summaries, nil
}

func extractContentsDiscord(messages []model.DiscordMessage) []string {
	var contents []string
	for _, msg := range messages {
		contents = append(contents, msg.Content)
	}
	return contents
}

func runBot(channelID, token string) ([]*discordgo.Message, error) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Printf("Discord : Error creating Discord session: %v", err)
		return nil, nil
	}

	dg.AddHandler(ready)

	err = dg.Open()
	if err != nil {
		log.Printf("Discord : Error opening connection: %v", err)
		return nil, nil
	}
	defer dg.Close()
	return getMessageHistory(dg, channelID)
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	fmt.Println("Discord : Bot is ready")
}

func getMessageHistory(s *discordgo.Session, channelID string) ([]*discordgo.Message, error) {
	var allMessages []*discordgo.Message
	var lastMessageID string
	for {
		messages, err := s.ChannelMessages(channelID, 100, lastMessageID, "", "")
		if err != nil {
			log.Printf("Discord : Error getting messages: %v", err)
			return nil, nil
		}
		if len(messages) == 0 {
			break
		}
		allMessages = append(allMessages, messages...)
		lastMessageID = messages[len(messages)-1].ID
	}
	return allMessages, nil
}

func categorizeMessages(messages []*discordgo.Message) map[string][]model.DiscordMessage {
	now := time.Now()
	startOfToday := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	startOfWeek := now.AddDate(0, 0, -int(now.Weekday()))
	startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)
	startOfYear := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, time.Local)

	startOfQuarter := getStartOfQuarter(now)
	endOfQuarter := startOfQuarter.AddDate(0, 3, -1)

	var todayMessages []model.DiscordMessage
	var weeklyMessages []model.DiscordMessage
	var monthlyMessages []model.DiscordMessage
	var quarterlyMessages []model.DiscordMessage
	var yearlyMessages []model.DiscordMessage

	for _, message := range messages {
		timestamp := message.Timestamp
		messageData := model.DiscordMessage{
			ID:        message.ID,
			Author:    message.Author.Username,
			Content:   message.Content,
			Timestamp: message.Timestamp,
		}

		// Periodの設定
		switch {
		case timestamp.After(startOfToday):
			messageData.Period = "Today"
			todayMessages = append(todayMessages, messageData)
		case timestamp.After(startOfWeek):
			messageData.Period = "This Week"
			weeklyMessages = append(weeklyMessages, messageData)
		case timestamp.After(startOfMonth):
			messageData.Period = "This Month"
			monthlyMessages = append(monthlyMessages, messageData)
		case timestamp.After(startOfQuarter) && timestamp.Before(endOfQuarter):
			messageData.Period = "This Quarter"
			quarterlyMessages = append(quarterlyMessages, messageData)
		case timestamp.After(startOfYear):
			messageData.Period = "This Year"
			yearlyMessages = append(yearlyMessages, messageData)
		}
	}

	return map[string][]model.DiscordMessage{
		"Today":        todayMessages,
		"This Week":    weeklyMessages,
		"This Month":   monthlyMessages,
		"Q1 (Jan-Mar)": filterMessagesByQuarter(quarterlyMessages, time.January, time.March),
		"Q2 (Apr-Jun)": filterMessagesByQuarter(quarterlyMessages, time.April, time.June),
		"Q3 (Jul-Sep)": filterMessagesByQuarter(quarterlyMessages, time.July, time.September),
		"Q4 (Oct-Dec)": filterMessagesByQuarter(quarterlyMessages, time.October, time.December),
		"This Year":    yearlyMessages,
	}
}

// 四半期の開始日を取得する関数
func getStartOfQuarter(now time.Time) time.Time {
	month := now.Month()
	var startOfQuarter time.Time
	switch {
	case month >= time.January && month <= time.March:
		startOfQuarter = time.Date(now.Year(), time.January, 1, 0, 0, 0, 0, time.Local)
	case month >= time.April && month <= time.June:
		startOfQuarter = time.Date(now.Year(), time.April, 1, 0, 0, 0, 0, time.Local)
	case month >= time.July && month <= time.September:
		startOfQuarter = time.Date(now.Year(), time.July, 1, 0, 0, 0, 0, time.Local)
	case month >= time.October && month <= time.December:
		startOfQuarter = time.Date(now.Year(), time.October, 1, 0, 0, 0, 0, time.Local)
	}
	return startOfQuarter
}

func filterMessagesByQuarter(messages []model.DiscordMessage, startMonth, endMonth time.Month) []model.DiscordMessage {
	var filteredMessages []model.DiscordMessage
	startDate := time.Date(time.Now().Year(), startMonth, 1, 0, 0, 0, 0, time.Local)
	endDate := time.Date(time.Now().Year(), endMonth+1, 1, 0, 0, 0, 0, time.Local).AddDate(0, 0, -1)

	for _, message := range messages {
		if isInRange(message.Timestamp, startDate, endDate) {
			filteredMessages = append(filteredMessages, message)
		}
	}

	return filteredMessages
}

func isInRange(timestamp, startDate, endDate time.Time) bool {
	return timestamp.After(startDate) && timestamp.Before(endDate)
}
