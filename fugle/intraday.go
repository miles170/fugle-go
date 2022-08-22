package fugle

import (
	"encoding/json"
	"fmt"
	"time"
)

// IntradayService handles communication with the intraday related
// methods of the Fugle API.
//
// Fugle API docs: https://developer.fugle.tw/docs/data/intraday/overview
type IntradayService struct {
	client *Client
}

type IntradyOptions struct {
	SymbolID string `url:"symbolId"` // 個股、指數識別代碼
	APIToken string `url:"apiToken"`
	OddLot   bool   `url:"oddLot"` // 是否回傳零股行情
}

type Timestamp struct {
	time.Time
}

// UnmarshalJSON handles incoming JSON.
func (p *Timestamp) UnmarshalJSON(bytes []byte) error {
	var i int64
	err := json.Unmarshal(bytes, &i)
	if err != nil {
		return err
	}
	// fugle api returns Unix timestamp in milliseconds
	p.Time = time.Unix(i/1000, (i % 1000 * 1000000)).In(time.UTC)
	return nil
}

type Date struct {
	Year  int
	Month time.Month
	Day   int
}

func (d Date) String() string {
	return fmt.Sprintf("%04d-%02d-%02d", d.Year, d.Month, d.Day)
}

// UnmarshalJSON handles incoming JSON.
func (d *Date) UnmarshalJSON(b []byte) error {
	t, err := time.Parse("\"2006-01-02\"", string(b))
	if err != nil {
		return err
	}
	year, month, day := t.Date()
	d.Year = year
	d.Month = month
	d.Day = day
	return nil
}

type Info struct {
	Date          Date      `json:"date"`          // 本筆資料所屬日期
	Type          string    `json:"type"`          // ticker 類別
	Exchange      string    `json:"exchange"`      // 交易所
	Market        string    `json:"market"`        // 市場別
	SymbolID      string    `json:"symbolId"`      // 股票代號
	CountryCode   string    `json:"countryCode"`   // 股票所屬國家ISO2代碼
	TimeZone      string    `json:"timeZone"`      // 股票所屬時區
	LastUpdatedAt time.Time `json:"lastUpdatedAt"` // 本筆資料最後更新時間
}
