package fugle

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

type BidAsk struct {
	Pirce  decimal.Decimal `json:"price"`  // 價格
	Volume int             `json:"volume"` // 數量
}

type Price struct {
	At    time.Time       `json:"at"`    // 時間
	Pirce decimal.Decimal `json:"price"` // 價格
}

type PriceLimit int

const (
	Normal    PriceLimit = 0 // 正常
	LimitDown PriceLimit = 1 // 跌停
	LimitUp   PriceLimit = 2 // 漲停
)

type Total struct {
	At               time.Time       `json:"at"`               // 最新一筆成交時間
	Transaction      int             `json:"transaction"`      // 總成交筆數
	TradeValue       decimal.Decimal `json:"tradeValue"`       // 總成交金額
	TradeVolume      int             `json:"tradeVolume"`      // 總成交數量
	TradeVolumeAtBid int             `json:"tradeVolumeAtBid"` // 個股內盤成交量
	TradeVolumeAtAsk int             `json:"tradeVolumeAtAsk"` // 個股外盤成交量
	BidOrders        int             `json:"bidOrders"`        // 總委買筆數 (僅加權、櫃買指數)
	AskOrders        int             `json:"askOrders"`        // 總委賣筆數 (僅加權、櫃買指數)
	BidVolume        int             `json:"bidVolume"`        // 總委買數量 (僅加權、櫃買指數)
	AskVolume        int             `json:"askVolume"`        // 總委賣數量 (僅加權、櫃買指數)
	Serial           int             `json:"serial"`           // 最新一筆成交之序號
}

type Last struct {
	At          time.Time       `json:"at"`          // 最新一筆成交時間
	Transaction int             `json:"transaction"` // 總成交筆數
	TradeValue  decimal.Decimal `json:"tradeValue"`  // 總成交金額
	TradeVolume int             `json:"tradeVolume"` // 總成交數量
	BidOrders   int             `json:"bidOrders"`   // 總委買筆數
	AskOrders   int             `json:"askOrders"`   // 總委賣筆數
	BidVolume   int             `json:"bidVolume"`   // 總委買數量
	AskVolume   int             `json:"askVolume"`   // 總委賣數量
	Serial      int             `json:"serial"`      // 最新一筆成交之序號
}

type Trial struct {
	At     time.Time       `json:"at"`     // 最新一筆試撮時間
	Bid    decimal.Decimal `json:"bid"`    // 最新一筆試撮買進價
	Ask    decimal.Decimal `json:"ask"`    // 最新一筆試撮賣出價
	Pirce  decimal.Decimal `json:"price"`  // 最新一筆試撮成交價
	Volume int             `json:"volume"` // 最新一筆試撮成交量
	Serial int             `json:"serial"` // 最新一筆試撮之序號
}

type Trade struct {
	At     time.Time       `json:"at"`     // 最新一筆成交時間
	Bid    decimal.Decimal `json:"bid"`    // 最新一筆成交買進價
	Ask    decimal.Decimal `json:"ask"`    // 最新一筆成交賣出價
	Pirce  decimal.Decimal `json:"price"`  // 最新一筆成交價
	Volume int             `json:"volume"` // 最新一筆成交量
	Serial int             `json:"serial"` // 最新一筆成交之序號
}

type Order struct {
	At   time.Time `json:"at"` // 最新一筆最佳五檔更新時間
	Bids []BidAsk  // 委買資料
	Asks []BidAsk  // 委賣資料
}

type Quote struct {
	IsCurbing      bool `json:"isCurbing"`      // 最近一次更新是否為瞬間價格穩定措施
	IsCurbingFall  bool `json:"isCurbingFall"`  // 最近一次更新是否為暫緩撮合且瞬間趨跌
	IsCurbingRise  bool `json:"isCurbingRise"`  // 最近一次更新是否為暫緩撮合且瞬間趨漲
	IsTrial        bool `json:"isTrial"`        // 最近一次更新是否為試算
	IsOpenDelayed  bool `json:"isOpenDelayed"`  // 最近一次更新是否為延後開盤狀態
	IsCloseDelayed bool `json:"isCloseDelayed"` // 最近一次更新是否為延後收盤狀態
	IsHalting      bool `json:"isHalting"`      // 最近一次更新是否為暫停交易
	IsDealt        bool `json:"isDealt"`        // 最近一次更新是否包含最新成交(試撮)價
	IsClosed       bool `json:"isClosed"`       // 當日是否為已收盤

	Total Total `json:"total"`
	Trial Trial `json:"trial"`
	Trade Trade `json:"trade"`
	Order Order `json:"order"`
	Last  Last  `json:"last"`

	PriceHigh    Price `json:"priceHigh"` // 當日之最高價，第一次到達當日最高價之時間
	PriceLow     Price `json:"priceLow"`  // 當日之最低價，第一次到達當日最低價之時間
	PriceOpen    Price `json:"priceOpen"` // 當日之開盤價，開盤定義：當天第一筆成交時才開盤，當日第一筆成交時間
	PriceAverage Price `json:"priceAvg"`  // 當日之成交均價，當日最後一筆成交時間

	Change        decimal.Decimal `json:"change"`        // 當日股價之漲跌
	ChangePercent decimal.Decimal `json:"changePercent"` // 當日股價之漲跌幅
	Amplitude     decimal.Decimal `json:"amplitude"`     // 當日股價之振幅
	PriceLimit    PriceLimit      `json:"priceLimit"`
}

type QuoteData struct {
	Info  `json:"info"`
	Quote `json:"quote"`
}

type QuoteResponse struct {
	APIVersion string    `json:"apiVersion"`
	Data       QuoteData `json:"data"`
}

// 提供盤中個股/指數逐筆交易金額、狀態、最佳五檔及統計資訊
// See https://developer.fugle.tw/docs/data/intraday/quote
func (s *IntradayService) Quote(symbolID string, oddLot bool) (*QuoteResponse, error) {
	url := fmt.Sprintf("realtime/v%s/intraday/quote", s.client.apiVersion)
	opts := IntradyOptions{SymbolID: symbolID, APIToken: s.client.apiToken, OddLot: oddLot}
	resp := &QuoteResponse{}
	err := s.client.Call(url, opts, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
