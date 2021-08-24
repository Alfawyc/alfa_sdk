package campaignnegative

import (
	"alfa_sdk/api"
	"alfa_sdk/types"
	"net/http"
	"net/url"
	"strconv"
)

var _ api.IParameter = (*GetCampaignNegativeKeyword)(nil)

type GetCampaignNegativeKeyword struct {
	Method     string
	Path       string
	PathParams map[string]string
	Body       interface{}
	Query      url.Values
}

func (c *GetCampaignNegativeKeyword) Validate() error {
	return nil
}

func (c *GetCampaignNegativeKeyword) Build() *api.Params {
	return &api.Params{
		Method:     c.Method,
		Path:       c.Path,
		PathParams: c.PathParams,
		Body:       c.Body,
		Query:      c.Query,
	}
}

func NewGetCampaignNegativeKeyword() *GetCampaignNegativeKeyword {
	return &GetCampaignNegativeKeyword{
		Method:     http.MethodGet,
		Path:       CampaignNegativeURLExtended,
		PathParams: map[string]string{},
		Body:       nil,
		Query:      url.Values{},
	}
}

// SetStartIndexAndCount record offset for the result set.Number of records to include in the paged response. Defaults to max page size.
func (c *GetCampaignNegativeKeyword) SetStartIndexAndCount(index, count int) *GetCampaignNegativeKeyword {
	if count <= 0 {
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

// SetCampaignIdFilter Restricts results to negative keywords associated with campaigns specified by identifier in the comma-delimited list.
func (c *GetCampaignNegativeKeyword) SetCampaignIdFilter(campaignId string) *GetCampaignNegativeKeyword {
	c.Query.Add("campaignIdFilter", campaignId)

	return c
}

//SetKeywordIdFilter Restricts results to negative keywords associated with campaigns specified by identifier in the comma-delimited list.
func (c *GetCampaignNegativeKeyword) SetKeywordIdFilter(keywordId string) *GetCampaignNegativeKeyword {
	c.Query.Add("keywordIdFilter", keywordId)

	return c
}

// SetKeywordText  Restricts results to negative keywords that match the specified text.
func (c *GetCampaignNegativeKeyword) SetKeywordText(keywordText string) *GetCampaignNegativeKeyword {
	c.Query.Add("keywordText", keywordText)

	return c
}

// SetMatchTypeFilter  Available values : negativePhrase, negativeExact
func (c *GetCampaignNegativeKeyword) SetMatchTypeFilter(matchType string) *GetCampaignNegativeKeyword {
	c.Query.Add("matchTypeFilter", matchType)

	return c
}
