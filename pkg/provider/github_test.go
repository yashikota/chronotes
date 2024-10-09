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

func TestGithubHandler(t *testing.T) {
	w := httptest.NewRecorder()

	token := os.Getenv("GITHUB_TOKEN")
	userID := os.Getenv("GITHUB_USER_ID")

	if token == "" {
		token = "GITHUB_TOKEN"
	}

	if userID == "" {
		userID = "taueikumi"
	}

	if token == "" {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, errors.New("GITHUB_TOKEN is not set"))
		return
	}

	if userID == "" {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, errors.New("GITHUB_USER_ID is not set"))
		return
	}
	summaries, err := provider.GitHubProvider(userID)
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
