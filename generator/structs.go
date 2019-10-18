package generator

type Generator struct {
	//*GlobalConfig
	//*PackageConfig
	//*TemplateConfig
	DBConfig
}

// 代码生成器：数据库信息
type DBConfig struct {
	DriverName string
	Host       string
	Port       int
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
	TableSchema            string
	TableName              string
	ColumnName             string
	OrdinalPosition        int64
	ColumnDefault          interface{}
	IsNullable             string
	DataType               string
	CharacterMaximumLength interface{}
	NumericPrecision       interface{}
	NumericScale           interface{}
	ColumnType             string
	ColumnComment          string
}

// 代码生成器：全局配置
type GlobalConfig struct {
	CamelCase bool
	Pascal    bool
	Override  bool
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
