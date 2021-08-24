package portfolios

import (
	"encoding/json"
	"strconv"
)

// request url
var (
	PortfoliosURL            = "/v2/portfolios"
	GetPortfoliosExtendedURL = "/v2/portfolios/extended"
)

//Portfolios ...
type Portfolios struct {
	PortfolioId        int64  `json:"portfolioId"`
	PortfoliosIdString string `json:"portfolios_id_string"`
	Name               string `json:"name"`
	Budget
	InBudget        bool   `json:"inBudget"`
	State           string `json:"state"`
	CreationDate    int64  `json:"creationDate"`
	LastUpdatedDate int64  `json:"lastUpdatedDate"`
	ServingStatus   string `json:"servingStatus"`
}

// Budget budget info ...
type Budget struct {
	Amount float64 `json:"amount,omitempty"`
	//CurrencyCode string  `json:"currencyCode"`
	Policy    string `json:"policy,omitempty"`
	StartDate string `json:"startDate,omitempty"`
	EndState  string `json:"endDate,omitempty"`
}

//CreatePortfoliosReq Portfolio budget policy is invalid for create operation.
type CreatePortfoliosReq struct {
	Name  string `json:"name"`
	State string `json:"state"`
}

type UpdatePortfoliosReq struct {
	PortfolioId int64  `json:"portfolioId,omitempty"`
	Name        string `json:"name,omitempty"`
	Budget      Budget `json:"budget,omitempty"`
}

// PortfoliosSuccessResp create portfolios success response
type PortfoliosSuccessResp struct {
	Code        string `json:"code"`
	PortfolioId int64  `json:"portfolioId"`
}

func (p *Portfolios) UnmarshalJSON(data []byte) error {
	type xPortfolios Portfolios
	x := &xPortfolios{}
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	x.PortfoliosIdString = strconv.Itoa(int(x.PortfolioId))
	*p = Portfolios(*x)

	return nil
}
