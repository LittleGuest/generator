package generate

import (
	"archive/zip"
	"generator/tool"
	"log"
	"os"
	"text/template"
)

// Create 生成代码
func Create(driver string, dbName string, tableNames string) error {
	zipFile, err := os.Create("code.zip")
	if err != nil {
		log.Panicln(err)
	}
	defer zipFile.Close()
	zw := zip.NewWriter(zipFile)
	defer zw.Close()

	done := make(chan struct{})
	listTables, err := ListTable(dbName, tableNames)
	if err != nil {
		return err
	}
	go func(zw *zip.Writer) {
		for _, v := range listTables {
			tableInfos, err := GetTableInfo(driver, dbName, v.Name)
			if err != nil {
				continue
			}
			CreateStruct(zw, v.Name, tableInfos)
		}
		done <- struct{}{}
	}(zw)
	<-done
	return nil
}

// CreateStruct 生成struct
func CreateStruct(zw *zip.Writer, tableName string, tableInfos []TableFieldInfo) {
	temp := template.Must(template.New(TempModelName).Parse(ModelTemp))
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
