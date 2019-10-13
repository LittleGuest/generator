package generator

import (
	"generator/common"
	"log"
	"net/http"
)

func ListDB(w http.ResponseWriter, r *http.Request) {
	page := common.Page{
		Curr:  1,
		Size:  20,
		Total: 10,
		Data:  CodeDB{}.List(),
	}
	common.RespPage(w, page)
}

func SaveDB(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Form)
}
