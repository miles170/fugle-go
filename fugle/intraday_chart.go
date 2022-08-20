package fugle

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type Chart struct {
	Averages   []decimal.Decimal `json:"a"`
	Opens      []decimal.Decimal `json:"o"`
	Highs      []decimal.Decimal `json:"h"`
	Lows       []decimal.Decimal `json:"l"`
	Closes     []decimal.Decimal `json:"c"`
	Volumes    []int             `json:"v"`
	Timestamps []Timestamp       `json:"t"`
}

type ChartData struct {
	Info  `json:"info"`
	Chart `json:"chart"`
}

// See https://developer.fugle.tw/docs/data/intraday/chart
type ChartResponse struct {
	APIVersion string    `json:"apiVersion"`
	Data       ChartData `json:"data"`
}

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
