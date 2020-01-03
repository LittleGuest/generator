package main

import (
	"fmt"
	"generator/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

// 记录请求时间
func RequestTimeHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler.ServeHTTP(w, r)
		end := time.Now()
		log.Printf("%v\t%v, 耗时：%v", r.Method, r.RequestURI, end.Sub(start))
	})
}

// 设置跨域请求
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
	log.Println("run at :65535")

	// 初始化路由
	router := mux.NewRouter()
	// 启用中间件
	router.Use(CorsHandler, RequestTimeHandler)

	// 路由
	router.HandleFunc("/api/v1/db/tables", service.ListTables).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/generate", service.Generate).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/download", service.Download).Methods(http.MethodGet)

	// 静态文件服务
	router.PathPrefix("").Handler(http.StripPrefix("", http.FileServer(http.Dir("views"))))

	server := &http.Server{
		Addr:         fmt.Sprintf(":65535"),
		Handler:      router,
		ReadTimeout:  time.Second * 60,
		WriteTimeout: time.Second * 60,
	}
	log.Fatalln(server.ListenAndServe())
}
