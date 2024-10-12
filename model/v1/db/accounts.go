package model

type Accounts struct {
	// ID			     uint   `json:"id" gorm:"primaryKey"`
	UserID           string `json:"user_id"`
	GitHubUserID     string `json:"github_user_id"`
	DiscordChannelID string `json:"discord_channel_id"`
	SlackChannelID   string `json:"slack_channel_id"`
	QiitaUserID      string `json:"qiita_user_id"`
	ZennUsername     string `json:"zenn_username"`
	WakatimeToken    string `json:"wakatime_token"`
}

func NewAccounts() *Accounts {
	return &Accounts{}
}
