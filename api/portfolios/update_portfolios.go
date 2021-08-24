package portfolios

import (
	"alfa_sdk/api"
	"net/http"
	"net/url"
)

var _ api.IParameter = (*GetPortfolios)(nil)

type UpdatePortfolios struct {
	Method     string
	Path       string
	PathParams map[string]string
	Body       []*UpdatePortfoliosReq
	Query      url.Values
}

func (p *UpdatePortfolios) Validate() error {
	return nil
}

func (p *UpdatePortfolios) Build() *api.Params {
	return &api.Params{
		Method:     p.Method,
		Path:       p.Path,
		PathParams: p.PathParams,
		Body:       p.Body,
		Query:      p.Query,
	}
}

func NewUpdatePortfolios() *UpdatePortfolios {
	body := make([]*UpdatePortfoliosReq, 0)
	return &UpdatePortfolios{
		Method:     http.MethodPut,
		Path:       PortfoliosURL,
		PathParams: map[string]string{},
		Body:       body,
		Query:      url.Values{},
	}
}

// SetBody set request body
func (p *UpdatePortfolios) SetBody(in []*UpdatePortfoliosReq) *UpdatePortfolios {
	for _, val := range in {
		p.Body = append(p.Body, val)
	}

	return p
}
