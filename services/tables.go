package services

import (
	"github.com/narrowizard/trino-dbt-docs/models"
)

func TransformSourceTable(table models.ShowTables, columns []models.ColumnInfo) (*models.SourceTable, error) {
	var data = models.SourceTable{
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

func TransformModelTable(table models.ShowTables, columns []models.ColumnInfo) (*models.ModelTable, error) {
	var data = models.ModelTable{
		Name:    table.TableName,
		Columns: make([]models.TableColumn, 0),
		Docs: models.ModelTableDocs{
			Show: true,
		},
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
