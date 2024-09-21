package provider

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ktsujichan/qiita-sdk-go/qiita"
)

func QiitaProvider(userID string) ([]string, error) {
	var todayItems []string
	token := os.Getenv("QIITA_TOKEN")
	if token == "" {
		log.Printf("Qiita: QIITA_TOKEN environment variable is not set")
		return nil, nil
	}
	config := qiita.NewConfig()
	c, err := qiita.NewClient(token, *config)
	if err != nil {
		log.Printf("Qiita: Error creating client: %v", err)
		return nil, nil
	}

	// 今日の日付を取得
	today := time.Now().Format("2006-01-02")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	items, err := c.ListItems(ctx, 1, 100, "user:"+userID)
	if err != nil {
		log.Printf("Qiita: Error fetching items: %v", err)
		return nil, nil
	}

	for _, item := range *items {
		createdAt, err := time.Parse(time.RFC3339, item.CreatedAt)
		if err != nil {
			log.Printf("Qiita: Error parsing created_at: %v", err)
			return nil, nil
		}

		// 今日の日付と比較
		if createdAt.Format("2006-01-02") == today {
			// タイトルとボディを含む文字列を作成
			todayItems = append(todayItems, fmt.Sprintf("Title: %s\nBody: %s", item.Title, item.RenderedBody))
		}
	}

	if len(todayItems) == 0 {
		log.Printf("Qiita: No items found for today")
		return nil, nil
	}

	return todayItems, nil
}
