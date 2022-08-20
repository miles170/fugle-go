package fugle

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func testIntradayService(t *testing.T, methodName string, raw string, want interface{}, f func(*Client) (interface{}, error)) {
	t.Helper()
	if methodName == "" {
		t.Error("testIntradayService: must supply method methodName")
	}

	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/realtime/v%s/intraday/%s", client.apiVersion, methodName), func(w http.ResponseWriter, r *http.Request) {
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

func testIntradayServiceError(t *testing.T, methodName string, statusCode int, raw string, want ErrorResponse, f func(*Client) (interface{}, error)) {
	t.Helper()
	if methodName == "" {
		t.Error("testIntradayService: must supply method methodName")
	}

	client, mux, teardown := setup()
	defer teardown()

	url := fmt.Sprintf("/realtime/v%s/intraday/%s", client.apiVersion, methodName)
	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.WriteHeader(statusCode)
		fmt.Fprint(w, raw)
	})

	_, err := f(client)
	if e, ok := err.(*ErrorResponse); ok {
		testErrorContains(t, e, url)
		testErrorContains(t, &e.Details, want.Details.Message)
		if !cmp.Equal(*e, want, cmpopts.IgnoreFields(ErrorResponse{}, "Response")) {
			t.Errorf("%s returned %v, want %v", methodName, *e, want)
		}
	} else {
		t.Errorf("%s returned %v", methodName, err)
	}
}

func testIntradayServiceInvalidParameterError(t *testing.T, methodName string, f func(*Client) (interface{}, error)) {
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
	testIntradayServiceError(t, methodName, http.StatusBadRequest, raw, want, f)
}

func testIntradayServiceUnauthorizedError(t *testing.T, methodName string, f func(*Client) (interface{}, error)) {
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
	testIntradayServiceError(t, methodName, http.StatusUnauthorized, raw, want, f)
}

func testIntradayServiceForbiddenError(t *testing.T, methodName string, f func(*Client) (interface{}, error)) {
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
	testIntradayServiceError(t, methodName, http.StatusForbidden, raw, want, f)
}

func testIntradayServiceNotFoundError(t *testing.T, methodName string, f func(*Client) (interface{}, error)) {
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
	testIntradayServiceError(t, methodName, http.StatusNotFound, raw, want, f)
}

func testIntradayServiceMetaUnmarshalError(t *testing.T, methodName string, f func(*Client) (interface{}, error)) {
	t.Helper()
	client, mux, teardown := setup()
	defer teardown()

	url := fmt.Sprintf("/realtime/v%s/intraday/%s", client.apiVersion, methodName)
	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
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

func testIntradayServiceErros(t *testing.T, methodName string, f func(*Client) (interface{}, error)) {
	t.Helper()
	testIntradayServiceMetaUnmarshalError(t, methodName, f)
	testIntradayServiceInvalidParameterError(t, methodName, f)
	testIntradayServiceUnauthorizedError(t, methodName, f)
	testIntradayServiceForbiddenError(t, methodName, f)
	testIntradayServiceNotFoundError(t, methodName, f)
}

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
				t.Errorf("InfoDate.UnmarshalJSON returned an error when we expected nil")
			}
			if err == nil && test.wantError {
				t.Errorf("InfoDate.UnmarshalJSON returned no error when we expected one")
			}
			if !cmp.Equal(test.want, date, cmp.AllowUnexported(InfoDate{})) {
				t.Errorf("InfoDate.UnmarshalJSON expected date %v, got %v", test.want, date)
			}
		})
	}
}

func TestTimestamp_UnmarshalJSON(t *testing.T) {
	var testCases = map[string]struct {
		data      []byte
		want      Timestamp
		wantError bool
	}{
		"Valid": {
			data:      []byte(`1640567145000`),
			want:      Timestamp{Time: time.Date(2021, 12, 27, 9, 5, 45, 0, time.Local)},
			wantError: false,
		},
		"Invalid": {
			data:      []byte(`\n`),
			want:      Timestamp{},
			wantError: true,
		},
	}
	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			date := Timestamp{}
			err := date.UnmarshalJSON(test.data)
			if err != nil && !test.wantError {
				t.Errorf("Timestamp.UnmarshalJSON returned an error when we expected nil")
			}
			if err == nil && test.wantError {
				t.Errorf("Timestamp.UnmarshalJSON returned no error when we expected one")
			}
			if !cmp.Equal(test.want, date) {
				t.Errorf("Timestamp.UnmarshalJSON expected date %v, got %v", test.want, date)
			}
		})
	}
}
