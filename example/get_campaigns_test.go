package example

import (
	"alfa_sdk/api/sp/campaigns"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCampaignsExtended(t *testing.T) {
	ctx := context.Background()
	t.Run("GetCampaignsExtended", func(t *testing.T) {
		req := getSellerRequest(t, "2500755174273302")
		getCampaignsExtended := campaigns.NewGetCampaignExtended().SetStartIndexAndCount(1, 5).SetCampaignIdFilter("9296337941131")
		res := make([]campaigns.Campaigns, 0)

		resp, err := req.Execute(ctx, getCampaignsExtended, &res)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
		t.Logf("%v", resp)
		t.Logf("%v", res)
	})
}
