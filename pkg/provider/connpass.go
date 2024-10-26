package provider

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"

	model "github.com/yashikota/chronotes/model/v1/provider"
	"github.com/yashikota/chronotes/pkg/utils"
)

func ConnpassProvider(userID string) ([]string, error) {
	basicName := os.Getenv("BASIC_NAME")
	if basicName == "" {
		slog.Error("Connpass: BasicName is not set")
	}
	basicPass := os.Getenv("BASIC_PASS")
	if basicPass == "" {
		slog.Error("Connpass: BasicPass is not set")
	}

	baseURL := "chronotes.yashikota.com/connpass/api/v1/user"
	url := fmt.Sprintf("https://%s:%s@%s/%s/attended_event/", basicName, basicPass, baseURL, userID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		slog.Error("Connpass : Failed to create request")
		return []string{}, err
	}


	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		slog.Error("Connpass : Failed to send request")
		return []string{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("Connpass : Failed to read response body")
		return []string{}, err
	}

	var response model.ConnpassResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		slog.Error("Connpass : Failed to unmarshal response body")
		return []string{}, err
	}

	if len(response.Events) == 0 {
		slog.Info("Connpass : No events found")
		return []string{}, nil
	}

	todayStr := utils.GetDateOnly()
	today, err := time.Parse("2006-01-02", todayStr)

	if err != nil {
		slog.Error("Connpass : Failed to parse today's date")
		return []string{}, err

	}
	var titles []string
	for _, event := range response.Events {
		// イベントの日付を解析
		eventDate, err := time.Parse("2006-01-02", event.StartedAt.Format("2006-01-02"))
		if err != nil {
			slog.Error("Connpass : Failed to parse event date")
			continue
		}

		// イベントの日付が今日と一致する場合のみ追加
		if eventDate.Equal(today) {
			titles = append(titles, event.Title)
		}
	}
	return titles, nil
}
