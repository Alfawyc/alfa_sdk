package example

import (
	"alfa_sdk/api/sp/campaignnegative"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateCampaignNegativeKeyword(t *testing.T) {
	t.Run("createCampaignNegativeKeyword", func(t *testing.T) {
		req := getSellerRequest(t, SandBxoScopeId)
		list := make([]*campaignnegative.CreateCampaignNegativeKeywordReq, 0)
		list = append(list, &campaignnegative.CreateCampaignNegativeKeywordReq{
			CampaignId:  109238869710315,
			State:       "enabled",
			KeywordText: "Chrome",
			MatchType:   "negativePhrase",
		})
		res := make([]campaignnegative.CampaignNegativeKeywordResponse, 0)
		createCampaign := campaignnegative.NewCreateCampaignNegativeKeyword().SetBody(list)
		resp, err := req.Execute(context.Background(), createCampaign, &res)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
		t.Logf("%v", res)
		t.Logf("%v", resp)
	})
}
