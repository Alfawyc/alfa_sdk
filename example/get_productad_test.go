package example

import (
	"alfa_sdk/api/sp/productad"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetProductAd(t *testing.T) {
	ctx := context.Background()
	t.Run("getProductAd", func(t *testing.T) {
		req := getSellerRequest(t, SandBxoScopeId)
		getProductAd := productad.NewGetProductAd().SetStartIndexAndCount(1, 1)
		res := make([]productad.ProductAd, 0)
		resp, err := req.Execute(ctx, getProductAd, &res)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
		t.Logf("%v", resp)
		t.Logf("%v", res)
	})
}
