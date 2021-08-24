package keywords

import (
	"alfa_sdk/api"
	"net/http"
	"net/url"
)

type UpdateKeywords struct {
	Method     string
	Path       string
	PathParams map[string]string
	Body       []*UpdateKeywordReq
	Query      url.Values
}

func (k *UpdateKeywords) Validate() error {
	return nil
}

func (k *UpdateKeywords) Build() *api.Params {
	return &api.Params{
		Method:     k.Method,
		Path:       k.Path,
		PathParams: k.PathParams,
		Body:       k.Body,
		Query:      k.Query,
	}
}

func NewUpdateKeywords() *UpdateKeywords {
	body := make([]*UpdateKeywordReq, 0)
	return &UpdateKeywords{
		Method:     http.MethodPut,
		Path:       KeywordURL,
		PathParams: map[string]string{},
		Body:       body,
		Query:      url.Values{},
	}
}

// SetBody ...
func (k *UpdateKeywords) SetBody(in []*UpdateKeywordReq) *UpdateKeywords {
	for _, val := range in {
		k.Body = append(k.Body, val)
	}

	return k
}
