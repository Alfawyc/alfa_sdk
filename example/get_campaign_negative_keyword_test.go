package example

import (
	"alfa_sdk/api/sp/campaignnegative"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCampaignNegativeKeyword(t *testing.T) {
	ctx := context.Background()
	t.Run("GetCampaignNegativeKeyword", func(t *testing.T) {
		req := getSellerRequest(t, SandBxoScopeId)
		getCampaignNegativeKeyword := campaignnegative.NewGetCampaignNegativeKeyword().SetStartIndexAndCount(1, 1)
		res := make([]campaignnegative.CampaignNegativeKeyword, 0)
		resp, err := req.Execute(ctx, getCampaignNegativeKeyword, &res)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
		t.Logf("%v", resp)
		t.Logf("%v", res)
	})
}
