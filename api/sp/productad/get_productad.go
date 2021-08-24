package productad

import (
	"alfa_sdk/api"
	"alfa_sdk/types"
	"net/http"
	"net/url"
	"strconv"
)

var _ api.IParameter = (*GetProductAd)(nil)

type GetProductAd struct {
	Method     string
	Path       string
	PathParams map[string]string
	Body       interface{}
	Query      url.Values
}

func (a *GetProductAd) Validate() error {
	return nil
}

func (a *GetProductAd) Build() *api.Params {
	return &api.Params{
		Method:     a.Method,
		Path:       a.Path,
		PathParams: a.PathParams,
		Body:       a.Body,
		Query:      a.Query,
	}
}

func NewGetProductAd() *GetProductAd {
	return &GetProductAd{
		Method:     http.MethodGet,
		Path:       ProductAdURLExtended,
		PathParams: map[string]string{},
		Body:       nil,
		Query:      url.Values{},
	}
}

// SetStartIndexAndCount ...
func (a *GetProductAd) SetStartIndexAndCount(index, count int) *GetProductAd {
	if index < 1 {
		index = 1
	}
	if count <= 0 {
		count = types.DefaultPageCount
	}
	startIndex := (index - 1) * count
	a.Query.Add("startIndex", strconv.Itoa(startIndex))
	a.Query.Add("count", strconv.Itoa(count))

	return a
}

// SetCampaignIdFilter A comma-delimited list of campaign identifiers
func (a *GetProductAd) SetCampaignIdFilter(campaignId string) *GetProductAd {
	a.Query.Add("campaignIdFilter", campaignId)

	return a
}

// SetAdGroupIdFilter Restricts results to keywords associated with ad groups specified by identifier in the comma-delimited list.
func (a *GetProductAd) SetAdGroupIdFilter(adGroupId string) *GetProductAd {
	a.Query.Add("adGroupIdFilter", adGroupId)

	return a
}

// SetAdIdFilter Restricts results to product ads associated with the product ad identifiers specified in the comma-delimited list.
func (a *GetProductAd) SetAdIdFilter(adId string) *GetProductAd {
	a.Query.Add("adIdFilter", adId)

	return a
}

// SetStateFilter  Restricts results to resources with state within the specified comma-separated list. Available values : enabled, paused, archived, enabled, paused, enabled, archived, paused, archived, enabled, paused, archived
func (a *GetProductAd) SetStateFilter(state string) *GetProductAd {
	a.Query.Add("stateFilter", state)

	return a
}
