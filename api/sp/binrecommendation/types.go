package binrecommendation

var (
	TargetBidRecommendationURL   = "/v2/sp/targets/bidRecommendations"
	AdGroupBidRecommendationURL  = "/v2/sp/adGroups/{adGroupId}/bidRecommendations"
	KeywordsBidRecommendationURL = "/v2/sp/keywords/bidRecommendations"
)

type TargetBindRecommendationReq struct {
	AdGroupID   string                `json:"adGroupId"`
	Expressions [][]*TargetExpression `json:"expressions"`
}

type TargetExpression struct {
	Value string `json:"value"`
	Type  string `json:"type,omitempty"`
}

type TargetBidRecommendationResp struct {
	AdGroupID       int64             `json:"adGroupId"`
	Recommendations []Recommendations `json:"recommendations"`
}

type Recommendations struct {
	SuggestedBid SuggestedBid     `json:"suggestedBid"`
	Expressions  TargetExpression `json:"expressions"`
	Code         string           `json:"code"`
}

type SuggestedBid struct {
	Suggested  float64 `json:"suggested"`
	RangeStart float64 `json:"rangeStart"`
	RangeEnd   float64 `json:"rangeEnd"`
}

type AdGroupBidRecommendation struct {
	AdGroupId    int64        `json:"adGroupId"`
	SuggestedBid SuggestedBid `json:"suggestedBid"`
}

type KeywordsBidRecommendationReq struct {
	AdGroupId string      `json:"adGroupId"`
	Keywords  []*Keywords `json:"keywords"`
}

type KeywordsBidRecommendationResp struct {
	AdGroupId       int64                     `json:"adGroupId"`
	Recommendations []*KeywordsRecommendation `json:"recommendations"`
}

type Keywords struct {
	Keyword   string `json:"keyword"`
	MatchType string `json:"matchType"`
}

type KeywordsRecommendation struct {
	Code         string       `json:"code"`
	Keyword      string       `json:"keyword"`
	MatchType    string       `json:"matchType"`
	SuggestedBid SuggestedBid `json:"suggestedBid"`
}
