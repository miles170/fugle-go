package fugle

import "time"

// MarketDataService handles communication with the marketdata related
// methods of the Fugle API.
//
// Fugle API docs: https://developer.fugle.tw/docs/data/marketdata/candles
type MarketDataService struct {
	client *Client
}

type MarketDataOptions struct {
	SymbolID string    `url:"symbolId"`
	APIToken string    `url:"apiToken"`
	From     time.Time `url:"from"`
	To       time.Time `url:"to"`
}
