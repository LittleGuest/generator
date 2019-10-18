package generator

import (
	"generator/config"
	"generator/response"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	response.Success(w, CodeDB{}.List())
}

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	response.Success(w, config.GetAppConfig())
}

func SingleGenerate(w http.ResponseWriter, r *http.Request) {
	codeDB := CodeDB{}.Get()
	g := Generator{DBConfig{
		DriverName: codeDB.Driver,
		Host:       codeDB.Host,
		Port:       codeDB.Port,
		Username:   codeDB.Username,
		Password:   codeDB.Password,
		DBName:     codeDB.DBName,
		Extra:      codeDB.Extra,
	}}
	g.SingleGenerate(r.URL.Query().Get("table_name"))
	response.Success(w, codeDB)
}

func ListDB(w http.ResponseWriter, r *http.Request) {
	page := response.PageInfo{
		Curr:  1,
		Size:  20,
		Total: 10,
		Data:  CodeDB{}.List(),
	}
	response.Page(w, page)
}
