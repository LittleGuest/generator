package handler

import (
	"encoding/json"
	"fmt"
	"generator/generate"
	"generator/repo"
	"generator/resp"
	"generator/tool"
	"io/ioutil"
	"net/http"
	"os"
)

type CodeDB struct {
	repo.ConnInfo
	TableNames []string `json:"table_names"` // 指定表
}

// ReadTemp 读取模板内容
func ReadTemp(w http.ResponseWriter, r *http.Request) {
	tempName := r.URL.Query().Get("temp_name")

	if tempName == "" {
		resp.WriteJSON(w, resp.Resp{
			Code: 1,
			Msg:  "没有找到对应的模板",
		})
		return
	}

	content, err := ioutil.ReadFile("./generate/templates/" + tempName + ".html")
	if err != nil {
		resp.WriteJSON(w, resp.Resp{
			Code: 1,
			Msg:  "读取模板文件失败",
		})
		return
	}
	resp.WriteJSON(w, resp.Resp{
		Data: string(content),
	})
}

// SaveTemp 修改生成模板文件
func SaveTemp(w http.ResponseWriter, r *http.Request) {
	params, _ := ioutil.ReadAll(r.Body)
	m := make(map[string]string)
	_ = json.Unmarshal(params, &m)
	tempName := m["temp_name"]
	content := m["content"]
	fmt.Println(m)

	if tool.IsBlank(tempName) {
		resp.WriteJSON(w, resp.Resp{
			Code: 1,
			Msg:  "模板文件名称为空",
		})
		return
	}
	if tool.IsBlank(content) {
		resp.WriteJSON(w, resp.Resp{
			Code: 1,
			Msg:  "模板内容为空",
		})
		return
	}

	path := "./generate/templates/" + tempName + ".html"
	err := os.Remove(path)
	if err == nil {
		tempFile, err := tool.CreateFile(path)
		if err != nil {
			resp.WriteJSON(w, resp.Resp{
				Code: 1,
				Msg:  "模板文件创建失败",
			})
			return
		}

		defer tempFile.Close()
		_, _ = tempFile.WriteString(content)
	}
	resp.WriteJSON(w, resp.Resp{})
}

// ListTables 获取指定数据库的所有表信息
func ListTables(w http.ResponseWriter, r *http.Request) {
	codeDB := CodeDB{}
	bytes, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(bytes, &codeDB)

	// 连接数据库
	connInfo := repo.ConnInfo{
		Driver:   codeDB.Driver,
		Username: codeDB.Username,
		Password: codeDB.Password,
		Host:     codeDB.Host,
		Port:     codeDB.Port,
		DBName:   codeDB.DBName,
		Extras:   nil,
	}

	// connInfo := repo.ConnInfo{}
	// err := tool.CopyStructProperty(codeDB, connInfo)
	// if err != nil {
	// 	resp.WriteJSON(w, resp.Resp{
	// 		Code: http.StatusInternalServerError,
	// 		Msg:  err.Error(),
	// 	})
	// 	return
	// }

	if connInfo.DBName == "" {
		return
	}

	err := repo.ConnectDB(connInfo)
	if err != nil {
		resp.WriteJSON(w, resp.Resp{
			Code: 1,
			Msg:  fmt.Sprintf("数据库连接失败：%v", err),
		})
		return
	}

	d, err := generate.ListTable(codeDB.DBName, "")
	if err != nil {
		resp.WriteJSON(w, resp.Resp{
			Code: 1,
			Msg:  err.Error(),
		})
		return
	}
	resp.WriteJSON(w, resp.Resp{
		Data: d,
	})
}

// Create 生成代码并打包下载代码
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
	connInfo := repo.ConnInfo{
		Driver:   driver,
		Username: username,
		Password: password,
		Host:     host,
		Port:     port,
		DBName:   dbName,
		Extras:   nil, // TODO 额外参数解析
	}

	if tool.IsBlank(connInfo.DBName) {
		resp.WriteJSON(w, resp.Resp{
			Code: 1,
			Msg:  "数据库为空",
		})
		return
	}

	err := repo.ConnectDB(connInfo)
	if err != nil {
		resp.WriteJSON(w, resp.Resp{
			Code: 1,
			Msg:  fmt.Sprintf("数据库连接失败：%v", err),
		})
		return
	}

	generate.Create(connInfo.Driver, connInfo.DBName, tableNames)

	codeFile, _ := ioutil.ReadFile("code.zip")
	w.Header().Set("Content-Disposition", "attachment;filename=code.zip")
	_, _ = w.Write(codeFile)

	if err := os.Remove("code.zip"); err != nil {
		resp.WriteJSON(w, resp.Resp{
			Code: 1,
			Msg:  "生成失败",
		})
		return
	}
}

// TestConnect test db connect
func TestConnect(w http.ResponseWriter, r *http.Request) {
	codeDB := CodeDB{}
	bytes, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(bytes, &codeDB)

	// 连接数据库
	connInfo := repo.ConnInfo{
		Driver:   codeDB.Driver,
		Username: codeDB.Username,
		Password: codeDB.Password,
		Host:     codeDB.Host,
		Port:     codeDB.Port,
		DBName:   codeDB.DBName,
		Extras:   nil,
	}

	if tool.IsBlank(connInfo.DBName) {
		resp.WriteJSON(w, resp.Resp{
			Code: 1,
			Msg:  "数据库为空",
		})
		return
	}

	err := repo.ConnectDB(connInfo)
	if err != nil {
		resp.WriteJSON(w, resp.Resp{
			Code: 1,
			Msg:  fmt.Sprintf("数据库连接失败：%v", err),
		})
		return
	}

	d, err := generate.ListTable(codeDB.DBName, "")
	if err != nil {
		resp.WriteJSON(w, resp.Resp{
			Code: 1,
			Msg:  err.Error(),
		})
		return
	}
	resp.WriteJSON(w, resp.Resp{
		Data: d,
	})
}
