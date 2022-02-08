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

type Table struct {
	Name        string
	Description string
	Columns     []TableColumn
}

type DbtSourceYamlFile struct {
	Version int
	Sources []DbtSourceDefination
}

type DbtSourceDefination struct {
	Name   string
	Tables []*Table
}
