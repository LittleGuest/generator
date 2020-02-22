// 文件工具包
package filetool

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

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

// 获取文件后缀
func GetFileSuffix(fileName string) string {
	lastIndex := strings.LastIndex(fileName, ".")
	return fileName[lastIndex+1:]
}

// IsExist 判断文件或文件夹是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsExist(err) {
		return true
	}
	return false
}
