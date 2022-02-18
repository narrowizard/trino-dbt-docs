package services

import (
	"fmt"

	"github.com/narrowizard/tinysql"
	"github.com/narrowizard/trino-dbt-docs/models"
)

type TableQuery interface {
	GetTables(schema string) ([]models.ShowTables, error)
	GetColumns(tableName string) ([]models.ColumnInfo, error)
}

type TrinoTableQuery struct {
	db *tinysql.DB
}

func NewTrinoTableQuery(dsn string) (*TrinoTableQuery, error) {
	db, err := InitTrinoDB(dsn)
	if err != nil {
		return nil, err
	}
	return &TrinoTableQuery{
		db,
	}, nil
}

func (q TrinoTableQuery) GetTables(schema string) ([]models.ShowTables, error) {
	rows := q.db.Query(fmt.Sprintf("SHOW tables FROM %s", schema))
	var dest = make([]models.ShowTables, 0)
	_, err := rows.Scan(&dest)
	return dest, err
}

func (q TrinoTableQuery) GetColumns(tableName string) ([]models.ColumnInfo, error) {
	rows := q.db.Query(fmt.Sprintf("SHOW columns FROM %s", tableName))
	var showColumns = make([]models.TrinoColumnDescribe, 0)
	_, err := rows.Scan(&showColumns)
	if err != nil {
		return nil, err
	}
	var columnsDetail = make([]models.TrinoColumn, 0)
	rows = q.db.Query(fmt.Sprintf("SELECT * FROM information_schema.columns WHERE table_name = '%s'", tableName))
	_, err = rows.Scan(&columnsDetail)
	if err != nil {
		return nil, err
	}
	var res = make([]models.ColumnInfo, 0)
	for i := 0; i < len(showColumns); i++ {
		var temp = models.ColumnInfo{
			Column:  showColumns[i].Column,
			Type:    showColumns[i].Type,
			Extra:   showColumns[i].Extra,
			Comment: showColumns[i].Comment,
		}
		for _, v := range columnsDetail {
			if showColumns[i].Column == v.ColumnName {
				if v.IsNullable == "YES" {
					temp.IsNullable = true
				}
				temp.ColumnDefault = v.ColumnDefault
			}
		}
		res = append(res, temp)
	}
	return res, err
}

type MysqlTableQuery struct {
	db *tinysql.DB
}

func NewMysqlTableQuery(dsn string) (*MysqlTableQuery, error) {
	db, err := InitMysqlDB(dsn)
	if err != nil {
		return nil, err
	}
	return &MysqlTableQuery{
		db,
	}, nil
}

func (q MysqlTableQuery) GetTables(schema string) ([]models.ShowTables, error) {
	rows := q.db.Query(fmt.Sprintf("SELECT table_name as `Table` from information_schema.tables WHERE TABLE_SCHEMA = '%s'", schema))
	var dest = make([]models.ShowTables, 0)
	_, err := rows.Scan(&dest)
	return dest, err
}

func (q MysqlTableQuery) GetColumns(tableName string) ([]models.ColumnInfo, error) {
	var columnInfo = make([]models.MysqlColumn, 0)
	rows := q.db.Query(fmt.Sprintf("SELECT * FROM information_schema.columns WHERE table_name = '%s'", tableName))
	_, err := rows.Scan(&columnInfo)
	if err != nil {
		return nil, err
	}
	var res = make([]models.ColumnInfo, 0)
	for _, v := range columnInfo {
		var temp = models.ColumnInfo{
			Column:        v.Name,
			Type:          v.ColumnType,
			Comment:       v.ColumnComment,
			ColumnDefault: v.ColumnDefault,
		}
		if v.IsNullable == "YES" {
			temp.IsNullable = true
		}
		res = append(res, temp)
	}
	return res, err
}
