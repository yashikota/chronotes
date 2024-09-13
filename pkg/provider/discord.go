package provider

import (
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

const (
	Token = "DISCORD_BOT_TOKEN"
)

func Discordprovider() {
	channelID := "channel_ID" // 取得したいチャンネルのIDを設定

	runBot(channelID)
}

// ボットを起動して指定されたチャンネルIDからメッセージを取得する関数
func runBot(channelID string) {
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(ready)

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL+C to exit.")

	// チャンネルIDからメッセージを取得する関数を呼び出す
	fetchMessagesFromChannel(dg, channelID)

	// ボットが動作し続けるようにするための待機
	select {}
}

// Botが準備完了したときに呼ばれるハンドラ
func ready(s *discordgo.Session, event *discordgo.Ready) {
	fmt.Println("Bot is ready")
}

// メッセージ履歴を取得する関数
func getMessageHistory(s *discordgo.Session, channelID string) {
	// 最新のメッセージから一定量を取得（ここでは100件）
	messages, err := s.ChannelMessages(channelID, 100, "", "", "")
	if err != nil {
		log.Println("Error getting messages:", err)
		return
	}

	// 現在の日時を取得
	now := time.Now()

	// 期間を定義
	startOfWeek := now.AddDate(0, 0, -int(now.Weekday()))
	startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)
	startOfQuarter := time.Date(now.Year(), (now.Month()-1)/3*3+1, 1, 0, 0, 0, 0, time.Local)
	startOfYear := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, time.Local)

	// 各期間ごとにメッセージをフィルタリングして表示
	var weeklyMessages []discordgo.Message
	var monthlyMessages []discordgo.Message
	var quarterlyMessages []discordgo.Message
	var yearlyMessages []discordgo.Message

	for _, message := range messages {
		timestamp := message.Timestamp

		if timestamp.After(startOfWeek) {
			weeklyMessages = append(weeklyMessages, *message)
		}
		if timestamp.After(startOfMonth) {
			monthlyMessages = append(monthlyMessages, *message)
		}
		if timestamp.After(startOfQuarter) {
			quarterlyMessages = append(quarterlyMessages, *message)
		}
		if timestamp.After(startOfYear) {
			yearlyMessages = append(yearlyMessages, *message)
		}
	}

	fmt.Println("Messages from this week:")
	printMessages(weeklyMessages)

	fmt.Println("Messages from this month:")
	printMessages(monthlyMessages)

	fmt.Println("Messages from this quarter:")
	printMessages(quarterlyMessages)

	fmt.Println("Messages from this year:")
	printMessages(yearlyMessages)
}

// メッセージを表示する関数
func printMessages(messages []discordgo.Message) {
	for _, message := range messages {
		fmt.Printf("Message ID: %s\nAuthor: %s\nContent: %s\nDate: %s\n\n", message.ID, message.Author.Username, message.Content, message.Timestamp.Format(time.RFC3339))
	}
}

// ここでチャンネルIDを直接指定してメッセージを取得
func fetchMessagesFromChannel(s *discordgo.Session, channelID string) {
	fmt.Printf("Fetching messages from channel: %s\n", channelID)
	go getMessageHistory(s, channelID)
}
