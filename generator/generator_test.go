package generator

import (
	"generator/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGenerator(t *testing.T) {
	generator := Generator{
		DBConfig: DBConfig{
			DriverName: "mysql",
			Host:       "localhost",
			Port:       3306,
			Username:   "root",
			Password:   "root",
			DBName:     "guest.org.cn",
			Extra:      "charset=utf8",
		},
	}
	generator.Generate()
}

func TestPascalUtil(t *testing.T) {
	t.Log(utils.CamelCaseUtil("sys_blog_Base", "_"))
	t.Log(utils.CamelCaseUtil("blog", "_"))
	t.Log(utils.PascalUtil("sys_blog_Base", "_"))
	t.Log(utils.PascalUtil("base", "_"))
}

func TestCreateDirectory(t *testing.T) {
	err := utils.CreateDirectory("./test/dd")
	if err != nil {
		t.Log(err)
	}
}

func TestCreateFile(t *testing.T) {
	err := utils.CreateFile("./test/test/struct.html")
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
