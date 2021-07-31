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
func (a *AcommonClient) GetAllSymbols() ([]common.Symbol, error) {
	return a.GetSymbolsByQuoteCurrency("")
}

func (a *AcommonClient) GetSymbolsByQuoteCurrency(quoteCurrency string) ([]common.Symbol, error) {
	resp, err := a.CommonClient.GetSymbols()
	var result []common.Symbol
	if err != nil {
		return []common.Symbol{}, err
	} else {
		if quoteCurrency == "" {
			return resp, nil
		}
		for _, res := range resp {
			if res.QuoteCurrency == quoteCurrency {
				result = append(result, res)
			}
		}
		return result, nil
	}
}
