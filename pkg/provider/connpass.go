package provider

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

	model "github.com/yashikota/chronotes/model/v1/provider"
)

func ConnpassProvider(userID string) ([]string, error) {

	baseURL := "https://chronotes:kota@chronotes.yashikota.com/connpass/api/v1/user"
	url := fmt.Sprintf("%s/%s/attended_event/", baseURL, userID)
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

	// startDate と endDate を Connpass のデータから取得
	if len(response.Events) == 0 {
		slog.Info("Connpass : No events found")
		return []string{}, nil
	}
	startDate := response.Events[0].StartedAt                  // 最初のイベントの開始日
	endDate := response.Events[len(response.Events)-1].EndedAt // 最後のイベントの終了日

	// 今日の日付を取得
	today := time.Now()

	// 今日が startDate と endDate の間にあるかチェック
	if today.Before(startDate) || today.After(endDate) {
		slog.Info("Connpass : No events found")
		return []string{}, nil
	}

	var titles []string
	for _, event := range response.Events {
		titles = append(titles, event.Title)
	}
	return titles, nil
}
