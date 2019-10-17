package generator

import (
	"generator/config"
	"generator/response"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	response.RespSuccess(w, CodeDB{}.List())
}

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	response.RespSuccess(w, config.GetAppConfig())
}

func ListDB(w http.ResponseWriter, r *http.Request) {
	page := response.Page{
		Curr:  1,
		Size:  20,
		Total: 10,
		Data:  CodeDB{}.List(),
	}
	response.RespPage(w, page)
}
