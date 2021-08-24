package campaigns

import (
	"encoding/json"
	"strconv"
)

var (
	CampaignsURL         string = "/v2/sp/campaigns"
	CampaignsURLExtended string = "/v2/sp/campaigns/extended"
)

// Campaigns campaigns extended info struct
type Campaigns struct {
	PortfolioId          int64   `json:"portfolioId"`
	PortfolioIdString    string  `json:"portfolio_id_string"`
	CampaignId           int64   `json:"campaignId"`
	CampaignIdString     string  `json:"campaign_id_string"`
	Tags                 Tags    `json:"tags"`
	Name                 string  `json:"name"`
	CampaignType         string  `json:"campaignType"`
	TargetingType        string  `json:"targetingType"`
	State                string  `json:"state"`
	DailyBudget          float64 `json:"dailyBudget"`
	StartDate            string  `json:"startDate"`
	EndDate              string  `json:"endDate"`
	PremiumBidAdjustment bool    `json:"premiumBidAdjustment"`
	Bidding              Bidding `json:"bidding"`
	Placement            string  `json:"placement"`
	CreationDate         int64   `json:"creationDate"`
	LastUpdatedDate      int64   `json:"lastUpdatedDate"`
	ServingStatus        string  `json:"servingStatus"`
}

// Tags campaigns tags , A list of advertiser-specified custom identifiers for the campaign. Each customer identifier is a key-value pair. You can specify a maximum of 50 identifiers.
type Tags struct {
	PONumber       string `json:"PONumber"`
	AccountManager string `json:"accountManager"`
}

// Bidding campaigns biding ,Specifies bidding controls.
type Bidding struct {
	Strategy    string        `json:"strategy"`
	Adjustments []Adjustments `json:"adjustments"`
}

// Adjustments campaigns adjustments
type Adjustments struct {
	Predicate  string `json:"predicate"`
	Percentage int64  `json:"percentage"`
}

// CreateCampaignReq campaign create request
type CreateCampaignReq struct {
	//The identifier of an existing campaign to update.
	PortfolioId int64 `json:"portfolioId,omitempty"`

	//The name of the campaign.
	Name string `json:"name"`

	Tags *Tags `json:"tags,omitempty"`

	//The advertising product managed by this campaign. sponsoredProducts
	CampaignType string `json:"campaignType"`

	//The type of targeting for the campaign
	TargetingType string `json:"targetingType"`

	//The current resource state.
	State string `json:"state"`

	//A daily budget for the campaign.
	DailyBudget float64 `json:"dailyBudget"`

	//A starting date for the campaign to go live. The format of the date is YYYYMMDD.
	StartDate string `json:"startDate"`

	//An ending date for the campaign to stop running. The format of the date is YYYYMMDD.
	EndDate string `json:"endDate"`

	PremiumBidAdjustment bool     `json:"premiumBidAdjustment,omitempty"`
	Bidding              *Bidding `json:"bidding,omitempty"`
}

// UpdateCampaignReq campaign update request
type UpdateCampaignReq struct {
	//The identifier of an existing campaign to update.
	CampaignId int64 `json:"campaignId"`

	//The identifier of an existing portfolio to which the campaign is associated.
	PortfoliosId int64 `json:"portfolioId,omitempty"`

	//The name of the campaign.
	Name string `json:"name"`

	Tags *Tags `json:"tags"`

	//The current resource state.
	State string `json:"state"`

	//The daily budget of the campaign. minimum 1
	DailyBudget float64 `json:"dailyBudget"`

	//The starting date of the campaign. The format of the date is YYYYMMDD.
	StartDate string `json:"startDate"`

	//The ending date of the campaign to stop running. The format of the date is YYYYMMDD.
	EndDate string `json:"endDate"`

	PremiumBidAdjustment bool     `json:"premiumBidAdjustment,omitempty"`
	Bidding              *Bidding `json:"bidding,omitempty"`
}

// CampaignResponse campaign create response
type CampaignResponse struct {
	CampaignId  int64  `json:"campaignId"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

func (p *Campaigns) UnmarshalJSON(data []byte) error {
	type xCampaigns Campaigns
	x := &xCampaigns{}
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	x.PortfolioIdString = strconv.Itoa(int(x.PortfolioId))
	x.CampaignIdString = strconv.Itoa(int(x.CampaignId))
	*p = Campaigns(*x)

	return nil
}
