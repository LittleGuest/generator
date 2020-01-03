package generate

type Generator struct {
	GlobalConfig
	//*PackageConfig
	//*TemplateConfig
	DBConfig
}

// 代码生成器：数据库信息
type DBConfig struct {
	DriverName string `json:"driver_name"`
	Host       string `json:"host"`
	Port       int64  `json:"port"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	DBName     string `json:"db_name"`
	Extra      string `json:"extra"`
}

// 代码生成器：表信息
type Table struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

// 代码生成器：表信息
type TableInfo struct {
	// 数据库名称
	TableSchema string `json:"table_schema"`
	// 表名
	TableName string `json:"table_name"`
	// 数据库字段名
	ColumnName string `json:"column_name"`
	// 默认值
	ColumnDefault interface{} `json:"column_default"`
	// 是否为空
	IsNullable string `json:"is_nullable"`
	// 数据库字段类型
	DataType string `json:"data_type"`
	// 数字精度
	NumericPrecision interface{} `json:"numeric_precision"`
	// 数值范围
	NumericScale interface{} `json:"numeric_scale"`
	// 字符最大长度
	CharacterMaximumLength interface{} `json:"character_maximum_length"`
	// 字段注释
	ColumnComment string `json:"column_comment"`
	// 驼峰名
	CamelName string `json:"camel_name"`
	// go 类型
	GoType string `json:"go_type"`
}

// 代码生成器：全局配置
type GlobalConfig struct {
	// 是否转驼峰
	CamelCase bool
	// 是否转大驼峰
	Pascal bool
	// 是否覆盖文件
	Override bool
}

func NewGlobalConfig() GlobalConfig {
	return GlobalConfig{
		CamelCase: true,
		Pascal:    true,
		Override:  true,
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
