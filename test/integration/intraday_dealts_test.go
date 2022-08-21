package integration

import "testing"

func TestIntrady_Dealts(t *testing.T) {
	m, err := client.Intrady.Dealts("2884", 50, 0, false)
	if err != nil {
		t.Fatalf("Intrady.Dealts returned error: %v", err)
	}
	if m.Data.Info.Type != "EQUITY" {
		t.Fatalf("Intrady.Dealts returned type: %s want %s", m.Data.Info.Type, "EQUITY")
	}
	if m.Data.Info.SymbolID != "2884" {
		t.Fatalf("Intrady.Dealts returned symbolId: %s want %s", m.Data.Info.SymbolID, "2884")
	}
	m, err = client.Intrady.Dealts("2884", 50, 0, true)
	if err != nil {
		t.Fatalf("Intrady.Dealts returned error: %v", err)
	}
	if m.Data.Info.Type != "ODDLOT" {
		t.Fatalf("Intrady.Dealts returned type: %s want %s", m.Data.Info.Type, "ODDLOT")
	}
	if m.Data.Info.SymbolID != "2884" {
		t.Fatalf("Intrady.Dealts returned symbolId: %s want %s", m.Data.Info.SymbolID, "2884")
	}
}
