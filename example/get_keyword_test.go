package example

import (
	"alfa_sdk/api/sp/keywords"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetKeyword(t *testing.T) {
	ctx := context.Background()
	t.Run("getKeyword", func(t *testing.T) {
		req := getSellerRequest(t, SandBxoScopeId)
		getKeyword := keywords.NewGetKeywords().SetStartIndexAndCount(1, 1)
		res := make([]keywords.Keywords, 0)
		resp, err := req.Execute(ctx, getKeyword, &res)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
		t.Logf("%v", resp)
		t.Logf("%v", res)
	})
}
