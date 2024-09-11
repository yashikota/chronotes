package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	token := "GITHUB TOKEN"
	if token == "" {
		log.Fatal("GITHUB_TOKEN environment variable is required")
	}
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	user := "GITHUB-USERNAME" // ここにGitHubのユーザー名を指定

	// リポジトリリストを取得
	repos, _, err := client.Repositories.List(ctx, user, nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, repo := range repos {
		fmt.Println("Repository:", *repo.Name)

		// 各リポジトリのコミット履歴を取得
		commits, _, err := client.Repositories.ListCommits(ctx, *repo.Owner.Login, *repo.Name, nil)
		if err != nil {
			log.Fatal(err)
		}

		for _, commit := range commits {
			if commit.Author != nil && *commit.Author.Login == user {
				fmt.Println("Commit SHA:", *commit.SHA)
				fmt.Println("Date:", commit.Commit.Author.Date.Format(time.RFC3339))
				fmt.Println("Message:", *commit.Commit.Message)

				// コミットの詳細を取得
				detailedCommit, _, err := client.Repositories.GetCommit(ctx, *repo.Owner.Login, *repo.Name, *commit.SHA)
				if err != nil {
					log.Fatalf("Error getting commit details for SHA %s: %v", *commit.SHA, err)
				}

				// コミットのファイルの情報を表示
				fmt.Println("Changed files:")
				if detailedCommit.Files != nil {
					for _, file := range detailedCommit.Files {
						if file.Filename == nil {
							fmt.Println("Filename is nil")
						}
						if file.Status == nil {
							fmt.Println("Status is nil")
						}
						if file.Additions == nil {
							fmt.Println("Additions is nil")
						}
						if file.Deletions == nil {
							fmt.Println("Deletions is nil")
						}
						if file.Changes == nil {
							fmt.Println("Changes is nil")
						}
						if file.Patch == nil {
							fmt.Println("Patch is nil")
						}
						// 正常な場合のみ表示
						if file.Filename != nil {
							fmt.Println("File:", *file.Filename)
						}
						if file.Status != nil {
							fmt.Println("Status:", *file.Status)
						}
						if file.Additions != nil {
							fmt.Println("Additions:", *file.Additions)
						}
						if file.Deletions != nil {
							fmt.Println("Deletions:", *file.Deletions)
						}
						if file.Changes != nil {
							fmt.Println("Changes:", *file.Changes)
						}
						if file.Patch != nil {
							fmt.Println("Patch:", *file.Patch)
						}
						fmt.Println()
					}
				} else {
					fmt.Println("No files changed or no detailed information available")
				}
				fmt.Println()
			}
		}
	}
}
