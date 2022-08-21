package integration

import (
	"testing"
	"time"
)

func TestMarketData_Candles(t *testing.T) {
	m, err := client.MarketData.Candles("2884",
		time.Date(2022, 8, 15, 0, 0, 0, 0, time.UTC),
		time.Date(2022, 8, 21, 0, 0, 0, 0, time.UTC))
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
