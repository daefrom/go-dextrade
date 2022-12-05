package dextrade

import (
	"net/url"
)

type TickerService struct {
	c *Client
}

type Ticker struct {
	Status       bool `json:"status,omitempty"`
	Error        string
	TickerStruct struct {
		Id        int     `json:"id"`
		Pair      string  `json:"pair"`
		Last      string  `json:"last"`
		Open      float64 `json:"open"`
		Close     float64 `json:"close"`
		High      string  `json:"high"`
		Low       string  `json:"low"`
		Volume24h string  `json:"volume_24H"`
		MinTrade  string  `json:"min_trade"`
	} `json:"data"`
}

//Get TickerService returns ticker
func (a *TickerService) Get(pair string) (Ticker, error) {
	params := url.Values{}
	params.Add("pair", pair)

	req, err := a.c.newRequest("GET", "ticker", params)

	if err != nil {
		return Ticker{}, err
	}

	var v Ticker

	_, err = a.c.performRequest(req, &v)

	if err != nil {
		return Ticker{}, err
	}

	return v, nil
}
