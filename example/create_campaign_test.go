package example

import (
	"alfa_sdk/api/sp/campaigns"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateCampaign(t *testing.T) {
	t.Run("CreateCampaign", func(t *testing.T) {
		req := getSellerRequest(t, "2500755174273302")
		adjustment := make([]campaigns.Adjustments, 0)
		adjustment = append(adjustment, campaigns.Adjustments{
			Predicate:  "placementTop",
			Percentage: 10,
		}, campaigns.Adjustments{
			Predicate:  "placementProductPage",
			Percentage: 15,
		})
		list := make([]*campaigns.CreateCampaignReq, 0)
		in := campaigns.CreateCampaignReq{
			PortfolioId: 0,
			Name:        "20210723_测试_2",
			Tags: &campaigns.Tags{
				PONumber:       "",
				AccountManager: "",
			},
			CampaignType:         "sponsoredProducts",
			TargetingType:        "manual",
			State:                "enabled",
			DailyBudget:          1,
			StartDate:            "20210725",
			EndDate:              "20210726",
			PremiumBidAdjustment: true,
			Bidding: &campaigns.Bidding{
				Strategy:    "manual",
				Adjustments: adjustment,
			},
		}
		list = append(list, &in)
		createCampaign := campaigns.NewCreateCampaigns().SetBody(list)
		res := make([]campaigns.CampaignResponse, 0)
		resp, err := req.Execute(context.Background(), createCampaign, &res)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
	})
}
