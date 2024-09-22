package provider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("ZennProvider: error reading response: %v\n", err)
		return []string{}, nil
	}

	var articles []model.Article
	if err := json.Unmarshal(body, &articles); err != nil {
		log.Printf("ZennProvider: error unmarshalling response: %v\n", err)
		return []string{}, nil
	}

	today := time.Now().Format("2006-01-02")

	var todaysArticles []string
	for _, article := range articles {
		publishedTime, err := time.Parse(time.RFC3339, article.PublishedAt)
		if err != nil {
			log.Printf("ZennProvider: error parsing date: %v\n", err)
			continue
		}
		publishedAt := publishedTime.Format("2006-01-02")

		if publishedAt == today {
			// 必要なデータをリストに追加
			todaysArticles = append(todaysArticles, fmt.Sprintf("Title: %s, URL: https://zenn.dev%s", article.Title, article.Path))
		}
	}
	if len(todaysArticles) == 0 {
		log.Printf("ZennProvider: no articles found for today")
		return []string{}, nil
	}

	return todaysArticles, nil
}
