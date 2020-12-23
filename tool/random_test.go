package tool

import (
	"testing"
)

func TestRandomNumbersToString(t *testing.T) {
	t.Logf("随机：%s", RandomNumbersToString(-1))
	t.Logf("随机：%s", RandomNumbersToString(0))
	t.Logf("随机：%s", RandomNumbersToString(6))
}
