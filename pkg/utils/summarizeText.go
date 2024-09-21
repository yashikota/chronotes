package utils

import (
	"context"
	"fmt"
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
		return nil, fmt.Errorf("error creating Gemini client: %v", err)
	}
	defer client.Close()

	// 使用するモデルを指定します。
	model := client.GenerativeModel("gemini-1.5-flash")

	// テキストを結合します。
	combinedText := strings.Join(texts, "\n\n") // 各テキストを2つの改行で区切る
	if combinedText == "" {
		summary := "進捗なし"
		return []string{summary}, nil
	}

	resp, err := model.GenerateContent(ctx, genai.Text(fmt.Sprintf("次の文章から要約を書いて 要約の量が200字数を超えたら重要な部分以外省いて 要約を作成することができない場合は進捗なしと出力:%s", combinedText)))
	if err != nil {
		fmt.Printf("Error generating content: %v\n", err) // エラーメッセージの出力
		return nil, fmt.Errorf("error generating content for text: %v", err)
	}

	summary := extractSummary(resp)

	if summary == "" {
		return nil, fmt.Errorf("summary is empty")
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
