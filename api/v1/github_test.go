package handler_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/yashikota/chronotes/pkg/provider"
)

func TestGithubHandler(t *testing.T) {
	// 環境変数を設定する
	githubToken := os.Getenv("GITHUB_TOKEN")
	if githubToken == "" {
		t.Fatal("GITHUB_TOKEN environment variable is not set")
	}
	userID := os.Getenv("GITHUB_USER_ID")
	if userID == "" {
		t.Fatal("GITHUB_USER_ID environment variable is not set")
	}
	categorizedCommits, err := provider.GitHubProvider(userID)
	if err != nil {
		t.Fatalf("Error fetching data: %v", err)
	}

	// リポジトリが空の場合のチェック
	if len(categorizedCommits) == 0 {
		t.Log("No commits found. Skipping category checks.")
		return
	}

	// カテゴリごとに確認する
	categories := []string{"Today", "This Week", "This Month", "Q1 (Jan-Mar)", "Q2 (Apr-Jun)", "Q3 (Jul-Sep)", "Q4 (Oct-Dec)", "Older"}

	for _, category := range categories {
		commits := categorizedCommits[category]

		// カテゴリ名を出力
		if len(commits) > 0 {
			fmt.Printf("\nCategory: %s\n", category)
		}

		// 各カテゴリにコミットがある場合、その内容を出力
		for _, commit := range commits {
			// コミットメッセージと期間を出力
			fmt.Printf("Commit Message: %s\n", commit.Message)
			fmt.Printf("Period: %s\n", commit.Period)

			// ファイルごとの変更を出力
			for _, file := range commit.Changes {
				fmt.Printf("File Name: %s\n", file.Filename)
				fmt.Printf("Status: %s\n", file.Status)
				fmt.Printf("Additions: %d\n", file.Additions)
				fmt.Printf("Deletions: %d\n", file.Deletions)
				fmt.Printf("Changes: %d\n", file.Changes)
				fmt.Printf("Patch: %s\n", file.Patch)
				fmt.Println()
			}
		}
	}
}
