package tool

import (
	"strings"
	"time"
)

const (
	dateTimeStr = "2006-01-02 15:04:05"
	dateStr     = "20060102"
)

// 格式化时间
func FormatDate(t time.Time) string {
	return t.Format(dateStr)
}

func FormatDateTime(t time.Time) string {
	return t.Format(dateTimeStr)
}

// 时间字符串转时间
func StringToTime(timeStr string) (time.Time, error) {
	timeStr = strings.Replace(timeStr, "+", "", -1)
	return time.Parse(dateTimeStr, timeStr)
}
