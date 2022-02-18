package main

import (
	"fmt"

	"github.com/narrowizard/trino-dbt-docs/models"
	"github.com/narrowizard/trino-dbt-docs/plugins"
	"github.com/narrowizard/trino-dbt-docs/services"
	_ "github.com/trinodb/trino-go-client/trino"
)

func main() {
	// var catalog = "vdev"
	// var schema = "public"
	// var query, err = services.NewTrinoTableQuery(fmt.Sprintf("https://user:password@host:port?custom_client=custom&catalog=%s&schema=%s", catalog, schema))
	// checkErr(err)
	// tables, err := getMetaData(query, fmt.Sprintf("%s.%s", catalog, schema), "./table_type.yml")
	// checkErr(err)
	query, err := services.NewMysqlTableQuery("user:password@tcp(host:port)/dbname?charset=utf8mb4&loc=UTC&parseTime=True")
	checkErr(err)
	tables, err := getMetaData(query, "schema", "./table_type.yml")
	checkErr(err)
	err = services.WriteToYaml(tables, "./dist/source.yml")
	checkErr(err)
}

func getMetaData(query services.TableQuery, schema string, typeConfig string) (*models.DbtSourceYamlFile, error) {
	tables, err := query.GetTables(schema)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Get %d tables\n", len(tables))
	tableTypes, err := plugins.ReadTableTypeMapping(typeConfig)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Read table type mappings successfully\n")

	var sourceTables = make([]*models.SourceTable, 0)
	var modelTables = make([]*models.ModelTable, 0)
	var seedTables = make([]*models.SeedTable, 0)
	for _, v := range tables {
		columns, err := query.GetColumns(v.TableName)
		if err != nil {
			return nil, err
		}
		var tableType, ok = tableTypes[v.TableName]
		if ok {
			if tableType == plugins.TableTypeModel {
				temp, err := services.TransformModelTable(v, columns)
				if err != nil {
					return nil, err
				}
				modelTables = append(modelTables, temp)
				fmt.Printf("Model table %s transformed\n", v.TableName)
				continue
			} else if tableType == plugins.TableTypeSeed {
				temp, err := services.TransformSeedTable(v, columns)
				if err != nil {
					return nil, err
				}
				seedTables = append(seedTables, temp)
				fmt.Printf("Seed table %s transformed\n", v.TableName)
				continue
			} else if tableType == plugins.TableTypeDeprecated {
				fmt.Printf("Table %s deprecated\n", v.TableName)
				continue
			}
		}
		temp, err := services.TransformSourceTable(v, columns)
		if err != nil {
			return nil, err
		}
		sourceTables = append(sourceTables, temp)
		fmt.Printf("Source table %s transformed\n", v.TableName)
	}
	if err != nil {
		return nil, err
	}
	fmt.Printf("Start writing data to yaml, find %d sources, %d models, %d seeds\n", len(sourceTables), len(modelTables), len(seedTables))
	var result = &models.DbtSourceYamlFile{
		Version: 2,
		Sources: []models.DbtSourceDefination{{
			Name:   schema,
			Tables: sourceTables,
		}},
		Models: modelTables,
		Seeds:  seedTables,
	}
	return result, nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
