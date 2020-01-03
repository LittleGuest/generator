// 代码生成器：
// 获取配置的数据库的所有表，从 information_schema.tables 获取
// 获取表字段信息, 从 information_schema.columns 获取
package generate

import (
	"archive/zip"
	"generator/util"
	"html/template"
	"log"
	"os"
)

// 单表生成
func (g Generator) SingleGenerate(tableName string) {
	zipFile, err := os.Create("code.zip")
	if err != nil {
		log.Panicln(err)
	}
	defer zipFile.Close()
	zw := zip.NewWriter(zipFile)
	defer zw.Close()

	tableInfos := g.GetTableInfo(tableName)
	g.CreateStruct(zw, tableName, tableInfos)
}

// 多表生成
func (g Generator) MultiGenerate(tableNames string) {
	zipFile, err := os.Create("code.zip")
	if err != nil {
		log.Panicln(err)
	}
	defer zipFile.Close()
	zw := zip.NewWriter(zipFile)
	defer zw.Close()

	done := make(chan bool, 0)
	listTables := g.ListTable(tableNames)
	go func(zw *zip.Writer) {
		for _, v := range listTables {
			tableInfos := g.GetTableInfo(v.Name)
			g.CreateStruct(zw, v.Name, tableInfos)
		}
		done <- true
	}(zw)
	select {
	case <-done:
	}
}

// 创建struct
func (g *Generator) CreateStruct(zw *zip.Writer, tableName string, tableInfos []TableInfo) {
	temp := template.Must(template.New(TempEntityName).ParseFiles(TempEntity))
	fw, err := zw.Create(tableName + ".go")
	if err != nil {
		log.Panicln(err)
	}
	m := make(map[string]interface{})
	m["tableName"] = tableName
	m["structName"] = util.PascalUtil(tableName, "_")
	m["tableInfos"] = tableInfos
	if err = temp.Execute(fw, m); err != nil {
		log.Panicln(err)
	}
}
