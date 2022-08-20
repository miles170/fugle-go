package fugle

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

type Dealt struct {
	At     time.Time       `json:"at"`
	Bid    decimal.Decimal `json:"bid"`
	Ask    decimal.Decimal `json:"ask"`
	Price  decimal.Decimal `json:"price"`
	Volume int             `json:"volume"`
	Serial int             `json:"serial"`
}

type DealtData struct {
	Info   `json:"info"`
	Dealts []Dealt `json:"dealts"`
}

// See https://developer.fugle.tw/docs/data/intraday/dealts
type DealtsResponse struct {
	APIVersion string    `json:"apiVersion"`
	Data       DealtData `json:"data"`
}

type DealtsOptions struct {
	SymbolID string `url:"symbolId"`
	APIToken string `url:"apiToken"`
	Limit    int    `url:"limit"`
	Offset   int    `url:"offset"`
	OddLot   bool   `url:"oddLot"`
}

func (s *IntradayService) Dealts(symbolID string, limit int, offset int, oddLot bool) (*DealtsResponse, error) {
	u := fmt.Sprintf("realtime/v%s/intraday/dealts", s.client.apiVersion)
	u, err := addOptions(u, DealtsOptions{
		SymbolID: symbolID,
		APIToken: s.client.apiToken,
		Limit:    limit,
		Offset:   offset,
		OddLot:   oddLot})
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", u)
	if err != nil {
		return nil, err
	}

	resp := &DealtsResponse{}
	_, err = s.client.Do(req, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
