package notes

import (
	"context"
	"log/slog"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

// NoteWithID - ノートのIDとコンテンツを保持する構造体
type NoteWithID struct {
	ID      string
	Content string
}

var ElasticTypedClient *elasticsearch.TypedClient

func SearchNote(userID string, word string) ([]NoteWithID, error) {
	// Elasticsearchのコンテキストを作成
	ctx := context.Background()

	// ユーザーのノートを検索するためのクエリを構築
	query := `{
		"query": {
			"bool": {
				"must": [
					{ "match": { "user_id": "` + userID + `" }},
					{ "match": { "content": "` + word + `" }}
				]
			}
		}
	}`

	// Elasticsearchに対して検索リクエストを送信
	req := esapi.SearchRequest{
		Index:          []string{"notes"},
		Body:           strings.NewReader(query),
		TrackTotalHits: true,
	}

	res, err := req.Do(ctx, ElasticTypedClient)
	if err != nil {
		slog.Error("Error searching notes", "error", err)
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		slog.Error("Error response from Elasticsearch", "status", res.Status)
		return nil, nil // エラー処理を追加することを検討してください
	}

	var matchingNotes []NoteWithID = make([]NoteWithID, 0)

	if len(matchingNotes) == 0 {
		slog.Info("No matching notes found", "userID", userID, "word", word)
	} else {
		for _, note := range matchingNotes {
			slog.Info("Matching note found", "userID", userID, "noteID", note.ID)
		}
	}

	return matchingNotes, nil
}
