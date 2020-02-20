package service

import (
	"encoding/json"
	"fmt"
	"generator/database"
	"generator/generate"
	"generator/resp"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type CodeDB struct {
	database.ConnInfo
	TableNames []string `json:"table_names"` // 指定表
}

func ListTables(w http.ResponseWriter, r *http.Request) {
	codeDB := CodeDB{}
	bytes, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(bytes, &codeDB)

	// 连接数据库
	connInfo := database.ConnInfo{
		Driver:   codeDB.Driver,
		Username: codeDB.Username,
		Password: codeDB.Password,
		Host:     codeDB.Host,
		Port:     codeDB.Port,
		DBName:   codeDB.DBName,
		Extras:   nil,
	}

	if connInfo.DBName == "" {
		return
	}

	err := database.ConnectDB(connInfo)
	if err != nil {
		resp.Error(w, 1, fmt.Sprintf("数据库连接失败：%v", err))
		return
	}

	resp.Success(w, generate.ListTable(codeDB.DBName, ""))
}

func Create(w http.ResponseWriter, r *http.Request) {
	driver := r.URL.Query().Get("driver")
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")
	host := r.URL.Query().Get("host")
	port := r.URL.Query().Get("port")
	dbName := r.URL.Query().Get("db_name")
	//extras := r.URL.Query().Get("extras")
	tableNames := r.URL.Query().Get("table_names")

	// 连接数据库
	connInfo := database.ConnInfo{
		Driver:   driver,
		Username: username,
		Password: password,
		Host:     host,
		Port:     port,
		DBName:   dbName,
		Extras:   nil, // TODO 额外参数解析
	}

	if connInfo.DBName == "" {
		return
	}

	err := database.ConnectDB(connInfo)
	if err != nil {
		resp.Error(w, 1, fmt.Sprintf("数据库连接失败：%v", err))
		return
	}

	generate.Create(connInfo.Driver, connInfo.DBName, tableNames)

	codeFile, _ := ioutil.ReadFile("code.zip")
	w.Header().Set("Content-Disposition", "attachment;filename=code.zip")
	_, _ = w.Write(codeFile)

	if err := os.Remove("code.zip"); err != nil {
		log.Println(err)
		resp.Error(w, 1, "生成失败")
		return
	}
}
