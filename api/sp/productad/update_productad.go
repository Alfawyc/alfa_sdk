package productad

import (
	"alfa_sdk/api"
	"net/http"
	"net/url"
)

var _ api.IParameter = (*UpdateProductAd)(nil)

type UpdateProductAd struct {
	Method     string
	Path       string
	PathParams map[string]string
	Body       []*UpdateProductAdReq
	Query      url.Values
}

func (c *UpdateProductAd) Validate() error {
	return nil
}

func (c *UpdateProductAd) Build() *api.Params {
	return &api.Params{
		Method:     c.Method,
		Path:       c.Path,
		PathParams: c.PathParams,
		Body:       c.Body,
		Query:      c.Query,
	}
}

func NewUpdateProductAd() *UpdateProductAd {
	body := make([]*UpdateProductAdReq, 0)
	return &UpdateProductAd{
		Method:     http.MethodPost,
		Path:       ProductAdURL,
		PathParams: map[string]string{},
		Body:       body,
		Query:      url.Values{},
	}
}

// SetBody ...
func (c *UpdateProductAd) SetBody(in []*UpdateProductAdReq) *UpdateProductAd {
	for _, val := range in {
		c.Body = append(c.Body, val)
	}

	return c
}
