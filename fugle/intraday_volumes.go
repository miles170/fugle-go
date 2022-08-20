package fugle

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type Volume struct {
	Price       decimal.Decimal `json:"price"`
	Volume      int             `json:"volume"`
	VolumeAtBid int             `json:"volumeAtBid"`
	VolumeAtAsk int             `json:"volumeAtAsk"`
}

type VolumesData struct {
	Info    Info     `json:"info"`
	Volumes []Volume `json:"volumes"`
}

// See https://developer.fugle.tw/docs/data/intraday/volumes
type VolumesResponse struct {
	APIVersion string      `json:"apiVersion"`
	Data       VolumesData `json:"data"`
}

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
