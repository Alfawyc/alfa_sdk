package example

import (
	"alfa_sdk/api/sp/keywords"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdateKeywords(t *testing.T) {
	t.Run("updateKeywords", func(t *testing.T) {
		req := getSellerRequest(t, SandBxoScopeId)
		list := make([]*keywords.UpdateKeywordReq, 0)
		list = append(list, &keywords.UpdateKeywordReq{
			KeywordId: 139655578095916,
			State:     "",
			Bid:       0.03,
		})
		res := make([]keywords.KeywordResponse, 0)
		updateKeywords := keywords.NewUpdateKeywords().SetBody(list)
		resp, err := req.Execute(context.Background(), updateKeywords, &res)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
		t.Logf("%v", resp)
		t.Logf("%v", res)
	})
}
