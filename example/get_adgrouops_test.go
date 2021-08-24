package example

import (
	"alfa_sdk/api/sp/adgroups"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAdGroups(t *testing.T) {
	t.Run("GetAdGroups", func(t *testing.T) {
		req := getSellerRequest(t, SandBxoScopeId)
		getAdGroups := adgroups.NewGetAdGroupsExtended().SetStartIndexAndCount(1, 1)
		res := make([]adgroups.AdGroups, 0)
		resp, err := req.Execute(context.Background(), getAdGroups, &res)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
		t.Logf("%v", resp)
		t.Logf("%v", res)
	})
}
