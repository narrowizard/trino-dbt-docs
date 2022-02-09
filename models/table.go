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
}

type DbtSourceDefination struct {
	Name   string
	Tables []*SourceTable
}
