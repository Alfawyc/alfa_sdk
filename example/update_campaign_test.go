package example

import (
	"alfa_sdk/api/sp/campaigns"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdateCampaign(t *testing.T) {
	t.Run("updateCampaign", func(t *testing.T) {
		req := getSellerRequest(t, SandBxoScopeId)
		adjustment := make([]campaigns.Adjustments, 0)
		adjustment = append(adjustment, campaigns.Adjustments{
			Predicate:  "placementTop",
			Percentage: 10,
		}, campaigns.Adjustments{
			Predicate:  "placementProductPage",
			Percentage: 15,
		})
		list := make([]*campaigns.UpdateCampaignReq, 0)
		list = append(list, &campaigns.UpdateCampaignReq{
			CampaignId:           9296337941131,
			PortfoliosId:         0,
			Name:                 "20210723_测试_change",
			DailyBudget:          1,
			StartDate:            "20210725",
			EndDate:              "20210726",
			PremiumBidAdjustment: false,
			Bidding: &campaigns.Bidding{
				Strategy:    "manual",
				Adjustments: adjustment,
			},
		})
		updateCampaign := campaigns.NewUpdateCampaigns().SetBody(list)
		res := make([]campaigns.CampaignResponse, 0)
		resp, err := req.Execute(context.Background(), updateCampaign, &res)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
		t.Logf("%v", resp)
		t.Logf("%v", res)
	})
}
