package example

import (
	"alfa_sdk/api/sp/campaignnegative"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdateCampaignNegativeKeyword(t *testing.T) {
	t.Run("updateCampaignNegativeKeyword", func(t *testing.T) {
		req := getSellerRequest(t, SandBxoScopeId)
		list := make([]*campaignnegative.UpdateCampaignNegativeKeywordReq, 0)
		list = append(list, &campaignnegative.UpdateCampaignNegativeKeywordReq{
			KeywordId: 5168080635869,
		})
		res := make([]campaignnegative.CampaignNegativeKeywordResponse, 0)
		createCampaign := campaignnegative.NewUpdateCampaignNegativeKeyword().SetBody(list)
		resp, err := req.Execute(context.Background(), createCampaign, &res)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
		t.Logf("%v", res)
		t.Logf("%v", resp)
	})
}
