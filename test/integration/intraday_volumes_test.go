package integration

import "testing"

func TestIntrady_Volumes(t *testing.T) {
	m, err := client.Intrady.Volumes("2884", false)
	if err != nil {
		t.Fatalf("Intrady.Volumes returned error: %v", err)
	}
	if m.Data.Info.Type != "EQUITY" {
		t.Fatalf("Intrady.Volumes returned type: %s want %s", m.Data.Info.Type, "EQUITY")
	}
	if m.Data.Info.SymbolID != "2884" {
		t.Fatalf("Intrady.Volumes returned symbolId: %s want %s", m.Data.Info.SymbolID, "2884")
	}
	m, err = client.Intrady.Volumes("2884", true)
	if err != nil {
		t.Fatalf("Intrady.Volumes returned error: %v", err)
	}
	if m.Data.Info.Type != "ODDLOT" {
		t.Fatalf("Intrady.Volumes returned type: %s want %s", m.Data.Info.Type, "ODDLOT")
	}
	if m.Data.Info.SymbolID != "2884" {
		t.Fatalf("Intrady.Volumes returned symbolId: %s want %s", m.Data.Info.SymbolID, "2884")
	}
}
