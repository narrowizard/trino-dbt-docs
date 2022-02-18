package models

type MysqlColumn struct {
	Name          string `col:"COLUMN_NAME"`
	IsNullable    string `col:"IS_NULLABLE"`
	DataType      string `col:"DATA_TYPE"`   // varchar
	ColumnType    string `col:"COLUMN_TYPE"` // varchar(64)
	ColumnComment string `col:"COLUMN_COMMENT"`
	ColumnDefault string `col:"COLUMN_DEFAULT"`
}
