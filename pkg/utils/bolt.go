package utils

import (
	"fmt"
	"net/http"
	"os"
)

type SlackResponse struct {
	Ok          bool   `json:"ok"`
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
	BotUserID   string `json:"bot_user_id"`
	AppID       string `json:"app_id"`
	Team        struct {
		Name string `json:"name"`
		ID   string `json:"id"`
	} `json:"team"`
}

func SlackOAuthFlow(w http.ResponseWriter, r *http.Request) {
	baseURL := "https://slack.com/oauth/v2"
	scope := "channels:join,channels:read,groups:read"
	clientID := os.Getenv("SLACK_CLIENT_ID")
	redirectURI := "https://localhost:8080/oauth/callback"
	url := fmt.Sprintf("%s/authorize?scope=%s&client_id=%s&redirect_uri=%s", baseURL, scope, clientID, redirectURI)
	http.Redirect(w, r, url, http.StatusFound)

	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Code is empty", http.StatusBadRequest)
		return
	}
}
