package fugle

import (
	"encoding/json"
	"time"
)

// IntradayService handles communication with the intraday related
// methods of the Fugle API.
//
// Fugle API docs: https://developer.fugle.tw/docs/data/intraday/overview
type IntradayService struct {
	client *Client
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
	p.Time = time.Unix(i/1000, i%1000).In(time.UTC)
	return nil
}

type InfoDate time.Time

// UnmarshalJSON handles incoming JSON.
func (d *InfoDate) UnmarshalJSON(b []byte) error {
	t, err := time.Parse("\"2006-01-02\"", string(b))
	if err != nil {
		return err
	}
	*d = InfoDate(t)
	return nil
}

type Info struct {
	Date          InfoDate  `json:"date"`
	Type          string    `json:"type"`
	Exchange      string    `json:"exchange"`
	Market        string    `json:"market"`
	SymbolID      string    `json:"symbolId"`
	CountryCode   string    `json:"countryCode"`
	TimeZone      string    `json:"timeZone"`
	LastUpdatedAt time.Time `json:"lastUpdatedAt"`
}
