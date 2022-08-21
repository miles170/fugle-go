package fugle

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

type Candle struct {
	Date   InfoDate        `json:"date"`
	Open   decimal.Decimal `json:"open"`
	High   decimal.Decimal `json:"high"`
	Low    decimal.Decimal `json:"low"`
	Close  decimal.Decimal `json:"close"`
	Volume int             `json:"volume"`
}

// See https://developer.fugle.tw/docs/data/marketdata/candles
type CandlesResponse struct {
	SymbolID string   `json:"symbolId"`
	Type     string   `json:"type"`
	Exchange string   `json:"exchange"`
	Market   string   `json:"market"`
	Candles  []Candle `json:"candles"`
}

func (s *MarketDataService) Candles(symbolID string, from time.Time, to time.Time) (*CandlesResponse, error) {
	url := fmt.Sprintf("marketdata/v%s/candles", s.client.apiVersion)
	opts := MarketDataOptions{SymbolID: symbolID, APIToken: s.client.apiToken, From: from.Format("2006-01-02"), To: to.Format("2006-01-02")}
	resp := &CandlesResponse{}
	err := s.client.Call(url, opts, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
