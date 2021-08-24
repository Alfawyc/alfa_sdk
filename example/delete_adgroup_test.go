package example

import (
	"alfa_sdk/api/sp/adgroups"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteAdGroup(t *testing.T) {
	t.Run("deleteAdGroup", func(t *testing.T) {
		req := getSellerRequest(t, SandBxoScopeId)
		deleteAdGroup := adgroups.NewDeleteAdGroup().SetAdGroupId(125214312528561)
		res := new(adgroups.AdGroupResponse)
		resp, err := req.Execute(context.Background(), deleteAdGroup, &res)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
		t.Logf("%v", res)
		t.Logf("%v", resp)
	})
}
