package productad

import (
	"alfa_sdk/api"
	"net/http"
	"net/url"
)

var _ api.IParameter = (*CreateProductAd)(nil)

type CreateProductAd struct {
	Method     string
	Path       string
	PathParams map[string]string
	Body       []*CreateProductAdReq
	Query      url.Values
}

func (c *CreateProductAd) Validate() error {
	return nil
}

func (c *CreateProductAd) Build() *api.Params {
	return &api.Params{
		Method:     c.Method,
		Path:       c.Path,
		PathParams: c.PathParams,
		Body:       c.Body,
		Query:      c.Query,
	}
}

func NewCreateProductAd() *CreateProductAd {
	body := make([]*CreateProductAdReq, 0)
	return &CreateProductAd{
		Method:     http.MethodPost,
		Path:       ProductAdURL,
		PathParams: map[string]string{},
		Body:       body,
		Query:      url.Values{},
	}
}

// SetBody ...
func (c *CreateProductAd) SetBody(in []*CreateProductAdReq) *CreateProductAd {
	for _, val := range in {
		c.Body = append(c.Body, val)
	}

	return c
}
