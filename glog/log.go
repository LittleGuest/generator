// 简单日志
package glog

import (
	"bufio"
	"fmt"
	"generator/util"
	"log"
	"os"
	"time"
)

const (
	SuffixLog = ".log"
	Flag      = log.Ldate | log.Lmicroseconds | log.Lshortfile
)

var infoLogger = log.New(os.Stderr, "[INFO]\t", Flag)
var errorLogger = log.New(os.Stderr, "[ERROR]\t", Flag)

// 打印日志信息
func Info(args ...interface{}) {
	_ = infoLogger.Output(2, fmt.Sprintln(args...))
}

// 打印日志信息
func Error(args ...interface{}) {
	_ = errorLogger.Output(2, fmt.Sprintln(args...))
}

// TODO fix 向文件中写入日志
func WriteToFile(args ...interface{}) {
	nowDate := util.FormatDate(time.Now())
	logPath := "./sys_log/"
	file, err := util.CreateFile(logPath + nowDate + SuffixLog)
	if err != nil {
		Error(err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	_, _ = writer.WriteString(fmt.Sprintln(args...))
	_ = writer.Flush()
}

// TODO 日志文件压缩：定时将日志文件压缩
