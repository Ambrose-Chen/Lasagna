package kline

import (
	"fmt"
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
	market.YEAR1: DBOperation.TABLES[9],
}

var period_second = map[string]int{
	market.MIN1:  60,
	market.MIN5:  300,
	market.MIN15: 900,
	market.MIN30: 1800,
	market.MIN60: 3600,
	market.HOUR4: 14400,
	market.DAY1:  86400,
	market.WEEK1: 604800,
	market.MON1:  2678400,
	market.YEAR1: 31622400,
}

type Kline struct {
	tablename string
	symbol    string
	period    string
}

func (k *Kline) Init(symbol string, period string) *Kline {
	k.symbol = symbol
	k.period = period
	k.tablename = period_table[period]
	return k
}

func (k *Kline) Insert([]market.Candlestick) {
	//maxId := k.getMaxId()
	//if maxId == 0|masId+period_second[period_second]>=
}

func (k *Kline) GetKline(StartTime int64, EndTime int64) {

}

func (k *Kline) getMaxId() int64 {
	dbo := new(DBOperation.DBO).Init()
	rows, err := dbo.DBC.Query(fmt.Sprintf("SELECT max(id) id FROM %s WHERE symbol = ? ", k.tablename), k.symbol)
	if err != nil {
		panic(err)
	}
	var id int64
	rows.Next()
	rows.Scan(&id)
	return id
}
