package dextrade

import (
	"testing"
)

func TestTradesServiceGet(t *testing.T) {

	client := NewClient()

	trades, err := client.Trades.Get("DAEUSDT")

	if err != nil {
		t.Error(err)
	}
	if len(trades.Data) == 0 {
		t.Error("Expected", 1)
	}
}
