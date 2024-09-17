package handler_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/yashikota/chronotes/pkg/provider"
	"github.com/yashikota/chronotes/pkg/utils"
)

func TestDiscordHandler(t *testing.T) {
	w := httptest.NewRecorder()
	err := godotenv.Load(fmt.Sprintf(".env.%s", os.Getenv("GO_ENV")))
	if err != nil {
		t.Fatalf("Failed to load .env file: %v", err)
	}

	token := os.Getenv("DISCORD_TOKEN")
	channelID := os.Getenv("DISCORD_CHANNEL_ID")

	if token == "" {
		token = "DISCORD_TOKEN"
	}

	if channelID == "" {
		channelID = "1241617406552445011"
	}
	if token == "" {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, errors.New("DISCORD_TOKEN is not set"))
		return
	}

	if channelID == "" {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, errors.New("DISCORD CHANNEL ID is not set"))
		return
	}

	categorizedCommits, err := provider.DiscordProvider(channelID)

	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	if categorizedCommits != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, errors.New("could not fetch commits"))
		return
	}

	categories := []string{
		"Today", "This Week", "This Month", "Q1 (Jan-Mar)", "Q2 (Apr-Jun)", "Q3 (Jul-Sep)", "Q4 (Oct-Dec)", "This Year"}
	var results []map[string]string

	for _, category := range categories {
		commits := categorizedCommits[category]
		if commits == nil {
			utils.ErrorJSONResponse(w, http.StatusBadRequest, errors.New("commits not found"))
			return
		}
		for _, commit := range commits {
			result := map[string]string{
				"content": commit.Content,
				"period":  commit.Period,
			}
			results = append(results, result)
		}
	}
	utils.SuccessJSONResponse(w, results)
}
