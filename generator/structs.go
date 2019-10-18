package generator

type Generator struct {
	GlobalConfig
	//*PackageConfig
	//*TemplateConfig
	DBConfig
}

// 代码生成器：数据库信息
type DBConfig struct {
	DriverName string
	Host       string
	Port       int64
	Username   string
	Password   string
	DBName     string
	Extra      string
}

// 代码生成器：表信息
type Table struct {
	Name    string
	Comment string
}

// 代码生成器：表信息
type TableInfo struct {
	// 数据库名称
	TableSchema string
	// 表名
	TableName string
	// 数据库字段名
	ColumnName string
	// 默认值
	ColumnDefault interface{}
	// 是否为空
	IsNullable string
	// 数据库字段类型
	DataType string
	// 数字精度
	NumericPrecision interface{}
	// 数值范围
	NumericScale interface{}
	// 字符最大长度
	CharacterMaximumLength interface{}
	// 字段注释
	ColumnComment string
	// 驼峰名
	CamelName string
	// go 类型
	GoType string
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
