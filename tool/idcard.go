package tool

import (
	"errors"
	"time"
)

// GetAgeByIdcard 根据身份证获取年龄
func GetAgeByIdcard(idcard string) int {
	if len(idcard) == 0 {
		return 0
	}
	b, err := GetBirthByIdcard(idcard)
	if err != nil {
		return 0
	}
	return time.Now().Year() - b.Year()
}

// GetBirthByIdcard 根据身份证获取生日
func GetBirthByIdcard(idcard string) (time.Time, error) {
	if len(idcard) == 0 {
		return time.Time{}, errors.New("身份证为空")
	}
	rs := []rune(idcard)
	return time.Parse(DateLayout, string(rs[6:14]))
}
