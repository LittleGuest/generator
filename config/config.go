// TODO 解析json文件
package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
)

// 数据库信息
type dataBase struct {
	Driver   string `json:"driver"`
	Url      string `json:"url"`
	Port     int64  `json:"port"`
	DbName   string `json:"db_name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Ext      struct {
		CharacterEncoding string `json:"character_encoding"`
		AutoReconnect     bool   `json:"auto_reconnect"`
		UseSSL            bool   `json:"use_ssl"`
		ServerTimezone    string `json:"server_timezone"`
	} `json:"ext"`
}

// 应用信息
type app struct {
	Name string `json:"name"`
}

// 服务信息
type server struct {
	Host         string `json:"host"`
	Port         int64  `json:"port"`
	ReadTimeout  time.Duration  `json:"read_timeout"`
	WriteTimeout time.Duration  `json:"write_timeout"`
}

// App配置
type appConfig struct {
	App      app      `json:"app"`
	Server   server   `json:"server"`
	DataBase dataBase `json:"data_base"`
}

var config *appConfig

func NewAppConfig() *appConfig {
	if config != nil {
		return config
	}
	// 解析json文件
	bytes, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatalf("解析配置文件 config.json 失败：%s", err)
	}
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		log.Fatalf("解析配置文件 config.json 失败：%s", err)
	}
	return config
}

func GetAppConfig() *appConfig {
	return config
}

func GetApp() *app {
	return &config.App
}

func GetServer() *server {
	return &config.Server
}

func GetDataBase() *dataBase {
	return &config.DataBase
}
