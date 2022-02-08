package services

import (
	"fmt"

	"github.com/narrowizard/trino-dbt-docs/models"
)

func GetTables(schema string) ([]models.ShowTables, error) {
	rows := db.Query(fmt.Sprintf("SHOW tables FROM %s", schema))
	var dest = make([]models.ShowTables, 0)
	_, err := rows.Scan(&dest)
	return dest, err
}

func GetColumns(tableName string) ([]models.ColumnInfo, error) {
	rows := db.Query(fmt.Sprintf("SHOW columns FROM %s", tableName))
	var dest = make([]models.ColumnInfo, 0)
	_, err := rows.Scan(&dest)
	if err != nil {
		return nil, err
	}
	var columnsDetail = make([]models.Columns, 0)
	rows = db.Query(fmt.Sprintf("SELECT * FROM information_schema.columns WHERE table_name = '%s'", tableName))
	_, err = rows.Scan(&columnsDetail)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(dest); i++ {
		for _, v := range columnsDetail {
			if dest[i].Column == v.ColumnName {
				if v.IsNullable == "YES" {
					dest[i].IsNullable = true
				}
				dest[i].ColumnDefault = v.ColumnDefault
			}
		}
	}
	return dest, err
}
