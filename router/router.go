package router

import (
	"generator/handler"
	"generator/pages"
	"net/http"

	_ "net/http/pprof"

	assetfs "github.com/elazarl/go-bindata-assetfs"

	"github.com/gorilla/mux"
)

func init() {
	rr := mux.NewRouter()
	rr.Use(CorsHandler, RequestTimeHandler)

	v1 := rr.PathPrefix("/api/v1").Subrouter()
	v1.HandleFunc("/db/tables", handler.ListTables).Methods(http.MethodPost)
	v1.HandleFunc("/create", handler.Create).Methods(http.MethodGet)
	v1.HandleFunc("/test", handler.TestConnect).Methods(http.MethodPost)
	// v1.HandleFunc("/temp", handler.ReadTemp).Methods(http.MethodGet)
	// v1.HandleFunc("/temp", handler.SaveTemp).Methods(http.MethodPost)

	// 静态文件服务
	// rr.PathPrefix("").Handler(http.StripPrefix("", http.FileServer(http.Dir("views"))))
	fs := assetfs.AssetFS{
		Asset:     pages.Asset,
		AssetDir:  pages.AssetDir,
		AssetInfo: pages.AssetInfo,
		Prefix:    "views",
	}
	rr.PathPrefix("/").Handler(http.FileServer(&fs))

	http.Handle("/", rr)
}
