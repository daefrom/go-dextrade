package dextrade

import (
	"net/url"
)

type OrderBooksService struct {
	c *Client
}

type OrderbookEntry struct {
	Volume float64 `json:"volume"`
	Count  int     `json:"count"`
	Rate   float64 `json:"rate"`
}

type OrderBooks struct {
	Status bool `json:"status,omitempty"`
	Data   struct {
		Bids []OrderbookEntry `json:"buy"`
		Asks []OrderbookEntry `json:"sell"`
	} `json:"data"`
}

//Get OrderBooksService returns order books for selected pairs
func (a *OrderBooksService) Get(pair string) (OrderBooks, error) {
	params := url.Values{}
	params.Add("pair", pair)

	req, err := a.c.newRequest("GET", "book", params)

	if err != nil {
		return OrderBooks{}, err
	}

	var v OrderBooks
	_, err = a.c.performRequest(req, &v)

	if err != nil {
		return OrderBooks{}, err
	}
	return v, nil
}
