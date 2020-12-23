package generate

const ModelTemp = `
package model

import (
    "encoding/json"
    "time"
)

type {{.structName}} struct {
{{range $k, $v := .tableInfos}}    {{$v.CamelName}} {{$v.GoType}} ` + "`" + `json:"{{$v.ColumnName}}" form:"{{$v.ColumnName}}"` + "`" + ` // {{.ColumnComment}}
{{end}}}

func (t {{.structName}}) TableName() string {
    return "{{.tableName}}"
}

`
