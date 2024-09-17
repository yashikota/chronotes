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

func TestGithubHandler(t *testing.T) {
	w := httptest.NewRecorder()
	err := godotenv.Load(fmt.Sprintf(".env.%s", os.Getenv("GO_ENV")))
	if err != nil {
		t.Fatalf("Failed to load .env file: %v", err)
	}

	token := os.Getenv("GITHUB_TOKEN")
	userID := os.Getenv("GITHUB_USER_ID")

	if token == "" {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, errors.New("GITHUB_TOKEN is not set"))
		return
	}

	if userID == "" {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, errors.New("USER_ID is not set"))
		return
	}

	categorizedCommits, err := provider.GitHubProvider(userID)

	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	if categorizedCommits != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, errors.New("could not fetch commits"))
		return
	}

	categories := []string{"Today", "This Week", "This Month", "Q1 (Jan-Mar)", "Q2 (Apr-Jun)", "Q3 (Jul-Sep)", "Q4 (Oct-Dec)", "Older"}
	var results []map[string]string

	for _, category := range categories {
		commits := categorizedCommits[category]

		if commits == nil {
			utils.ErrorJSONResponse(w, http.StatusBadRequest, errors.New("commits not found"))
			return
		}
		for _, commit := range commits {
			for _, file := range commit.Changes {
				if file.Filename == "" {
					utils.ErrorJSONResponse(w, http.StatusBadRequest, errors.New("filename not found"))
					return
				}
				if file.Status == "" {
					utils.ErrorJSONResponse(w, http.StatusBadRequest, errors.New("status not found"))
					return
				}
				if file.Additions < 0 {
					utils.ErrorJSONResponse(w, http.StatusBadRequest, errors.New("additions not found"))
					return
				}
				if file.Deletions < 0 {
					utils.ErrorJSONResponse(w, http.StatusBadRequest, errors.New("deletions not found"))
					return
				}
				if file.Changes < 0 {
					utils.ErrorJSONResponse(w, http.StatusBadRequest, errors.New("changes not found"))
					return
				}
				if file.Patch == "" {
					utils.ErrorJSONResponse(w, http.StatusBadRequest, errors.New("patch not found"))
					return
				}

				result := map[string]string{
					"period":   commit.Period,
					"filename": file.Filename,
					"patch":    file.Patch,
				}

				results = append(results, result)
			}
		}
	}
	utils.SuccessJSONResponse(w, results)
}
