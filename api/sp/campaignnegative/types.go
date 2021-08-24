package campaignnegative

import (
	"encoding/json"
	"strconv"
)

var (
	CampaignNegativeURL         string = "/v2/sp/campaignNegativeKeywords"
	CampaignNegativeURLExtended string = "/v2/sp/campaignNegativeKeywords/extended"
)

type CampaignNegativeKeyword struct {
	KeywordId        int64  `json:"keywordId"`
	KeywordIdString  string `json:"keyword_string_id"`
	CampaignId       int64  `json:"campaignId"`
	CampaignIdString string `json:"campaign_string_id"`
	State            string `json:"state"`
	KeywordText      string `json:"keywordText"`
	MatchType        string `json:"matchType"`
	CreationDate     int64  `json:"creationDate"`
	LastUpdatedDate  int64  `json:"lastUpdatedDate"`
	ServingStatus    string `json:"servingStatus"`
}

// CreateCampaignNegativeKeywordReq create campaign negative keyword req
type CreateCampaignNegativeKeywordReq struct {
	CampaignId  int64  `json:"campaignId"`
	State       string `json:"state"`
	KeywordText string `json:"keywordText"`
	MatchType   string `json:"matchType"`
}

// UpdateCampaignNegativeKeywordReq update campaign negative keyword req
type UpdateCampaignNegativeKeywordReq struct {
	KeywordId int64  `json:"keywordId"`
	State     string `json:"state"`
}

type CampaignNegativeKeywordResponse struct {
	KeywordId   int64  `json:"keywordId"`
	Code        string `json:"code"`
	Details     string `json:"details"`
	Description string `json:"description"`
}

func (p *CampaignNegativeKeyword) UnmarshalJSON(data []byte) error {
	type xCampaignNegativeKeyword CampaignNegativeKeyword
	x := &xCampaignNegativeKeyword{}
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	x.CampaignIdString = strconv.Itoa(int(x.CampaignId))
	x.KeywordIdString = strconv.Itoa(int(x.KeywordId))
	*p = CampaignNegativeKeyword(*x)

	return nil
}
