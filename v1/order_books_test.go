package dextrade

import (
	"testing"
)

func TestOrderBooksServiceGet(t *testing.T) {

	client := NewClient()

	orderBooks, err := client.OrderBooks.Get("DAEUSDT")

	if err != nil {
		t.Error(err)
	}
	if false == orderBooks.Status {
		t.Error("Expected", 1)
	}
}
