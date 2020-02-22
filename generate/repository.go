package generate

import (
	"generator/database"
	"generator/tool/strtool"
	"log"
	"strings"
)

// ListTable 获取指定数据库中所有表信息
func ListTable(dbName string, tableNames string) (tables []TableInfo) {
	tablesSql := "SELECT t.TABLE_NAME,t.TABLE_COMMENT FROM information_schema.`TABLES` t WHERE t.TABLE_SCHEMA = ?"
	if tableNames != "" {
		tablesSql += " AND FIND_IN_SET(t.TABLE_NAME, ?)"
	}
	stmt, err := database.DB.Prepare(tablesSql)
	if err != nil {
		log.Panicln(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(dbName)
	if tableNames != "" {
		rows, err = stmt.Query(dbName, tableNames)
	}
	if err != nil {
		log.Panicln(err)
	}
	defer rows.Close()
	for rows.Next() {
		table := TableInfo{}
		err = rows.Scan(&table.Name, &table.Comment)
		if err != nil {
			log.Println(err)
			continue
		}
		tables = append(tables, table)
	}
	return
}

// GetTableInfo 获取指定数据库中指定表字段信息
func GetTableInfo(driver string, dbName string, tableName string) (tableInfos []TableFieldInfo) {
	tableInfoSql := "SELECT c.TABLE_SCHEMA,c.TABLE_NAME,c.COLUMN_NAME,c.COLUMN_DEFAULT,c.IS_NULLABLE,c.DATA_TYPE,c.NUMERIC_PRECISION,c.NUMERIC_SCALE,c.CHARACTER_MAXIMUM_LENGTH,c.COLUMN_COMMENT FROM information_schema.`COLUMNS` c WHERE c.TABLE_SCHEMA = ? AND c.TABLE_NAME = ?"
	stmt, err := database.DB.Prepare(tableInfoSql)
	if err != nil {
		log.Panicln(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(dbName, tableName)
	if err != nil {
		log.Panicln(err)
	}
	defer rows.Close()
	for rows.Next() {
		tableInfo := TableFieldInfo{}
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
				tableInfo.CamelName = strtool.ToPascal(tableInfo.ColumnName, "_")
			} else {
				tableInfo.CamelName = strtool.ToCamelCase(tableInfo.ColumnName, "_")
			}
		}
		// 类型转换
		switch strings.ToUpper(driver) {
		case "MYSQL":
			// mysql类型 => golang类型
			tableInfo.GoType = MysqlToGo[tableInfo.DataType]
		default:
			tableInfo.GoType = "interface{}"
		}
		tableInfos = append(tableInfos, tableInfo)
	}
	return
}
