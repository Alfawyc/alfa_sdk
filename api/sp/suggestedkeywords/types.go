package suggestedkeywords

import (
	"encoding/json"
	"strconv"
)

var (
	AdGroupsSuggestedKeywordsExtendURL = "/v2/sp/adGroups/{adGroupId}/suggested/keywords/extended"
)

type AdgroupsSuggestedKeywords struct {
	AdGroupID        int64   `json:"adGroupId"`
	AdGroupIdString  string  `json:"ad_group_id_string"`
	CampaignID       int64   `json:"campaignId"`
	CampaignIdString string  `json:"campaign_id_string"`
	KeywordText      string  `json:"keywordText"`
	MatchType        string  `json:"matchType"`
	State            string  `json:"state"`
	Bid              float64 `json:"bid"`
}

func (p *AdgroupsSuggestedKeywords) UnmarshalJSON(data []byte) error {
	type xAdgroupsSuggestedKeywords AdgroupsSuggestedKeywords
	x := &xAdgroupsSuggestedKeywords{}
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	x.AdGroupIdString = strconv.Itoa(int(x.AdGroupID))
	x.CampaignIdString = strconv.Itoa(int(x.CampaignID))
	*p = AdgroupsSuggestedKeywords(*x)

	return nil
}
