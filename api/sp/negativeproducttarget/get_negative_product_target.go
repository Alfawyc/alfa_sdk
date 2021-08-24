package negativeproducttarget

import (
	"alfa_sdk/api"
	"alfa_sdk/types"
	"net/http"
	"net/url"
	"strconv"
)

type GetNegativeProductTarget struct {
	Method     string
	Path       string
	PathParams map[string]string
	Body       interface{}
	Query      url.Values
}

func (t *GetNegativeProductTarget) Validate() error {
	return nil
}

func (t *GetNegativeProductTarget) Build() *api.Params {
	return &api.Params{
		Method:     t.Method,
		Path:       t.Path,
		PathParams: t.PathParams,
		Body:       t.Body,
		Query:      t.Query,
	}
}

func NewGetNegativeProductTarget() *GetNegativeProductTarget {
	return &GetNegativeProductTarget{
		Method:     http.MethodGet,
		Path:       NegativeProductTargetExtendURL,
		PathParams: map[string]string{},
		Body:       nil,
		Query:      url.Values{},
	}
}

func (t *GetNegativeProductTarget) SetStartIndexAndCount(index, count int) *GetNegativeProductTarget {
	if index < 1 {
		index = 1
	}
	if count <= 0 {
		count = types.DefaultPageCount
	}
	startIndex := (index - 1) * count
	t.Query.Add("startIndex", strconv.Itoa(startIndex))
	t.Query.Add("count", strconv.Itoa(count))

	return t
}

func (t *GetNegativeProductTarget) SetStateFilter(state string) *GetNegativeProductTarget {
	t.Query.Add("stateFilter", state)

	return t
}

func (t *GetNegativeProductTarget) SetCampaignIdFilter(campaignID string) *GetNegativeProductTarget {
	t.Query.Add("campaignIdFilter", campaignID)

	return t
}

func (t *GetNegativeProductTarget) SetAdGroupIdFilter(adGroupID string) *GetNegativeProductTarget {
	t.Query.Add("adGroupIdFilter", adGroupID)

	return t
}

func (t *GetNegativeProductTarget) SetTargetIdFilter(targetID string) *GetNegativeProductTarget {
	t.Query.Add("targetIdFilter", targetID)

	return t
}
