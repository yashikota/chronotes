package utils

import (
	"time"
)

func GetDay() string {
	// 現在の日付を取得
	return time.Now().Format(time.RFC3339)
}

func GetDateOnly() string {
	return time.Now().Format(time.DateOnly)
}


// IsDateBefore は date1 の日付が date2 の日付より前かどうかを判定します。
//日付が同じであれば、false.
func IsDateBefore(date1, date2 time.Time) bool{
	d1 := date1.Format(time.DateOnly)
	d2 := date2.Format(time.DateOnly)
	return d1 < d2
}


// IsDateAfter は date1 の日付が date2 の日付より後かどうかを判定します。
//日付が同じであれば、false.
func IsDateAfter(date1, date2 time.Time) bool{
	d1 := date1.Format(time.DateOnly)
	d2 := date2.Format(time.DateOnly)
	return d1 > d2
}