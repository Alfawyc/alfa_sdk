package keywords

import (
	"alfa_sdk/api"
	"net/http"
	"net/url"
)

var _ api.IParameter = (*CreateKeywords)(nil)

type CreateKeywords struct {
	Method     string
	Path       string
	PathParams map[string]string
	Body       []*CreateKeywordsReq
	Query      url.Values
}

func (k *CreateKeywords) Validate() error {
	return nil
}

func (k *CreateKeywords) Build() *api.Params {
	return &api.Params{
		Method:     k.Method,
		Path:       k.Path,
		PathParams: k.PathParams,
		Body:       k.Body,
		Query:      k.Query,
	}
}

func NewCreateKeywords() *CreateKeywords {
	body := make([]*CreateKeywordsReq, 0)
	return &CreateKeywords{
		Method:     http.MethodPost,
		Path:       KeywordURL,
		PathParams: map[string]string{},
		Body:       body,
		Query:      url.Values{},
	}
}

// SetBody ...
func (k *CreateKeywords) SetBody(in []*CreateKeywordsReq) *CreateKeywords {
	for _, val := range in {
		k.Body = append(k.Body, val)
	}

	return k
}
