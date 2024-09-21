package utils

import (
	"time"
)

func GetDay() string {
	// 現在の日付を取得
	return time.Now().Format(time.RFC3339)
}
