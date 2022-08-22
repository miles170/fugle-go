package fugle

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

func testIntradayServiceChart(t *testing.T, raw string, want ChartResponse) {
	testIntradayService(t, "chart", raw, &want, func(client *Client) (interface{}, error) {
		resp, err := client.Intrady.Chart("", false)
		return resp, err
	})
}

func TestIntradayService_Chart_2330(t *testing.T) {
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
			"chart": {
				"a": [
					519.06,
					519.04,
					519.04,
					519.03,
					519.02
				],
				"o": [
					519,
					519,
					519,
					519,
					519
				],
				"h": [
					520,
					519,
					519,
					519,
					519
				],
				"l": [
					518,
					518,
					518,
					518,
					518
				],
				"c": [
					519,
					519,
					519,
					519,
					519
				],
				"v": [
					1289,
					76,
					70,
					203,
					85
				],
				"t": [
					1660870860000,
					1660870920000,
					1660870980000,
					1660871040000,
					1660871100000
				]
			}
		}
	}`
	want := ChartResponse{
		APIVersion: "0.3.0",
		Data: ChartData{
			Info: Info{
				Date:          InfoDate{2022, 8, 19},
				Type:          "EQUITY",
				Exchange:      "TWSE",
				Market:        "TSE",
				SymbolID:      "2330",
				CountryCode:   "TW",
				TimeZone:      "Asia/Taipei",
				LastUpdatedAt: time.Date(2022, 8, 19, 13, 30, 0, 0, location),
			},
			Chart: Chart{
				Averages: []decimal.Decimal{
					decimal.NewFromFloat(519.06),
					decimal.NewFromFloat(519.04),
					decimal.NewFromFloat(519.04),
					decimal.NewFromFloat(519.03),
					decimal.NewFromFloat(519.02),
				},
				Opens: []decimal.Decimal{
					decimal.NewFromFloat(519),
					decimal.NewFromFloat(519),
					decimal.NewFromFloat(519),
					decimal.NewFromFloat(519),
					decimal.NewFromFloat(519),
				},
				Highs: []decimal.Decimal{
					decimal.NewFromFloat(520),
					decimal.NewFromFloat(519),
					decimal.NewFromFloat(519),
					decimal.NewFromFloat(519),
					decimal.NewFromFloat(519),
				},
				Lows: []decimal.Decimal{
					decimal.NewFromFloat(518),
					decimal.NewFromFloat(518),
					decimal.NewFromFloat(518),
					decimal.NewFromFloat(518),
					decimal.NewFromFloat(518),
				},
				Closes: []decimal.Decimal{
					decimal.NewFromFloat(519),
					decimal.NewFromFloat(519),
					decimal.NewFromFloat(519),
					decimal.NewFromFloat(519),
					decimal.NewFromFloat(519),
				},
				Volumes: []int{1289, 76, 70, 203, 85},
				Timestamps: []Timestamp{
					{time.Date(2022, 8, 19, 9, 1, 0, 0, location)},
					{time.Date(2022, 8, 19, 9, 2, 0, 0, location)},
					{time.Date(2022, 8, 19, 9, 3, 0, 0, location)},
					{time.Date(2022, 8, 19, 9, 4, 0, 0, location)},
					{time.Date(2022, 8, 19, 9, 5, 0, 0, location)},
				},
			},
		},
	}
	testIntradayServiceChart(t, raw, want)
}

func TestIntradayService_Chart_IX0001(t *testing.T) {
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
			"chart": {
				"o": [
					15394.36,
					15382.25,
					15379.91,
					15377.69,
					15381.13
				],
				"h": [
					15394.36,
					15382.42,
					15381.14,
					15380.08,
					15381.13
				],
				"l": [
					15371.83,
					15369.69,
					15369.09,
					15376.38,
					15370.94
				],
				"c": [
					15382.85,
					15379.37,
					15369.09,
					15380.08,
					15380.52
				],
				"v": [
					7420596020,
					2793601820,
					3110558330,
					2677265310,
					2270303390
				],
				"t": [
					1660870860000,
					1660870920000,
					1660870980000,
					1660871040000,
					1660871100000
				]
			}
		}
	}`
	want := ChartResponse{
		APIVersion: "0.3.0",
		Data: ChartData{
			Info: Info{
				Date:          InfoDate{2022, 8, 19},
				Type:          "INDEX",
				Exchange:      "TWSE",
				Market:        "TSE",
				SymbolID:      "IX0001",
				CountryCode:   "TW",
				TimeZone:      "Asia/Taipei",
				LastUpdatedAt: time.Date(2022, 8, 19, 13, 30, 0, 0, location),
			},
			Chart: Chart{
				Opens: []decimal.Decimal{
					decimal.NewFromFloat(15394.36),
					decimal.NewFromFloat(15382.25),
					decimal.NewFromFloat(15379.91),
					decimal.NewFromFloat(15377.69),
					decimal.NewFromFloat(15381.13),
				},
				Highs: []decimal.Decimal{
					decimal.NewFromFloat(15394.36),
					decimal.NewFromFloat(15382.42),
					decimal.NewFromFloat(15381.14),
					decimal.NewFromFloat(15380.08),
					decimal.NewFromFloat(15381.13),
				},
				Lows: []decimal.Decimal{
					decimal.NewFromFloat(15371.83),
					decimal.NewFromFloat(15369.69),
					decimal.NewFromFloat(15369.09),
					decimal.NewFromFloat(15376.38),
					decimal.NewFromFloat(15370.94),
				},
				Closes: []decimal.Decimal{
					decimal.NewFromFloat(15382.85),
					decimal.NewFromFloat(15379.37),
					decimal.NewFromFloat(15369.09),
					decimal.NewFromFloat(15380.08),
					decimal.NewFromFloat(15380.52),
				},
				Volumes: []int{
					7420596020,
					2793601820,
					3110558330,
					2677265310,
					2270303390},
				Timestamps: []Timestamp{
					{time.Date(2022, 8, 19, 9, 1, 0, 0, location)},
					{time.Date(2022, 8, 19, 9, 2, 0, 0, location)},
					{time.Date(2022, 8, 19, 9, 3, 0, 0, location)},
					{time.Date(2022, 8, 19, 9, 4, 0, 0, location)},
					{time.Date(2022, 8, 19, 9, 5, 0, 0, location)},
				},
			},
		},
	}
	testIntradayServiceChart(t, raw, want)
}

func TestIntradayService_ChartErrors(t *testing.T) {
	testIntradayServiceErros(t, "chart", func(client *Client) (interface{}, error) {
		resp, err := client.Intrady.Chart("", false)
		return resp, err
	})
}
