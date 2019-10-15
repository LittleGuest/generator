package generator

import (
	"generator/common"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	common.RespSuccess(w, CodeDB{}.List())
}

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	common.RespSuccess(w, CodeDB{}.List())
}

func ListDB(w http.ResponseWriter, r *http.Request) {
	page := common.Page{
		Curr:  1,
		Size:  20,
		Total: 10,
		Data:  CodeDB{}.List(),
	}
	common.RespPage(w, page)
}
