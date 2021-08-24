package campaigns

import (
	"alfa_sdk/api"
	"fmt"
	"net/http"
	"net/url"
)

var _ api.IParameter = (*DeleteCampaign)(nil)

type DeleteCampaign struct {
	Method     string
	Path       string
	PathParams map[string]string
	Body       interface{}
	Query      url.Values
}

func (c *DeleteCampaign) Validate() error {
	return nil
}

func (c *DeleteCampaign) Build() *api.Params {
	return &api.Params{
		Method:     c.Method,
		Path:       c.Path,
		PathParams: c.PathParams,
		Body:       c.Body,
		Query:      c.Query,
	}
}

func NewDeleteCampaign() *DeleteCampaign {
	return &DeleteCampaign{
		Method:     http.MethodDelete,
		Path:       fmt.Sprintf("%s/{campaignId}", CampaignsURL),
		PathParams: map[string]string{},
		Body:       nil,
		Query:      url.Values{},
	}
}

// SetCampaignId ...
func (c *DeleteCampaign) SetCampaignId(campaignId string) *DeleteCampaign {
	c.PathParams["campaignId"] = campaignId

	return c
}
