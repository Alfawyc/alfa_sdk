package suggestedkeywords

import (
	"alfa_sdk/api"
	"net/http"
	"net/url"
	"strconv"
)

type GetAdGroupSuggestedKeyword struct {
	Method     string
	Path       string
	PathParams map[string]string
	Body       interface{}
	Query      url.Values
}

func (k *GetAdGroupSuggestedKeyword) Validate() error {
	return nil
}

func (k *GetAdGroupSuggestedKeyword) Build() *api.Params {
	return &api.Params{
		Method:     k.Method,
		Path:       k.Path,
		PathParams: k.PathParams,
		Body:       k.Body,
		Query:      k.Query,
	}
}

func NewGetAdGroupSuggestedKeyword() *GetAdGroupSuggestedKeyword {
	return &GetAdGroupSuggestedKeyword{
		Method:     http.MethodGet,
		Path:       AdGroupsSuggestedKeywordsExtendURL,
		PathParams: map[string]string{},
		Body:       nil,
		Query:      url.Values{},
	}
}

// SetAdGroupID The identifier of a valid ad group.
func (k *GetAdGroupSuggestedKeyword) SetAdGroupID(v string) *GetAdGroupSuggestedKeyword {
	k.PathParams["adGroupId"] = v

	return k
}

// SetMaxNumSuggestions The maxiumum number of suggested keywords for the response.Default value : 100
func (k *GetAdGroupSuggestedKeyword) SetMaxNumSuggestions(number int) *GetAdGroupSuggestedKeyword {
	k.Query.Add("maxNumSuggestions", strconv.Itoa(number))

	return k
}

// EnabledSuggestBids Set to yes to include a suggest bid for the suggested keyword in the response. Otherwise, set to no.
func (k *GetAdGroupSuggestedKeyword) EnabledSuggestBids() *GetAdGroupSuggestedKeyword {
	k.Query.Add("suggestBids", "yes")

	return k
}
