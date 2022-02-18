package models

type TrinoColumn struct {
	TableCatalog    string `sql:"table_catalog"`
	TableSchema     string `sql:"table_schema"`
	TableName       string `sql:"table_name"`
	ColumnName      string `sql:"column_name"`
	OrdinalPosition int    `sql:"ordinal_position"`
	ColumnDefault   string `sql:"column_default"`
	IsNullable      string `sql:"is_nullable"`
	DataType        string `sql:"data_type"`
}

type TrinoColumnDescribe struct {
	Column  string `col:"Column"`
	Type    string `col:"Type"`
	Extra   string `col:"Extra"`
	Comment string `col:"Comment"`
}
