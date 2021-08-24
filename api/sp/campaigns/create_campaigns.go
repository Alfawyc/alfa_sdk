package campaigns

import (
	"alfa_sdk/api"
	"net/http"
	"net/url"
)

var _ api.IParameter = (*CreateCampaigns)(nil)

type CreateCampaigns struct {
	Method     string
	Path       string
	PathParams map[string]string
	Body       []*CreateCampaignReq
	Query      url.Values
}

func (c *CreateCampaigns) Validate() error {
	return nil
}

func (c *CreateCampaigns) Build() *api.Params {
	return &api.Params{
		Method:     c.Method,
		Path:       c.Path,
		PathParams: c.PathParams,
		Body:       c.Body,
		Query:      c.Query,
	}
}

func NewCreateCampaigns() *CreateCampaigns {
	body := make([]*CreateCampaignReq, 0)
	return &CreateCampaigns{
		Method:     http.MethodPost,
		Path:       CampaignsURL,
		PathParams: map[string]string{},
		Body:       body,
		Query:      url.Values{},
	}
}

// SetBody ...
func (c *CreateCampaigns) SetBody(in []*CreateCampaignReq) *CreateCampaigns {
	for _, val := range in {
		temp := &CreateCampaignReq{
			PortfolioId: val.PortfolioId,
			Name:        val.Name,
			Tags: &Tags{
				PONumber:       val.Tags.PONumber,
				AccountManager: val.Tags.AccountManager,
			},
			CampaignType:  val.CampaignType,
			TargetingType: val.TargetingType,
			State:         val.State,
			DailyBudget:   val.DailyBudget,
			StartDate:     val.StartDate,
			EndDate:       val.EndDate,
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
