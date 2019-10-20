package main

import (
	"archive/zip"
	"bytes"
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

		// 写zip
		buf := new(bytes.Buffer)
		zw := zip.NewWriter(buf)
		var files = []struct {
			Name string
			Body string
		}{
			{"readme.txt", "thi archive contains some text files;"},
			{"gopher.txt", "Gopher name:\nGeorge\nGeoffrey\nGonzo"},
			{"todo.txt", "get animal handing licence.\nWrite more examples"},
		}

		for _, file := range files {
			f, err := zw.Create(file.Name)
			if err != nil {
				log.Println(err)
				return
			}
			_, err = f.Write([]byte(file.Body))
			if err != nil {
				log.Println(err)
				return
			}

		}

		zw.Flush()
		zw.Close()

		//fileInfos, _ := ioutil.ReadDir("./test/")
		//log.Println(fileInfos)
		//for _, v := range fileInfos {
		//	fileHeader, _ := zip.FileInfoHeader(v)
		//	log.Println(fileHeader)
		//}
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
