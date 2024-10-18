package provider_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/yashikota/chronotes/pkg/provider"
	"github.com/yashikota/chronotes/pkg/utils"
)

func TestSlackHandler(t *testing.T) {
	w := httptest.NewRecorder()
	token := os.Getenv("SLACK_TOKEN")
	channelID := os.Getenv("SLACK_CHANNEL_ID")

	if token == "" {
		token = "SLACK_TOKEN"
	}

	if channelID == "" {
		channelID = "SLACK_CHANNEL_ID"
	}

	if token == "" {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, errors.New("SLACK_TOKEN is not set"))
		return
	}

	if channelID == "" {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, errors.New("SLACK_CHANNEL_ID is not set"))
		return
	}

	// provider.SlackProvider の呼び出し
	summaries, err := provider.SlackProvider(channelID)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	if summaries == nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, errors.New("could not fetch commits"))
		return
	}

	fmt.Println(summaries)

	utils.SuccessJSONResponse(w, summaries)
}
