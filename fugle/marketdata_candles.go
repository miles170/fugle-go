package fugle

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

type Candle struct {
	Date   InfoDate        `json:"date"`   // 本筆資料所屬日期
	Open   decimal.Decimal `json:"open"`   // 開盤價
	High   decimal.Decimal `json:"high"`   // 最高價
	Low    decimal.Decimal `json:"low"`    // 最低價
	Close  decimal.Decimal `json:"close"`  // 收盤價
	Volume int             `json:"volume"` // 成交量
}

type CandlesResponse struct {
	SymbolID string   `json:"symbolId"` // 個股、指數識別代碼
	Type     string   `json:"type"`     // ticker 類別
	Exchange string   `json:"exchange"` // 交易所
	Market   string   `json:"market"`   // 市場別
	Candles  []Candle `json:"candles"`  // 歷史股價資料
}

// 提供歷史股價資料，包含開高低收量 (OHLCV)。歷史資料目前設計單次呼叫的資料區間以一年為限；資料區間最遠可回溯至 2010 年！
// See https://developer.fugle.tw/docs/data/marketdata/candles
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
