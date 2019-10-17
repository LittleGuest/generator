package main

import (
	"flag"
	"fmt"
	"generator/config"
	"generator/generator"
	"generator/middleware"
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

	router := mux.NewRouter()
	middle := middleware.NewMiddleware()
	router.Use(middle.CorsHandler, middle.RequestTimeHandler)

	router.HandleFunc("/api/v1/login", generator.Login).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/users", generator.GetUserInfo).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/db", generator.ListDB).Methods(http.MethodGet)
	router.PathPrefix("").Handler(http.StripPrefix("", http.FileServer(http.Dir("views"))))

	appConfig := config.NewAppConfig()
	host := appConfig.Server.Host
	port := appConfig.Server.Port
	log.Printf("run at %s:%d", host, port)

	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", host, port),
		Handler:      router,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}
	log.Fatalln(server.ListenAndServe())
}
