package notes

import (
	"log/slog"
	"strings"
)

// NoteWithID - ノートのIDとコンテンツを保持する構造体
type NoteWithID struct {
	ID      string
	Content string
}

// SearchWord - 指定されたユーザーのノートから指定された単語を含むノートを検索
func Search(userID string, word string) ([]NoteWithID, error) {
	// ユーザーのすべてのノートを取得
	notes, err := GetUSerAllNotes(userID, []string{"id", "content"})
	if err != nil {
		slog.Error("Error getting user notes")
		return nil, err
	}

	var matchingNotes []NoteWithID

	// 各ノートの内容をチェックし、指定された単語を含むノートIDとコンテンツを収集
	for _, note := range notes {
		content := note["content"]
		if strings.Contains(content, word) {
			matchingNotes = append(matchingNotes, NoteWithID{
				ID:      note["id"],
				Content: content,
			})
		}
	}

	return matchingNotes, nil
}
