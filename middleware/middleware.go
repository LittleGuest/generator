package middleware

import (
	"log"
	"net/http"
	"time"
)

type Middleware struct{}

func NewMiddleware() Middleware {
	return Middleware{}
}

// 记录请求时间
func (m Middleware) RequestTimeHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler.ServeHTTP(w, r)
		end := time.Now()
		log.Printf("%v\t%v, 耗时：%v", r.Method, r.RequestURI, end.Sub(start))
	})
}

// 设置跨域请求
func (m Middleware) CorsHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 允许所有的请求
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// 设置允许的Header类型
		w.Header().Set("Access-Control-Allow-Headers", "*")
		// options 请求直接返回
		if r.Method == http.MethodOptions {
			return
		}
		handler.ServeHTTP(w, r)
	})
}
