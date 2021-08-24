package negativekeywords

import (
	"alfa_sdk/api"
	"alfa_sdk/types"
	"net/http"
	"net/url"
	"strconv"
)

type GetNegativeKeywords struct {
	Method     string
	Path       string
	PathParams map[string]string
	Body       interface{}
	Query      url.Values
}

func (k *GetNegativeKeywords) Validate() error {
	return nil
}

func (k *GetNegativeKeywords) Build() *api.Params {
	return &api.Params{
		Method:     k.Method,
		Path:       k.Path,
		PathParams: k.PathParams,
		Body:       k.Body,
		Query:      k.Query,
	}
}

func NewGetNegativeKeyword() *GetNegativeKeywords {
	return &GetNegativeKeywords{
		Method:     http.MethodGet,
		Path:       NegativeKeywordsURL,
		PathParams: map[string]string{},
		Body:       nil,
		Query:      url.Values{},
	}
}

// SetIndexAndCount 0-indexed record offset for the result set.Number of records to include in the paged response. Defaults to max page size.
func (k *GetNegativeKeywords) SetIndexAndCount(index, count int) *GetNegativeKeywords {
	if count == 0 {
		count = types.DefaultPageCount
	}
	if index < 1 {
		index = 1
	}
	startIndex := (index - 1) * count
	k.Query.Add("startIndex", strconv.Itoa(startIndex))
	k.Query.Add("count", strconv.Itoa(count))

	return k
}

// SetMatchTypeFilter set negative keyword matchType
func (k *GetNegativeKeywords) SetMatchTypeFilter(matchType string) *GetNegativeKeywords {
	k.Query.Add("matchTypeFilter", matchType)

	return k
}

// SetKeywordText set keywordText
func (k *GetNegativeKeywords) SetKeywordText(keywordText string) *GetNegativeKeywords {
	k.Query.Add("keywordText", keywordText)

	return k
}

// SetStateFilter set state
func (k *GetNegativeKeywords) SetStateFilter(state string) *GetNegativeKeywords {
	k.Query.Add("state", state)

	return k
}

// SetCampaignIdFilter set campaignId
func (k *GetNegativeKeywords) SetCampaignIdFilter(campaignId string) *GetNegativeKeywords {
	k.Query.Add("campaignIdFilter", campaignId)

	return k
}

// SetAdGroupIdFilter set adgroupId
func (k *GetNegativeKeywords) SetAdGroupIdFilter(adGroupId string) *GetNegativeKeywords {
	k.Query.Add("adGroupId", adGroupId)

	return k
}

// SetKeywordIdFilter set keywordId
func (k *GetNegativeKeywords) SetKeywordIdFilter(keywordId string) *GetNegativeKeywords {
	k.Query.Add("keywordIdFilter", keywordId)

	return k
}
