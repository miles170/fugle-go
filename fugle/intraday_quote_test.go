package fugle

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

func testIntradayServiceQuote(t *testing.T, raw string, want QuoteResponse) {
	testIntradayService(t, "quote", raw, &want, func(client *Client) (interface{}, error) {
		resp, err := client.Intrady.Quote("", false)
		return resp, err
	})
}

func TestIntradayService_Quote_2330(t *testing.T) {
	location, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		t.Errorf("time.LoadLocation returned error: %v", err)
	}
	raw := `
	{
		"apiVersion": "0.3.0",
		"data": {
			"info": {
				"date": "2022-08-19",
				"type": "EQUITY",
				"exchange": "TWSE",
				"market": "TSE",
				"symbolId": "2330",
				"countryCode": "TW",
				"timeZone": "Asia/Taipei",
				"lastUpdatedAt": "2022-08-19T13:30:00.000+08:00"
			},
			"quote": {
				"isCurbing": false,
				"isCurbingFall": false,
				"isCurbingRise": false,
				"isTrial": false,
				"isDealt": true,
				"isOpenDelayed": false,
				"isCloseDelayed": false,
				"isHalting": false,
				"isClosed": true,
				"total": {
					"at": "2022-08-19T13:30:00.000+08:00",
					"transaction": 3438,
					"tradeValue": 7254815,
					"tradeVolume": 13950,
					"tradeVolumeAtBid": 5477,
					"tradeVolumeAtAsk": 7478,
					"serial": 6152613
				},
				"trial": {
					"at": "2022-08-19T13:29:55.357+08:00",
					"bid": 518,
					"ask": 519,
					"price": 519,
					"volume": 1925,
					"serial": 6151969
				},
				"trade": {
					"at": "2022-08-19T13:30:00.000+08:00",
					"bid": 518,
					"ask": 519,
					"price": 519,
					"volume": 2127,
					"serial": 6152613
				},
				"order": {
					"at": "2022-08-19T13:30:00.000+08:00",
					"bids": [
						{
							"price": 519,
							"volume": 53
						},
						{
							"price": 518,
							"volume": 602
						},
						{
							"price": 517,
							"volume": 808
						},
						{
							"price": 516,
							"volume": 429
						},
						{
							"price": 515,
							"volume": 722
						}
					],
					"asks": [
						{
							"price": 520,
							"volume": 16
						},
						{
							"price": 521,
							"volume": 314
						},
						{
							"price": 522,
							"volume": 498
						},
						{
							"price": 523,
							"volume": 1131
						},
						{
							"price": 524,
							"volume": 432
						}
					]
				},
				"priceHigh": {
					"price": 523,
					"at": "2022-08-19T10:24:03.296+08:00"
				},
				"priceLow": {
					"price": 517,
					"at": "2022-08-19T09:35:03.109+08:00"
				},
				"priceOpen": {
					"price": 519,
					"at": "2022-08-19T09:00:03.495+08:00"
				},
				"priceAvg": {
					"price": 520.06,
					"at": "2022-08-19T13:30:00.000+08:00"
				},
				"change": -1,
				"changePercent": -0.19,
				"amplitude": 1.15,
				"priceLimit": 0
			}
		}
	}`
	want := QuoteResponse{
		APIVersion: "0.3.0",
		Data: QuoteData{
			Info: Info{
				Date:          InfoDate(time.Date(2022, 8, 19, 0, 0, 0, 0, time.UTC)),
				Type:          "EQUITY",
				Exchange:      "TWSE",
				Market:        "TSE",
				SymbolID:      "2330",
				CountryCode:   "TW",
				TimeZone:      "Asia/Taipei",
				LastUpdatedAt: time.Date(2022, 8, 19, 13, 30, 0, 0, location),
			},
			Quote: Quote{
				IsCurbing:      false,
				IsCurbingFall:  false,
				IsCurbingRise:  false,
				IsTrial:        false,
				IsDealt:        true,
				IsOpenDelayed:  false,
				IsCloseDelayed: false,
				IsHalting:      false,
				IsClosed:       true,
				Total: Total{
					At:               time.Date(2022, 8, 19, 13, 30, 0, 0, location),
					Transaction:      3438,
					TradeValue:       decimal.NewFromInt(7254815),
					TradeVolume:      13950,
					TradeVolumeAtBid: 5477,
					TradeVolumeAtAsk: 7478,
					Serial:           6152613,
				},
				Trial: Trial{
					At:     time.Date(2022, 8, 19, 13, 29, 55, 357000000, location),
					Bid:    decimal.NewFromInt(518),
					Ask:    decimal.NewFromInt(519),
					Pirce:  decimal.NewFromInt(519),
					Volume: 1925,
					Serial: 6151969,
				},
				Trade: Trade{
					At:     time.Date(2022, 8, 19, 13, 30, 0, 0, location),
					Bid:    decimal.NewFromInt(518),
					Ask:    decimal.NewFromInt(519),
					Pirce:  decimal.NewFromInt(519),
					Volume: 2127,
					Serial: 6152613,
				},
				Order: Order{
					At: time.Date(2022, 8, 19, 13, 30, 0, 0, location),
					Bids: []BidAsk{
						{Pirce: decimal.NewFromInt(519), Volume: 53},
						{Pirce: decimal.NewFromInt(518), Volume: 602},
						{Pirce: decimal.NewFromInt(517), Volume: 808},
						{Pirce: decimal.NewFromInt(516), Volume: 429},
						{Pirce: decimal.NewFromInt(515), Volume: 722},
					},
					Asks: []BidAsk{
						{Pirce: decimal.NewFromInt(520), Volume: 16},
						{Pirce: decimal.NewFromInt(521), Volume: 314},
						{Pirce: decimal.NewFromInt(522), Volume: 498},
						{Pirce: decimal.NewFromInt(523), Volume: 1131},
						{Pirce: decimal.NewFromInt(524), Volume: 432},
					},
				},
				PriceHigh: Price{
					At:    time.Date(2022, 8, 19, 10, 24, 3, 296000000, location),
					Pirce: decimal.NewFromInt(523),
				},
				PriceLow: Price{
					At:    time.Date(2022, 8, 19, 9, 35, 3, 109000000, location),
					Pirce: decimal.NewFromInt(517),
				},
				PriceOpen: Price{
					At:    time.Date(2022, 8, 19, 9, 0, 3, 495000000, location),
					Pirce: decimal.NewFromInt(519),
				},
				PriceAverage: Price{
					At:    time.Date(2022, 8, 19, 13, 30, 0, 0, location),
					Pirce: decimal.NewFromFloat(520.06),
				},
				Change:        decimal.NewFromInt(-1),
				ChangePercent: decimal.NewFromFloat(-0.19),
				Amplitude:     decimal.NewFromFloat(1.15),
				PriceLimit:    Normal,
			},
		},
	}
	testIntradayServiceQuote(t, raw, want)
}

func TestIntradayService_Quote_IX0001(t *testing.T) {
	location, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		t.Errorf("time.LoadLocation returned error: %v", err)
	}
	raw := `
	{
		"apiVersion": "0.3.0",
		"data": {
			"info": {
				"date": "2022-08-19",
				"type": "INDEX",
				"exchange": "TWSE",
				"market": "TSE",
				"symbolId": "IX0001",
				"countryCode": "TW",
				"timeZone": "Asia/Taipei",
				"lastUpdatedAt": "2022-08-19T13:30:00.000+08:00"
			},
			"quote": {
				"isClosed": true,
				"priceHigh": {
					"price": 15458.45,
					"at": "2022-08-19T11:09:05.000+08:00"
				},
				"priceLow": {
					"price": 15346.26,
					"at": "2022-08-19T09:35:45.000+08:00"
				},
				"priceOpen": {
					"price": 15394.36,
					"at": "2022-08-19T09:00:05.000+08:00"
				},
				"trade": {
					"at": "2022-08-19T13:30:00.000+08:00",
					"price": 15408.78,
					"serial": 3242
				},
				"last": {
					"at": "2022-08-19T13:30:00.000+08:00",
					"transaction": 57085,
					"tradeValue": 14328555470,
					"tradeVolume": 220091,
					"bidOrders": 4519,
					"askOrders": 3075,
					"bidVolume": 211792,
					"askVolume": 34744,
					"serial": 3242
				},
				"total": {
					"at": "2022-08-19T13:30:00.000+08:00",
					"transaction": 1554797,
					"tradeValue": 218664552040,
					"tradeVolume": 5662269,
					"bidOrders": 8519380,
					"askOrders": 9497041,
					"bidVolume": 19500851,
					"askVolume": 9579438,
					"serial": 3242
				},
				"change": 12.02,
				"changePercent": 0.08,
				"amplitude": 0.73
			}
		}
	}`
	want := QuoteResponse{
		APIVersion: "0.3.0",
		Data: QuoteData{
			Info: Info{
				Date:          InfoDate(time.Date(2022, 8, 19, 0, 0, 0, 0, time.UTC)),
				Type:          "INDEX",
				Exchange:      "TWSE",
				Market:        "TSE",
				SymbolID:      "IX0001",
				CountryCode:   "TW",
				TimeZone:      "Asia/Taipei",
				LastUpdatedAt: time.Date(2022, 8, 19, 13, 30, 0, 0, location),
			},
			Quote: Quote{
				IsClosed: true,
				PriceHigh: Price{
					At:    time.Date(2022, 8, 19, 11, 9, 5, 0, location),
					Pirce: decimal.NewFromFloat(15458.45),
				},
				PriceLow: Price{
					At:    time.Date(2022, 8, 19, 9, 35, 45, 0, location),
					Pirce: decimal.NewFromFloat(15346.26),
				},
				PriceOpen: Price{
					At:    time.Date(2022, 8, 19, 9, 0, 5, 0, location),
					Pirce: decimal.NewFromFloat(15394.36),
				},
				Trade: Trade{
					At:     time.Date(2022, 8, 19, 13, 30, 0, 0, location),
					Pirce:  decimal.NewFromFloat(15408.78),
					Serial: 3242,
				},
				Total: Total{
					At:          time.Date(2022, 8, 19, 13, 30, 0, 0, location),
					Transaction: 1554797,
					TradeValue:  decimal.NewFromInt(218664552040),
					TradeVolume: 5662269,
					BidOrders:   8519380,
					AskOrders:   9497041,
					BidVolume:   19500851,
					AskVolume:   9579438,
					Serial:      3242,
				},
				Last: Last{
					At:          time.Date(2022, 8, 19, 13, 30, 0, 0, location),
					Transaction: 57085,
					TradeValue:  decimal.NewFromInt(14328555470),
					TradeVolume: 220091,
					BidOrders:   4519,
					AskOrders:   3075,
					BidVolume:   211792,
					AskVolume:   34744,
					Serial:      3242,
				},
				Change:        decimal.NewFromFloat(12.02),
				ChangePercent: decimal.NewFromFloat(0.08),
				Amplitude:     decimal.NewFromFloat(0.73),
			},
		},
	}
	testIntradayServiceQuote(t, raw, want)
}

func TestIntradayService_QuoteErrors(t *testing.T) {
	testIntradayServiceErros(t, "quote", func(client *Client) (interface{}, error) {
		resp, err := client.Intrady.Quote("", false)
		return resp, err
	})
}
