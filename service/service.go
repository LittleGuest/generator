package service

import (
	"encoding/json"
	"generator/generate"
	"generator/response"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type CodeDB struct {
	// 数据库类型
	Driver string `json:"driver"`
	// 数据库主机
	Host string `json:"host"`
	// 数据库端口
	Port int64 `json:"port"`
	// 数据库名称
	DBName string `json:"db_name"`
	// 用户名
	Username string `json:"username"`
	// 密码
	Password string `json:"password"`
	// 额外参数
	Extras map[string]string `json:"extras"`
	// 是否单表
	Single bool `json:"single"`
	// 指定表
	TableNames []string `json:"table_names"`
}

// 获取指定数据库的表
func ListTables(w http.ResponseWriter, r *http.Request) {
	codeDB := CodeDB{}
	bytes, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(bytes, &codeDB)

	g := generate.Generator{
		DBConfig: generate.DBConfig{
			DriverName: codeDB.Driver,
			Host:       codeDB.Host,
			Port:       codeDB.Port,
			Username:   codeDB.Username,
			Password:   codeDB.Password,
			DBName:     codeDB.DBName,
			// Extra:      codeDB.Extra,
		},
	}
	response.Success(w, g.ListTable(""))
}

func Generate(w http.ResponseWriter, r *http.Request) {
	codeDB := CodeDB{}
	bytes, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(bytes, &codeDB)

	g := generate.Generator{
		DBConfig: generate.DBConfig{
			DriverName: codeDB.Driver,
			Host:       codeDB.Host,
			Port:       codeDB.Port,
			Username:   codeDB.Username,
			Password:   codeDB.Password,
			DBName:     codeDB.DBName,
			//Extra:      codeDB.Extra,
		},
	}
	switch codeDB.Single {
	case false:
		// 多表
		g.MultiGenerate(strings.Join(codeDB.TableNames, ","))
		response.Success(w, "成功")
	case true:
		// 单表
		if len(codeDB.TableNames) <= 0 {
			response.Error(w, 1, "table_name必填")
			return
		}
		g.SingleGenerate(codeDB.TableNames[0])
		response.Success(w, codeDB)
	default:
		// 无
		response.Error(w, 1, "无此类型")
	}
}

func Download(w http.ResponseWriter, r *http.Request) {
	bytes, _ := ioutil.ReadFile("code.zip")
	w.Header().Set("Content-Disposition", "attachment;filename=code.zip")
	_, _ = w.Write(bytes)

	if err := os.Remove("code.zip"); err != nil {
		log.Println(err)
		response.Error(w, 1, "删除失败")
		return
	}
}
