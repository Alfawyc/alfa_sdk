package adgroups

import (
	"encoding/json"
	"strconv"
)

var (
	AdGroupsURL         string = "/v2/sp/adGroups"
	AdGroupsURLExtended string = "/v2/sp/adGroups/extended"
)

// AdGroups ...
type AdGroups struct {
	AdGroupId        int64   `json:"adGroupId"`
	AdGroupIdString  string  `json:"ad_group_id_string"`
	Name             string  `json:"name"`
	CampaignId       int64   `json:"campaignId"`
	CampaignIdString string  `json:"campaign_id_string"`
	DefaultBid       float64 `json:"defaultBid"`
	State            string  `json:"state"`
	CreationDate     int64   `json:"creationDate"`
	LastUpdatedDate  int64   `json:"lastUpdatedDate"`
	ServingStatus    string  `json:"servingStatus"`
}

// CreateAdGroupReq create request
type CreateAdGroupReq struct {
	Name       string  `json:"name,omitempty"`
	CampaignId int64   `json:"campaignId,omitempty"`
	DefaultBid float64 `json:"defaultBid,omitempty"`
	State      string  `json:"state,omitempty"`
}

// UpdateAdGroupReq update request
type UpdateAdGroupReq struct {
	AdGroupId  int64   `json:"adGroupId,omitempty"`
	Name       string  `json:"name,omitempty"`
	DefaultBid float64 `json:"defaultBid,omitempty"`
	State      string  `json:"state,omitempty"`
}

// AdGroupResponse response
type AdGroupResponse struct {
	AdGroupId   int64  `json:"adGroupId"`
	Code        string `json:"code"`
	Details     string `json:"details"`
	Description string `json:"description"`
}

func (p *AdGroups) UnmarshalJSON(data []byte) error {
	type xAdGroups AdGroups
	x := &xAdGroups{}
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	x.AdGroupIdString = strconv.Itoa(int(x.AdGroupId))
	x.CampaignIdString = strconv.Itoa(int(x.CampaignId))
	*p = AdGroups(*x)

	return nil
}
