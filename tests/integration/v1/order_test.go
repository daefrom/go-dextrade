package tests

import (
	"testing"
)

func TestOrderCreateAndCancel(t *testing.T) {
	order, err := client.Order.Create("DAEUSDT", 3, 0.0043, "sell", "0")

	if err != nil {
		t.Fatalf("Order.Create() returned error: %v", err)
	}
	orderId := order.Data.Id
	_, err = client.Order.GetOrder(orderId)
	if err != nil {
		t.Logf("Order id: %d is not canceled!", orderId)
		t.Fatalf("Order.Trades() returned error: %v", err)
	}
	err = client.Order.Cancel(orderId)

	if err != nil {
		t.Fatalf("Order.Cancel() returned error: %v", err)
	}
}

func TestGetOpenOrders(t *testing.T) {
	openOrdersResponse, err := client.Order.GetOpenOrders()

	if err != nil {
		t.Fatalf("Order.GetOpenOrders() returned error: %v", err)
	}
	if openOrdersResponse.Status {
		t.Logf("orders : %v", len(openOrdersResponse.Data.List))
	}

}
