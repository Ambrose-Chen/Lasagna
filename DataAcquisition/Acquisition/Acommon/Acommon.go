package Acommon

import (
	"github.com/huobirdcenter/huobi_golang/config"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/huobirdcenter/huobi_golang/pkg/model/common"
)

type AcommonClient struct {
	CommonClient *client.CommonClient
}

func (a *AcommonClient) Init() *AcommonClient {
	a.CommonClient = new(client.CommonClient).Init(config.Host)
	return a
}
func (a *AcommonClient) GetSymbols() ([]common.Symbol, error) {
	return a.GetSymbolsByQuoteCurrency("")
}

func (a *AcommonClient) GetSymbolsByQuoteCurrency(quote_currency string) ([]common.Symbol, error) {
	resp, err := a.CommonClient.GetSymbols()
	var result []common.Symbol
	if err != nil {
		return []common.Symbol{}, err
	} else {
		if quote_currency == "" {
			return resp, nil
		}
		for _, res := range resp {
			if res.QuoteCurrency == quote_currency {
				result = append(result, res)
			}
		}
		return result, nil
	}
}
