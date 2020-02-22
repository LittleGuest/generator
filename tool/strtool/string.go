// 字符工具包
package strtool

import (
	"encoding/json"
	"strconv"
	"strings"
)

// 小驼峰
func ToCamelCase(source string, sep string) (result string) {
	sources := strings.Split(source, sep)
	for k, v := range sources {
		if k == 0 {
			result += v
			continue
		}
		result += FirstLetterToUpper(v)
	}
	return
}

// 大驼峰
func ToPascal(source string, sep string) (result string) {
	sources := strings.Split(source, sep)
	for _, v := range sources {
		result += FirstLetterToUpper(v)
	}
	return
}

// 首字母大写
func FirstLetterToUpper(target string) (result string) {
	chars := []rune(target)
	for key, value := range chars {
		if key == 0 {
			if value >= 97 && value <= 122 {
				value -= 32
			}
		}
		result += string(value)
	}
	return
}

// ToString
func ToString(arg interface{}) string {
	switch arg.(type) {
	case int:
		return strconv.Itoa(arg.(int))
	case int64:
		return strconv.FormatInt(arg.(int64), 10)
	default:
		res, _ := json.Marshal(arg)
		return string(res)
	}
}

// 判断字符串是否为空
func IsBlank(str string) bool {
	if len(str) <= 0 {
		return true
	}
	if str == "" {
		return true
	}
	return false
}

// 判断字符串是否不为空
func IsNotBlank(str string) bool {
	return !IsBlank(str)
}
