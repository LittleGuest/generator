// 代码生成器：
// 获取配置的数据库的所有表，从 information_schema.tables 获取
// 获取表字段信息, 从 information_schema.columns 获取
package generate

import (
	"archive/zip"
	"html/template"
	"log"
	"os"

	"github.com/LittleGuest/tool"
)

// Create 生成代码
func Create(driver string, dbName string, tableNames string) {
	zipFile, err := os.Create("code.zip")
	if err != nil {
		log.Panicln(err)
	}
	defer zipFile.Close()
	zw := zip.NewWriter(zipFile)
	defer zw.Close()

	done := make(chan bool, 0)
	listTables := ListTable(dbName, tableNames)
	go func(zw *zip.Writer) {
		for _, v := range listTables {
			tableInfos := GetTableInfo(driver,dbName, v.Name)
			CreateStruct(zw, v.Name, tableInfos)
		}
		done <- true
	}(zw)
	<-done
}

// CreateStruct 生成struct
func CreateStruct(zw *zip.Writer, tableName string, tableInfos []TableFieldInfo) {
	temp := template.Must(template.New(TempEntityName).ParseFiles(TempEntity))
	fw, err := zw.Create(tableName + ".go")
	if err != nil {
		log.Panicln(err)
	}
	m := make(map[string]interface{})
	m["tableName"] = tableName
	m["structName"] = tool.ToPascal(tableName, "_")
	m["tableInfos"] = tableInfos
	if err = temp.Execute(fw, m); err != nil {
		log.Panicln(err)
	}
}
