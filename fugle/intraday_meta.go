package fugle

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type Meta struct {
	Market                 string          `json:"market"`
	NameZhTw               string          `json:"nameZhTw"`
	IndustryZhTw           string          `json:"industryZhTw"`
	PreviousClose          decimal.Decimal `json:"previousClose"`
	PriceReference         decimal.Decimal `json:"priceReference"`
	PriceHighLimit         decimal.Decimal `json:"priceHighLimit"`
	PriceLowLimit          decimal.Decimal `json:"priceLowLimit"`
	CanDayBuySell          bool            `json:"canDayBuySell"`
	CanDaySellBuy          bool            `json:"canDaySellBuy"`
	CanShortMargin         bool            `json:"canShortMargin"`
	CanShortLend           bool            `json:"canShortLend"`
	TradingUnit            int             `json:"tradingUnit"`
	Currency               string          `json:"currency"`
	IsTerminated           bool            `json:"isTerminated"`
	IsSuspended            bool            `json:"isSuspended"`
	TypeZhTw               string          `json:"typeZhTw"`
	Abnormal               string          `json:"abnormal"`
	IsUnusuallyRecommended bool            `json:"isUnusuallyRecommended"`
	IsNewlyCompiled        bool            `json:"isNewlyCompiled"`
}

type MetaData struct {
	Info Info `json:"info"`
	Meta Meta `json:"meta"`
}

// See https://developer.fugle.tw/docs/data/intraday/meta
type MetaResponse struct {
	APIVersion string   `json:"apiVersion"`
	Data       MetaData `json:"data"`
}

func (s *IntradayService) Meta(symbolID string, oddLot bool) (*MetaResponse, error) {
	u := fmt.Sprintf("realtime/v%s/intraday/meta", s.client.apiVersion)
	u, err := addOptions(u, IntradyOptions{SymbolID: symbolID, APIToken: s.client.apiToken, OddLot: oddLot})
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", u)
	if err != nil {
		return nil, err
	}

	resp := &MetaResponse{}
	_, err = s.client.Do(req, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
