package dextrade

import (
	"errors"
	"strconv"
)

type OrderService struct {
	c *Client
}

type Order struct {
	Id int
}

type OrderCreateResponse struct {
	Status  bool
	Message string
	Error   string
	Data    Order
}

type OrderCancelResponse struct {
	Status  bool
	Message string
	Error   string
}

type OrderDataResponse struct {
	Id         int         `json:"id"`
	Type       int         `json:"type"`
	Status     int         `json:"status"`
	TypeTrade  int         `json:"type_trade"`
	Pair       string      `json:"pair"`
	Volume     float64     `json:"volume"`
	VolumeDone int         `json:"volume_done"`
	Rate       float64     `json:"rate"`
	Price      float64     `json:"price"`
	PriceDone  float64     `json:"price_done"`
	TimeCreate int         `json:"time_create"`
	TimeDone   interface{} `json:"time_done"`
	Commission interface{} `json:"commission"`
}

type OrderResponse struct {
	Status bool
	Error  string
	Data   OrderDataResponse
}

type OpenOrdersResponse struct {
	Status bool
	Error  string
	Data   OpenOrderDataResponse
}

type OpenOrderDataResponse struct {
	List []OpenOrderListResponse
}

type OpenOrderListResponse struct {
	Id         int         `json:"id"`
	Type       int         `json:"type"`
	Status     int         `json:"status"`
	TypeTrade  int         `json:"type_trade"`
	Pair       string      `json:"pair"`
	Volume     string      `json:"volume"`
	VolumeDone interface{} `json:"volume_done"`
	Rate       string      `json:"rate"`
	Price      string      `json:"price"`

	TimeCreate int         `json:"time_create"`
	TimeDone   interface{} `json:"time_done"`
	PriceDone  interface{} `json:"price_done"`
}

//Create OrderService creates order and return orderId
func (a *OrderService) Create(pair string, quantity float64, price float64, orderType string, typeTrade string) (OrderCreateResponse, error) {

	if orderType == "buy" {
		orderType = "0"
	} else {
		orderType = "1"
	}
	params := make(map[string]string)
	params["pair"] = pair
	params["volume"] = strconv.FormatFloat(quantity, 'f', 6, 64)
	params["rate"] = strconv.FormatFloat(price, 'f', 6, 64)
	params["type"] = orderType
	params["type_trade"] = typeTrade

	req, err := a.c.newAuthenticatedRequest("create-order", params)

	if err != nil {
		return OrderCreateResponse{}, err
	}

	var v = OrderCreateResponse{}

	_, err = a.c.performRequest(req, &v)

	if err != nil {
		return v, err
	}

	return v, nil
}

func (a *OrderService) GetOrder(orderID int) (OrderResponse, error) {
	params := make(map[string]string)

	params["order_id"] = strconv.Itoa(orderID)

	req, err := a.c.newAuthenticatedRequest("get-order", params)

	if err != nil {
		return OrderResponse{}, err
	}

	var v OrderResponse

	_, err = a.c.performRequest(req, &v)

	if err != nil {
		return OrderResponse{}, err
	}

	return v, nil
}

//Cancel OrderService cancel order by given id
func (a *OrderService) Cancel(orderID int) error {
	params := make(map[string]string)

	params["order_id"] = strconv.Itoa(orderID)

	req, err := a.c.newAuthenticatedRequest("delete-order", params)

	if err != nil {
		return err
	}

	var v OrderCancelResponse

	_, err = a.c.performRequest(req, &v)

	if err != nil {
		return err
	}

	if !v.Status {
		return errors.New(v.Error)
	}

	return nil
}

func (a *OrderService) GetOpenOrders() (OpenOrdersResponse, error) {
	params := make(map[string]string)

	req, err := a.c.newAuthenticatedRequest("orders", params)

	if err != nil {
		return OpenOrdersResponse{}, err
	}

	var v OpenOrdersResponse

	_, err = a.c.performRequest(req, &v)

	if err != nil {
		return OpenOrdersResponse{}, err
	}

	return v, nil
}
