package provider

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
	model "github.com/yashikota/chronotes/model/v1/provider"
)

func DiscordProvider(channelID string) (map[string][]model.Message, error) {
	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		log.Fatal("DISCORD_TOKEN environment variable is not set")
	}

	messages, err := runBot(channelID, token)
	if err != nil {
		log.Fatalf("Error running bot: %v", err)
	}

	categorizedMessages := categorizeMessages(messages)
	printCategorizedMessages(categorizedMessages)
	return categorizedMessages, nil
}

func runBot(channelID, token string) ([]*discordgo.Message, error) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, fmt.Errorf("error creating Discord session: %w", err)
	}

	dg.AddHandler(ready)

	err = dg.Open()
	if err != nil {
		return nil, fmt.Errorf("error opening connection: %w", err)
	}
	defer dg.Close()

	fmt.Println("Bot is now running. Press CTRL+C to exit.")

	return getMessageHistory(dg, channelID)
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	fmt.Println("Bot is ready")
}

func getMessageHistory(s *discordgo.Session, channelID string) ([]*discordgo.Message, error) {
	messages, err := s.ChannelMessages(channelID, 100, "", "", "")
	if err != nil {
		return nil, fmt.Errorf("error getting messages: %w", err)
	}
	return messages, nil
}

func categorizeMessages(messages []*discordgo.Message) map[string][]model.Message {
	now := time.Now()

	startOfWeek := now.AddDate(0, 0, -int(now.Weekday()))
	startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)
	startOfQuarter := time.Date(now.Year(), (now.Month()-1)/3*3+1, 1, 0, 0, 0, 0, time.Local)
	startOfYear := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, time.Local)

	var weeklyMessages []model.Message
	var monthlyMessages []model.Message
	var quarterlyMessages []model.Message
	var yearlyMessages []model.Message

	for _, message := range messages {
		messageData := model.Message{
			ID:        message.ID,
			Author:    message.Author.Username,
			Content:   message.Content,
			Timestamp: message.Timestamp,
		}

		timestamp := message.Timestamp
		if timestamp.After(startOfWeek) {
			weeklyMessages = append(weeklyMessages, messageData)
		}
		if timestamp.After(startOfMonth) {
			monthlyMessages = append(monthlyMessages, messageData)
		}
		if timestamp.After(startOfQuarter) {
			quarterlyMessages = append(quarterlyMessages, messageData)
		}
		if timestamp.After(startOfYear) {
			yearlyMessages = append(yearlyMessages, messageData)
		}
	}

	return map[string][]model.Message{
		"This Week":    weeklyMessages,
		"This Month":   monthlyMessages,
		"This Quarter": quarterlyMessages,
		"This Year":    yearlyMessages,
	}
}

func printCategorizedMessages(messages map[string][]model.Message) {
	for period, msgs := range messages {
		fmt.Printf("Messages from %s:\n", period)
		for _, message := range msgs {
			fmt.Printf("Message ID: %s\nAuthor: %s\nContent: %s\nDate: %s\n\n",
				message.ID, message.Author, message.Content, message.Timestamp.Format(time.RFC3339))
		}
	}
}
