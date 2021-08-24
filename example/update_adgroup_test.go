package example

import (
	"alfa_sdk/api/sp/adgroups"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdateAdGroup(t *testing.T) {
	t.Run("updateAdGroup", func(t *testing.T) {
		req := getSellerRequest(t, SandBxoScopeId)
		list := make([]*adgroups.UpdateAdGroupReq, 0)
		list = append(list, &adgroups.UpdateAdGroupReq{
			AdGroupId: 125214312528561,
			Name:      "广告活动20210722_2",
			State:     "paused",
		})
		res := make([]adgroups.AdGroupResponse, 0)
		updateAdGroup := adgroups.NewUpdateAdGroup().SetBody(list)
		resp, err := req.Execute(context.Background(), updateAdGroup, &res)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
		t.Logf("%v", resp)
		t.Logf("%v", res)
	})
}
