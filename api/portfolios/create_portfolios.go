package portfolios

import (
	"alfa_sdk/api"
	"net/http"
	"net/url"
)

var _ api.IParameter = (*CreatePortfolios)(nil)

type CreatePortfolios struct {
	Method     string
	Path       string
	PathParams map[string]string
	Body       []*CreatePortfoliosReq
	Query      url.Values
}

func (p *CreatePortfolios) Validate() error {
	return nil
}

func (p *CreatePortfolios) Build() *api.Params {
	return &api.Params{
		Method:     p.Method,
		Path:       p.Path,
		PathParams: p.PathParams,
		Body:       p.Body,
		Query:      p.Query,
	}
}

// NewCreatePortfolios ...
func NewCreatePortfolios() *CreatePortfolios {
	body := make([]*CreatePortfoliosReq, 0)
	return &CreatePortfolios{
		Method:     http.MethodPost,
		Path:       PortfoliosURL,
		PathParams: map[string]string{},
		Body:       body,
		Query:      url.Values{},
	}
}

// SetBody set request body
func (c *CreatePortfolios) SetBody(in []*CreatePortfoliosReq) *CreatePortfolios {
	for _, val := range in {
		c.Body = append(c.Body, val)
	}

	return c
}
