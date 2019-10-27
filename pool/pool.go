package pool

import (
	"database/sql"
	"fmt"
	"log"
)

var (
	pool     *sql.DB
	driver   = "mysql"
	host     = "localhost"
	port     = 3306
	dbName   = "blog"
	username = "root"
	password = "root"
	extras   = "?charset=UTF8&parseTime=true"
)

func GetPool() *sql.DB {
	if pool != nil {
		return pool
	}
	db, err := sql.Open(driver, fmt.Sprintf("%s:%s@(%s:%d)/%s%s", username, password, host, port, dbName, extras))
	if err != nil {
		log.Panicf("获取数据库连接失败：%v", err)
	}
	return db
}
