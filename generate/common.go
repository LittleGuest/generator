package generate

type Generator struct {
	GlobalConfig
	//*PackageConfig
	//*TemplateConfig
	TableInfo
}

// Table 代码生成器：表信息
type TableInfo struct {
	Name    string           `json:"name"`    // 表名称
	Comment string           `json:"comment"` // 表注释
	Fields  []TableFieldInfo `json:"fields"`  // 表字段
}

// TableInfo 代码生成器：表字段信息
type TableFieldInfo struct {
	TableSchema            string      `json:"table_schema"`             // 数据库名称
	TableName              string      `json:"table_name"`               // 表名
	ColumnName             string      `json:"column_name"`              // 数据库字段名
	ColumnDefault          interface{} `json:"column_default"`           // 默认值
	IsNullable             string      `json:"is_nullable"`              // 是否为空
	DataType               string      `json:"data_type"`                // 数据库字段类型
	NumericPrecision       interface{} `json:"numeric_precision"`        // 数字精度
	NumericScale           interface{} `json:"numeric_scale"`            // 数值范围
	CharacterMaximumLength interface{} `json:"character_maximum_length"` // 字符最大长度
	ColumnComment          string      `json:"column_comment"`           // 字段注释
	CamelName              string      `json:"camel_name"`               // 驼峰名
	GoType                 string      `json:"go_type"`                  // go 类型
}

// GlobalConfig 代码生成器：全局配置
type GlobalConfig struct {
	CamelCase bool // 是否转驼峰
	Pascal    bool // 是否转大驼峰
}

func NewGlobalConfig() GlobalConfig {
	return GlobalConfig{
		CamelCase: true,
		Pascal:    true,
	}
}

//type PackageConfig struct {
//	PkgEntity      string
//	PkgService     string
//	PkgServiceImpl string
//	PkgController  string
//	Path           string
//}
//
//func NewPackageConfig() *PackageConfig {
//	return &PackageConfig{
//		PkgEntity:      PkgEntity,
//		PkgService:     PkgService,
//		PkgServiceImpl: PkgServiceImpl,
//		PkgController:  PkgController,
//		Path:           Path,
//	}
//}
//
//type TemplateConfig struct {
//	TempEntity      string
//	TempService     string
//	TempServiceImpl string
//	TempController  string
//}
//
//func NewTemplateConfig() *TemplateConfig {
//	return &TemplateConfig{
//		TempEntity:      TempEntity,
//		TempService:     TempService,
//		TempServiceImpl: TempServiceImpl,
//		TempController:  TempController,
//	}
//}
