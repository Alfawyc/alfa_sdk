package producttarget

import (
	"alfa_sdk/api"
	"alfa_sdk/types"
	"net/http"
	"net/url"
	"strconv"
)

type GetProductTargetExtend struct {
	Method     string
	Path       string
	PathParams map[string]string
	Body       interface{}
	Query      url.Values
}

func (c *GetProductTargetExtend) Validate() error {
	return nil
}

func (c *GetProductTargetExtend) Build() *api.Params {
	return &api.Params{
		Method:     c.Method,
		Path:       c.Path,
		PathParams: c.PathParams,
		Body:       c.Body,
		Query:      c.Query,
	}
}

func NewGetProductTargetExtend() *GetProductTargetExtend {
	return &GetProductTargetExtend{
		Method:     http.MethodGet,
		Path:       ProductTargetExtendURL,
		PathParams: map[string]string{},
		Body:       nil,
		Query:      url.Values{},
	}
}

//SetStartIndexAndCount indexed record offset for the result set.Number of records to include in the paged response. Defaults to max page size.
func (c *GetProductTargetExtend) SetStartIndexAndCount(index, count int) *GetProductTargetExtend {
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

// SetStateFilter Restricts results to resources with state within the specified comma-separated list.
func (c *GetProductTargetExtend) SetStateFilter(state string) *GetProductTargetExtend {
	c.Query.Add("stateFilter", state)

	return c
}

// SetCampaignIdFilter A comma-delimited list of campaign identifiers.
func (c *GetProductTargetExtend) SetCampaignIdFilter(campaignId string) *GetProductTargetExtend {
	c.Query.Add("campaignIdFilter", campaignId)

	return c
}

// SetAdGroupIdFilter Restricts results to keywords associated with ad groups specified by identifier in the comma-delimited list.
func (c *GetProductTargetExtend) SetAdGroupIdFilter(adGroupId string) *GetProductTargetExtend {
	c.Query.Add("adGroupIdFilter", adGroupId)

	return c
}

// SetTargetIdFilter  A comma-delimited list of target identifiers.
func (c *GetProductTargetExtend) SetTargetIdFilter(targetId string) *GetProductTargetExtend {
	c.Query.Add("targetIdFilter", targetId)

	return c
}
