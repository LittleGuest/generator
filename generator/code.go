package generator

import (
	"generator/pool"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type CodeDB struct {
	// ID
	Id int64 `json:"id"`
	// 数据库类型
	Driver string `json:"driver"`
	// 数据库主机
	Host string `json:"host"`
	// 数据库端口
	Port int64 `json:"port"`
	// 数据库名称
	DBName string `json:"db_name"`
	// 用户名
	Username string `json:"username"`
	// 密码
	Password string `json:"password"`
	// 其他参数
	Extra string `json:"extra"`
}

func (t *CodeDB) save() {

}

// 获取配置的数据库列表
func (t CodeDB) Get() CodeDB {
	stmt, err := pool.GetPool().Prepare("SELECT * FROM code_db")
	if err != nil {
		log.Panicln(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Panicln(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&t.Id, &t.Driver, &t.Host, &t.Port, &t.DBName, &t.Username, &t.Password, &t.Extra)
		if err != nil {
			log.Println(err)
			continue
		}
	}
	return t
}

// 获取配置的数据库列表
func (t CodeDB) List() []CodeDB {
	stmt, err := pool.GetPool().Prepare("SELECT * FROM code_db")
	if err != nil {
		log.Panicln(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Panicln(err)
	}
	defer rows.Close()
	codes := make([]CodeDB, 0)
	for rows.Next() {
		err = rows.Scan(&t.Id, &t.Driver, &t.Host, &t.Port, &t.DBName, &t.Username, &t.Password, &t.Extra)
		if err != nil {
			log.Println(err)
			continue
		}
		codes = append(codes, t)
	}
	return codes
}
