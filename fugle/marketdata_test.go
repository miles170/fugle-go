package fugle

import (
	"fmt"
	"testing"
)

func testMarketDataService(t *testing.T, methodName string, raw string, want interface{}, f func(*Client) (interface{}, error)) {
	t.Helper()
	testService(t, fmt.Sprintf("/marketdata/v%s/%s", defaultAPIVersion, methodName), methodName, raw, want, f)
}

func testMarketDataServiceErros(t *testing.T, methodName string, f func(*Client) (interface{}, error)) {
	t.Helper()
	testServiceErros(t, fmt.Sprintf("/marketdata/v%s/%s", defaultAPIVersion, methodName), methodName, f)
}
