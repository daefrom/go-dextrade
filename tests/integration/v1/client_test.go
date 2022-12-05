package tests

import (
	dextrade "github.com/daefrom/go-dextrade/v1"
	"os"
)

var (
	client *dextrade.Client
)

func init() {
	key := os.Getenv("DEXTRADE_API_KEY")
	secret := os.Getenv("DEXTRADE_API_SECRET")
	client = dextrade.NewClient().Auth(key, secret)
}
