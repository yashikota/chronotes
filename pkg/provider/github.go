package provider

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/go-github/github"
	model "github.com/yashikota/chronotes/model/v1/provider"
	"golang.org/x/oauth2"
)

func GitHubProvider(userID string) (map[string][]model.CommitInfo, error) {
	ctx := context.Background()
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("GITHUB_TOKEN environment variable is required")
	}
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	filterCategories := []string{"This Week", "This Month", "This Year"}

	categorizedCommits := make(map[string][]model.CommitInfo)

	// リポジトリリストを取得
	repos, _, err := client.Repositories.List(ctx, userID, nil)
	if err != nil {
		return nil, err
	}

	for _, repo := range repos {
		fmt.Println("Repository:", *repo.Name)

		// 各リポジトリのコミット履歴を取得
		commits, _, err := client.Repositories.ListCommits(ctx, *repo.Owner.Login, *repo.Name, nil)
		if err != nil {
			return nil, err
		}

		// 指定したカテゴリのコミットのみをフィルタリング
		filteredCommits, err := filterCommitsByCategories(commits, filterCategories, client, repo)
		if err != nil {
			return nil, err
		}

		for category, commits := range filteredCommits {
			categorizedCommits[category] = commits
		}
	}

	return categorizedCommits, nil
}

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
func filterCommitsByCategories(commits []*github.RepositoryCommit, categories []string, client *github.Client, repo *github.Repository) (map[string][]model.CommitInfo, error) {
	filteredCommits := make(map[string][]model.CommitInfo)
	ctx := context.Background()

	for _, commit := range commits {
		if commit.Author != nil && commit.Commit.Author.Date != nil {
			date := *commit.Commit.Author.Date
			commitCategories := categorizeCommitDate(date)
			categoryMap := make(map[string]bool)
			for _, c := range commitCategories {
				categoryMap[c] = true
			}

			for _, filterCat := range categories {
				if categoryMap[filterCat] {
					if _, exists := filteredCommits[filterCat]; !exists {
						filteredCommits[filterCat] = []model.CommitInfo{}
					}

					// コミットの詳細を取得
					detailedCommit, _, err := client.Repositories.GetCommit(ctx, *repo.Owner.Login, *repo.Name, *commit.SHA)
					if err != nil {
						log.Printf("Error getting commit details for SHA %s: %v", *commit.SHA, err)
						continue // Proceed with other commits
					}

					// コミットのファイルの情報を収集
					fileChanges := []model.FileChange{}
					if detailedCommit.Files != nil {
						for _, file := range detailedCommit.Files {
							fileChange := model.FileChange{
								Filename:  *file.Filename,
								Status:    *file.Status,
								Additions: *file.Additions,
								Deletions: *file.Deletions,
								Changes:   *file.Changes,
								Patch:     *file.Patch,
							}
							fileChanges = append(fileChanges, fileChange)
						}
					}

					commitInfo := model.CommitInfo{
						Message: *commit.Commit.Message,
						Changes: fileChanges,
					}

					filteredCommits[filterCat] = append(filteredCommits[filterCat], commitInfo)
					break
				}
			}
		}
	}
	return filteredCommits, nil
}
