package generator

import (
	"database/sql"
	"fmt"
	"generator/utils"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var generatorPool *sql.DB

// 获取指定数据库连接
func (g Generator) OpenGeneratorPool() *sql.DB {
	if generatorPool != nil {
		return generatorPool
	}
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", g.Username, g.Password, g.Host, g.Port, g.DBName)
	if g.Extra != "" {
		dataSource += fmt.Sprintf("?%s", g.Extra)
	}
	db, err := sql.Open(g.DriverName, dataSource)
	if err != nil {
		log.Fatalf("获取指定数据库连接失败：%s", err.Error())
	}
	return db
}

// 获取指定数据库中所有表信息
func (g Generator) ListTable() (tables []Table) {
	tablesSql := "SELECT t.TABLE_NAME,t.TABLE_COMMENT FROM information_schema.`TABLES` t WHERE t.TABLE_SCHEMA = ?"
	stmt, err := generatorPool.Prepare(tablesSql)
	if err != nil {
		log.Panicln(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(g.DBName)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		table := Table{}
		err = rows.Scan(&table.Name, &table.Comment)
		if err != nil {
			log.Println(err)
			continue
		}
		tables = append(tables, table)
	}
	return
}

// 获取指定数据库中指定表字段信息
func (g Generator) GetTableInfo(tableName string) (tableInfos []TableInfo) {
	tableInfoSql := "SELECT c.TABLE_SCHEMA,c.TABLE_NAME,c.COLUMN_NAME,c.COLUMN_DEFAULT,c.IS_NULLABLE,c.DATA_TYPE,c.NUMERIC_PRECISION,c.NUMERIC_SCALE,c.CHARACTER_MAXIMUM_LENGTH,c.COLUMN_COMMENT FROM information_schema.`COLUMNS` c WHERE c.TABLE_SCHEMA = ? AND c.TABLE_NAME = ?"
	stmt, err := g.OpenGeneratorPool().Prepare(tableInfoSql)
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(g.DBName, tableName)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()
	for rows.Next() {
		tableInfo := TableInfo{}
		err := rows.Scan(
			&tableInfo.TableSchema,
			&tableInfo.TableName,
			&tableInfo.ColumnName,
			&tableInfo.ColumnDefault,
			&tableInfo.IsNullable,
			&tableInfo.DataType,
			&tableInfo.NumericPrecision,
			&tableInfo.NumericScale,
			&tableInfo.CharacterMaximumLength,
			&tableInfo.ColumnComment,
		)
		if err != nil {
			log.Println(err)
			continue
		}
		globalConfig := NewGlobalConfig()
		// 转换大小写
		if globalConfig.CamelCase {
			if globalConfig.Pascal {
				tableInfo.CamelName = utils.PascalUtil(tableInfo.ColumnName, "_")
			} else {
				tableInfo.CamelName = utils.CamelCaseUtil(tableInfo.ColumnName, "_")
			}
		}
		// mysql类型 => golang类型
		tableInfo.GoType = convert[tableInfo.DataType]
		tableInfos = append(tableInfos, tableInfo)
	}
	return
}
