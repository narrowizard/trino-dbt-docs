package main

import (
	"fmt"

	"github.com/narrowizard/trino-dbt-docs/models"
	"github.com/narrowizard/trino-dbt-docs/plugins"
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
	fmt.Printf("Get %d tables\n", len(tables))
	tableTypes, err := plugins.ReadTableTypeMapping("./table_type.yml")
	checkErr(err)
	fmt.Printf("Read table type mappings successfully\n")

	var sourceTables = make([]*models.SourceTable, 0)
	var modelTables = make([]*models.ModelTable, 0)
	var seedTables = make([]*models.ModelTable, 0)
	for _, v := range tables {
		columns, err := services.GetColumns(v.TableName)
		checkErr(err)
		var tableType, ok = tableTypes[v.TableName]
		if ok {
			if tableType == plugins.TableTypeModel {
			temp, err := services.TransformModelTable(v, columns)
			checkErr(err)
			modelTables = append(modelTables, temp)
				fmt.Printf("Model table %s transformed\n", v.TableName)
				continue
			} else if tableType == plugins.TableTypeSeed {
				temp, err := services.TransformModelTable(v, columns)
				checkErr(err)
				seedTables = append(seedTables, temp)
				fmt.Printf("Seed table %s transformed\n", v.TableName)
				continue
			} else if tableType == plugins.TableTypeDeprecated {
				fmt.Printf("Table %s deprecated\n", v.TableName)
				continue
			}
		}
			temp, err := services.TransformSourceTable(v, columns)
			checkErr(err)
			sourceTables = append(sourceTables, temp)
		fmt.Printf("Source table %s transformed\n", v.TableName)
	}
	checkErr(err)
	fmt.Printf("Start writing data to yaml, find %d sources, %d models, %d seeds\n", len(sourceTables), len(modelTables), len(seedTables))
	err = services.WriteToYaml(&models.DbtSourceYamlFile{
		Version: 2,
		Sources: []models.DbtSourceDefination{{
			Name:   schema,
			Tables: sourceTables,
		}},
		Models: modelTables,
		Seeds:  seedTables,
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
