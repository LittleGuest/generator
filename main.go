package main

import (
	"flag"
	"fmt"
	"generator/config"
	"generator/middleware"
	"generator/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os/exec"
	"time"
)

var (
	open bool
)

func init() {
	flag.BoolVar(&open, "o", true, "open with browser")
	flag.Parse()
}

func main() {
	s := config.GetServer()
	log.Printf("run at %s:%d", s.Host, s.Port)

	if open {
		cmd := exec.Command("cmd", "/C", fmt.Sprintf("start http://%s:%d", s.Host, s.Port))
		log.Println(cmd.Run())
	}

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
		Addr:         fmt.Sprintf("%s:%d", s.Host, s.Port),
		Handler:      router,
		ReadTimeout:  time.Second * s.ReadTimeout,
		WriteTimeout: time.Second * s.ReadTimeout,
	}
	log.Fatalln(server.ListenAndServe())
}
