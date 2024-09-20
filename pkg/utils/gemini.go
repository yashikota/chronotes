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

	// Gemini API クライアントを作成します。
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_TOKEN")))
	if err != nil {
		return nil, fmt.Errorf("error creating Gemini client: %v", err)
	}
	defer client.Close()

	// 使用するモデルを指定します。
	model := client.GenerativeModel("gemini-1.5-flash")

	// テキストを結合します。
	combinedText := strings.Join(texts, "\n\n") // 各テキストを2つの改行で区切る

	// 要約リクエストを作成し、API を呼び出します。
	resp, err := model.GenerateContent(ctx, genai.Text(fmt.Sprintf("次の文章から日記を書いて: %s", combinedText)))
	if err != nil {
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
