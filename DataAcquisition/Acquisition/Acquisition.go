package Acquisition

import (
	"github.com/Ambrose-Chen/Lasagna/DataAcquisition/Acquisition/Acommon"
	"github.com/Ambrose-Chen/Lasagna/DataAcquisition/Acquisition/Amarket"
	"github.com/Ambrose-Chen/Lasagna/DataAcquisition/DBOperation/kline"
	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"sync"
)

var ch chan int

const maxRoutineNum = 500

func acqBySymbolPeriod(symbol string, period string, wg *sync.WaitGroup) {
	am := new(Amarket.AmarketClient).Init()
	cl, err := am.GetCandlestick(symbol, period, 2000)
	if err != nil {
		applogger.Error("%v", err)
	}
	kl := new(kline.Kline).Init(symbol, period)
	if len(cl) != 0 {
		kl.Insert(cl)
	}
	<-ch
	wg.Done()
}
func AcqBySymbolPeriod(symbol string, period string) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go acqBySymbolPeriod(symbol, period, &wg)
	wg.Wait()
}

func AcqBySymbol(symbol string) {
	ch = make(chan int, maxRoutineNum)
	wg := sync.WaitGroup{}
	for period := range kline.PeriodTablePrefix {
		ch <- 1
		wg.Add(1)
		go acqBySymbolPeriod(symbol, period, &wg)
	}
	wg.Wait()
}
func AcqAllSymbol() {
	ch = make(chan int, maxRoutineNum)
	wg := sync.WaitGroup{}
	ac := new(Acommon.AcommonClient).Init()
	symbols, err := ac.GetSymbolsByQuoteCurrency("usdt")
	if err != nil {
		applogger.Error("%v", err)
	}
	for _, symbol := range symbols {
		for period := range kline.PeriodTablePrefix {
			ch <- 1
			wg.Add(1)
			go acqBySymbolPeriod(symbol.Symbol, period, &wg)
		}
	}
	wg.Wait()
}
