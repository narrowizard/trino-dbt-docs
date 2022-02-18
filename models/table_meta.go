package models

type ColumnInfo struct {
	Column        string
	Type          string
	Extra         string
	Comment       string
	IsNullable    bool
	ColumnDefault string
}

type ShowTables struct {
	TableName string `col:"Table"`
}
