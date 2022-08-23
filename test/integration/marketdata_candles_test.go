package integration

import (
	"testing"

	"github.com/miles170/fugle-go/fugle"
)

func TestMarketData_Candles(t *testing.T) {
	m, err := client.MarketData.Candles("2884",
		fugle.Date{Year: 2022, Month: 8, Day: 15},
		fugle.Date{Year: 2022, Month: 8, Day: 21})
	if err != nil {
		t.Fatalf("MarketData.Candles returned error: %v", err)
	}
	if m.Type != "EQUITY" {
		t.Fatalf("MarketData.Candles returned type: %s want %s", m.Type, "EQUITY")
	}
	if m.SymbolID != "2884" {
		t.Fatalf("MarketData.Candles returned symbolId: %s want %s", m.SymbolID, "2884")
	}
}
