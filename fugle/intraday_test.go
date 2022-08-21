package fugle

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func testIntradayService(t *testing.T, methodName string, raw string, want interface{}, f func(*Client) (interface{}, error)) {
	t.Helper()
	testService(t, fmt.Sprintf("/realtime/v%s/intraday/%s", defaultAPIVersion, methodName), methodName, raw, want, f)
}

func testIntradayServiceErros(t *testing.T, methodName string, f func(*Client) (interface{}, error)) {
	t.Helper()
	testServiceErros(t, fmt.Sprintf("/realtime/v%s/intraday/%s", defaultAPIVersion, methodName), methodName, f)
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
			if test.want.String() != date.String() {
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
			want:      Timestamp{Time: time.Date(2021, 12, 27, 1, 5, 45, 0, time.UTC)},
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
