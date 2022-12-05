package dextrade

import (
	"testing"
)

func TestTickerServiceGet(t *testing.T) {

	client := NewClient()

	_, err := client.Ticker.Get("DAEUSDT")

	if err != nil {
		t.Error(err)
	}
}
