package fugle

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

// setup sets up a test HTTP server along with a fugle.Client that is
// configured to talk to that test server. Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup() (client *Client, mux *http.ServeMux, teardown func()) {
	// mux is the HTTP request multiplexer used with the test server.
	mux = http.NewServeMux()

	// server is a test HTTP server used to provide mock API responses.
	server := httptest.NewServer(mux)

	// client is the fugle client being tested and is
	// configured to use test server.
	client = NewClient("")
	url, _ := url.Parse(server.URL + "/")
	client.baseURL = url

	return client, mux, server.Close
}

func testMethod(t *testing.T, r *http.Request, want string) {
	t.Helper()
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

// Test how bad options are handled. Method f under test should
// return an error.
func testBadOptions(t *testing.T, methodName string, f func() error) {
	t.Helper()
	if methodName == "" {
		t.Error("testBadOptions: must supply method methodName")
	}
	if err := f(); err == nil {
		t.Errorf("bad options %v err = nil, want error", methodName)
	}
}

// Test function under NewRequest failure and then s.client.Do failure.
// Method f should be a regular call that would normally succeed, but
// should return an error when NweRequest or s.client.Do fails.
func testNewRequestAndDoFailure(t *testing.T, methodName string, client *Client, f func() error) {
	t.Helper()
	if methodName == "" {
		t.Error("testNewRequestAndDoFailure: must supply method methodName")
	}

	client.baseURL.Path = ""
	err := f()
	if err == nil {
		t.Errorf("client.baseURL.Path='' %v err = nil, want error", methodName)
	}
}

func testURLParseError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err, ok := err.(*url.Error); !ok || err.Op != "parse" {
		t.Errorf("Expected URL parse error, got %+v", err)
	}
}

func testErrorContains(t *testing.T, e error, want string) {
	t.Helper()
	if !strings.Contains(e.Error(), want) {
		t.Errorf("testErrorContains: err message = %s, want %s", e.Error(), want)
	}
}

func testService(t *testing.T, pattern string, methodName string, raw string, want interface{}, f func(*Client) (interface{}, error)) {
	t.Helper()
	if methodName == "" {
		t.Error("testService: must supply method methodName")
	}

	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, raw)
	})

	resp, err := f(client)
	if err != nil {
		t.Errorf("%s returned error: %v", methodName, err)
	}
	if !cmp.Equal(resp, want, cmp.AllowUnexported(InfoDate{})) {
		t.Errorf("%s returned %v, want %v", methodName, resp, want)
	}
	testNewRequestAndDoFailure(t, methodName, client, func() error {
		_, err = f(client)
		return err
	})
	testBadOptions(t, methodName, func() error {
		client.apiVersion = "\n"
		_, err = f(client)
		return err
	})
}

func testServiceError(t *testing.T, pattern string, methodName string, statusCode int, raw string, want ErrorResponse, f func(*Client) (interface{}, error)) {
	t.Helper()
	if methodName == "" {
		t.Error("testServiceError: must supply method methodName")
	}

	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.WriteHeader(statusCode)
		fmt.Fprint(w, raw)
	})

	_, err := f(client)
	if e, ok := err.(*ErrorResponse); ok {
		testErrorContains(t, e, pattern)
		testErrorContains(t, &e.Details, want.Details.Message)
		if !cmp.Equal(*e, want, cmpopts.IgnoreFields(ErrorResponse{}, "Response")) {
			t.Errorf("%s returned %v, want %v", methodName, *e, want)
		}
	} else {
		t.Errorf("%s returned %v", methodName, err)
	}
}

func testServiceInvalidParameterError(t *testing.T, pattern string, methodName string, f func(*Client) (interface{}, error)) {
	t.Helper()
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
	testServiceError(t, pattern, methodName, http.StatusBadRequest, raw, want, f)
}

func testServiceUnauthorizedError(t *testing.T, pattern string, methodName string, f func(*Client) (interface{}, error)) {
	t.Helper()
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
	testServiceError(t, pattern, methodName, http.StatusUnauthorized, raw, want, f)
}

func testServiceForbiddenError(t *testing.T, pattern string, methodName string, f func(*Client) (interface{}, error)) {
	t.Helper()
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
	testServiceError(t, pattern, methodName, http.StatusForbidden, raw, want, f)
}

func testServiceNotFoundError(t *testing.T, pattern string, methodName string, f func(*Client) (interface{}, error)) {
	t.Helper()
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
	testServiceError(t, pattern, methodName, http.StatusNotFound, raw, want, f)
}

func testServiceUnmarshalError(t *testing.T, pattern string, f func(*Client) (interface{}, error)) {
	t.Helper()
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{\n}`)
	})

	_, err := f(client)
	if e, ok := err.(*ErrorResponse); ok {
		if e.Details.Code != 0 || e.Details.Message != "" {
			t.Errorf("Intrady.Meta returned %v, want nil", *e)
		}
	} else {
		t.Errorf("Intrady.Meta returned %v", err)
	}
}

func testServiceErros(t *testing.T, pattern string, methodName string, f func(*Client) (interface{}, error)) {
	t.Helper()
	testServiceUnmarshalError(t, pattern, f)
	testServiceInvalidParameterError(t, pattern, methodName, f)
	testServiceUnauthorizedError(t, pattern, methodName, f)
	testServiceForbiddenError(t, pattern, methodName, f)
	testServiceNotFoundError(t, pattern, methodName, f)
}

func TestAddOptions_QueryValues(t *testing.T) {
	if _, err := addOptions("yo", ""); err == nil {
		t.Error("addOptions err = nil, want error")
	}
	if _, err := addOptions("yo", (*IntradyOptions)(nil)); err != nil {
		t.Errorf("addOptions returned %v, want nil", err)
	}
}

func TestNewRequest_BadURL(t *testing.T) {
	c := NewClient("")
	_, err := c.NewRequest("GET", ":")
	testURLParseError(t, err)
}

func TestNewRequest_BadMethod(t *testing.T) {
	c := NewClient("")
	if _, err := c.NewRequest("BOGUS\nMETHOD", "."); err == nil {
		t.Fatal("NewRequest returned nil; expected error")
	}
}

// ensure that no User-Agent header is set if the client's UserAgent is empty.
func TestNewRequest_EmptyUserAgent(t *testing.T) {
	c := NewClient("")
	c.userAgent = ""
	req, err := c.NewRequest("GET", ".")
	if err != nil {
		t.Fatalf("NewRequest returned unexpected error: %v", err)
	}
	if _, ok := req.Header["User-Agent"]; ok {
		t.Fatal("constructed request contains unexpected User-Agent header")
	}
}

func TestNewRequest_ErrorForNoTrailingSlash(t *testing.T) {
	tests := []struct {
		rawurl    string
		wantError bool
	}{
		{rawurl: "https://example.com/api/v3", wantError: true},
		{rawurl: "https://example.com/api/v3/", wantError: false},
	}
	c := NewClient("")
	for _, test := range tests {
		u, err := url.Parse(test.rawurl)
		if err != nil {
			t.Fatalf("url.Parse returned unexpected error: %v.", err)
		}
		c.baseURL = u
		if _, err := c.NewRequest(http.MethodGet, "test"); test.wantError && err == nil {
			t.Fatalf("Expected error to be returned.")
		} else if !test.wantError && err != nil {
			t.Fatalf("NewRequest returned unexpected error: %v.", err)
		}
	}
}

func TestDo_BadRequestURL(t *testing.T) {
	client, _, teardown := setup()
	defer teardown()

	req, err := client.NewRequest("GET", "test-url")
	if err != nil {
		t.Fatalf("client.NewRequest returned error: %v", err)
	}

	req.URL = nil
	resp, err := client.Do(req, nil)
	if resp != nil {
		t.Errorf("client.Do resp = %#v, want nil", resp)
	}
	if err == nil {
		t.Error("client.Do err = nil, want error")
	}
}
