package keywords

import (
	"alfa_sdk/api"
	"alfa_sdk/types"
	"net/http"
	"net/url"
	"strconv"
)

var _ api.IParameter = (*GetKeywords)(nil)

type GetKeywords struct {
	Method     string
	Path       string
	PathParams map[string]string
	Body       interface{}
	Query      url.Values
}

func (k *GetKeywords) Validate() error {
	return nil
}

func (k *GetKeywords) Build() *api.Params {
	return &api.Params{
		Method:     k.Method,
		Path:       k.Path,
		PathParams: k.PathParams,
		Body:       k.Body,
		Query:      k.Query,
	}
}

func NewGetKeywords() *GetKeywords {
	return &GetKeywords{
		Method:     http.MethodGet,
		Path:       KeywordURLExtended,
		PathParams: map[string]string{},
		Body:       nil,
		Query:      url.Values{},
	}
}

// SetStartIndexAndCount record offset for the result set. Number of records to include in the paged response. Defaults to max page size.
func (k *GetKeywords) SetStartIndexAndCount(index, count int) *GetKeywords {
	if index < 1 {
		index = 1
	}
	if count <= 0 {
		count = types.DefaultPageCount
	}
	startIndex := (index - 1) * count
	k.Query.Add("startIndex", strconv.Itoa(startIndex))
	k.Query.Add("count", strconv.Itoa(count))

	return k
}

// SetMatchTypeFilter Restricts results to keywords with match types within the specified comma-separated list. For more information, see match types in the Amazon Advertising support center.
func (k *GetKeywords) SetMatchTypeFilter(matchType string) *GetKeywords {
	k.Query.Add("matchTypeFilter", matchType)

	return k
}

// SetKeywordText Restricts results to keywords that match the specified text exactly.
func (k *GetKeywords) SetKeywordText(keywordText string) *GetKeywords {
	k.Query.Add("keywordText", keywordText)

	return k
}

// SetStateFilter Restricts results to resources with state within the specified comma-separated list.Available values : enabled, paused, archived, enabled, paused, enabled, archived, paused, archived, enabled, paused, archived
func (k *GetKeywords) SetStateFilter(stateFilter string) *GetKeywords {
	k.Query.Add("stateFilter", stateFilter)

	return k
}

// SetCampaignFilter A comma-delimited list of campaign identifiers.
func (k *GetKeywords) SetCampaignFilter(campaignId string) *GetKeywords {
	k.Query.Add("campaignIdFilter", campaignId)

	return k
}

// SetAdGroupIdFilter Restricts results to keywords associated with ad groups specified by identifier in the comma-delimited list.
func (k *GetKeywords) SetAdGroupIdFilter(adGroupId string) *GetKeywords {
	k.Query.Add("adGroupIdFilter", adGroupId)

	return k
}

// SetKeywordIdFilter Restricts results to keywords associated with campaigns specified by identifier in the comma-delimited list.
func (k *GetKeywords) SetKeywordIdFilter(keywordId string) *GetKeywords {
	k.Query.Add("keywordIdFilter", keywordId)

	return k
}

// SetLocale Restricts results to keywords associated with locale.
func (k *GetKeywords) SetLocale(locale string) *GetKeywords {
	k.Query.Add("locale", locale)

	return k
}
