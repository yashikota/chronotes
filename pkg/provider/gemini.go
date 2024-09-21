package provider

import (
	"fmt"

	model "github.com/yashikota/chronotes/model/v1/provider"
	"github.com/yashikota/chronotes/pkg/utils"
)

func Gemini(input model.Gemini) (model.Response, error) {
	var text []string
	var result []string
	var day string
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
	day = utils.GetDay()
	fmt.Println(day)
	result, err := utils.SummarizeText(text)
	if err != nil {
		return model.Response{}, fmt.Errorf("error summarizing text: %v", err)
	}
	return model.Response{
		Result: result,
		Day:    day,
	}, nil
}
