package models

type TableColumn struct {
	Name        string
	Type        string `yaml:"type"`
	Description string
	Extra       string
	Quote       bool
	Tests       []string
	Tags        []string
	Meta        map[string]interface{}
}

type SourceTable struct {
	Name        string
	Description string
	Columns     []TableColumn
}

type ModelTableDocs struct {
	Show bool `yaml:"show"`
}

type ModelTableConfig struct {
}

type ModelTable struct {
	Name        string
	Description string
	Docs        ModelTableDocs
	Config      ModelTableConfig
	Tests       []string
	Columns     []TableColumn
}

type DbtSourceYamlFile struct {
	Version int
	Sources []DbtSourceDefination
	Models  []*ModelTable
	Seeds   []*SeedTable
}

type DbtSourceDefination struct {
	Name   string
	Tables []*SourceTable
}

type SeedConfig struct {
	QuoteColumns bool              `yaml:"quote_columns"`
	ColumnTypes  map[string]string `yaml:"column_types"`
}

type SeedTable struct {
	Name        string
	Description string
	Docs        ModelTableDocs
	Tests       []string
	Columns     []TableColumn
	Config      SeedConfig
}
