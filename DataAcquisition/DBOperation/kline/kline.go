package kline

import (
	"github.com/huobirdcenter/huobi_golang/pkg/model/market"
	"github.com/kuianchen/Lasagna/DataAcquisition/DBOperation"
)

var period_table = map[string]string{
	market.MIN1:  DBOperation.TABLES[0],
	market.MIN5:  DBOperation.TABLES[1],
	market.MIN15: DBOperation.TABLES[2],
	market.MIN30: DBOperation.TABLES[3],
	market.MIN60: DBOperation.TABLES[4],
	market.HOUR4: DBOperation.TABLES[5],
	market.DAY1:  DBOperation.TABLES[6],
	market.WEEK1: DBOperation.TABLES[7],
	market.MON1:  DBOperation.TABLES[8],
}

type kline struct {
	tablename string
	symbol    string
	period    string
}

func (k *kline) Init(symbol string, period string) *kline {
	k.symbol = symbol
	k.period = period
	return k
}

func (k *kline) Insert([]market.Candlestick) {

}

func (k *kline) GetKline(StartTime int64, EndTime int64) {

}

func (k *kline) getMaxId() int {
	//dbo := new(DBOperation.DBO).Init()
	//dbo.dbc.Query("SELECT max(id) FROM ? WHERE symbol = ? ")
	return 0
}
