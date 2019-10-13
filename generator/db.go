package generator

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"log"
)

var Pool *sql.DB

// 获取数据库连接
func init() {
	if Pool == nil {
		db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/guest.org.cn")
		if err != nil {
			log.Panicln("数据库连接失败。。。", err.Error())
		}
		Pool = db
	}
}

// 获取数据库连接
func (g *Generator) Open() (db *sql.DB) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		g.Username, g.Password, g.Host, g.Port, g.DBName, g.OtherParam)
	db, err := sql.Open(g.DriverName, dataSourceName)
	if err != nil {
		log.Panicln("数据库连接失败", err)
	}
	return
}

// 获取数据库所有表信息
func (g *Generator) ListTables() []Tables {
	tablesSql := "SELECT t.TABLE_NAME,t.TABLE_COMMENT " +
		"FROM information_schema.`TABLES` t " +
		"WHERE t.TABLE_SCHEMA = ?"
	stmt, err := g.Prepare(tablesSql)
	if err != nil {
		log.Panicln(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(g.DBName)
	if err != nil {
		log.Panicln(err)
	}
	defer rows.Close()
	tables := make([]Tables, 0)
	for rows.Next() {
		table := Tables{}
		err = rows.Scan(&table.TableName, &table.TableComment)
		if err != nil {
			log.Println(err)
			continue
		}
		tables = append(tables, table)
	}
	return tables
}

// 获取数据库表信息
func (g *Generator) GetTableInfo(tableName string) []TableInfo {
	tableInfoSql := "SELECT c.TABLE_SCHEMA,c.TABLE_NAME,c.COLUMN_NAME,c.ORDINAL_POSITION,c.COLUMN_DEFAULT,c.IS_NULLABLE," +
		"c.DATA_TYPE,c.CHARACTER_MAXIMUM_LENGTH,c.NUMERIC_PRECISION,c.NUMERIC_SCALE,c.COLUMN_TYPE,c.COLUMN_COMMENT " +
		"FROM information_schema.`COLUMNS` c " +
		"WHERE c.TABLE_SCHEMA = ? " +
		"AND c.TABLE_NAME = ?"
	stmt, err := g.Prepare(tableInfoSql)
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(g.DBName, tableName)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()
	tableInfos := make([]TableInfo, 0)
	for rows.Next() {
		tableInfo := TableInfo{}
		err := rows.Scan(&tableInfo.TableSchema, &tableInfo.TableName, &tableInfo.ColumnName, &tableInfo.OrdinalPosition,
			&tableInfo.ColumnDefault, &tableInfo.IsNullable, &tableInfo.DataType, &tableInfo.CharacterMaximumLength,
			&tableInfo.NumericPrecision, &tableInfo.NumericScale, &tableInfo.ColumnType, &tableInfo.ColumnComment)
		if err != nil {
			log.Println(err)
			continue
		}
		tableInfos = append(tableInfos, tableInfo)
	}
	return tableInfos
}
