package DBOperation

import (
	"database/sql"
	"fmt"
	"github.com/Ambrose-Chen/Lasagna/config"
)

const DBNAME = "lasagna"

var TABLESPREFIX = []string{
	"kline1min",
	"kline5min",
	"kline15min",
	"kline30min",
	"kline60min",
	"kline4hour",
	"kline1day",
	"kline1week",
	"kline1mon",
	"kline1year",
}

type DBO struct {
	DBC *sql.DB
}

func (dbo *DBO) Init() *DBO {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%v)/%s",
		config.DBUser, config.DBPass, config.DBHost, config.DBPort, DBNAME))
	if err != nil {
		panic(err)
	}
	dbo.DBC = db
	return dbo
}

func CreateDBIfNotExist() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%v)/sys",
		config.DBUser, config.DBPass, config.DBHost, config.DBPort))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlCreateDB := "create database " + DBNAME
	db.Exec(sqlCreateDB)

	//CreateTableIfNotExist()
}
func CreateTableIfNotExist(tableName string) {
	dbo := new(DBO).Init()
	defer dbo.DBC.Close()

	sqlCreateTable := fmt.Sprintf("create table %s(id int,open double,close double,low double,high double,count int,vol double,amount double,PRIMARY KEY (`id`))", tableName)
	dbo.DBC.Exec(sqlCreateTable)

}
