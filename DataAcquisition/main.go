package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/huobirdcenter/huobi_golang/pkg/model/market"
	"github.com/kuianchen/Lasagna/DataAcquisition/DBOperation"
	"github.com/kuianchen/Lasagna/DataAcquisition/DBOperation/kline"
)

func main() {
	DBOperation.CreateDBIfNotExist()

	kl := new(kline.Kline).Init("btcusdt", market.DAY1)
	kl.Insert([]market.Candlestick{})
}
