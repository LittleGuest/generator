package main

import (
	"generator/common"
	"generator/generator"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	middle := common.NewMiddleware()
	router.Use(middle.RequestTimeHandler, middle.CorsHandler)
	router.HandleFunc("/api/db/save", generator.SaveDB)
	router.HandleFunc("/api/db", generator.ListDB).Methods(http.MethodGet)
	//router.PathPrefix("/views/").Handler(http.StripPrefix("/views/", http.FileServer(http.Dir("views"))))

	log.Println("run at localhost:65535")
	log.Println(http.ListenAndServe(":65535", router))
}
