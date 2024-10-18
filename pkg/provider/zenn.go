package provider

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"time"

	model "github.com/yashikota/chronotes/model/v1/provider"
)

func ZennProvider(username string) ([]string, error) {
	url := fmt.Sprintf("https://zenn.dev/api/articles?username=%s", username)
	resp, err := http.Get(url)
	if err != nil {
		slog.Error("ZennProvider: error getting articles" + err.Error())
		return []string{}, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		slog.Error("ZennProvider: error getting articles" + resp.Status)
		return []string{}, nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("ZennProvider: error reading response" + err.Error())
		return []string{}, nil
	}

	var result struct {
		Articles []model.Article `json:"articles"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		slog.Error("ZennProvider: error unmarshalling response" + err.Error())
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
			slog.Error("ZennProvider: error parsing date" + err.Error())
			continue
		}
		// タイムゾーンを考慮して比較
		publishedAt := publishedTime.In(loc).Format("2006-01-02")
		if publishedAt == today {
			todaysArticles = append(todaysArticles, fmt.Sprintf("Title: %s", article.Title))
		}
	}
	if len(todaysArticles) == 0 {
		slog.Error("ZennProvider: no articles found for today")
		return []string{}, nil
	}
	slog.Debug("ZennProvider: today's articles", "articles", todaysArticles)
	todaysArticles[0] = strings.Replace(todaysArticles[0], "Title: ", "", 1)
	return todaysArticles, nil
}
