package handler_test

import (
	"os"
	"testing"
	"time"

	"github.com/yashikota/chronotes/pkg/provider"
)

func TestDiscordHandler(t *testing.T) {
	// 環境変数を設定する
	discordToken := os.Getenv("DISCORD_TOKEN")
	if discordToken == "" {
		t.Fatal("DISCORD_TOKEN environment variable is not set")
	}
	channelID := os.Getenv("CHANNEL_ID")
	if channelID == "" {
		t.Fatal("DISCORD_CHANNEL_ID environment variable is not set")
	}
	categorizedMessages, err := provider.DiscordProvider(channelID)
	if err != nil {
		t.Fatalf("Error fetching data: %v", err)
	}

	// カテゴリごとに確認する
	categories := []string{"Today", "This Week", "This Month", "Q1 (Jan-Mar)", "Q2 (Apr-Jun)", "Q3 (Jul-Sep)", "Q4 (Oct-Dec)", "This Year"}

	for _, category := range categories {
		messages := categorizedMessages[category]

		// カテゴリにメッセージがある場合、その内容を出力
		if len(messages) > 0 {
			t.Logf("\nCategory: %s\n", category)
			for _, message := range messages {
				// メッセージの作成日付を整形
				period := formatPeriod(message.Timestamp)

				// メッセージ内容と期間を出力
				t.Logf("Message Content: %s\n", message.Content)
				t.Logf("Period: %s\n", period)
			}
		} else {
			t.Logf("\nCategory: %s - No messages found.\n", category)
		}
	}
}

// メッセージのタイムスタンプから、指定したフォーマットで期間を取得
func formatPeriod(timestamp time.Time) string {
	return timestamp.Format("2006-01-02 15:04:05")
}
