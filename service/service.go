package service

import (
	"encoding/json"
	"generator/config"
	"generator/generator"
	"generator/response"
	"io/ioutil"
	"net/http"
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

func SingleGenerate(w http.ResponseWriter, r *http.Request) {
	tableName := r.URL.Query().Get("table_name")
	if tableName == "" {
		response.Fatal(w, "table_name必填")
		return
	}
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
	g.SingleGenerate(tableName)
	response.Success(w, codeDB)
}

func MultiGenerate(w http.ResponseWriter, r *http.Request) {
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
	g.MultiGenerate()
	response.Success(w, codeDB)
}
