package provider

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	model "github.com/yashikota/chronotes/model/v1/provider"
)

func ZennProvider(username string) ([]string, error) {
	url := fmt.Sprintf("https://zenn.dev/api/articles?username=%s", username)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("ZennProvider: error getting articles: %v\n", err)
		return []string{}, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("ZennProvider: error getting articles: %v\n", resp.Status)
		return []string{}, nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("ZennProvider: error reading response: %v\n", err)
		return []string{}, nil
	}

	var result struct {
		Articles []model.Article `json:"articles"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Printf("ZennProvider: error unmarshalling response: %v\n", err)
		return []string{}, nil
	}

	// タイムゾーンを考慮して現在の日付を取得
	loc, _ := time.LoadLocation("Asia/Tokyo")
	now := time.Now().In(loc)
	today := now.Format("2006-01-02")

	var todaysArticles []string
	for _, article := range result.Articles {
		publishedTime, err := time.Parse(time.RFC3339, article.PublishedAt)
		if err != nil {
			log.Printf("ZennProvider: error parsing date: %v\n", err)
			continue
		}
		// タイムゾーンを考慮して比較
		publishedAt := publishedTime.In(loc).Format("2006-01-02")
		if publishedAt == today {
			todaysArticles = append(todaysArticles, fmt.Sprintf("Title: %s", article.Title))
		}
	}
	if len(todaysArticles) == 0 {
		log.Printf("ZennProvider: no articles found for today")
		return []string{}, nil
	}
	// fmt.Println("Today's articles on Zenn", todaysArticles)
	return todaysArticles, nil
}
