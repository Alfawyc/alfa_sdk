package negativeproducttarget

import (
	"encoding/json"
	"strconv"
)

var (
	NegativeProductTargetExtendURL = "/v2/sp/negativeTargets/extended"
)

type NegativeProductTarget struct {
	TargetID         int64         `json:"targetId"`
	TargetIdString   string        `json:"target_id_string"`
	CampaignID       int64         `json:"campaignId"`
	CampaignIdString string        `json:"campaign_id_string"`
	AdGroupID        int64         `json:"adGroupId"`
	AdGroupIdString  string        `json:"ad_group_id_string"`
	State            string        `json:"state"`
	Expression       []*Expression `json:"expression"`
	ExpressionType   string        `json:"expressionType"`
	CreationDate     int64         `json:"creationDate"`
	LastUpdatedDate  int64         `json:"lastUpdatedDate"`
	ServingStatus    string        `json:"servingStatus"`
}

func (p *NegativeProductTarget) UnmarshalJSON(data []byte) error {
	type xNegativeProductTarget NegativeProductTarget
	x := &xNegativeProductTarget{}
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	x.TargetIdString = strconv.Itoa(int(x.TargetID))
	x.CampaignIdString = strconv.Itoa(int(x.CampaignID))
	x.AdGroupIdString = strconv.Itoa(int(x.AdGroupID))
	*p = NegativeProductTarget(*x)

	return nil
}

type Expression struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}
