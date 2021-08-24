package producttarget

import (
	"alfa_sdk/api"
	"fmt"
	"net/http"
	"net/url"
)

type GetProductTargetSpecified struct {
	Method     string
	Path       string
	PathParams map[string]string
	Body       interface{}
	Query      url.Values
}

func (c *GetProductTargetSpecified) Validate() error {
	return nil
}

func (c *GetProductTargetSpecified) Build() *api.Params {
	return &api.Params{
		Method:     c.Method,
		Path:       c.Path,
		PathParams: c.PathParams,
		Body:       c.Body,
		Query:      c.Query,
	}
}

func NewGetProductTargetSpecified() *GetProductTargetSpecified {
	return &GetProductTargetSpecified{
		Method:     http.MethodGet,
		Path:       fmt.Sprintf("%s/{targetId}", ProductTargetURL),
		PathParams: map[string]string{},
		Body:       nil,
		Query:      url.Values{},
	}
}

// SetTargetId ...
func (c *GetProductTargetSpecified) SetTargetId(targetId string) *GetProductTargetSpecified {
	c.PathParams["targetId"] = targetId

	return c
}
