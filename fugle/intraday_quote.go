package fugle

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

type BidAsk struct {
	Pirce  decimal.Decimal `json:"price"`
	Volume int             `json:"volume"`
}

type Price struct {
	At    time.Time       `json:"at"`
	Pirce decimal.Decimal `json:"price"`
}

type PriceLimit int

const (
	Normal    PriceLimit = 0
	LimitDown PriceLimit = 1
	LimitUp   PriceLimit = 2
)

type Total struct {
	At               time.Time       `json:"at"`
	Transaction      int             `json:"transaction"`
	TradeValue       decimal.Decimal `json:"tradeValue"`
	TradeVolume      int             `json:"tradeVolume"`
	TradeVolumeAtBid int             `json:"tradeVolumeAtBid"`
	TradeVolumeAtAsk int             `json:"tradeVolumeAtAsk"`
	BidOrders        int             `json:"bidOrders"`
	AskOrders        int             `json:"askOrders"`
	BidVolume        int             `json:"bidVolume"`
	AskVolume        int             `json:"askVolume"`
	Serial           int             `json:"serial"`
}

type Last struct {
	At          time.Time       `json:"at"`
	Transaction int             `json:"transaction"`
	TradeValue  decimal.Decimal `json:"tradeValue"`
	TradeVolume int             `json:"tradeVolume"`
	BidOrders   int             `json:"bidOrders"`
	AskOrders   int             `json:"askOrders"`
	BidVolume   int             `json:"bidVolume"`
	AskVolume   int             `json:"askVolume"`
	Serial      int             `json:"serial"`
}

type Trial struct {
	At     time.Time       `json:"at"`
	Bid    decimal.Decimal `json:"bid"`
	Ask    decimal.Decimal `json:"ask"`
	Pirce  decimal.Decimal `json:"price"`
	Volume int             `json:"volume"`
	Serial int             `json:"serial"`
}

type Trade struct {
	At     time.Time       `json:"at"`
	Bid    decimal.Decimal `json:"bid"`
	Ask    decimal.Decimal `json:"ask"`
	Pirce  decimal.Decimal `json:"price"`
	Volume int             `json:"volume"`
	Serial int             `json:"serial"`
}

type Order struct {
	At   time.Time `json:"at"`
	Bids []BidAsk
	Asks []BidAsk
}

type Quote struct {
	IsCurbing      bool `json:"isCurbing"`
	IsCurbingFall  bool `json:"isCurbingFall"`
	IsCurbingRise  bool `json:"isCurbingRise"`
	IsTrial        bool `json:"isTrial"`
	IsOpenDelayed  bool `json:"isOpenDelayed"`
	IsCloseDelayed bool `json:"isCloseDelayed"`
	IsHalting      bool `json:"isHalting"`
	IsDealt        bool `json:"isDealt"`
	IsClosed       bool `json:"isClosed"`

	Total Total `json:"total"`
	Trial Trial `json:"trial"`
	Trade Trade `json:"trade"`
	Order Order `json:"order"`
	Last  Last  `json:"last"`

	PriceHigh    Price `json:"priceHigh"`
	PriceLow     Price `json:"priceLow"`
	PriceOpen    Price `json:"priceOpen"`
	PriceAverage Price `json:"priceAvg"`

	Change        decimal.Decimal `json:"change"`
	ChangePercent decimal.Decimal `json:"changePercent"`
	Amplitude     decimal.Decimal `json:"amplitude"`
	PriceLimit    PriceLimit      `json:"priceLimit"`
}

type QuoteData struct {
	Info  `json:"info"`
	Quote `json:"quote"`
}

// See https://developer.fugle.tw/docs/data/intraday/quote
type QuoteResponse struct {
	APIVersion string    `json:"apiVersion"`
	Data       QuoteData `json:"data"`
}

func (s *IntradayService) Quote(symbolID string, oddLot bool) (*QuoteResponse, error) {
	u := fmt.Sprintf("realtime/v%s/intraday/quote", s.client.apiVersion)
	u, err := addOptions(u, IntradyOptions{SymbolID: symbolID, APIToken: s.client.apiToken, OddLot: oddLot})
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", u)
	if err != nil {
		return nil, err
	}

	resp := &QuoteResponse{}
	_, err = s.client.Do(req, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
