package fugle

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type Chart struct {
	Averages   []decimal.Decimal `json:"a"` // 當日個股於此分鐘的成交均價
	Opens      []decimal.Decimal `json:"o"` // 此分鐘的開盤價
	Highs      []decimal.Decimal `json:"h"` // 此分鐘的最高價
	Lows       []decimal.Decimal `json:"l"` // 此分鐘的最低價
	Closes     []decimal.Decimal `json:"c"` // 此分鐘的收盤價
	Volumes    []int             `json:"v"` // 此分鐘的成交量 (指數：金額；個股：張數；興櫃股票及零股：股數)
	Timestamps []Timestamp       `json:"t"`
}

type ChartData struct {
	Info  `json:"info"`
	Chart `json:"chart"`
}

type ChartResponse struct {
	APIVersion string    `json:"apiVersion"`
	Data       ChartData `json:"data"`
}

// 提供盤中個股/指數 線圖時所需的各項即時資訊
// See https://developer.fugle.tw/docs/data/intraday/chart
func (s *IntradayService) Chart(symbolID string, oddLot bool) (*ChartResponse, error) {
	url := fmt.Sprintf("realtime/v%s/intraday/chart", s.client.apiVersion)
	opts := IntradyOptions{SymbolID: symbolID, APIToken: s.client.apiToken, OddLot: oddLot}
	resp := &ChartResponse{}
	err := s.client.Call(url, opts, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
