package provider

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/ktsujichan/qiita-sdk-go/qiita"
)

func QiitaProvider(userID string) ([]string, error) {
	var todayItems []string
	token := os.Getenv("QIITA_TOKEN")
	if token == "" {
		slog.Error("Qiita: QIITA_TOKEN environment variable is not set")
		return []string{}, nil
	}
	config := qiita.NewConfig()
	c, err := qiita.NewClient(token, *config)
	if err != nil {
		slog.Error("Qiita: Error creating client" + err.Error())
		return []string{}, nil
	}

	// 今日の日付を取得
	today := time.Now().Format("2006-01-02")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	items, err := c.ListItems(ctx, 1, 100, "user:"+userID)
	if err != nil {
		slog.Error("Qiita: Error fetching items" + err.Error())
		return []string{}, nil
	}

	for _, item := range *items {
		createdAt, err := time.Parse(time.RFC3339, item.CreatedAt)
		if err != nil {
			slog.Error("Qiita: Error parsing created_at" + err.Error())
			return []string{}, nil
		}

		// 今日の日付と比較
		if createdAt.Format("2006-01-02") == today {
			// タイトルとボディを含む文字列を作成
			todayItems = append(todayItems, fmt.Sprintf("Title: %s\nBody: %s", item.Title, item.RenderedBody))
		}
	}

	if len(todayItems) == 0 {
		slog.Error("Qiita: No items found for today")
		return []string{}, nil
	}

	return todayItems, nil
}
