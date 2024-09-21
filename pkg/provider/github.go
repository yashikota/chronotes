package provider

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"

	model "github.com/yashikota/chronotes/model/v1/provider"
	"github.com/yashikota/chronotes/pkg/utils"
)

func GitHubProvider(userID string) ([]string, error) {
	ctx := context.Background()
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Printf("GitHub : GITHUB_TOKEN environment variable is not set")
		return []string{}, nil
	}

	if userID == "" {
		log.Printf("GitHub : userID is not set")
		return []string{}, nil
	}
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	filterCategories := []string{"Today"}

	var summaries []string

	repos, _, err := client.Repositories.List(ctx, userID, nil)
	if err != nil {
		log.Printf("GitHub : Error fetching repositories: %v", err)
		return []string{}, nil
	}

	for _, repo := range repos {
		if repo == nil || repo.Owner == nil || repo.Name == nil {
			log.Printf("GitHub : Skipping repository due to nil Owner or Name")
			continue
		}

		commits, _, err := client.Repositories.ListCommits(ctx, *repo.Owner.Login, *repo.Name, nil)
		if err != nil {
			continue
		}

		filteredCommits, err := filterCommitsByCategories(commits, filterCategories, client, repo)
		if err != nil {
			log.Printf("GitHub : Error filtering commits by categories: %v", err)
			continue
		}

		summaries = append(summaries, filteredCommits...)
	}

	if len(summaries) == 0 {
		return []string{}, nil
	}

	finalSummary, err := utils.SummarizeText(summaries)
	if err != nil {
		log.Printf("GitHub : Error summarizing text: %v", err)
		return []string{}, nil
	}
	return finalSummary, nil
}

func filterCommitsByCategories(commits []*github.RepositoryCommit, categories []string, client *github.Client, repo *github.Repository) ([]string, error) {
	var commitMessages []string
	ctx := context.Background()

	for _, commit := range commits {
		if commit == nil || commit.Author == nil || commit.Commit == nil || commit.Commit.Author == nil || commit.Commit.Author.Date == nil {
			log.Println("GitHub : Skipping invalid commit")
			continue
		}
		date := *commit.Commit.Author.Date
		commitCategory := utils.CategorizeCommitDate(date)
		for _, filterCat := range categories {
			if filterCat == commitCategory {

				detailedCommit, _, err := client.Repositories.GetCommit(ctx, *repo.Owner.Login, *repo.Name, *commit.SHA)
				if err != nil {
					log.Printf("GitHub : Error getting commit details for SHA %s: %v", *commit.SHA, err)
					return nil, nil
				}

				fileChanges := []model.FileChange{}
				if detailedCommit.Files != nil {
					for _, file := range detailedCommit.Files {
						patch := ""
						if file.Patch != nil {
							patch = *file.Patch
						}

						fileChange := model.FileChange{
							Filename:  *file.Filename,
							Status:    *file.Status,
							Additions: *file.Additions,
							Deletions: *file.Deletions,
							Changes:   *file.Changes,
							Patch:     patch,
						}

						fileChanges = append(fileChanges, fileChange)
					}
				}

				var changesSummary string
				for _, change := range fileChanges {
					changesSummary += fmt.Sprintf("File: %s, Status: %s, Additions: %d, Deletions: %d, Changes: %d, Patch: %s\n",
						change.Filename, change.Status, change.Additions, change.Deletions, change.Changes, change.Patch)
				}

				// コミットメッセージと変更内容を結合した文字列を作成
				commitSummary := fmt.Sprintf("Message: %s\nChanges:\n%s", *commit.Commit.Message, changesSummary)
				commitMessages = append(commitMessages, commitSummary)
				break
			}
		}
	}
	return commitMessages, nil
}
