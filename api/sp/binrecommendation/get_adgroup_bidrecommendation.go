package binrecommendation

import (
	"alfa_sdk/api"
	"net/http"
	"net/url"
)

type GetAdGroupBidRecommendation struct {
	Method     string
	Path       string
	PathParams map[string]string
	Body       interface{}
	Query      url.Values
}

func (t *GetAdGroupBidRecommendation) Validate() error {
	return nil
}

func (t *GetAdGroupBidRecommendation) Build() *api.Params {
	return &api.Params{
		Method:     t.Method,
		Path:       t.Path,
		PathParams: t.PathParams,
		Body:       t.Body,
		Query:      t.Query,
	}
}

func NewGetAdGroupBidRecommendation() *GetAdGroupBidRecommendation {
	return &GetAdGroupBidRecommendation{
		Method:     http.MethodGet,
		Path:       AdGroupBidRecommendationURL,
		PathParams: map[string]string{},
		Body:       nil,
		Query:      url.Values{},
	}
}

// SetAdGroupId set request adgroupId
func (t *GetAdGroupBidRecommendation) SetAdGroupId(v string) *GetAdGroupBidRecommendation {
	t.PathParams["adGroupId"] = v

	return t
}
