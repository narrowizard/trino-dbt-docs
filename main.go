package main

import (
	"fmt"

	"github.com/narrowizard/trino-dbt-docs/models"
	"github.com/narrowizard/trino-dbt-docs/services"
	_ "github.com/trinodb/trino-go-client/trino"
)

func main() {
	var catalog = "vdev"
	var schema = "public"
	dsn := fmt.Sprintf("https://user:password@host:port?custom_client=custom&catalog=%s&schema=%s", catalog, schema)
	err := services.InitDB(dsn)
	checkErr(err)

	tables, err := services.GetTables(fmt.Sprintf("%s.%s", catalog, schema))
	checkErr(err)
	var transformedTables = make([]*models.Table, 0)
	for _, v := range tables {
		columns, err := services.GetColumns(v.TableName)
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
