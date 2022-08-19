package fugle

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/google/go-querystring/query"
)

const (
	Version = "0.1.0"

	defaultBaseURL    = "https://api.fugle.tw/"
	defaultUserAgent  = "fugle-go" + "/" + Version
	defaultAPIVersion = "0.3"
)

// A Client manages communication with the Fugle API.
type Client struct {
	// HTTP client used to communicate with the API.
	client *http.Client

	// base URL for API requests.
	baseURL *url.URL

	// Api token used when communicating with the Fugle API.
	apiToken string

	// Api version used when communicating with the Fugle API.
	apiVersion string

	// User agent used when communicating with the Fugle API.
	userAgent string

	// Services used for talking to different parts of the API
	Intrady *IntradayService
}

type BasicOptions struct {
	SymbolID string `url:"symbolId"`
	APIToken string `url:"apiToken"`
}

type OddLotOptions struct {
	OddLot bool `url:"oddLot"`
}

// addOptions adds the parameters in opts as URL query parameters to s. opts
// must be a struct whose fields may contain "url" tags.
func addOptions(s string, opts interface{}) (string, error) {
	v := reflect.ValueOf(opts)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opts)
	if err != nil {
		return s, err
	}

	q := u.Query()
	for k, v := range qs {
		q[k] = v
	}
	u.RawQuery = q.Encode()
	return u.String(), nil
}

// NewClient returns a new Fugle API client.
func NewClient(apiToken string) *Client {
	httpClient := &http.Client{}
	baseURL, _ := url.Parse(defaultBaseURL)
	c := &Client{
		client:     httpClient,
		baseURL:    baseURL,
		apiToken:   apiToken,
		apiVersion: defaultAPIVersion,
		userAgent:  defaultUserAgent,
	}
	c.Intrady = &IntradayService{client: c}
	return c
}

// NewRequest creates an API request.
func (c *Client) NewRequest(method, urlStr string) (*http.Request, error) {
	u, err := c.baseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", c.userAgent)
	return req, nil
}

// Do sends an API request and returns the API response.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	err = CheckResponse(resp)
	if err != nil {
		return resp, err
	}

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}
	return resp, err
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
	Date          InfoDate   `json:"date"`
	Type          string     `json:"type"`
	Exchange      string     `json:"exchange"`
	Market        string     `json:"market"`
	SymbolID      string     `json:"symbolId"`
	CountryCode   string     `json:"countryCode"`
	TimeZone      string     `json:"timeZone"`
	LastUpdatedAt *time.Time `json:"lastUpdatedAt,omitempty"` // (Optional.)
}

type ErrorResponse struct {
	Response *http.Response // HTTP response that caused this error
	Details  Error          `json:"error"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Details.Message)
}

type Error struct {
	Code    int    `json:"code"`    // error code
	Message string `json:"message"` // error message
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d error: %s", e.Code, e.Message)
}

// CheckResponse checks the API response for errors
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	errorResponse := &ErrorResponse{Response: r}
	data, err := io.ReadAll(r.Body)
	if err == nil && data != nil {
		err = json.Unmarshal(data, errorResponse)
		if err != nil {
			return errorResponse
		}
	}
	return errorResponse
}
