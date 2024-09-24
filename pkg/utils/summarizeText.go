package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// SummarizeText は、複数のテキストを要約し、その要約を []string 型で返します。
func SummarizeText(texts []string) ([]string, error) {
	ctx := context.Background()
	token := os.Getenv("GEMINI_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("GEMINI_TOKEN が環境変数に設定されていません")
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(token))
	if err != nil {
		log.Printf("SummarizeText : error creating Gemini client: %v\n", err) // エラーメッセージの出力
		summary := "進捗なし"
		return []string{summary}, nil
	}
	defer client.Close()

	// 使用するモデルを指定します。
	model := client.GenerativeModel("gemini-1.5-flash")

	// テキストを結合します。
	combinedText := strings.Join(texts, "\n\n") // 各テキストを2つの改行で区切る
	if combinedText == "" {
		return []string{"進捗なし"}, nil
	}

	resp, err := model.GenerateContent(ctx, genai.Text(fmt.Sprintf("次の文章から要約を書いて 要約の量が200字数を超えたら重要な部分以外省いて また「プルリクエスト」文言が含まれている場合は省いて:%s", combinedText)))
	if err != nil {
		fmt.Printf("SummarizeText : Error generating content: %v\n", err) // エラーメッセージの出力
		return []string{"進捗なし"}, nil
	}

	summary := extractSummary(resp)

	if summary == "" {
		log.Printf("SummarizeText : summary is empty") // エラーメッセージの出力
		return []string{"進捗なし"}, nil
	}

	return []string{summary}, nil
}

func extractSummary(resp *genai.GenerateContentResponse) string {
	var summaryParts []string
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				if textPart, ok := part.(genai.Text); ok {
					summaryParts = append(summaryParts, string(textPart)) // genai.Textを文字列に変換
				}
			}
		}
	}
	return strings.Join(summaryParts, "\n")
}
