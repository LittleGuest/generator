package pool

import (
	"database/sql"
	"fmt"
	"generator/config"
	"log"
)

var pool *sql.DB

func GetPool() *sql.DB {
	if pool != nil {
		return pool
	}
	dataBase := config.GetDataBase()
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dataBase.Username, dataBase.Password, dataBase.Host, dataBase.Port, dataBase.DBName)
	db, err := sql.Open(dataBase.Driver, dataSource)
	if err != nil {
		log.Fatalf("获取数据库连接失败：%v", err)
	}
	return db
}
