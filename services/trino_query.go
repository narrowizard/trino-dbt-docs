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

func GetColumns(tableName string) ([]models.DescribeColumn, error) {
	rows := db.Query(fmt.Sprintf("SHOW columns FROM %s", tableName))
	var dest = make([]models.DescribeColumn, 0)
	_, err := rows.Scan(&dest)
	return dest, err
}
