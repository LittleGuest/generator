package generator

import (
	"html/template"
	"log"
	"os"
)

// TODO 代码生成器待优化，时间待定
// TODO golang类型 <=> mysql类型
// TODO 字段名大小写、创建目录
// TODO ...
func (g *Generator) Generate() {
	if g.DB == nil {
		g.DB = g.Open()
	}
	if g.GlobalConfig == nil {
		g.GlobalConfig = NewGlobalConfig()
	}

	if g.PackageConfig == nil {
		g.PackageConfig = NewPackageConfig()
	}
	if g.TemplateConfig == nil {
		g.TemplateConfig = NewTemplateConfig()
	}

	defer g.DB.Close()
	listTables := g.ListTables()
	log.Println(listTables)
	g.Create(listTables)
}

// 创建struct
func (g *Generator) CreateEntity(tableName string, tableInfos []TableInfo) {
	temp := template.Must(template.New(TempEntityName).ParseFiles(TempEntity))
	file, err := os.Create("./test/" + tableName + ".go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	m := make(map[string]interface{})
	m["tableName"] = PascalUtil(tableName, "_")
	m["tableInfos"] = tableInfos
	if err = temp.Execute(file, m); err != nil {
		log.Fatalln(err)
	}
}

func (g *Generator) Create(tables []Tables) {
	done := make(chan bool)
	go func() {
		for _, v := range tables {
			tableInfos := g.GetTableInfo(v.TableName)
			g.CreateEntity(v.TableName, tableInfos)
		}
		done <- true
	}()
	select {
	case <-done:
	}
}
