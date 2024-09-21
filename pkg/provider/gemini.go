package provider

import (
	model "github.com/yashikota/chronotes/model/v1/provider"
	"github.com/yashikota/chronotes/pkg/utils"
)

func Gemini(input model.Gemini) ([]string, error) {
	var text []string
	var result []string

	// GitHubプロバイダーからのテキストを追加
	if githubText, err := GitHubProvider(input.GitHubUserID); err == nil {
		text = append(text, githubText...)
	}

	// Slackプロバイダーからのテキストを追加
	if slackText, err := SlackProvider(input.SlackChannelID); err == nil {
		text = append(text, slackText...)
	}

	// Discordプロバイダーからのテキストを追加
	if discordText, err := DiscordProvider(input.DiscordChannelID); err == nil {
		text = append(text, discordText...)
	}

	result, err := utils.SummarizeText(text)
	return result, err
}
