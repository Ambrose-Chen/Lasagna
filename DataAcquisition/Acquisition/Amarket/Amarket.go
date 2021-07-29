package Amarket

import (
	"github.com/huobirdcenter/huobi_golang/config"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/huobirdcenter/huobi_golang/pkg/model/market"
)

type AmarketClient struct {
	MarketClient *client.CommonClient
}

func (a *AmarketClient) Init() *AmarketClient {
	a.MarketClient = new(client.CommonClient).Init(config.Host)
	return a
}

func (a *AmarketClient) GetCandlestick(Symbols string, Period string, Size int) ([]market.Candlestick, error) {
	client := new(client.MarketClient).Init(config.Host)

	optionalRequest := market.GetCandlestickOptionalRequest{Period: Period, Size: Size}
	resp, err := client.GetCandlestick(Symbols, optionalRequest)
	if err != nil {
		return []market.Candlestick{}, err
	} else {
		return resp, nil
	}
}
