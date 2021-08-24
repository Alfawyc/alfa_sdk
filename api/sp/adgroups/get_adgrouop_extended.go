package adgroups

import (
	"alfa_sdk/api"
	"alfa_sdk/types"
	"net/http"
	"net/url"
	"strconv"
)

var _ api.IParameter = (*GetAdGroupsExtended)(nil)

type GetAdGroupsExtended struct {
	Method     string
	Path       string
	PathParams map[string]string
	Body       interface{}
	Query      url.Values
}

func (a *GetAdGroupsExtended) Validate() error {
	return nil
}

func (a *GetAdGroupsExtended) Build() *api.Params {
	return &api.Params{
		Method:     a.Method,
		Path:       a.Path,
		PathParams: a.PathParams,
		Body:       a.Body,
		Query:      a.Query,
	}
}

func NewGetAdGroupsExtended() *GetAdGroupsExtended {
	return &GetAdGroupsExtended{
		Method:     http.MethodGet,
		Path:       AdGroupsURLExtended,
		PathParams: map[string]string{},
		Body:       nil,
		Query:      url.Values{},
	}
}

// SetStartIndexAndCount record offset for the result set.Number of records to include in the paged response. Defaults to max page size.
func (a *GetAdGroupsExtended) SetStartIndexAndCount(index, count int) *GetAdGroupsExtended {
	if count <= 0 {
		count = types.DefaultPageCount
	}
	if index < 1 {
		index = 1
	}
	startIndex := (index - 1) * count
	a.Query.Add("startIndex", strconv.Itoa(startIndex))

	a.Query.Add("count", strconv.Itoa(count))

	return a
}

// SetCampaignType Restricts results to ad groups within campaign types specified in comma-separated list.
func (a *GetAdGroupsExtended) SetCampaignType(v string) *GetAdGroupsExtended {
	a.Query.Add("campaignType", v)

	return a
}

// SetStateFilter Restricts results to resources with state within the specified comma-separated list. Available values : enabled, paused, archived, enabled, paused, enabled, archived, paused, archived, enabled, paused, archived
func (a *GetAdGroupsExtended) SetStateFilter(state string) *GetAdGroupsExtended {
	a.Query.Add("stateFilter", state)

	return a
}

// SetName Restricts results to campaigns with the specified name.
func (a *GetAdGroupsExtended) SetName(name string) *GetAdGroupsExtended {
	a.Query.Add("name", name)

	return a
}

// SetCampaignIdFilter A comma-delimited list of campaign identifiers.
func (a *GetAdGroupsExtended) SetCampaignIdFilter(campaignId string) *GetAdGroupsExtended {
	a.Query.Add("campaignIdFilter", campaignId)

	return a
}
