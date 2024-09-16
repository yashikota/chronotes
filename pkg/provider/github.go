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

	// 上位カテゴリから順にフィルタリングするように順序を設定
	filterCategories := []string{"Today", "This Week", "This Month", "This Year", "Q1 (Jan-Mar)", "Q2 (Apr-Jun)", "Q3 (Jul-Sep)", "Q4 (Oct-Dec)"}

	categorizedCommits := make(map[string][]model.CommitInfo)

	// リポジトリリストを取得
	repos, _, err := client.Repositories.List(ctx, userID, nil)
	if err != nil {
		return nil, fmt.Errorf("error fetching repositories: %v", err)
	}

	for _, repo := range repos {
		if repo == nil || repo.Owner == nil || repo.Name == nil {
			log.Printf("Skipping repository due to nil Owner or Name")
			continue
		}

		// 各リポジトリのコミット履歴を取得
		commits, _, err := client.Repositories.ListCommits(ctx, *repo.Owner.Login, *repo.Name, nil)
		if err != nil {
			continue
		}
		if len(commits) == 0 {
			// リポジトリが空である場合はスキップ
			log.Printf("Skipping empty repository %s", *repo.Name)
			continue
		}

		// 指定したカテゴリのコミットのみをフィルタリング（重複を避ける）
		filteredCommits, err := filterCommitsByCategories(commits, filterCategories, client, repo)
		if err != nil {
			return nil, fmt.Errorf("error filtering commits for repository %s: %v", *repo.Name, err)
		}

		// フィルタリングされたコミットのデバッグ出力
		for category, commits := range filteredCommits {
			if _, exists := categorizedCommits[category]; !exists {
				categorizedCommits[category] = []model.CommitInfo{}
			}
			categorizedCommits[category] = append(categorizedCommits[category], commits...)
		}
	}

	return categorizedCommits, nil
}

func categorizeCommitDate(date time.Time) string {
	now := time.Now().UTC() // Current UTC time for comparison

	// 今日のコミットかどうか
	if date.Year() == now.Year() && date.YearDay() == now.YearDay() {
		return "Today"
	}

	// 今日でない場合は今週かどうか判定
	nowYear, nowWeek := now.ISOWeek()
	commitYear, commitWeek := date.ISOWeek()
	if commitYear == nowYear && commitWeek == nowWeek {
		return "This Week"
	}

	// 今週でない場合は今月かどうか判定
	if date.Month() == now.Month() && date.Year() == now.Year() {
		return "This Month"
	}

	// 今月でない場合は今年かどうか判定
	if date.Year() == now.Year() {
		// 四半期を判定
		month := int(date.Month())
		quarter := (month-1)/3 + 1
		switch quarter {
		case 1:
			return "Q1 (Jan-Mar)"
		case 2:
			return "Q2 (Apr-Jun)"
		case 3:
			return "Q3 (Jul-Sep)"
		case 4:
			return "Q4 (Oct-Dec)"
		}
		return "This Year" // 四半期に該当しない場合は今年
	}

	// それ以外は古いコミット
	return "Older"
}

func filterCommitsByCategories(commits []*github.RepositoryCommit, categories []string, client *github.Client, repo *github.Repository) (map[string][]model.CommitInfo, error) {
	filteredCommits := make(map[string][]model.CommitInfo)
	ctx := context.Background()

	for _, commit := range commits {
		if commit == nil || commit.Author == nil || commit.Commit == nil || commit.Commit.Author == nil || commit.Commit.Author.Date == nil {
			log.Println("Skipping invalid commit")
			continue
		}

		date := *commit.Commit.Author.Date
		commitCategory := categorizeCommitDate(date)

		for _, filterCat := range categories {
			if filterCat == commitCategory {
				if _, exists := filteredCommits[filterCat]; !exists {
					filteredCommits[filterCat] = []model.CommitInfo{}
				}

				// コミットの詳細を取得する
				detailedCommit, _, err := client.Repositories.GetCommit(ctx, *repo.Owner.Login, *repo.Name, *commit.SHA)
				if err != nil {
					log.Printf("Error getting commit details for SHA %s: %v", *commit.SHA, err)
					return nil, err
				}

				fileChanges := []model.FileChange{}
				if detailedCommit.Files != nil {
					for _, file := range detailedCommit.Files {
						fileChange := model.FileChange{
							Filename:  *file.Filename,
							Status:    *file.Status,    // ファイルのステータス（追加、削除、変更）
							Additions: *file.Additions, // 追加された行数
							Deletions: *file.Deletions, // 削除された行数
							Changes:   *file.Changes,   // 変更された行数
							Patch:     "",              // Patch が nil の場合に空の文字列を設定
						}
						if file.Patch != nil {
							fileChange.Patch = *file.Patch
						}

						fileChanges = append(fileChanges, fileChange)
					}
				}

				commitInfo := model.CommitInfo{
					Message: *commit.Commit.Message,
					Changes: fileChanges,
					Period:  commitCategory, // 追加: コミットが属する期間
				}

				filteredCommits[filterCat] = append(filteredCommits[filterCat], commitInfo)
				break
			}
		}
	}

	return filteredCommits, nil
}
