package tool

import (
	"testing"
	"time"
)

func TestFormatDate(t *testing.T) {
	t.Log(FormatDate(time.Now()))
}

func TestFormatDateTime(t *testing.T) {
	t.Log(FormatDateTime(time.Now()))
}

func TestStringToTime(t *testing.T) {
	timeStr := "2006-01-02 15:04:05"
	stt, _ := StringToTime(timeStr)
	t.Log(stt)
}
