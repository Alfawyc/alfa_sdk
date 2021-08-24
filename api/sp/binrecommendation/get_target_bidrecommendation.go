package binrecommendation

import (
	"alfa_sdk/api"
	"net/http"
	"net/url"
)

type GetTargetBidRecommendation struct {
	Method     string
	Path       string
	PathParams map[string]string
	Body       interface{}
	Query      url.Values
}

func (t *GetTargetBidRecommendation) Validate() error {
	return nil
}

func (t *GetTargetBidRecommendation) Build() *api.Params {
	return &api.Params{
		Method:     t.Method,
		Path:       t.Path,
		PathParams: t.PathParams,
		Body:       t.Body,
		Query:      t.Query,
	}
}

func NewGetTargetBidRecommendation() *GetTargetBidRecommendation {
	return &GetTargetBidRecommendation{
		Method:     http.MethodPost,
		Path:       TargetBidRecommendationURL,
		PathParams: map[string]string{},
		Body:       &TargetBindRecommendationReq{},
		Query:      url.Values{},
	}
}

// SetBody set request body
func (t *GetTargetBidRecommendation) SetBody(v *TargetBindRecommendationReq) *GetTargetBidRecommendation {
	t.Body = v

	return t
}
