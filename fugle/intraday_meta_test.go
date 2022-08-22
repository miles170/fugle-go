package fugle

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

func testIntradayServiceMeta(t *testing.T, raw string, want MetaResponse) {
	testIntradayService(t, "meta", raw, &want, func(client *Client) (interface{}, error) {
		resp, err := client.Intrady.Meta("", false)
		return resp, err
	})
}

func TestIntradayService_Meta_2330(t *testing.T) {
	location, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		t.Errorf("time.LoadLocation returned error: %v", err)
	}
	raw := `
	{
		"apiVersion":"0.3.0",
		"data":{
		   "info":{
			  "date":"2022-08-19",
			  "type":"EQUITY",
			  "exchange":"TWSE",
			  "market":"TSE",
			  "symbolId":"2330",
			  "countryCode":"TW",
			  "timeZone":"Asia/Taipei",
			  "lastUpdatedAt":"2022-08-19T13:14:59.185+08:00"
		   },
		   "meta":{
			  "market":"TSE",
			  "nameZhTw":"台積電",
			  "industryZhTw":"半導體業",
			  "previousClose":520,
			  "priceReference":520,
			  "priceHighLimit":572,
			  "priceLowLimit":468,
			  "canDayBuySell":true,
			  "canDaySellBuy":true,
			  "canShortMargin":true,
			  "canShortLend":true,
			  "tradingUnit":1000,
			  "currency":"TWD",
			  "isTerminated":false,
			  "isSuspended":false,
			  "typeZhTw":"一般股票",
			  "abnormal":"正常",
			  "isUnusuallyRecommended":false
		   }
		}
	 }`
	want := MetaResponse{
		APIVersion: "0.3.0",
		Data: MetaData{
			Info: Info{
				Date:          Date{2022, 8, 19},
				Type:          "EQUITY",
				Exchange:      "TWSE",
				Market:        "TSE",
				SymbolID:      "2330",
				CountryCode:   "TW",
				TimeZone:      "Asia/Taipei",
				LastUpdatedAt: time.Date(2022, 8, 19, 13, 14, 59, 185000000, location),
			}, Meta: Meta{
				Market:                 "TSE",
				NameZhTw:               "台積電",
				IndustryZhTw:           "半導體業",
				PreviousClose:          decimal.NewFromInt(520),
				PriceReference:         decimal.NewFromInt(520),
				PriceHighLimit:         decimal.NewFromInt(572),
				PriceLowLimit:          decimal.NewFromInt(468),
				CanDayBuySell:          true,
				CanDaySellBuy:          true,
				CanShortMargin:         true,
				CanShortLend:           true,
				TradingUnit:            1000,
				Currency:               "TWD",
				IsTerminated:           false,
				IsSuspended:            false,
				TypeZhTw:               "一般股票",
				Abnormal:               "正常",
				IsUnusuallyRecommended: false,
			},
		},
	}
	testIntradayServiceMeta(t, raw, want)
}

func TestIntradayService_Meta_IX0001(t *testing.T) {
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
				"lastUpdatedAt": "2022-08-19T07:31:34.909+08:00"
			},
			"meta": {
				"nameZhTw": "發行量加權股價指數",
				"priceReference": 15396.76,
				"typeZhTw": "一般指數",
				"isNewlyCompiled": false
			}
		}
	}`
	want := MetaResponse{
		APIVersion: "0.3.0",
		Data: MetaData{
			Info: Info{
				Date:          Date{2022, 8, 19},
				Type:          "INDEX",
				Exchange:      "TWSE",
				Market:        "TSE",
				SymbolID:      "IX0001",
				CountryCode:   "TW",
				TimeZone:      "Asia/Taipei",
				LastUpdatedAt: time.Date(2022, 8, 19, 07, 31, 34, 909000000, location),
			}, Meta: Meta{
				NameZhTw:        "發行量加權股價指數",
				PriceReference:  decimal.NewFromFloat(15396.76),
				TypeZhTw:        "一般指數",
				IsNewlyCompiled: false,
			},
		},
	}
	testIntradayServiceMeta(t, raw, want)
}

func TestIntradayService_MetaErrors(t *testing.T) {
	testIntradayServiceErros(t, "meta", func(client *Client) (interface{}, error) {
		resp, err := client.Intrady.Meta("", false)
		return resp, err
	})
}
