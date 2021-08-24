package campaigns

import (
	"alfa_sdk/api"
	"alfa_sdk/types"
	"net/http"
	"net/url"
	"strconv"
)

var _ api.IParameter = (*GetCampaignsExtended)(nil)

type GetCampaignsExtended struct {
	Method     string
	Path       string
	PathParams map[string]string
	Body       interface{}
	Query      url.Values
}

func (c *GetCampaignsExtended) Validate() error {
	return nil
}

func (c *GetCampaignsExtended) Build() *api.Params {
	return &api.Params{
		Method:     c.Method,
		Path:       c.Path,
		PathParams: c.PathParams,
		Body:       c.Body,
		Query:      c.Query,
	}
}

func NewGetCampaignExtended() *GetCampaignsExtended {
	return &GetCampaignsExtended{
		Method:     http.MethodGet,
		Path:       CampaignsURLExtended,
		PathParams: map[string]string{},
		Body:       nil,
		Query:      url.Values{},
	}
}

//SetStartIndexAndCount indexed record offset for the result set.Number of records to include in the paged response. Defaults to max page size.
func (c *GetCampaignsExtended) SetStartIndexAndCount(index, count int) *GetCampaignsExtended {
	if count == 0 {
		count = types.DefaultPageCount
	}
	if index < 1 {
		index = 1
	}
	startIndex := (index - 1) * count
	c.Query.Add("startIndex", strconv.Itoa(startIndex))
	c.Query.Add("count", strconv.Itoa(count))

	return c
}

//SetStateFilter Restricts results to resources with state within the specified comma-separated list.Available values : enabled, paused, archived, enabled, paused, enabled, archived, paused, archived, enabled, paused, archived
func (c *GetCampaignsExtended) SetStateFilter(state string) *GetCampaignsExtended {
	c.Query.Add("stateFilter", state)

	return c
}

// SetName Restricts results to campaigns with the specified name.
func (c *GetCampaignsExtended) SetName(name string) *GetCampaignsExtended {
	c.Query.Add("name", name)

	return c
}

// SetPortfolioIdFilter A comma-delimited list of portfolio identifiers.
func (c *GetCampaignsExtended) SetPortfolioIdFilter(portfolios string) *GetCampaignsExtended {
	c.Query.Add("portfolioIdFilter", portfolios)

	return c
}

// SetCampaignIdFilter A comma-delimited list of campaign identifiers.
func (c *GetCampaignsExtended) SetCampaignIdFilter(campaignId string) *GetCampaignsExtended {
	c.Query.Add("campaignIdFilter", campaignId)

	return c
}
