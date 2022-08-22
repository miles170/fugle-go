package fugle

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

func testMarketDataServiceCandles(t *testing.T, raw string, want CandlesResponse) {
	testMarketDataService(t, "candles", raw, &want, func(client *Client) (interface{}, error) {
		resp, err := client.MarketData.Candles("", time.Now(), time.Now())
		return resp, err
	})
}

func TestMarketDataService_Candles_2330(t *testing.T) {
	raw := `
	{
		"symbolId": "2330",
		"type": "EQUITY",
		"exchange": "TWSE",
		"market": "TSE",
		"candles": [
			{
				"date": "2022-08-19",
				"open": 519,
				"high": 523,
				"low": 517,
				"close": 519,
				"volume": 14235983
			},
			{
				"date": "2022-08-18",
				"open": 520,
				"high": 521,
				"low": 519,
				"close": 520,
				"volume": 18721898
			},
			{
				"date": "2022-08-17",
				"open": 524,
				"high": 527,
				"low": 521,
				"close": 527,
				"volume": 28461939
			},
			{
				"date": "2022-08-16",
				"open": 526,
				"high": 526,
				"low": 523,
				"close": 525,
				"volume": 21234122
			},
			{
				"date": "2022-08-15",
				"open": 520,
				"high": 524,
				"low": 519,
				"close": 523,
				"volume": 22519886
			}
		]
	}`
	want := CandlesResponse{
		SymbolID: "2330",
		Type:     "EQUITY",
		Exchange: "TWSE",
		Market:   "TSE",
		Candles: []Candle{
			{
				Date:   Date{2022, 8, 19},
				Open:   decimal.NewFromInt(519),
				High:   decimal.NewFromInt(523),
				Low:    decimal.NewFromInt(517),
				Close:  decimal.NewFromInt(519),
				Volume: 14235983,
			},
			{
				Date:   Date{2022, 8, 18},
				Open:   decimal.NewFromInt(520),
				High:   decimal.NewFromInt(521),
				Low:    decimal.NewFromInt(519),
				Close:  decimal.NewFromInt(520),
				Volume: 18721898,
			},
			{
				Date:   Date{2022, 8, 17},
				Open:   decimal.NewFromInt(524),
				High:   decimal.NewFromInt(527),
				Low:    decimal.NewFromInt(521),
				Close:  decimal.NewFromInt(527),
				Volume: 28461939,
			},
			{
				Date:   Date{2022, 8, 16},
				Open:   decimal.NewFromInt(526),
				High:   decimal.NewFromInt(526),
				Low:    decimal.NewFromInt(523),
				Close:  decimal.NewFromInt(525),
				Volume: 21234122,
			},
			{
				Date:   Date{2022, 8, 15},
				Open:   decimal.NewFromInt(520),
				High:   decimal.NewFromInt(524),
				Low:    decimal.NewFromInt(519),
				Close:  decimal.NewFromInt(523),
				Volume: 22519886,
			},
		},
	}
	testMarketDataServiceCandles(t, raw, want)
}

func TestMarketDataService_Candles_IX0001(t *testing.T) {
	raw := `
	{
		"symbolId": "IX0001",
		"type": "INDEX",
		"exchange": "TWSE",
		"market": "TSE",
		"candles": [
			{
				"date": "2022-08-19",
				"open": 15394.36,
				"high": 15458.45,
				"low": 15346.26,
				"close": 15408.78,
				"volume": 5696021842
			},
			{
				"date": "2022-08-18",
				"open": 15384.73,
				"high": 15396.76,
				"low": 15311.22,
				"close": 15396.76,
				"volume": 5017166848
			},
			{
				"date": "2022-08-17",
				"open": 15423.44,
				"high": 15475.89,
				"low": 15390.84,
				"close": 15465.45,
				"volume": 5511618302
			},
			{
				"date": "2022-08-16",
				"open": 15435.04,
				"high": 15451.83,
				"low": 15392.49,
				"close": 15420.57,
				"volume": 5723526118
			},
			{
				"date": "2022-08-15",
				"open": 15332.92,
				"high": 15437.23,
				"low": 15315.12,
				"close": 15417.35,
				"volume": 6430023322
			}
		]
	}`
	want := CandlesResponse{
		SymbolID: "IX0001",
		Type:     "INDEX",
		Exchange: "TWSE",
		Market:   "TSE",
		Candles: []Candle{
			{
				Date:   Date{2022, 8, 19},
				Open:   decimal.NewFromFloat(15394.36),
				High:   decimal.NewFromFloat(15458.45),
				Low:    decimal.NewFromFloat(15346.26),
				Close:  decimal.NewFromFloat(15408.78),
				Volume: 5696021842,
			},
			{
				Date:   Date{2022, 8, 18},
				Open:   decimal.NewFromFloat(15384.73),
				High:   decimal.NewFromFloat(15396.76),
				Low:    decimal.NewFromFloat(15311.22),
				Close:  decimal.NewFromFloat(15396.76),
				Volume: 5017166848,
			},
			{
				Date:   Date{2022, 8, 17},
				Open:   decimal.NewFromFloat(15423.44),
				High:   decimal.NewFromFloat(15475.89),
				Low:    decimal.NewFromFloat(15390.84),
				Close:  decimal.NewFromFloat(15465.45),
				Volume: 5511618302,
			},
			{
				Date:   Date{2022, 8, 16},
				Open:   decimal.NewFromFloat(15435.04),
				High:   decimal.NewFromFloat(15451.83),
				Low:    decimal.NewFromFloat(15392.49),
				Close:  decimal.NewFromFloat(15420.57),
				Volume: 5723526118,
			},
			{
				Date:   Date{2022, 8, 15},
				Open:   decimal.NewFromFloat(15332.92),
				High:   decimal.NewFromFloat(15437.23),
				Low:    decimal.NewFromFloat(15315.12),
				Close:  decimal.NewFromFloat(15417.35),
				Volume: 6430023322,
			},
		},
	}
	testMarketDataServiceCandles(t, raw, want)
}

func TestMarketDataService_CandlesErrors(t *testing.T) {
	testMarketDataServiceErros(t, "candles", func(client *Client) (interface{}, error) {
		resp, err := client.MarketData.Candles("", time.Now(), time.Now())
		return resp, err
	})
}
