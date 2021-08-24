package campaignnegative

import (
	"alfa_sdk/api"
	"net/http"
	"net/url"
)

var _ api.IParameter = (*UpdateCampaignNegativeKeyword)(nil)

type UpdateCampaignNegativeKeyword struct {
	Method     string
	Path       string
	PathParams map[string]string
	Body       []*UpdateCampaignNegativeKeywordReq
	Query      url.Values
}

func (c *UpdateCampaignNegativeKeyword) Validate() error {
	return nil
}

func (c *UpdateCampaignNegativeKeyword) Build() *api.Params {
	return &api.Params{
		Method:     c.Method,
		Path:       c.Path,
		PathParams: c.PathParams,
		Body:       c.Body,
		Query:      c.Query,
	}
}

func NewUpdateCampaignNegativeKeyword() *UpdateCampaignNegativeKeyword {
	body := make([]*UpdateCampaignNegativeKeywordReq, 0)
	return &UpdateCampaignNegativeKeyword{
		Method:     http.MethodPut,
		Path:       CampaignNegativeURL,
		PathParams: map[string]string{},
		Body:       body,
		Query:      url.Values{},
	}
}

// SetBody ...
func (c *UpdateCampaignNegativeKeyword) SetBody(in []*UpdateCampaignNegativeKeywordReq) *UpdateCampaignNegativeKeyword {
	for _, val := range in {

		//state must be deleted
		val.State = "deleted"

		c.Body = append(c.Body, val)
	}

	return c
}
