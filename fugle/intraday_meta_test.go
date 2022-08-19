package fugle

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/shopspring/decimal"
)

func TestInfoDate_UnmarshalJSON(t *testing.T) {
	var testCases = map[string]struct {
		data      []byte
		want      InfoDate
		wantError bool
	}{
		"Valid": {
			data:      []byte(`"2022-08-19"`),
			want:      InfoDate(time.Date(2022, 8, 19, 0, 0, 0, 0, time.UTC)),
			wantError: false,
		},
		"Invalid": {
			data:      []byte(`"022-08-19`),
			want:      InfoDate{},
			wantError: true,
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			date := InfoDate{}
			err := date.UnmarshalJSON(test.data)
			if err != nil && !test.wantError {
				t.Errorf("RequiredReviewer.UnmarshalJSON returned an error when we expected nil")
			}
			if err == nil && test.wantError {
				t.Errorf("RequiredReviewer.UnmarshalJSON returned no error when we expected one")
			}
			if !cmp.Equal(test.want, date, cmp.AllowUnexported(InfoDate{})) {
				t.Errorf("RequiredReviewer.UnmarshalJSON expected date %v, got %v", test.want, date)
			}
		})
	}
}

func testIntradayServiceMeta(t *testing.T, raw string, want interface{}) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/realtime/v%s/intraday/meta", client.apiVersion), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, raw)
	})

	meta, err := client.Intrady.Meta("2330", nil)
	if err != nil {
		t.Errorf("Intrady.Meta returned error: %v", err)
	}
	if !cmp.Equal(*meta, want, cmp.AllowUnexported(InfoDate{})) {
		t.Errorf("Intrady.Meta returned %v, want %v", *meta, want)
	}
	const methodName = "Meta"
	testBadOptions(t, methodName, func() (err error) {
		client.apiVersion = "\n"
		_, err = client.Intrady.Meta("2330", nil)
		return err
	})
}

func TestIntradayService_Meta_2330(t *testing.T) {
	location, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		t.Errorf("time.LoadLocation returned error: %v", err)
	}
	date := time.Date(2022, 8, 19, 13, 14, 59, 185000000, location)
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
				Date:          InfoDate(time.Date(2022, 8, 19, 0, 0, 0, 0, time.UTC)),
				Type:          "EQUITY",
				Exchange:      "TWSE",
				Market:        "TSE",
				SymbolID:      "2330",
				CountryCode:   "TW",
				TimeZone:      "Asia/Taipei",
				LastUpdatedAt: &date,
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
	date := time.Date(2022, 8, 19, 07, 31, 34, 909000000, location)
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
				Date:          InfoDate(time.Date(2022, 8, 19, 0, 0, 0, 0, time.UTC)),
				Type:          "INDEX",
				Exchange:      "TWSE",
				Market:        "TSE",
				SymbolID:      "IX0001",
				CountryCode:   "TW",
				TimeZone:      "Asia/Taipei",
				LastUpdatedAt: &date,
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

func testIntradayServiceMetaError(t *testing.T, statusCode int, raw string, want ErrorResponse) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/realtime/v%s/intraday/meta", client.apiVersion), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.WriteHeader(statusCode)
		fmt.Fprint(w, raw)
	})

	_, err := client.Intrady.Meta("", nil)
	if e, ok := err.(*ErrorResponse); ok {
		if !cmp.Equal(*e, want, cmpopts.IgnoreFields(ErrorResponse{}, "Response")) {
			t.Errorf("Intrady.Meta returned %v, want %v", *e, want)
		}
	} else {
		t.Errorf("Intrady.Meta returned %v", err)
	}
}

func TestIntradayService_Meta_InvalidParameterError(t *testing.T) {
	raw := `
	{
		"apiVersion": "0.3.0",
		"error": {
			"code": 400,
			"message": "invalid parameters"
		}
	}`
	want := ErrorResponse{Details: Error{
		Code:    400,
		Message: "invalid parameters",
	}}
	testIntradayServiceMetaError(t, http.StatusBadRequest, raw, want)
}

func TestIntradayService_Meta_UnauthorizedError(t *testing.T) {
	raw := `
	{
		"apiVersion": "0.3.0",
		"error": {
			"code": 401,
			"message": "Unauthorized"
		}
	}`
	want := ErrorResponse{Details: Error{
		Code:    401,
		Message: "Unauthorized",
	}}
	testIntradayServiceMetaError(t, http.StatusUnauthorized, raw, want)
}

func TestIntradayService_Meta_ForbiddenError(t *testing.T) {
	raw := `
	{
		"apiVersion": "0.3.0",
		"error": {
			"code": 403,
			"message": "Forbidden"
		}
	}`
	want := ErrorResponse{Details: Error{
		Code:    403,
		Message: "Forbidden",
	}}
	testIntradayServiceMetaError(t, http.StatusForbidden, raw, want)
}

func TestIntradayService_Meta_NotFoundError(t *testing.T) {
	raw := `
	{
		"apiVersion": "0.3.0",
		"error": {
			"code": 404,
			"message": "Resource Not Found"
		}
	}`
	want := ErrorResponse{Details: Error{
		Code:    404,
		Message: "Resource Not Found",
	}}
	testIntradayServiceMetaError(t, http.StatusNotFound, raw, want)
}
