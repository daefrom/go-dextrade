package dextrade

type BalancesService struct {
	c *Client
}

type Currency struct {
	Iso3 string
	Name string
}
type BalanceDataResponse struct {
	List []BalanceListResponse
}

type BalanceListResponse struct {
	Balance          float64 `json:"balance"`
	BalanceAvailable float64 `json:"balance_available"`
	Currency         Currency
}

type BalanceResponse struct {
	Status bool
	Error  string
	Data   BalanceDataResponse
}

func (a *BalancesService) GetBalances() (BalanceResponse, error) {
	params := make(map[string]string)

	req, err := a.c.newAuthenticatedRequest("balances", params)

	if err != nil {
		return BalanceResponse{}, err
	}
	var v BalanceResponse

	_, err = a.c.performRequest(req, &v)

	if err != nil {
		return BalanceResponse{}, err
	}

	return v, nil
}
