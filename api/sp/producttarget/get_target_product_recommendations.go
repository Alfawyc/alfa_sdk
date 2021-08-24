package producttarget

import (
	"alfa_sdk/api"
	"net/http"
	"net/url"
)

type GetTargetProductRecommendation struct {
	Method     string
	Path       string
	PathParams map[string]string
	Body       interface{}
	Query      url.Values
}

func (c *GetTargetProductRecommendation) Validate() error {
	return nil
}

func (c *GetTargetProductRecommendation) Build() *api.Params {
	return &api.Params{
		Method:     c.Method,
		Path:       c.Path,
		PathParams: c.PathParams,
		Body:       c.Body,
		Query:      c.Query,
	}
}

func NewGetTargetProductRecommendation() *GetTargetProductRecommendation {
	return &GetTargetProductRecommendation{
		Method:     http.MethodPost,
		Path:       TargetProductRecommendationURL,
		PathParams: map[string]string{},
		Body:       &GetTargetProductRecommendationReq{},
		Query:      url.Values{},
	}
}

// SetBody asins is required ,pageNumber must start from 0
func (c *GetTargetProductRecommendation) SetBody(param *GetTargetProductRecommendationReq) *GetTargetProductRecommendation {
	c.Body = param

	return c
}
