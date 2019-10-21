package service

import (
	"encoding/json"
	"generator/generator"
	"generator/response"
	"io/ioutil"
	"log"
	"net/http"
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
	Single int64 `json:"single"`
	// 指定表
	TableNames string `json:"table_names"`
}

// 获取指定数据库的表
func ListTables(w http.ResponseWriter, r *http.Request) {
	codeDB := CodeDB{}
	bytes, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(bytes, &codeDB)
	log.Println(codeDB)

	// 获取配置的数据库信息
	g := generator.Generator{
		DBConfig: generator.DBConfig{
			DriverName: codeDB.Driver,
			Host:       codeDB.Host,
			Port:       codeDB.Port,
			Username:   codeDB.Username,
			Password:   codeDB.Password,
			DBName:     codeDB.DBName,
			// TODO Extra:      codeDB.Extra,
		},
	}
	response.Success(w, g.ListTable(""))
}

func Generate(w http.ResponseWriter, r *http.Request) {
	codeDB := CodeDB{}
	bytes, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(bytes, &codeDB)

	//single, err := strconv.Atoi(r.URL.Query().Get("single"))
	//if err != nil {
	//	log.Println(err)
	//	response.Error(w, 1, err.Error())
	//	return
	//}
	//tableNames := r.URL.Query().Get("table_names")

	// TODO 文件输出到浏览器，打包下载

	g := generator.Generator{
		DBConfig: generator.DBConfig{
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
	case 0:
		// 多表
		g.MultiGenerate(codeDB.TableNames)
		response.Success(w, "成功")
	case 1:
		// 单表
		if codeDB.TableNames == "" {
			response.Error(w, 1, "table_name必填")
			return
		}

		g.SingleGenerate(codeDB.TableNames)
		response.Success(w, codeDB)
	default:
		// 无
		response.Error(w, 1, "无此类型")
	}
}
