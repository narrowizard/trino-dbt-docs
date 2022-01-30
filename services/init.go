package services

import (
	"github.com/narrowizard/tinysql"
	"github.com/trinodb/trino-go-client/trino"
)

var db *tinysql.DB

func InitDB(dsn string) error {
	trino.RegisterCustomClient("custom", CustomerHttpClient)
	err := tinysql.RegisterDB("default", "trino", dsn, 10)
	if err != nil {
		return err
	}
	db = tinysql.OpenDefault()
	return nil
}
