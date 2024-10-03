package model

type Gemini struct {
	UserID           string `json:"user_id"`
	GitHubUserID     string `json:"github_user_id"`
	DiscordChannelID string `json:"discord_channel_id"`
	SlackChannelID   string `json:"slack_channel_id"`
	QiitaUserID      string `json:"qiita_user_id"`
	ZennUsername     string `json:"zenn_username"`
	WakatimeToken    string `json:"wakatime_token"`
}

type Response struct {
	Result string `json:"result"`
	Title  string `json:"title"`
	Day    string `json:"day"`
	Tag    string `json:"tag"`
}
