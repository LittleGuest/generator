package glog

import (
	"testing"
	"time"
)

func TestInfo(t *testing.T) {
	Info("this is log info test")
}

func TestError(t *testing.T) {
	Error("this is log error test")
}

func TestWriteToFile(t *testing.T) {
	for i := 0; i < 10; i++ {
		go WriteToFile("this is log file test : ", i)
	}
	time.Sleep(time.Second)
}
