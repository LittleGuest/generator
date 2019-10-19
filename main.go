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
	// TODO 调用浏览器打开页面
	//if open {
	//	cmd := exec.Command("cmd", "/C", "start http://localhost:65535")
	//	log.Println(cmd.Run())
	//}

	// 初始化路由
	router := mux.NewRouter()
	middle := middleware.NewMiddleware()
	// 启用中间件
	router.Use(middle.CorsHandler, middle.RequestTimeHandler)

	// 路由
	router.HandleFunc("/api/v1/login", service.Login).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/users", service.GetUserInfo).Methods(http.MethodGet)

	router.HandleFunc("/api/v1/db", service.GetCodeDB).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/db", service.SaveCodeDB).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/db/list", service.ListCodeDB).Methods(http.MethodGet)

	router.HandleFunc("/api/v1/generator/single", service.SingleGenerate).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/generator/multi", service.MultiGenerate).Methods(http.MethodGet)
	// 静态文件服务
	router.PathPrefix("").Handler(http.StripPrefix("", http.FileServer(http.Dir("views"))))

	s := config.GetServer()
	log.Printf("run at %s:%d", s.Host, s.Port)

	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", s.Host, s.Port),
		Handler:      router,
		ReadTimeout:  time.Second * s.ReadTimeout,
		WriteTimeout: time.Second * s.ReadTimeout,
	}
	log.Fatalln(server.ListenAndServe())
}
