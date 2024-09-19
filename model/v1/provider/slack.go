package model

import "time"

type SlackMessage struct {
	ID        string    `json:"id"`         // メッセージのID
	User      string    `json:"user"`       // メッセージのユーザーID
	Text      string    `json:"text"`       // メッセージのテキスト内容
	Timestamp float64   `json:"timestamp"`  // メッセージのタイムスタンプ（UNIXタイム）
	Channel   string    `json:"channel"`    // メッセージが投稿されたチャンネルのID
	CreatedAt time.Time `json:"created_at"` // メッセージの作成日時
}

type CategorizedMessages struct {
	Category string         `json:"category"` // カテゴリ名（例: "daily", "weekly"）
	Messages []SlackMessage `json:"messages"` // カテゴリに属するメッセージのリスト
}
