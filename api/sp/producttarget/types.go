package producttarget

import (
	"encoding/json"
	"strconv"
)

var (
	ProductTargetURL               = "/v2/sp/targets"
	ProductTargetExtendURL         = "/v2/sp/targets/extended"
	TargetProductRecommendationURL = "/v2/sp/targets/productRecommendations"
)

type ProductTarget struct {
	TargetId           int64                `json:"targetId"`
	TargetStringId     string               `json:"target_string_id"`
	CampaignId         int64                `json:"campaignId"`
	CampaignStringId   string               `json:"campaign_string_id"`
	AdGroupId          int64                `json:"adGroupId"`
	AdGroupStringId    string               `json:"ad_group_string_id"`
	State              string               `json:"state"`
	Expression         []Expression         `json:"expression"`
	ResolvedExpression []ResolvedExpression `json:"resolvedExpression"`
	ExpressionType     string               `json:"expressionType"`
	Bid                float64              `json:"bid"`
	CreationDate       int64                `json:"creationDate"`
	LastUpdatedDate    int64                `json:"lastUpdatedDate"`
	ServingStatus      string               `json:"servingStatus"`
}

type Expression struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}

type ResolvedExpression struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}

type TargetProductRecommendation struct {
	TotalResultCount    int64                 `json:"totalResultCount"`
	RecommendedProducts []RecommendedProducts `json:"recommendedProducts"`
}

type RecommendedProducts struct {
	RecommendedTargetAsin string `json:"recommendedTargetAsin"`
}

type GetTargetProductRecommendationReq struct {
	PageSize   int64    `json:"pageSize"`
	PageNumber int64    `json:"pageNumber"`
	Asins      []string `json:"asins"`
}

func (p *ProductTarget) UnmarshalJSON(data []byte) error {
	type xProductTarget ProductTarget
	x := &xProductTarget{}
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	x.CampaignStringId = strconv.Itoa(int(x.CampaignId))
	x.AdGroupStringId = strconv.Itoa(int(x.AdGroupId))
	x.TargetStringId = strconv.Itoa(int(x.TargetId))
	*p = ProductTarget(*x)

	return nil
}
