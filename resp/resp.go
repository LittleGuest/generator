// 响应信息封装
package resp

import (
	"encoding/json"
	"log"
	"net/http"
)

// PageInfo 分页信息
type PageInfo struct {
	Curr  int64       `json:"curr"`
	Size  int64       `json:"size"`
	Total int64       `json:"total"`
	Data  interface{} `json:"data"`
}

// Success 返回成功信息
func Success(w http.ResponseWriter, data interface{}) {
	m := make(map[string]interface{})
	m["code"] = 0
	m["data"] = data
	jsonData, err := json.Marshal(m)
	if err != nil {
		log.Panicln("json编码错误", err)
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(jsonData)
}

// Error 返回错误信息
func Error(w http.ResponseWriter, code int64, msg string) {
	m := make(map[string]interface{})
	m["code"] = code
	m["msg"] = msg
	jsonData, err := json.Marshal(m)
	if err != nil {
		log.Panicln("json编码错误", err)
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(jsonData)
}

// Page 返回分页成功信息
func Page(w http.ResponseWriter, page PageInfo) {
	m := make(map[string]interface{})
	m["code"] = 0
	m["data"] = page
	jsonData, err := json.Marshal(m)
	if err != nil {
		log.Panicln("json编码错误", err)
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(jsonData)
}
