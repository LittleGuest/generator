package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// 分页信息结
type Page struct {
	Curr  int         `json:"curr"`
	Size  int         `json:"size"`
	Total int         `json:"total"`
	Data  interface{} `json:"data"`
}

// 返回成功信息
func RespSuccess(w http.ResponseWriter, data interface{}) {
	m := make(map[string]interface{})
	m["code"] = 0
	m["data"] = data
	jsonData, err := json.Marshal(m)
	if err != nil {
		log.Fatalln("json编码错误", err)
	}
	_, _ = w.Write(jsonData)
}

// 返回失败信息
func RespFail(w http.ResponseWriter, msg string) {
	m := make(map[string]interface{})
	m["code"] = 1
	m["msg"] = msg
	jsonData, err := json.Marshal(m)
	if err != nil {
		log.Fatalln("json编码错误", err)
	}
	_, _ = w.Write(jsonData)
}

// 返回错误信息
func RespError(w http.ResponseWriter, code uint, msg string) {
	m := make(map[string]interface{})
	m["code"] = code
	m["msg"] = msg
	jsonData, err := json.Marshal(m)
	if err != nil {
		log.Fatalln("json编码错误", err)
	}
	_, _ = w.Write(jsonData)
}

// 返回分页成功信息
func RespPage(w http.ResponseWriter, page Page) {
	m := make(map[string]interface{})
	m["code"] = 0
	m["data"] = page
	jsonData, err := json.Marshal(m)
	if err != nil {
		log.Fatalln("json编码错误", err)
	}
	_, _ = w.Write(jsonData)
}
