// 字符工具包
package util

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

// 小驼峰
func CamelCaseUtil(source string, sep string) (result string) {
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
func PascalUtil(source string, sep string) (result string) {
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

// 创建目录
func CreateDirectory(path string) (err error) {
	err = os.MkdirAll(path, 0666)
	if err != nil {
		log.Printf("创建目录失败：目录：%s, %v", path, err)
	}
	return
}

// 创建文件
func CreateFile(name string) (file *os.File, err error) {
	// 先创建目录
	p := filepath.Dir(name)
	err = CreateDirectory(p)
	if err != nil {
		return
	}
	// 再创建文件
	file, err = os.OpenFile(name, os.O_CREATE, 0666)
	if err != nil {
		log.Printf("创建文件失败：文件名为：%s, %v", name, err)
	}
	return
}

// 判断字符串是否为空
func IsBlank(str string) bool {
	if len(str) <= 0 {
		return true
	}
	if str != "" {
		return false
	}
	return true
}
