package fugle

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type Volume struct {
	Price       decimal.Decimal `json:"price"`       // 成交價
	Volume      int             `json:"volume"`      // 成交量
	VolumeAtBid int             `json:"volumeAtBid"` // 內盤成交量
	VolumeAtAsk int             `json:"volumeAtAsk"` // 外盤成交量
}

type VolumesData struct {
	Info    Info     `json:"info"`
	Volumes []Volume `json:"volumes"`
}

type VolumesResponse struct {
	APIVersion string      `json:"apiVersion"`
	Data       VolumesData `json:"data"`
}

// 提供盤中個股即時分價量
// See https://developer.fugle.tw/docs/data/intraday/volumes
func (s *IntradayService) Volumes(symbolID string, oddLot bool) (*VolumesResponse, error) {
	url := fmt.Sprintf("realtime/v%s/intraday/volumes", s.client.apiVersion)
	opts := IntradyOptions{SymbolID: symbolID, APIToken: s.client.apiToken, OddLot: oddLot}
	resp := &VolumesResponse{}
	err := s.client.Call(url, opts, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
