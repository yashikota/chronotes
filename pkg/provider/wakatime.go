package provider

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	model "github.com/yashikota/chronotes/model/v1/provider"
)

func WakatimeProvider(apiKey string, startDate string, endDate string) ([]model.LanguageSummary, error) {
	baseURL := "https://wakatime.com/api/v1"
	url := fmt.Sprintf("%s/users/current/summaries?start=%s&end=%s", baseURL, startDate, endDate)
	fmt.Println("URL:", url)

	encodedAPIKey := base64.StdEncoding.EncodeToString([]byte(apiKey))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		slog.Error("Wakatime: Error creating request" + err.Error())
		return nil, err
	}

	req.Header.Add("Authorization", "Basic "+encodedAPIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		slog.Error("Wakatime: Error sending request" + err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			slog.Error("Wakatime: Error reading response body" + err.Error())
			return nil, err
		}

		var stats model.WakatimeResponse
		if err := json.Unmarshal(body, &stats); err != nil {
			slog.Error("Wakatime: Error unmarshalling response body" + err.Error())
			return nil, err
		}

		// 言語とその作業時間のリストを作成
		var languageSummaries []model.LanguageSummary
		for _, lang := range stats.Data[0].Languages {
			languageSummaries = append(languageSummaries, model.LanguageSummary{
				Name: lang.Name,
				Time: lang.TotalSeconds / 60,
			})
		}

		return languageSummaries, nil
	}
	return []model.LanguageSummary{}, nil
}
