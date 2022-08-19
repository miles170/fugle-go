package integration

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/miles170/fugle-go/0.3/fugle"
)

func TestIntrady_Meta(t *testing.T) {
	m, err := client.Intrady.Meta("2884", nil)
	if err != nil {
		t.Fatalf("Meta returned error: %v", err)
	}
	if m.Data.Info.Type != "EQUITY" {
		t.Fatalf("Intrady.Meta returned type: %s want %s", m.Data.Info.Type, "EQUITY")
	}
	if m.Data.Info.SymbolID != "2884" {
		t.Fatalf("Intrady.Meta returned symbolId: %s want %s", m.Data.Info.SymbolID, "2884")
	}
	m, err = client.Intrady.Meta("2884", &fugle.OddLotOptions{OddLot: true})
	if err != nil {
		t.Fatalf("Meta returned error: %v", err)
	}
	if m.Data.Info.Type != "ODDLOT" {
		t.Fatalf("Intrady.Meta returned type: %s want %s", m.Data.Info.Type, "ODDLOT")
	}
	if m.Data.Info.SymbolID != "2884" {
		t.Fatalf("Intrady.Meta returned symbolId: %s want %s", m.Data.Info.SymbolID, "2884")
	}
}

func testIntradyMetaError(t *testing.T, c *fugle.Client, symbolID string, want fugle.ErrorResponse) {
	_, err := c.Intrady.Meta(symbolID, nil)
	if e, ok := err.(*fugle.ErrorResponse); ok {
		if diff := cmp.Diff(*e, want, cmpopts.IgnoreFields(fugle.ErrorResponse{}, "Response")); diff != "" {
			t.Errorf("Intrady.Meta mismatch (-want +got):\n%s", diff)
		}
	} else {
		t.Errorf("Intrady.Meta returned %v", err)
	}
}

func TestIntrady_Meta_InvalidParameterError(t *testing.T) {
	want := fugle.ErrorResponse{Details: fugle.Error{
		Code:    400,
		Message: "invalid parameters",
	}}
	testIntradyMetaError(t, client, "", want)
}

func TestIntrady_Meta_UnauthorizedError(t *testing.T) {
	want := fugle.ErrorResponse{Details: fugle.Error{
		Code:    401,
		Message: "Unauthorized",
	}}
	testIntradyMetaError(t, unauthorizedClient, "2884", want)
}

func TestIntrady_Meta_ForbiddenError(t *testing.T) {
	want := fugle.ErrorResponse{Details: fugle.Error{
		Code:    403,
		Message: "Forbidden",
	}}
	testIntradyMetaError(t, client, "2884a", want)
}
