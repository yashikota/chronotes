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

func TestGithubHandler(t *testing.T) {
	w := httptest.NewRecorder()
	err := godotenv.Load(fmt.Sprintf(".env.%s", os.Getenv("GO_ENV")))
	if err != nil && !os.IsNotExist(err) {
		t.Error(err)
	}

	token := os.Getenv("GITHUB_TOKEN")
	userID := os.Getenv("GITHUB_USER_ID")

	if token == "" {
		token = "GITHUB_TOKEN"
	}

	if userID == "" {
		userID = "TaueIkumi"
	}

	if token == "" {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, errors.New("GITHUB_TOKEN is not set"))
		return
	}

	if userID == "" {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, errors.New("GITHUB_USER_ID is not set"))
		return
	}
	categorizedCommits, err := provider.GitHubProvider(userID)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	if categorizedCommits == nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, errors.New("could not fetch commits"))
		return
	}
	categoryResults := make(map[string][]map[string]string)
	categories := []string{"Today", "This Week", "This Month", "Q1 (Jan-Mar)", "Q2 (Apr-Jun)", "Q3 (Jul-Sep)", "Q4 (Oct-Dec)", "Older"}

	for _, category := range categories {
		commits := categorizedCommits[category]
		if commits == nil {
			// カテゴリが空でもスキップ
			continue
		}
		fmt.Println("commits:", commits)
		for _, commit := range commits {
			for _, file := range commit.Changes {
				if file.Filename == "" {
					// filenameが空の場合はスキップ
					fmt.Println("Skipping commit due to empty filename:", commit)
					continue
				}
				if file.Status == "" {
					// statusが空の場合はスキップ
					fmt.Println("Skipping commit due to empty status:", commit)
					continue
				}
				if file.Additions < 0 {
					// additionsが不正な場合はスキップ
					fmt.Println("Skipping commit due to negative additions:", commit)
					continue
				}
				if file.Deletions < 0 {
					// deletionsが不正な場合はスキップ
					fmt.Println("Skipping commit due to negative deletions:", commit)
					continue
				}
				if file.Changes < 0 {
					// changesが不正な場合はスキップ
					fmt.Println("Skipping commit due to negative changes:", commit)
					continue
				}
				if file.Patch == "" {
					// patchが空の場合はスキップ
					fmt.Println("Skipping commit due to empty patch:", commit)
					continue
				}

				result := map[string]string{
					"period":   commit.Period,
					"filename": file.Filename,
					"patch":    file.Patch,
					"message":  commit.Message, // コミットメッセージを追加
				}

				categoryResults[category] = append(categoryResults[category], result)
			}
		}
	}

	// 指定した期間の結果を出力する
	periodToPrint := "Today" // 例: "Today" を指定
	resultsToPrint, exists := categoryResults[periodToPrint]
	if exists {
		fmt.Printf("Results for period '%s':\n", periodToPrint)
		for _, result := range resultsToPrint {
			fmt.Printf("Period: %s, Filename: %s, Patch: %s, Message: %s\n", result["period"], result["filename"], result["patch"], result["message"])
		}
	} else {
		fmt.Printf("No results found for period '%s'.\n", periodToPrint)
	}

	if len(categoryResults) == 0 {
		utils.ErrorJSONResponse(w, http.StatusNoContent, errors.New("no results found"))
		return
	}

	utils.SuccessJSONResponse(w, categoryResults)
}
