package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"generator/config"
	"generator/middleware"
	"generator/service"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
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
	router.HandleFunc("/api/v1/db/tables", service.ListTables).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/generate", service.Generate).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/zip", func(w http.ResponseWriter, r *http.Request) {
		// 读zip
		//rc, err := zip.OpenReader("./test.zip")
		//if err != nil {
		//	log.Println(err)
		//	return
		//}
		//defer rc.Close()
		//log.Println(rc.Comment)
		//log.Println(rc.File)

		log.Println("=======================================")
		zipFile, err := os.Create("test.zip")
		if err != nil {
			log.Panicln(err)
		}
		defer zipFile.Close()
		zw := zip.NewWriter(zipFile)
		defer zw.Close()

		//fileInfos, err := ioutil.ReadDir(`D:\coding\workspaces\gopher-go`)
		//for _, file := range fileInfos {
		//	fw, err := zw.Create(file.Name())
		//	if err != nil {
		//		log.Println(err)
		//		return
		//	}
		//
		//	_, err = fw.Write()
		//	if err != nil {
		//		log.Println(err)
		//		return
		//	}
		//}

		bytes, err := ioutil.ReadFile(`D:\coding\workspaces\generator\main.go`)
		if err != nil {
			log.Panic(err)
		}
		fw, err := zw.Create("main.go")
		if err != nil {
			log.Panic(err)
		}
		_, _ = fw.Write(bytes)
		w.Header().Set("Content-Disposition", "attachment;filename=test.zip")
		//w.Header().Set("Content-Type", "application/octet-stream")
		out, _ := ioutil.ReadFile("test.zip")
		w.Write(out)
	}).Methods(http.MethodGet)

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
