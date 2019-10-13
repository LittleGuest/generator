package common

import (
	"log"
	"net/http"
	"time"
)

type Middleware struct{}

func NewMiddleware() Middleware {
	return Middleware{}
}

// 设置跨域请求
func (m Middleware) CorsHandler(handler http.Handler) http.Handler {
	log.Println("CorsHandler")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		//if r.Method == http.MethodOptions {
		//	return
		//}
		handler.ServeHTTP(w, r)
		log.Println("设置了跨域请求")
	})
}

// 记录请求时间
func (m Middleware) RequestTimeHandler(handler http.Handler) http.Handler {
	log.Println("RequestTimeHandler")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler.ServeHTTP(w, r)
		end := time.Now()
		log.Printf("%s\t%s, 耗时：%v", r.Method, r.URL, end.Sub(start))
	})
}
