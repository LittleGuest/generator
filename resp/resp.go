package resp

import (
	"encoding/json"
	"io"
)

// Resp 响应
type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// PageInfo 分页信息
type PageInfo struct {
	Offset int64       `json:"offset"`
	Limit  int64       `json:"limit"`
	Total  int64       `json:"total"`
	Data   interface{} `json:"data"`
}

func encode(resp Resp) string {
	res, err := json.Marshal(resp)
	if err != nil {
		return ""
	}
	return string(res)
}

// WriteJSON write json with w
func WriteJSON(w io.Writer, resp Resp) {
	io.WriteString(w, encode(resp))
}
