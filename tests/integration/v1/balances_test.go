package tests

import (
	"testing"
)

func TestBalancesServiceGet(t *testing.T) {

	balances, err := client.Balances.GetBalances()

	if err != nil {
		t.Error(err)
	}
	if false == balances.Status {
		t.Error("Expected", 1)
	}
}
