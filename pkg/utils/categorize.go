package utils

import (
	"time"
)

// CategorizeCommitDate は、コミットの日付をカテゴリに分類します。
// 優先度: Today > This Week > This Month > Q1 Q2 Q3 Q4 > Older
func CategorizeCommitDate(date time.Time) string {
	now := time.Now().UTC() // 現在のUTC時刻を取得

	// 今日かどうかを判定
	if date.Year() == now.Year() && date.YearDay() == now.YearDay() {
		return "Today"
	}

	// 今週かどうかを判定
	nowYear, nowWeek := now.ISOWeek()
	commitYear, commitWeek := date.ISOWeek()
	if commitYear == nowYear && commitWeek == nowWeek {
		return "This Week"
	}

	// 今月かどうかを判定
	if date.Month() == now.Month() && date.Year() == now.Year() {
		return "This Month"
	}

	// 今年かどうかを判定し、四半期に分類
	if date.Year() == now.Year() {
		month := int(date.Month())
		quarter := (month-1)/3 + 1
		switch quarter {
		case 1:
			return "Q1 (Jan-Mar)"
		case 2:
			return "Q2 (Apr-Jun)"
		case 3:
			return "Q3 (Jul-Sep)"
		case 4:
			return "Q4 (Oct-Dec)"
		}
		// 上記の四半期判定は必ずどれかにマッチするはずなので、ここには到達しない
	}

	// それ以外の場合は古いメッセージとして分類
	return "Older"
}
