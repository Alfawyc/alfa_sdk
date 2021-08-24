package keywords

import (
	"encoding/json"
	"strconv"
)

var (
	KeywordURL         = "/v2/sp/keywords"
	KeywordURLExtended = "/v2/sp/keywords/extended"
)

type Keywords struct {
	KeywordId             int64   `json:"keywordId"`
	KeywordStringId       string  `json:"keyword_string_id"`
	CampaignId            int64   `json:"campaignId"`
	CampaignStringId      string  `json:"campaign_string_id"`
	AdGroupId             int64   `json:"adGroupId"`
	AdGroupStringId       string  `json:"ad_group_string_id"`
	State                 string  `json:"state"`
	KeywordText           string  `json:"keywordText"`
	NativeLanguageKeyword string  `json:"nativeLanguageKeyword"`
	MatchType             string  `json:"matchType"`
	Bid                   float64 `json:"bid"`
	CreationDate          int64   `json:"creationDate"`
	LastUpdatedDate       int64   `json:"lastUpdatedDate"`
	ServingStatus         string  `json:"servingStatus"`
}

// CreateKeywordsReq createKeyword request
type CreateKeywordsReq struct {
	CampaignId            int64   `json:"campaignId"`
	AdGroupId             int64   `json:"adGroupId"`
	State                 string  `json:"state"`
	KeywordText           string  `json:"keywordText"`
	NativeLanguageKeyword string  `json:"nativeLanguageKeyword,omitempty"`
	NativeLanguageLocale  string  `json:"nativeLanguageLocale,omitempty"`
	MatchType             string  `json:"matchType"`
	Bid                   float64 `json:"bid,omitempty"`
}

// UpdateKeywordReq updateKeyword request
type UpdateKeywordReq struct {
	KeywordId int64   `json:"keywordId"`
	State     string  `json:"state,omitempty"`
	Bid       float64 `json:"bid,omitempty"`
}

// KeywordResponse edit response
type KeywordResponse struct {
	KeywordId   int64  `json:"keywordId"`
	Code        string `json:"code"`
	Details     string `json:"details"`
	Description string `json:"description"`
}

func (p *Keywords) UnmarshalJSON(data []byte) error {
	type xKeywords Keywords
	x := &xKeywords{}
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	x.CampaignStringId = strconv.Itoa(int(x.CampaignId))
	x.KeywordStringId = strconv.Itoa(int(x.KeywordId))
	x.AdGroupStringId = strconv.Itoa(int(x.AdGroupId))
	*p = Keywords(*x)

	return nil
}
