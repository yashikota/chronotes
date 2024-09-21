package utils

import (
	"fmt"
)

func Summary(message []string, token string) ([]string, error) {
	if len(message) < 2 {
		return message, nil
	}

	// 最初の2つのメッセージを要約
	result, err := SummarizeText(message[:2])
	if err != nil {
		return nil, fmt.Errorf("最初の要約でエラーが発生しました: %v", err)
	}

	// 3番目以降のメッセージに対して繰り返し要約処理を行う
	for i := 2; i < len(message); i++ {
		// 現在の要約結果と次のメッセージを要約
		newResult, err := SummarizeText([]string{result[0], message[i]})
		if err != nil {
			return nil, fmt.Errorf("要約中にエラーが発生しました: %v", err)
		}
		result = newResult // 新しい要約結果を更新
	}

	// 最終結果を返す
	return result, nil
}
