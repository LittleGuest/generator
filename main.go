package main

import (
	"fmt"
	"generator/middleware"
	"generator/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("run at :65535")

	// 初始化路由
	router := mux.NewRouter()
	middle := middleware.NewMiddleware()
	// 启用中间件
	router.Use(middle.CorsHandler, middle.RequestTimeHandler)

	// 路由
	router.HandleFunc("/api/v1/db/tables", service.ListTables).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/generate", service.Generate).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/download", service.Download).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/remove", service.Remove).Methods(http.MethodPut)

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
