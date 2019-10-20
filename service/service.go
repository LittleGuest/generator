package service

import (
	"encoding/json"
	"generator/config"
	"generator/generator"
	"generator/response"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func Login(w http.ResponseWriter, r *http.Request) {
	response.Success(w, CodeDB{}.List())
}

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	response.Success(w, config.GetAppConfig())
}

func SaveCodeDB(w http.ResponseWriter, r *http.Request) {
	codeDB := CodeDB{}
	bytes, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(bytes, &codeDB)
	codeDB.update()
	response.Success(w, codeDB)
}

func GetCodeDB(w http.ResponseWriter, r *http.Request) {
	response.Success(w, CodeDB{}.Get())
}

func ListCodeDB(w http.ResponseWriter, r *http.Request) {
	page := response.PageInfo{
		Curr:  1,
		Size:  20,
		Total: 10,
		Data:  CodeDB{}.List(),
	}
	response.Page(w, page)
}

func ListTables(w http.ResponseWriter, r *http.Request) {
	// 获取配置的数据库信息
	codeDB := CodeDB{}.Get()
	g := generator.Generator{
		DBConfig: generator.DBConfig{
			DriverName: codeDB.Driver,
			Host:       codeDB.Host,
			Port:       codeDB.Port,
			Username:   codeDB.Username,
			Password:   codeDB.Password,
			DBName:     codeDB.DBName,
			Extra:      codeDB.Extra,
		},
	}
	response.Success(w, g.ListTable(""))
}

func Generate(w http.ResponseWriter, r *http.Request) {
	single, err := strconv.Atoi(r.URL.Query().Get("single"))
	if err != nil {
		log.Println(err)
		response.Fatal(w, err.Error())
		return
	}
	tableNames := r.URL.Query().Get("table_names")

	// TODO 文件输出到浏览器，打包下载

	// 获取配置的数据库信息
	codeDB := CodeDB{}.Get()
	g := generator.Generator{
		DBConfig: generator.DBConfig{
			DriverName: codeDB.Driver,
			Host:       codeDB.Host,
			Port:       codeDB.Port,
			Username:   codeDB.Username,
			Password:   codeDB.Password,
			DBName:     codeDB.DBName,
			Extra:      codeDB.Extra,
		},
	}
	switch single {
	case 0:
		// 多表
		g.MultiGenerate(tableNames)
		response.Success(w, "成功")
	case 1:
		// 单表
		if tableNames == "" {
			response.Fatal(w, "table_name必填")
			return
		}

		g.SingleGenerate(tableNames)
		response.Success(w, codeDB)
	default:
		// 无
		response.Fatal(w, "无此类型")
	}
}
