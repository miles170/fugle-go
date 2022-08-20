package integration

import "testing"

func TestIntrady_Quote(t *testing.T) {
	m, err := client.Intrady.Quote("2884", false)
	if err != nil {
		t.Fatalf("Intrady.Quote returned error: %v", err)
	}
	if m.Data.Info.Type != "EQUITY" {
		t.Fatalf("Intrady.Quote returned type: %s want %s", m.Data.Info.Type, "EQUITY")
	}
	if m.Data.Info.SymbolID != "2884" {
		t.Fatalf("Intrady.Quote returned symbolId: %s want %s", m.Data.Info.SymbolID, "2884")
	}
	m, err = client.Intrady.Quote("2884", true)
	if err != nil {
		t.Fatalf("Intrady.Quote returned error: %v", err)
	}
	if m.Data.Info.Type != "ODDLOT" {
		t.Fatalf("Intrady.Quote returned type: %s want %s", m.Data.Info.Type, "ODDLOT")
	}
	if m.Data.Info.SymbolID != "2884" {
		t.Fatalf("Intrady.Quote returned symbolId: %s want %s", m.Data.Info.SymbolID, "2884")
	}
}
