package provider

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ktsujichan/qiita-sdk-go/qiita"
)

func QiitaProvider(userID string) ([]string, error) {
	var todayItems []string
	config := qiita.NewConfig()
	c, _ := qiita.NewClient("<qiita access token>", *config)
	today := time.Now().Format("2006-01-02")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	items, err := c.ListItems(ctx, 1, 100, "user:"+userID)
	if err != nil {
		return nil, fmt.Errorf("error fetching items: %v", err)
	}

	for _, item := range *items {
		createdAt := strings.Split(item.CreatedAt, "T")[0]

		if createdAt == today {
			todayItems = append(todayItems, item.Title)
		}
	}

	if len(todayItems) == 0 {
		return nil, fmt.Errorf("no items found for today")
	}
	fmt.Println("todayItems: ", todayItems)
	return todayItems, nil
}
