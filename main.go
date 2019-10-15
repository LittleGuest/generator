package main

import (
	"generator/common"
	"generator/generator"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	router := mux.NewRouter()

	middle := common.NewMiddleware()
	router.Use(middle.CorsHandler, middle.RequestTimeHandler)

	router.HandleFunc("/api/v1/login", generator.Login).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/users", generator.GetUserInfo).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/db", generator.ListDB).Methods(http.MethodGet)
	router.PathPrefix("").Handler(http.StripPrefix("", http.FileServer(http.Dir("views"))))

	log.Println("run at localhost:65535")
	server := &http.Server{
		Addr:         ":65535",
		Handler:      router,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}
	log.Println(server.ListenAndServe())
}
