// 时间工具包
package util

import (
	"strings"
	"time"
)

const (
	DateTimeStr = "2006-01-02 15:04:05"
	DateStr     = "20060102"
)

// 格式化时间
func FormatDate(t time.Time) string {
	return t.Format(DateStr)
}

func FormatDateTime(t time.Time) string {
	return t.Format(DateTimeStr)
}

// 时间字符串转时间
func StringToTime(timeStr string) (time.Time, error) {
	timeStr = strings.Replace(timeStr, "+", "", -1)
	return time.Parse(DateTimeStr, timeStr)
}
