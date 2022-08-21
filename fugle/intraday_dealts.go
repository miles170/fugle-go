package fugle

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

type Dealt struct {
	At     time.Time       `json:"at"`     // 此筆交易的成交時間
	Bid    decimal.Decimal `json:"bid"`    // 此筆交易的買進價
	Ask    decimal.Decimal `json:"ask"`    // 此筆交易的賣出價
	Price  decimal.Decimal `json:"price"`  // 此筆交易的成交價
	Volume int             `json:"volume"` // 此筆交易的成交量
	Serial int             `json:"serial"` // 此筆交易的序號
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
	SymbolID string `url:"symbolId"` // 個股、指數識別代碼
	APIToken string `url:"apiToken"`
	Limit    int    `url:"limit"`  // 限制最多回傳的資料筆數
	Offset   int    `url:"offset"` // 指定從第幾筆後開始回傳
	OddLot   bool   `url:"oddLot"` // 是否回傳零股行情
}

// 取得個股當日所有成交資訊（ex: 個股價量、大盤總量）
// See https://developer.fugle.tw/docs/data/intraday/dealts
func (s *IntradayService) Dealts(symbolID string, limit int, offset int, oddLot bool) (*DealtsResponse, error) {
	url := fmt.Sprintf("realtime/v%s/intraday/dealts", s.client.apiVersion)
	opts := DealtsOptions{
		SymbolID: symbolID,
		APIToken: s.client.apiToken,
		Limit:    limit,
		Offset:   offset,
		OddLot:   oddLot}
	resp := &DealtsResponse{}
	err := s.client.Call(url, opts, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
