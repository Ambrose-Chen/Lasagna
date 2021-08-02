package kline

import (
	"fmt"
	"github.com/Ambrose-Chen/Lasagna/DataAcquisition/DBOperation"
	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/huobirdcenter/huobi_golang/pkg/model/market"
	"time"
)

var PeriodTablePrefix = map[string]string{
	market.MIN1:  DBOperation.TABLESPREFIX[0],
	market.MIN5:  DBOperation.TABLESPREFIX[1],
	market.MIN15: DBOperation.TABLESPREFIX[2],
	market.MIN30: DBOperation.TABLESPREFIX[3],
	market.MIN60: DBOperation.TABLESPREFIX[4],
	market.HOUR4: DBOperation.TABLESPREFIX[5],
	market.DAY1:  DBOperation.TABLESPREFIX[6],
	market.WEEK1: DBOperation.TABLESPREFIX[7],
	market.MON1:  DBOperation.TABLESPREFIX[8],
	market.YEAR1: DBOperation.TABLESPREFIX[9],
}

type Kline struct {
	TablePrefix string
	symbol      string
	period      string
}

func (k *Kline) Init(symbol string, period string) *Kline {
	k.symbol = symbol
	k.period = period
	k.TablePrefix = PeriodTablePrefix[period]
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

	DBOperation.CreateTableIfNotExist(k.TablePrefix + k.symbol)
	//maxId := k.getMaxId()
	//
	//if maxId == 0 || (maxId != 0 && candlestick[len(candlestick)-1].Id == getNextTimestamp(maxId, k.period)) {
	dbo := new(DBOperation.DBO).Init()
	for i := len(candlestick) - 1; i >= 0; i-- {
		v := candlestick[i]
		_, err := dbo.DBC.Exec(
			fmt.Sprintf("INSERT INTO `lasagna`.`%s%s`(`id`,`open`,`close`,`low`,`high`,`count`,`vol`,`amount`)VALUES(?,?,?,?,?,?,?,?)", k.TablePrefix, k.symbol),
			v.Id, v.Open, v.Close, v.Low, v.High, v.Count, v.Vol, v.Amount)
		if err != nil {
			applogger.Warn("%v", err)
		}
	}

	err := dbo.DBC.Close()
	if err != nil {
		applogger.Error("%v", err)
	}
	//}
}

func (k *Kline) GetKline(StartTime int64, EndTime int64) {

}

func (k *Kline) getMaxId() int64 {
	dbo := new(DBOperation.DBO).Init()
	rows, err := dbo.DBC.Query(fmt.Sprintf("SELECT max(id) id FROM %s%s WHERE symbol = ? ", k.TablePrefix, k.symbol), k.symbol)
	if err != nil {
		applogger.Error("%v", err)
	}
	var id int64
	rows.Next()

	err = rows.Scan(&id)
	if err != nil {
		applogger.Warn("%v", err)
	}

	err = rows.Close()
	if err != nil {
		applogger.Error("%v", err)
	}

	err = dbo.DBC.Close()
	if err != nil {
		applogger.Error("%v", err)
	}

	return id
}
