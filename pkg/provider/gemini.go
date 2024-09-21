package provider

import (
	"fmt"
	"log"

	model "github.com/yashikota/chronotes/model/v1/provider"
	"github.com/yashikota/chronotes/pkg/utils"
)

func Gemini(input model.Gemini) (model.Response, error) {
	var text []string
	var result []string
	var day string

	if githubText, err := GitHubProvider(input.GitHubUserID); err == nil {
		text = append(text, githubText...)
	} else {
		log.Printf("GitHubProvider error for user %s: %v\n", input.GitHubUserID, err)
	}

	if slackText, err := SlackProvider(input.SlackChannelID); err == nil {
		text = append(text, slackText...)
	} else {
		log.Printf("SlackProvider error for channel %s: %v\n", input.SlackChannelID, err)
	}

	if discordText, err := DiscordProvider(input.DiscordChannelID); err == nil {
		text = append(text, discordText...)
	} else {
		log.Printf("DiscordProvider error for channel %s: %v\n", input.DiscordChannelID, err)
	}

	if qiitaText, err := QiitaProvider(input.QiitaUserID); err == nil {
		text = append(text, qiitaText...)
	} else {
		log.Printf("QiitaProvider error for user %s: %v\n", input.QiitaUserID, err)
	}
	day = utils.GetDay()
	fmt.Println(day)

	if len(text) == 0 {
		return model.Response{
			Result: []string{"進捗なし"},
			Day:    day,
		}, nil
	}

	result, err := utils.SummarizeText(text)
	if err != nil {
		return model.Response{}, fmt.Errorf("error summarizing text: %v", err)
	}

	return model.Response{
		Result: result,
		Day:    day,
	}, nil
}
