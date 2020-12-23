package repo

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// 全局变量
var db *sql.DB

// DB 数据库连接信息
type ConnInfo struct {
	Driver   string            `json:"driver"`   // 数据库驱动
	Username string            `json:"username"` // 账号
	Password string            `json:"password"` // 密码
	Host     string            `json:"host"`     // 数据库地址
	Port     string            `json:"port"`     // 端口
	DBName   string            `json:"db_name"`  // 数据库名称
	Extras   map[string]string `json:"extras"`   // 额外参数
}

// ConnectDB 连接数据库
func ConnectDB(info ConnInfo) error {
	var err error

	driverName := info.Driver
	extras := "?"
	if info.Extras != nil {
		for k, v := range info.Extras {
			extras += k + "=" + v
		}
	}

	dataSourceName := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v%v", info.Username, info.Password, info.Host, info.Port, info.DBName, extras)

	db, err = sql.Open(driverName, dataSourceName)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	return nil
}

// GetDB return sql.DB's pointer
func GetDB() *sql.DB {
	return db
}
