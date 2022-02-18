package services

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/narrowizard/tinysql"
	"github.com/trinodb/trino-go-client/trino"
)

func InitTrinoDB(dsn string) (*tinysql.DB, error) {
	trino.RegisterCustomClient("custom", CustomerHttpClient)
	err := tinysql.RegisterDB("default", "trino", dsn, 10)
	if err != nil {
		return nil, err
	}
	db := tinysql.OpenDefault()
	return db, nil
}

func InitMysqlDB(dsn string) (*tinysql.DB, error) {
	err := tinysql.RegisterDB("default", "mysql", dsn, 10)
	if err != nil {
		return nil, err
	}
	db := tinysql.OpenDefault()
	return db, nil
}
