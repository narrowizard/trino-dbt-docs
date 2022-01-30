package services

import "github.com/narrowizard/trino-dbt-docs/models"

func TransformTable(table models.ShowTables, columns []models.DescribeColumn) (*models.Table, error) {
	var data = models.Table{
		Name:    table.TableName,
		Columns: make([]models.TableColumn, 0),
	}
	for _, v := range columns {
		var column = models.TableColumn{
			Name:        v.Column,
			DataType:    v.Type,
			Description: v.Comment,
			Extra:       v.Extra,
		}
		data.Columns = append(data.Columns, column)
	}
	return &data, nil
}
