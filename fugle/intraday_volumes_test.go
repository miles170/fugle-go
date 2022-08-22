package fugle

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

func testIntradayServiceVolumes(t *testing.T, raw string, want VolumesResponse) {
	testIntradayService(t, "volumes", raw, &want, func(client *Client) (interface{}, error) {
		resp, err := client.Intrady.Volumes("", false)
		return resp, err
	})
}

func TestIntradayService_Volumes_2330(t *testing.T) {
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
			"volumes": [
				{
					"price": 523,
					"volume": 914,
					"volumeAtBid": 0,
					"volumeAtAsk": 914
				},
				{
					"price": 522,
					"volume": 1985,
					"volumeAtBid": 978,
					"volumeAtAsk": 1007
				},
				{
					"price": 521,
					"volume": 2136,
					"volumeAtBid": 962,
					"volumeAtAsk": 1174
				},
				{
					"price": 520,
					"volume": 2164,
					"volumeAtBid": 1287,
					"volumeAtAsk": 877
				},
				{
					"price": 519,
					"volume": 5485,
					"volumeAtBid": 1246,
					"volumeAtAsk": 3244
				}
			]
		}
	}`
	want := VolumesResponse{
		APIVersion: "0.3.0",
		Data: VolumesData{
			Info: Info{
				Date:          Date{2022, 8, 19},
				Type:          "EQUITY",
				Exchange:      "TWSE",
				Market:        "TSE",
				SymbolID:      "2330",
				CountryCode:   "TW",
				TimeZone:      "Asia/Taipei",
				LastUpdatedAt: time.Date(2022, 8, 19, 13, 30, 0, 0, location),
			},
			Volumes: []Volume{
				{
					Price:       decimal.NewFromInt(523),
					Volume:      914,
					VolumeAtBid: 0,
					VolumeAtAsk: 914,
				},
				{
					Price:       decimal.NewFromInt(522),
					Volume:      1985,
					VolumeAtBid: 978,
					VolumeAtAsk: 1007,
				},
				{
					Price:       decimal.NewFromInt(521),
					Volume:      2136,
					VolumeAtBid: 962,
					VolumeAtAsk: 1174,
				},
				{
					Price:       decimal.NewFromInt(520),
					Volume:      2164,
					VolumeAtBid: 1287,
					VolumeAtAsk: 877,
				},
				{
					Price:       decimal.NewFromInt(519),
					Volume:      5485,
					VolumeAtBid: 1246,
					VolumeAtAsk: 3244,
				},
			},
		},
	}
	testIntradayServiceVolumes(t, raw, want)
}

func TestIntradayService_VolumesErrors(t *testing.T) {
	testIntradayServiceErros(t, "volumes", func(client *Client) (interface{}, error) {
		resp, err := client.Intrady.Volumes("", false)
		return resp, err
	})
}
