# fugle-go

[![GoDoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/miles170/fugle-go/fugle)
[![Test Status](https://github.com/miles170/fugle-go/workflows/tests/badge.svg)](https://github.com/miles170/fugle-go/actions?query=workflow%3Atests)
[![Test Coverage](https://codecov.io/gh/miles170/fugle-go/branch/main/graph/badge.svg)](https://codecov.io/gh/miles170/fugle-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/miles170/fugle-go)](https://goreportcard.com/report/github.com/miles170/fugle-go)
[![Code Climate](https://codeclimate.com/github/miles170/fugle-go/badges/gpa.svg)](https://codeclimate.com/github/miles170/fugle-go)

fugle-go is Go library for accessing [Fugle API v0.3](https://developer.fugle.tw/docs/data/intro)

## Installation

fugle-go is compatible with modern Go releases in module mode, with Go installed:

```bash
go get github.com/miles170/fugle-go/fugle
```

## Usage

### HTTP API

```go
import "github.com/miles170/fugle-go/fugle"

client := fugle.NewClient("demo")
```

### META (提供盤中個股/指數當日基本資訊)

```go
meta, err := client.Intrady.Meta("2884", false)
```

### QUOTE (提供盤中個股/指數逐筆交易金額、狀態、最佳五檔及統計資訊)

```go
quote, err := client.Intrady.Quote("2884", false)
```

### CHART (提供盤中個股/指數 線圖時所需的各項即時資訊)

```go
chart, err := client.Intrady.Chart("2884", false)
```

### DEALTS (取得個股當日所有成交資訊)

```go
dealts, err := client.Intrady.Dealts("2884", 50, 0, false)
```

### VOLUMES (提供盤中個股即時分價量)

```go
volumes, err := client.Intrady.Volumes("2884", false)
```

### CANDLES (提供歷史股價資料，包含開高低收量)

```go
candles, err := client.MarketData.Candles("2884",
    time.Date(2022, 8, 15, 0, 0, 0, 0, time.UTC),
    time.Date(2022, 8, 21, 0, 0, 0, 0, time.UTC))
```

## License

[BSD-3-Clause](LICENSE)
