package campaignnegative

import (
	"alfa_sdk/api"
	"fmt"
	"net/http"
	"net/url"
)

var _ api.IParameter = (*DeleteCampaignNegativeKeyword)(nil)

type DeleteCampaignNegativeKeyword struct {
	Method     string
	Path       string
	PathParams map[string]string
	Body       interface{}
	Query      url.Values
}

func (c *DeleteCampaignNegativeKeyword) Validate() error {
	return nil
}

func (c *DeleteCampaignNegativeKeyword) Build() *api.Params {
	return &api.Params{
		Method:     c.Method,
		Path:       c.Path,
		PathParams: c.PathParams,
		Body:       c.Body,
		Query:      c.Query,
	}
}

func NewDeleteCampaignNegativeKeyword() *DeleteCampaignNegativeKeyword {
	return &DeleteCampaignNegativeKeyword{
		Method:     http.MethodDelete,
		Path:       fmt.Sprintf("%v/{keywordId}", CampaignNegativeURL),
		PathParams: map[string]string{},
		Body:       nil,
		Query:      url.Values{},
	}
}

// SetKeywordId ...
func (c *DeleteCampaignNegativeKeyword) SetKeywordId(keywordId string) *DeleteCampaignNegativeKeyword {
	c.PathParams["keywordId"] = keywordId

	return c
}
