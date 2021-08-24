package keywords

import (
	"alfa_sdk/api"
	"fmt"
	"net/http"
	"net/url"
)

type DeleteKeywords struct {
	Method     string
	Path       string
	PathParams map[string]string
	Body       interface{}
	Query      url.Values
}

func (k *DeleteKeywords) Validate() error {
	return nil
}

func (k *DeleteKeywords) Build() *api.Params {
	return &api.Params{
		Method:     k.Method,
		Path:       k.Path,
		PathParams: k.PathParams,
		Body:       k.Body,
		Query:      k.Query,
	}
}

func NewDeleteKeywords() *DeleteKeywords {
	return &DeleteKeywords{
		Method:     http.MethodDelete,
		Path:       fmt.Sprintf("%s/{keywordId}", KeywordURL),
		PathParams: map[string]string{},
		Body:       nil,
		Query:      url.Values{},
	}
}

// SetKeywordsId ...
func (k *DeleteKeywords) SetKeywordsId(keywordId string) *DeleteKeywords {
	k.PathParams["keywordId"] = keywordId

	return k
}
