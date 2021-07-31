package kline

import (
	"fmt"
	"github.com/huobirdcenter/huobi_golang/pkg/model/market"
	"github.com/kuianchen/Lasagna/DataAcquisition/DBOperation"
	"time"
)

var periodTable = map[string]string{
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

type Kline struct {
	tableName string
	symbol    string
	period    string
}

func (k *Kline) Init(symbol string, period string) *Kline {
	k.symbol = symbol
	k.period = period
	k.tableName = periodTable[period]
	return k
}

func getNextTimestamp(maxId int64, period string) int64 {
	periodSecond := map[string]int64{
		market.MIN1:  60,
		market.MIN5:  300,
		market.MIN15: 900,
		market.MIN30: 1800,
		market.MIN60: 3600,
		market.HOUR4: 14400,
		market.DAY1:  86400,
		market.WEEK1: 604800,
	}
	var resultId int64
	if _, ok := periodSecond[period]; ok {
		resultId = maxId + periodSecond[period]
	} else if period == market.MON1 {
		TmaxId := time.Unix(maxId, 0)
		NId := TmaxId.AddDate(0, 1, 0)
		resultId = NId.Unix()
	} else if period == market.YEAR1 {
		TmaxId := time.Unix(maxId, 0)
		NId := TmaxId.AddDate(1, 0, 0)
		resultId = NId.Unix()
	}
	return resultId
}

func (k *Kline) Insert(candlestick []market.Candlestick) {
	fmt.Printf("%v\n", getNextTimestamp(1606752000, market.MON1))
	maxId := k.getMaxId()
	if maxId == 0 || (maxId != 0 && candlestick[len(candlestick)-1].Id == getNextTimestamp(maxId, k.period)) {
		dbo := new(DBOperation.DBO).Init()
		defer dbo.DBC.Close()
		//for i, v := range candlestick {
		//	rows, err := dbo.DBC.Query(fmt.Sprintf("SELECT max(id) id FROM %s WHERE symbol = ? ", k.tableName), k.symbol)
		//	if err != nil {
		//		panic(err)
		//	}
		//}
	}
	//if maxId == 0|masId+period_second[period_second]>=
}

func (k *Kline) GetKline(StartTime int64, EndTime int64) {

}

func (k *Kline) getMaxId() int64 {
	dbo := new(DBOperation.DBO).Init()
	rows, err := dbo.DBC.Query(fmt.Sprintf("SELECT max(id) id FROM %s WHERE symbol = ? ", k.tableName), k.symbol)
	if err != nil {
		panic(err)
	}
	var id int64
	rows.Next()
	rows.Scan(&id)
	return id
}
