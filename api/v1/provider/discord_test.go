package provider_test

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
	if err != nil && !os.IsNotExist(err) {
		t.Error(err)
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

	if categorizedCommits == nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, errors.New("could not fetch commits"))
		return
	}

	// Todayカテゴリのみを取り出す
	category := "Today"
	commits := categorizedCommits[category]
	if commits == nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, errors.New("commits not found for Today"))
		return
	}

	var results []map[string]string
	for _, commit := range commits {
		result := map[string]string{
			"content": commit.Content,
			"period":  commit.Period,
		}
		results = append(results, result)
	}

	fmt.Println("result[period]:", results[0]["period"])
	fmt.Println("result[content]:", results[0]["content"])
	utils.SuccessJSONResponse(w, results)
}
