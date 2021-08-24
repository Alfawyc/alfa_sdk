package productad

import (
	"alfa_sdk/api"
	"fmt"
	"net/http"
	"net/url"
)

type DeleteProductAd struct {
	Method     string
	Path       string
	PathParams map[string]string
	Body       interface{}
	Query      url.Values
}

func (c *DeleteProductAd) Validate() error {
	return nil
}

func (c *DeleteProductAd) Build() *api.Params {
	return &api.Params{
		Method:     c.Method,
		Path:       c.Path,
		PathParams: c.PathParams,
		Body:       c.Body,
		Query:      c.Query,
	}
}

func NewDeleteProductAd() *DeleteProductAd {
	return &DeleteProductAd{
		Method:     http.MethodDelete,
		Path:       fmt.Sprintf("%s/{adId}", ProductAdURL),
		PathParams: map[string]string{},
		Body:       nil,
		Query:      url.Values{},
	}
}

// SetAdId ...
func (c *DeleteProductAd) SetAdId(adId string) *DeleteProductAd {
	c.PathParams["adId"] = adId

	return c
}
