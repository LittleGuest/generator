// 时间工具包
package tool

import (
	"strings"
	"time"
)

const (
	DateLayout      = "2006-01-02"
	DateTimeLayout  = "2006-01-02 15:04:05"
	DateShortLayout = "20060102"
)

// FormatDate 格式化时间
func FormatDate(t time.Time) string {
	return t.Format(DateShortLayout)
}

// FormatDateTime 格式化时间
func FormatDateTime(t time.Time) string {
	return t.Format(DateTimeLayout)
}

// StringToTime 时间字符串转时间
func StringToTime(timeStr string) (time.Time, error) {
	timeStr = strings.Replace(timeStr, "+", "", -1)
	return time.Parse(DateTimeLayout, timeStr)
}
