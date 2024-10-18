package provider_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/yashikota/chronotes/pkg/provider"
	"github.com/yashikota/chronotes/pkg/utils"
)

func TestQiitaHandler(t *testing.T) {
	w := httptest.NewRecorder()

	token := os.Getenv("QIITA_TOKEN")
	userID := os.Getenv("QIITA_USER_ID")

	if token == "" {
		token = "QIIITA_TOKEN"
	}

	if userID == "" {
		userID = "loverboy"
	}

	if token == "" {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, errors.New("QIITA_TOKEN is not set"))
		return
	}

	if userID == "" {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, errors.New("QIITA_USER_ID is not set"))
		return
	}

	summaries, err := provider.QiitaProvider(userID)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	if summaries == nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, errors.New("could not fetch commits"))
		return
	}
	utils.SuccessJSONResponse(w, summaries)
}
