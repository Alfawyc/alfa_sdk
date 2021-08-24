package binrecommendation

import (
	"alfa_sdk/api"
	"net/http"
	"net/url"
)

type GetKeywordsBidRecommendation struct {
	Method     string
	Path       string
	PathParams map[string]string
	Body       interface{}
	Query      url.Values
}

func (t *GetKeywordsBidRecommendation) Validate() error {
	return nil
}

func (t *GetKeywordsBidRecommendation) Build() *api.Params {
	return &api.Params{
		Method:     t.Method,
		Path:       t.Path,
		PathParams: t.PathParams,
		Body:       t.Body,
		Query:      t.Query,
	}
}

func NewGetKeywordsBidRecommendation() *GetKeywordsBidRecommendation {
	return &GetKeywordsBidRecommendation{
		Method:     http.MethodPost,
		Path:       KeywordsBidRecommendationURL,
		PathParams: map[string]string{},
		Body:       &KeywordsBidRecommendationReq{},
		Query:      url.Values{},
	}
}

// SetBody set request body
func (t *GetKeywordsBidRecommendation) SetBody(v *KeywordsBidRecommendationReq) *GetKeywordsBidRecommendation {
	t.Body = v

	return t
}
