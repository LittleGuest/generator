package main

import (
	"generator/service"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// RequestTimeHandler 记录请求时间
func RequestTimeHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler.ServeHTTP(w, r)
		end := time.Now()
		log.Printf("%v\t%v, 耗时：%v", r.Method, r.RequestURI, end.Sub(start))
	})
}

// CorsHandler 设置跨域请求
func CorsHandler(handler http.Handler) http.Handler {
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

func main() {
	// 初始化路由
	router := mux.NewRouter()
	// 启用中间件
	router.Use(CorsHandler, RequestTimeHandler)

	// 路由
	router.HandleFunc("/api/v1/db/tables", service.ListTables).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/create", service.Create).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/temp", service.ReadTemp).Methods(http.MethodGet)

	// 静态文件服务
	router.PathPrefix("").Handler(http.StripPrefix("", http.FileServer(http.Dir("views"))))

	log.Println("run at :65535")

	server := &http.Server{
		Addr:         ":65535",
		Handler:      router,
		ReadTimeout:  time.Second * 60,
		WriteTimeout: time.Second * 60,
	}
	log.Fatalln(server.ListenAndServe())
}
