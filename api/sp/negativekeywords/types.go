package negativekeywords

import (
	"encoding/json"
	"strconv"
)

var (
	NegativeKeywordsURL = "/v2/sp/negativeKeywords"
)

type NegativeKeyword struct {
	KeywordID        int64  `json:"keywordId"`
	KeywordIdString  string `json:"keyword_id_string"`
	CampaignID       int64  `json:"campaignId"`
	CampaignIdString string `json:"campaign_id_string"`
	AdGroupID        int64  `json:"adGroupId"`
	AdGroupIdString  string `json:"adgroup_id_string"`
	State            string `json:"state"`
	KeywordText      string `json:"keywordText"`
	MatchType        string `json:"matchType"`
}

func (p *NegativeKeyword) UnmarshalJSON(data []byte) error {
	type xNegativeKeyword NegativeKeyword
	x := &xNegativeKeyword{}
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	x.CampaignIdString = strconv.Itoa(int(x.CampaignID))
	x.AdGroupIdString = strconv.Itoa(int(x.AdGroupID))
	x.KeywordIdString = strconv.Itoa(int(x.KeywordID))
	*p = NegativeKeyword(*x)

	return nil
}

type CreateNegativeKeywordReq struct {
	CampaignID  int64  `json:"campaignId"`
	AdGroupID   int64  `json:"adGroupId"`
	State       string `json:"state"`
	KeywordText string `json:"keywordText"`
	MatchType   string `json:"matchType"`
}

type UpdateNegativeKeywordReq struct {
	KeywordID string `json:"keywordId"`
	State     string `json:"state"`
}

type NegativeKeywordResp struct {
	KeywordID   int64  `json:"keywordId"`
	Code        string `json:"code"`
	Details     string `json:"details"`
	Description string `json:"description"`
}
