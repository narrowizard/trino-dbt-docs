package services

import (
	"github.com/narrowizard/trino-dbt-docs/models"
)

func TransformTable(table models.ShowTables, columns []models.ColumnInfo) (*models.Table, error) {
	var data = models.Table{
		Name:    table.TableName,
		Columns: make([]models.TableColumn, 0),
	}
	for _, v := range columns {
		var tests = make([]string, 0)
		if !v.IsNullable {
			tests = append(tests, "not_null")
		}
		var column = models.TableColumn{
			Name:        v.Column,
			Type:        v.Type,
			Description: v.Comment,
			Extra:       v.Extra,
			Tests:       tests,
		}
		data.Columns = append(data.Columns, column)
	}
	return &data, nil
}
