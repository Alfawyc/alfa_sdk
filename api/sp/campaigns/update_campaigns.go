package campaigns

import (
	"alfa_sdk/api"
	"net/http"
	"net/url"
)

type UpdateCampaigns struct {
	Method     string
	Path       string
	PathParams map[string]string
	Body       []*UpdateCampaignReq
	Query      url.Values
}

func (c *UpdateCampaigns) Validate() error {

	return nil
}

func (c *UpdateCampaigns) Build() *api.Params {
	return &api.Params{
		Method:     c.Method,
		Path:       c.Path,
		PathParams: c.PathParams,
		Body:       c.Body,
		Query:      c.Query,
	}
}

func NewUpdateCampaigns() *UpdateCampaigns {
	body := make([]*UpdateCampaignReq, 0)
	return &UpdateCampaigns{
		Method:     http.MethodPut,
		Path:       CampaignsURL,
		PathParams: map[string]string{},
		Body:       body,
		Query:      url.Values{},
	}
}

// SetBody ...
func (c *UpdateCampaigns) SetBody(in []*UpdateCampaignReq) *UpdateCampaigns {
	for _, val := range in {
		temp := &UpdateCampaignReq{
			CampaignId:   val.CampaignId,
			PortfoliosId: val.PortfoliosId,
			Name:         val.Name,
			Tags:         val.Tags,
			State:        val.State,
			DailyBudget:  val.DailyBudget,
			StartDate:    val.StartDate,
			EndDate:      val.EndDate,
		}
		if val.PremiumBidAdjustment {
			temp.PremiumBidAdjustment = val.PremiumBidAdjustment
		} else {
			temp.Bidding = val.Bidding
		}
		c.Body = append(c.Body, temp)
	}

	return c
}
