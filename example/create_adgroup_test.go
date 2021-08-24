package example

import (
	"alfa_sdk/api/sp/adgroups"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateAdGroup(t *testing.T) {
	t.Run("createAdGroup", func(t *testing.T) {

		req := getSellerRequest(t, SandBxoScopeId)
		list := make([]*adgroups.CreateAdGroupReq, 0)
		list = append(list, &adgroups.CreateAdGroupReq{
			Name:       "广告活动20210722_2",
			CampaignId: 109238869710315,
			DefaultBid: 0.02,
			State:      "enabled",
		})
		res := make([]adgroups.AdGroupResponse, 0)
		createAdGroup := adgroups.NewCreateAdGroup().SetBody(list)
		resp, err := req.Execute(context.Background(), createAdGroup, &res)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
		t.Logf("%v", res)
		t.Logf("%v", resp)
	})
}
