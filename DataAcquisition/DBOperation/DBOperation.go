package DBOperation

import (
	"database/sql"
	"fmt"
	"github.com/kuianchen/Lasagna/config"
)

const DBNAME = "lasagna"

var TABLES = []string{
	"kline1min",
	"kline5min",
	"kline15min",
	"kline30min",
	"kline60min",
	"kline4hour",
	"kline1day",
	"kline1mon",
	"kline1week",
	"kline1year",
}

type DBO struct {
	dbc *sql.DB
}

func (dbo *DBO) Init() *DBO {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%v)/%s",
		config.DBUser, config.DBPass, config.DBHost, config.DBPort, DBNAME))
	if err != nil {
		panic(err)
	}
	dbo.dbc = db
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

	CreateTableIfNotExist()
}
func CreateTableIfNotExist() {
	dbo := new(DBO).Init()
	defer dbo.dbc.Close()

	for _, s := range TABLES {
		sqlCreateTable := fmt.Sprintf(`
create table %s
(
	symbol VARCHAR(20) primary key,
    id int,
	open double,
	close double,
	low double,
	high double,
	count int,
	vol double,
	amount double,
	UNIQUE index symbol_id(id desc)
)
`, s)
		dbo.dbc.Exec(sqlCreateTable)
	}

}
