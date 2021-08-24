package campaignnegative

import (
	"alfa_sdk/api"
	"net/http"
	"net/url"
)

var _ api.IParameter = (*CreateCampaignNegativeKeyword)(nil)

type CreateCampaignNegativeKeyword struct {
	Method     string
	Path       string
	PathParams map[string]string
	Body       []*CreateCampaignNegativeKeywordReq
	Query      url.Values
}

func (c *CreateCampaignNegativeKeyword) Validate() error {
	return nil
}

func (c *CreateCampaignNegativeKeyword) Build() *api.Params {
	return &api.Params{
		Method:     c.Method,
		Path:       c.Path,
		PathParams: c.PathParams,
		Body:       c.Body,
		Query:      c.Query,
	}
}

func NewCreateCampaignNegativeKeyword() *CreateCampaignNegativeKeyword {
	body := make([]*CreateCampaignNegativeKeywordReq, 0)
	return &CreateCampaignNegativeKeyword{
		Method:     http.MethodPost,
		Path:       CampaignNegativeURL,
		PathParams: map[string]string{},
		Body:       body,
		Query:      url.Values{},
	}
}

// SetBody ...
func (c *CreateCampaignNegativeKeyword) SetBody(in []*CreateCampaignNegativeKeywordReq) *CreateCampaignNegativeKeyword {
	for _, val := range in {
		c.Body = append(c.Body, val)
	}

	return c
}
