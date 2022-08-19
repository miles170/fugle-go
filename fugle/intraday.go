package fugle

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

// IntradayService handles communication with the intraday related
// methods of the Fugle API.
//
// Fugle API docs: https://developer.fugle.tw/docs/data/intraday/overview
type IntradayService struct {
	client *Client
}

type BasicOptions struct {
	SymbolID string `url:"symbolId"`
	APIToken string `url:"apiToken"`
}

type OddLotOptions struct {
	OddLot bool `url:"oddLot"`
}

type InfoDate time.Time

// UnmarshalJSON handles incoming JSON.
func (d *InfoDate) UnmarshalJSON(b []byte) error {
	t, err := time.Parse("\"2006-01-02\"", string(b))
	if err != nil {
		return err
	}
	*d = InfoDate(t)
	return nil
}

type Info struct {
	Date          InfoDate   `json:"date"`
	Type          string     `json:"type"`
	Exchange      string     `json:"exchange"`
	Market        string     `json:"market"`
	SymbolID      string     `json:"symbolId"`
	CountryCode   string     `json:"countryCode"`
	TimeZone      string     `json:"timeZone"`
	LastUpdatedAt *time.Time `json:"lastUpdatedAt,omitempty"` // (Optional.)
}

type Meta struct {
	Market                 string          `json:"market,omitempty"` // (Optional.)
	NameZhTw               string          `json:"nameZhTw"`
	IndustryZhTw           string          `json:"industryZhTw,omitempty"`  // (Optional.)
	PreviousClose          decimal.Decimal `json:"previousClose,omitempty"` // (Optional.)
	PriceReference         decimal.Decimal `json:"priceReference"`
	PriceHighLimit         decimal.Decimal `json:"priceHighLimit,omitempty"` // (Optional.)
	PriceLowLimit          decimal.Decimal `json:"priceLowLimit,omitempty"`  // (Optional.)
	CanDayBuySell          bool            `json:"canDayBuySell,omitempty"`  // (Optional.)
	CanDaySellBuy          bool            `json:"canDaySellBuy,omitempty"`  // (Optional.)
	CanShortMargin         bool            `json:"canShortMargin,omitempty"` // (Optional.)
	CanShortLend           bool            `json:"canShortLend,omitempty"`   // (Optional.)
	TradingUnit            int             `json:"tradingUnit,omitempty"`    // (Optional.)
	Currency               string          `json:"currency,omitempty"`       // (Optional.)
	IsTerminated           bool            `json:"isTerminated,omitempty"`   // (Optional.)
	IsSuspended            bool            `json:"isSuspended,omitempty"`    // (Optional.)
	TypeZhTw               string          `json:"typeZhTw"`
	Abnormal               string          `json:"abnormal,omitempty"`               // (Optional.)
	IsUnusuallyRecommended bool            `json:"isUnusuallyRecommended,omitempty"` // (Optional.)
	IsNewlyCompiled        bool            `json:"isNewlyCompiled,omitempty"`        // (Optional.)
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

func (s *IntradayService) Meta(symbolID string, opts *OddLotOptions) (*MetaResponse, error) {
	u := fmt.Sprintf("realtime/v%s/intraday/meta", apiVersion)
	u, err := addOptions(u, BasicOptions{SymbolID: symbolID, APIToken: s.client.apiToken})
	if err != nil {
		return nil, err
	}
	u, err = addOptions(u, opts)
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
