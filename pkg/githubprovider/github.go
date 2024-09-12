package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// Function to categorize the commit date
func categorizeCommitDate(date time.Time) []string {
	now := time.Now()
	var categories []string

	// Check for this week
	nowYear, nowWeek := now.ISOWeek()
	commitYear, commitWeek := date.ISOWeek()
	if commitYear == nowYear && commitWeek == nowWeek {
		categories = append(categories, "This Week")
	}

	// Check for this month
	if date.Month() == now.Month() && date.Year() == now.Year() {
		categories = append(categories, "This Month")
	}

	// Check for this year
	if date.Year() == nowYear {
		categories = append(categories, "This Year")
	}

	// Check for previous weeks of the current year
	if commitYear == nowYear-1 && commitWeek < nowWeek {
		categories = append(categories, "Previous Weeks")
	}

	// If no category matched, add "Older"
	if len(categories) == 0 {
		categories = append(categories, "Older")
	}

	return categories
}

// Function to filter commits by multiple categories
func filterCommitsByCategories(commits []*github.RepositoryCommit, categories []string) []*github.RepositoryCommit {
	var filteredCommits []*github.RepositoryCommit
	for _, commit := range commits {
		if commit.Author != nil && commit.Commit.Author.Date != nil {
			date := *commit.Commit.Author.Date
			commitCategories := categorizeCommitDate(date)
			for _, c := range commitCategories {
				for _, filterCat := range categories {
					if c == filterCat {
						filteredCommits = append(filteredCommits, commit)
						break
					}
				}
			}
		}
	}
	return filteredCommits
}

func main() {
	ctx := context.Background()
	token := "GITHUB_TOKEN" // ここにGitHubのアクセストークンを指定
	if token == "" {
		log.Fatal("GITHUB_TOKEN environment variable is required")
	}
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	user := "TaueIkumi"
	filterCategories := []string{"This Week", "This Month", "This Year"}

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

		// 指定したカテゴリのコミットのみをフィルタリング
		filteredCommits := filterCommitsByCategories(commits, filterCategories)

		for _, commit := range filteredCommits {
			if commit.Author != nil && commit.Commit.Author.Date != nil {
				date := commit.Commit.Author.Date
				fmt.Println("Commit SHA:", *commit.SHA)
				fmt.Println("Date:", date.Format(time.RFC3339))
				fmt.Println("Message:", *commit.Commit.Message)
				fmt.Println("Category:", categorizeCommitDate(*date)) // Show all categories for the commit

				// コミットの詳細を取得
				detailedCommit, _, err := client.Repositories.GetCommit(ctx, *repo.Owner.Login, *repo.Name, *commit.SHA)
				if err != nil {
					log.Fatalf("Error getting commit details for SHA %s: %v", *commit.SHA, err)
				}

				// コミットのファイルの情報を表示
				fmt.Println("Changed files:")
				if detailedCommit.Files != nil {
					for _, file := range detailedCommit.Files {
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
