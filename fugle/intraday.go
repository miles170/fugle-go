package fugle

import "time"

// IntradayService handles communication with the intraday related
// methods of the Fugle API.
//
// Fugle API docs: https://developer.fugle.tw/docs/data/intraday/overview
type IntradayService struct {
	client *Client
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
