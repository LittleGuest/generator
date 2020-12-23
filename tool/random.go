package tool

import (
	"math/rand"
	"strings"
)

const (
	BaseNumber = "0123456789"
)

// RandomNumbersToString return a random string of l
func RandomNumbersToString(l int) string {
	return randomString(BaseNumber, l)
}

// randomString return a bs based random string of l
func randomString(bs string, l int) string {
	if bs == "" {
		return ""
	}

	if l < 1 {
		l = 1
	}

	bsl := len(bs)
	sb := strings.Builder{}
	for i := 0; i < l; i++ {
		r := rand.Intn(bsl)
		_, _ = sb.WriteString(string(bs[r]))
	}
	return sb.String()
}
