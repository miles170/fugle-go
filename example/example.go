package main

import (
	"fmt"

	"github.com/miles170/fugle-go/0.3/fugle"
)

func main() {
	client := fugle.NewClient("demo")
	meta, err := client.Intrady.Meta("2884", nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("date: %v\n", meta.Data.Info.Date)
	fmt.Printf("api version: %s\n", meta.APIVersion)
	fmt.Printf("symbol id: %s\n", meta.Data.Info.SymbolID)
}
