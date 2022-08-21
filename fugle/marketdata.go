package fugle

// MarketDataService handles communication with the marketdata related
// methods of the Fugle API.
//
// Fugle API docs: https://developer.fugle.tw/docs/data/marketdata/candles
type MarketDataService struct {
	client *Client
}

type MarketDataOptions struct {
	SymbolID string `url:"symbolId"` // 個股、指數識別代碼
	APIToken string `url:"apiToken"`
	From     string `url:"from"` // 開始日期
	To       string `url:"to"`   // 結束日期
}
