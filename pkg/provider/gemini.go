package provider

import (
	"fmt"
	"log"
	"strings"

	model "github.com/yashikota/chronotes/model/v1/provider"
	"github.com/yashikota/chronotes/pkg/utils"
)

func Gemini(input model.Gemini) (model.Response, error) {
	var text []string
	var summary []string
	var result string
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

	if len(text) == 0 {
		return model.Response{
			Result: "",
			Title:  "",
			Day:    day,
		}, nil
	}

	summary, err := utils.SummarizeText(text)
	if err != nil {
		log.Printf("Gemini : error summarizing text: %v\n", err)
		return model.Response{
			Result: "",
			Title:  "",
			Day:    day,
			Tag:    "",
		}, nil
	}
	result = strings.Join(summary, "\n")
	title, err := utils.MakeTitle(summary)
	if err != nil {
		log.Printf("Gemini : error making title: %v\n", err)
		return model.Response{
			Result: result,
			Title:  day,
			Day:    day,
			Tag:    "",
		}, nil
	}

	tag, err := utils.MakeTag(summary)
	if err != nil {
		log.Printf("Gemini : error making tag: %v\n", err)
		return model.Response{
			Result: result,
			Title:  title,
			Day:    day,
			Tag:    "",
		}, nil
	}
	fmt.Printf("Gemini : result: %s, title: %s, day: %s, tag: %s\n", result, title, day, tag)
	return model.Response{
		Result: result,
		Title:  title,
		Day:    day,
		Tag:    tag,
	}, nil
}
