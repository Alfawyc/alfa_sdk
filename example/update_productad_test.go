package example

import (
	"alfa_sdk/api/sp/productad"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdateProductAdTest(t *testing.T) {
	t.Run("updateProductAdTest", func(t *testing.T) {
		req := getSellerRequest(t, SandBxoScopeId)
		list := make([]*productad.UpdateProductAdReq, 0)
		list = append(list, &productad.UpdateProductAdReq{
			AdId:  0,
			State: "enabled",
		})
		updateProductAd := productad.NewUpdateProductAd().SetBody(list)
		res := make([]productad.ProductAdResponse, 0)
		resp, err := req.Execute(context.Background(), updateProductAd, &res)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
		t.Logf("%v", res)
		t.Logf("%v", resp)
	})
}
