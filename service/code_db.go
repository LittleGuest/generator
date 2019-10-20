package service

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

// 保存配置的数据库信息
func (t CodeDB) save() (id int64, affected int64) {
	stmt, err := pool.GetPool().Prepare("INSERT INTO code_db(driver,`host`,`port`,db_name,username,`password`,extra) VALUES(?,?,?,?,?,?,?)")
	if err != nil {
		log.Printf("保存配置的数据库信息失败==>%v", err)
		return
	}
	defer stmt.Close()
	if stmt == nil {
		log.Printf("保存配置的数据库信息失败==>%v", err)
		return
	}
	result, err := stmt.Exec(t.Driver, t.Host, t.Port, t.DBName, t.Username, t.Password, t.Extra)
	if err != nil {
		log.Printf("保存配置的数据库信息失败==>%v", err)
		return
	}
	affected, err = result.RowsAffected()
	id, err = result.LastInsertId()
	return
}

// 更新配置的数据库信息
func (t CodeDB) update() (affected int64) {
	updateSql := "UPDATE code_db SET id = ?"
	if t.Driver != "" {
		updateSql += ",driver = ?"
	}
	if t.Host != "" {
		updateSql += ",`host` = ?"
	}
	if t.Port >= 0 || t.Port <= 65535 {
		updateSql += ",`port` = ?"
	}
	if t.DBName != "" {
		updateSql += ",db_name = ?"
	}
	if t.Username != "" {
		updateSql += ",username = ?"
	}
	if t.Password != "" {
		updateSql += ",`password` = ?"
	}
	if t.Extra != "" {
		updateSql += ",extra = ?"
	}
	updateSql += " WHERE id = ?"

	stmt, err := pool.GetPool().Prepare(updateSql)
	if err != nil {
		log.Printf("更新配置的数据库信息失败==>%v", err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(t.Id, t.Driver, t.Host, t.Port, t.DBName, t.Username, t.Password, t.Extra, t.Id)
	if err != nil {
		log.Printf("更新配置的数据库信息失败==>%v", err)
		return
	}
	affected, err = result.RowsAffected()
	return
}

// 获取配置的数据库列表
func (t CodeDB) Get() CodeDB {
	stmt, err := pool.GetPool().Prepare("SELECT * FROM code_db")
	if err != nil {
		log.Printf("获取配置的数据库列表失败==>%v", err)
		return CodeDB{}
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Printf("获取配置的数据库列表失败==>%v", err)
		return CodeDB{}
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
		log.Printf("获取配置的数据库列表失败==>%v", err)
		return nil
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Printf("获取配置的数据库列表失败==>%v", err)
		return nil
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
