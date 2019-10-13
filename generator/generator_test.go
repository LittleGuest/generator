package generator

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGenerator(t *testing.T) {
	generator := Generator{
		DBConfig: &DBConfig{
			DriverName: "mysql",
			Host:       "localhost",
			Port:       3306,
			Username:   "root",
			Password:   "root",
			DBName:     "guest.org.cn",
			OtherParam: "charset=utf8",
		},
	}
	generator.Generate()
}

func TestPascalUtil(t *testing.T) {
	t.Log(CamelCaseUtil("sys_blog_Base", "_"))
	t.Log(CamelCaseUtil("blog", "_"))
	t.Log(PascalUtil("sys_blog_Base", "_"))
	t.Log(PascalUtil("base", "_"))
}

func TestCreateDirectory(t *testing.T) {
	err := CreateDirectory("./test/dd")
	if err != nil {
		t.Log(err)
	}
}

func TestCreateFile(t *testing.T) {
	err := CreateFile("./test/test/struct.html")
	if err != nil {
		t.Log(err)
	}
}

func TestConvey(t *testing.T) {
	_, _ = convey.Println("hello")
}

func TestConvert(t *testing.T) {
	for k, v := range conv {
		t.Log(k, v)
	}
}
