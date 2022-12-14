package integration

import "github.com/miles170/fugle-go/fugle"

var (
	client             *fugle.Client
	unauthorizedClient *fugle.Client
)

func init() {
	client = fugle.NewClient("demo")
	unauthorizedClient = fugle.NewClient("")
}
