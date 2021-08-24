package portfolios

import (
	"alfa_sdk/api"
	"net/http"
	"net/url"
)

var _ api.IParameter = (*GetPortfolios)(nil)

type GetPortfolios struct {
	Method     string
	Path       string
	PathParams map[string]string
	Body       interface{}
	Query      url.Values
}

func (p *GetPortfolios) Validate() error {
	return nil
}

func (p *GetPortfolios) Build() *api.Params {
	return &api.Params{
		Method:     p.Method,
		Path:       p.Path,
		PathParams: p.PathParams,
		Body:       p.Body,
		Query:      p.Query,
	}
}

func NewGetPortfolios() *GetPortfolios {
	return &GetPortfolios{
		Method:     http.MethodGet,
		Path:       GetPortfoliosExtendedURL,
		PathParams: map[string]string{},
		Body:       nil,
		Query:      url.Values{},
	}
}
