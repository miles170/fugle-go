package fugle

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type Meta struct {
	Market                 string          `json:"market"`                 // 股票所屬市場別
	NameZhTw               string          `json:"nameZhTw"`               // 股票中文簡稱
	IndustryZhTw           string          `json:"industryZhTw"`           // 股票所屬產業別
	PreviousClose          decimal.Decimal `json:"previousClose"`          // 上次收盤價
	PriceReference         decimal.Decimal `json:"priceReference"`         // 今日參考價
	PriceHighLimit         decimal.Decimal `json:"priceHighLimit"`         // 漲停價
	PriceLowLimit          decimal.Decimal `json:"priceLowLimit"`          // 跌停價
	CanDayBuySell          bool            `json:"canDayBuySell"`          // 是否可先買後賣現股當沖
	CanDaySellBuy          bool            `json:"canDaySellBuy"`          // 是否可先賣後買現股當沖
	CanShortMargin         bool            `json:"canShortMargin"`         // 是否豁免平盤下融券賣出
	CanShortLend           bool            `json:"canShortLend"`           // 是否豁免平盤下借券賣出
	TradingUnit            int             `json:"tradingUnit"`            // 交易單位
	Currency               string          `json:"currency"`               // 交易幣別代號
	IsTerminated           bool            `json:"isTerminated"`           // 今日是否已終止上市
	IsSuspended            bool            `json:"isSuspended"`            // 今日是否暫停買賣
	TypeZhTw               string          `json:"typeZhTw"`               // 股票類別
	Abnormal               string          `json:"abnormal"`               // 警示或處置股標示 (正常、注意、處置、注意及處置、再次處置、注意及再次處置、彈性處置、注意及彈性處置)
	IsUnusuallyRecommended bool            `json:"isUnusuallyRecommended"` // 是否為投資理財節目異常推介個股
	IsNewlyCompiled        bool            `json:"isNewlyCompiled"`        // 是否為新編指數 (僅指數類別)
}

type MetaData struct {
	Info Info `json:"info"`
	Meta Meta `json:"meta"`
}

type MetaResponse struct {
	APIVersion string   `json:"apiVersion"`
	Data       MetaData `json:"data"`
}

// 提供盤中個股/指數當日基本資訊
// See https://developer.fugle.tw/docs/data/intraday/meta
func (s *IntradayService) Meta(symbolID string, oddLot bool) (*MetaResponse, error) {
	url := fmt.Sprintf("realtime/v%s/intraday/meta", s.client.apiVersion)
	opts := IntradyOptions{SymbolID: symbolID, APIToken: s.client.apiToken, OddLot: oddLot}
	resp := &MetaResponse{}
	err := s.client.Call(url, opts, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
