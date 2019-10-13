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
	router.Use(middle.RequestTimeHandler, middle.CorsHandler)
	router.HandleFunc("/api/db/save", generator.SaveDB)
	router.HandleFunc("/api/db", generator.ListDB).Methods(http.MethodGet)
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
