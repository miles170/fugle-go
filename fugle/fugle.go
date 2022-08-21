package fugle

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/google/go-querystring/query"
)

const (
	Version = "0.3.1"

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
	Intrady    *IntradayService
	MarketData *MarketDataService
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
	c.MarketData = &MarketDataService{client: c}
	return c
}

// NewRequest creates an API request.
func (c *Client) NewRequest(method, urlStr string) (*http.Request, error) {
	if !strings.HasSuffix(c.baseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.baseURL)
	}

	u, err := c.baseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}

	if c.userAgent != "" {
		req.Header.Set("User-Agent", c.userAgent)
	}
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

// Call first creates an API request then sends it and returns the API response.
func (c *Client) Call(url string, opts interface{}, resp interface{}) error {
	url, err := addOptions(url, opts)
	if err != nil {
		return err
	}

	req, err := c.NewRequest("GET", url)
	if err != nil {
		return err
	}

	_, err = c.Do(req, &resp)
	if err != nil {
		return err
	}
	return nil
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
