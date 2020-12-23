package tool

import (
	"testing"
)

func TestCreateDirectory(t *testing.T) {
	//t.Log(CreateDirectory("test"))
	t.Log(CreateDirectory("test/test"))
	//t.Log(CreateDirectory("test/test.log"))
}

func TestCreateFile(t *testing.T) {
	t.Log(CreateFile("test1"))
	t.Log(CreateFile("test.log"))
}

func TestGetFileSuffix(t *testing.T) {
	t.Log(GetFileSuffix("test.log"))
	t.Log(GetFileSuffix("test.log.png"))
}

func TestIsExist(t *testing.T) {
	t.Log(IsExist("test"))
	t.Log(IsExist("test-file"))
	t.Log(IsExist("test/test.log"))
	t.Log(IsExist("test/test.txt"))
}
