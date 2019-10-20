// 代码生成器：
// 获取配置的数据库的所有表，从 information_schema.tables 获取
// 获取表字段信息, 从 information_schema.columns 获取
package generator

import (
	"generator/util"
	"html/template"
	"log"
)

// 单表生成
func (g Generator) SingleGenerate(tableName string) {
	// 获取表信息
	tableInfos := g.GetTableInfo(tableName)
	g.CreateStruct(tableName, tableInfos)
}

// 多表生成
func (g Generator) MultiGenerate(tableNames string) {
	listTables := g.ListTable(tableNames)
	g.Run(listTables)
}

// 创建struct
func (g *Generator) CreateStruct(tableName string, tableInfos []TableInfo) {
	temp := template.Must(template.New(TempEntityName).ParseFiles(TempEntity))
	file, err := util.CreateFile("./test/" + tableName + ".go")
	if err != nil {
		log.Printf("创建 struct 失败==>%v", err)
		return
	}
	defer file.Close()
	m := make(map[string]interface{})
	m["tableName"] = util.PascalUtil(tableName, "_")
	m["tableInfos"] = tableInfos
	if err = temp.Execute(file, m); err != nil {
		log.Fatalln(err)
	}
}

// 代码生成启动
func (g *Generator) Run(table []Table) {
	go func() {
		for _, v := range table {
			tableInfos := g.GetTableInfo(v.Name)
			g.CreateStruct(v.Name, tableInfos)
		}
	}()
}
