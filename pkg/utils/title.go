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

// MakeTitle は、与えられたテキストからタイトルを生成し、string 型で返します。
func MakeTitle(texts []string) (string, error) {
	ctx := context.Background()
	token := os.Getenv("GEMINI_TOKEN")
	if token == "" {
		log.Printf("MakeTitle : GEMINI_TOKEN が環境変数に設定されていません")
		return "", nil
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(token))
	if err != nil {
		log.Printf("MakeTitle : error creating Gemini client: %v\n", err)
		return "", nil
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")

	resp, err := model.GenerateContent(ctx, genai.Text(fmt.Sprintf("次の文章からシンプルなタイトルを１つ考えて、タイトルのみを出力して 大体13字程度に:%s", texts)))
	if err != nil {
		fmt.Printf("MakeTitle : Error generating content: %v\n", err)
		return "", nil
	}

	var titleParts []string
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				if textPart, ok := part.(genai.Text); ok {
					titleParts = append(titleParts, string(textPart))
				}
			}
		}
	}
	title := strings.Join(titleParts, "\n")

	if title == "" {
		log.Printf("MakeTitle : title is empty")
		return "", nil
	}

	return title, nil
}
