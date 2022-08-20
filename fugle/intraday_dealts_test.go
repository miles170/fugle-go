package fugle

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

func testIntradayServiceDealts(t *testing.T, raw string, want DealtsResponse) {
	testIntradayService(t, "dealts", raw, &want, func(client *Client) (interface{}, error) {
		resp, err := client.Intrady.Dealts("", 50, 0, false)
		return resp, err
	})
}

func TestIntradayService_Dealts_2330(t *testing.T) {
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
				"timeZone": "Asia/Taipei"
			},
			"dealts": [
				{
					"at": "2022-08-19T14:30:00.000+08:00",
					"price": 519,
					"volume": 51,
					"serial": 99999999
				},
				{
					"at": "2022-08-19T13:30:00.000+08:00",
					"bid": 518,
					"ask": 519,
					"price": 519,
					"volume": 2127,
					"serial": 6152613
				},
				{
					"at": "2022-08-19T13:24:55.357+08:00",
					"bid": 520,
					"ask": 521,
					"price": 521,
					"volume": 1,
					"serial": 6122455
				},
				{
					"at": "2022-08-19T13:24:55.188+08:00",
					"bid": 520,
					"ask": 521,
					"price": 520,
					"volume": 1,
					"serial": 6122261
				},
				{
					"at": "2022-08-19T13:24:55.175+08:00",
					"bid": 520,
					"ask": 521,
					"price": 521,
					"volume": 6,
					"serial": 6122251
				}
			]
		}
	}`
	want := DealtsResponse{
		APIVersion: "0.3.0",
		Data: DealtData{
			Info: Info{
				Date:        InfoDate(time.Date(2022, 8, 19, 0, 0, 0, 0, time.UTC)),
				Type:        "EQUITY",
				Exchange:    "TWSE",
				Market:      "TSE",
				SymbolID:    "2330",
				CountryCode: "TW",
				TimeZone:    "Asia/Taipei",
			},
			Dealts: []Dealt{
				{
					At:     time.Date(2022, 8, 19, 14, 30, 0, 0, location),
					Price:  decimal.NewFromInt(519),
					Volume: 51,
					Serial: 99999999,
				},
				{
					At:     time.Date(2022, 8, 19, 13, 30, 0, 0, location),
					Bid:    decimal.NewFromInt(518),
					Ask:    decimal.NewFromInt(519),
					Price:  decimal.NewFromInt(519),
					Volume: 2127,
					Serial: 6152613,
				},
				{
					At:     time.Date(2022, 8, 19, 13, 24, 55, 357000000, location),
					Bid:    decimal.NewFromInt(520),
					Ask:    decimal.NewFromInt(521),
					Price:  decimal.NewFromInt(521),
					Volume: 1,
					Serial: 6122455,
				},
				{
					At:     time.Date(2022, 8, 19, 13, 24, 55, 188000000, location),
					Bid:    decimal.NewFromInt(520),
					Ask:    decimal.NewFromInt(521),
					Price:  decimal.NewFromInt(520),
					Volume: 1,
					Serial: 6122261,
				},
				{
					At:     time.Date(2022, 8, 19, 13, 24, 55, 175000000, location),
					Bid:    decimal.NewFromInt(520),
					Ask:    decimal.NewFromInt(521),
					Price:  decimal.NewFromInt(521),
					Volume: 6,
					Serial: 6122251,
				},
			},
		},
	}
	testIntradayServiceDealts(t, raw, want)
}

func TestIntradayService_Dealts_IX0001(t *testing.T) {
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
				"timeZone": "Asia/Taipei"
			},
			"dealts": [
				{
					"at": "2022-08-19T13:49:01.519+08:00",
					"price": 15408.78,
					"volume": 218664552040,
					"serial": 3242
				},
				{
					"at": "2022-08-19T13:30:00.000+08:00",
					"price": 15425.32,
					"volume": 204335996570,
					"serial": 3241
				},
				{
					"at": "2022-08-19T13:29:55.000+08:00",
					"price": 15425.32,
					"volume": 204335996570,
					"serial": 3240
				},
				{
					"at": "2022-08-19T13:29:50.000+08:00",
					"price": 15425.32,
					"volume": 204335996570,
					"serial": 3239
				},
				{
					"at": "2022-08-19T13:29:45.000+08:00",
					"price": 15425.32,
					"volume": 204335996570,
					"serial": 3238
				}
			]
		}
	}`
	want := DealtsResponse{
		APIVersion: "0.3.0",
		Data: DealtData{
			Info: Info{
				Date:        InfoDate(time.Date(2022, 8, 19, 0, 0, 0, 0, time.UTC)),
				Type:        "INDEX",
				Exchange:    "TWSE",
				Market:      "TSE",
				SymbolID:    "IX0001",
				CountryCode: "TW",
				TimeZone:    "Asia/Taipei",
			},
			Dealts: []Dealt{
				{
					At:     time.Date(2022, 8, 19, 13, 49, 1, 519000000, location),
					Price:  decimal.NewFromFloat(15408.78),
					Volume: 218664552040,
					Serial: 3242,
				},
				{
					At:     time.Date(2022, 8, 19, 13, 30, 0, 0, location),
					Price:  decimal.NewFromFloat(15425.32),
					Volume: 204335996570,
					Serial: 3241,
				},
				{
					At:     time.Date(2022, 8, 19, 13, 29, 55, 0, location),
					Price:  decimal.NewFromFloat(15425.32),
					Volume: 204335996570,
					Serial: 3240,
				},
				{
					At:     time.Date(2022, 8, 19, 13, 29, 50, 0, location),
					Price:  decimal.NewFromFloat(15425.32),
					Volume: 204335996570,
					Serial: 3239,
				},
				{
					At:     time.Date(2022, 8, 19, 13, 29, 45, 0, location),
					Price:  decimal.NewFromFloat(15425.32),
					Volume: 204335996570,
					Serial: 3238,
				},
			},
		},
	}
	testIntradayServiceDealts(t, raw, want)
}

func TestIntradayService_DealtsErrors(t *testing.T) {
	testIntradayServiceErros(t, "dealts", func(client *Client) (interface{}, error) {
		resp, err := client.Intrady.Dealts("", 50, 0, false)
		return resp, err
	})
}
