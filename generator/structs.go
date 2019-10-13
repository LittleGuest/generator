package generator

import "database/sql"

type Generator struct {
	*sql.DB
	*GlobalConfig
	*PackageConfig
	*TemplateConfig
	*DBConfig
}

type DBConfig struct {
	DriverName string
	Host       string
	Port       int
	Username   string
	Password   string
	DBName     string
	OtherParam string
}

type Tables struct {
	TableName    string
	TableComment string
}

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

type GlobalConfig struct {
	CamelCase bool
	Pascal    bool
	Override  bool
}

func NewGlobalConfig() *GlobalConfig {
	return &GlobalConfig{
		CamelCase: true,
		Pascal:    true,
		Override:  true,
	}
}

type PackageConfig struct {
	PkgEntity      string
	PkgService     string
	PkgServiceImpl string
	PkgController  string
	Path           string
}

func NewPackageConfig() *PackageConfig {
	return &PackageConfig{
		PkgEntity:      PkgEntity,
		PkgService:     PkgService,
		PkgServiceImpl: PkgServiceImpl,
		PkgController:  PkgController,
		Path:           Path,
	}
}

type TemplateConfig struct {
	TempEntity      string
	TempService     string
	TempServiceImpl string
	TempController  string
}

func NewTemplateConfig() *TemplateConfig {
	return &TemplateConfig{
		TempEntity:      TempEntity,
		TempService:     TempService,
		TempServiceImpl: TempServiceImpl,
		TempController:  TempController,
	}
}
