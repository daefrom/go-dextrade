package dextrade

import (
	"net/url"
)

type TradesService struct {
	c *Client
}

type TradeStruct struct {
	Status    bool
	Volume    float64
	Rate      float64
	Price     float64
	Timestamp int64
	Type      string
}

type Trades struct {
	Status bool `json:"status,omitempty"`
	Error  string
	Data   []TradeStruct `json:"data"`
}

//Get TradesService returns trades for selected pairs
func (a *TradesService) Get(pair string) (Trades, error) {
	params := url.Values{}
	params.Add("pair", pair)

	req, err := a.c.newRequest("GET", "trades", params)

	if err != nil {
		return Trades{}, err
	}

	var v Trades

	_, err = a.c.performRequest(req, &v)

	if err != nil {
		return Trades{}, err
	}

	return v, nil
}
