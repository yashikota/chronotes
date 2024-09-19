package handler_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	model "github.com/yashikota/chronotes/model/v1/provider"
	"github.com/yashikota/chronotes/pkg/provider"
	"github.com/yashikota/chronotes/pkg/utils"
)

func TestSlackHandler(t *testing.T) {
	w := httptest.NewRecorder()
	err := godotenv.Load(fmt.Sprintf(".env.%s", os.Getenv("GO_ENV")))
	if err != nil && !os.IsNotExist(err) {
		t.Error(err)
	}
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
	categorizedCommits, err := provider.SlackProvider(channelID)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	if categorizedCommits == nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, errors.New("could not fetch commits"))
		return
	}

	// 今日のメッセージを取得
	var todayMessages []model.SlackMessage

	// マップのキーと値を処理する
	if messages, ok := categorizedCommits["Today"]; ok {
		todayMessages = messages
	}

	if len(todayMessages) == 0 {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, errors.New("no commits found for Today"))
		return
	}

	var todayResults []string

	for _, commit := range todayMessages {
		todayResults = append(todayResults, commit.Text) // テキスト部分のみを追加
	}

	fmt.Println(todayResults)
	// 今日のデータをレスポンスとして送信
	utils.SuccessJSONResponse(w, todayResults)
}
