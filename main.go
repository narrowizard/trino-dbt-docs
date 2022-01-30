package main

import (
	"fmt"

	"github.com/narrowizard/trino-dbt-docs/models"
	"github.com/narrowizard/trino-dbt-docs/services"
	_ "github.com/trinodb/trino-go-client/trino"
)

func main() {
	dsn := "https://user:password@host:port?custom_client=custom"
	err := services.InitDB(dsn)
	checkErr(err)
	var catalog = "catalog"
	var schema = "schema"
	tables, err := services.GetTables(fmt.Sprintf("%s.%s", catalog, schema))
	checkErr(err)
	var transformedTables = make([]*models.Table, 0)
	for _, v := range tables {
		columns, err := services.GetColumns(fmt.Sprintf("%s.%s.%s", catalog, schema, v.TableName))
		checkErr(err)
		temp, err := services.TransformTable(v, columns)
		checkErr(err)
		transformedTables = append(transformedTables, temp)
	}
	checkErr(err)
	err = services.WriteToYaml(&models.DbtSourceYamlFile{
		Version: 2,
		Sources: []models.DbtSourceDefination{{
			Name:   schema,
			Tables: transformedTables,
		}},
	}, "./dist/source.yml")
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
