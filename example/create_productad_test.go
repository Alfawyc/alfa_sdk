package example

import (
	"alfa_sdk/api/sp/productad"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateProductAd(t *testing.T) {
	t.Run("createProductAd", func(t *testing.T) {
		req := getSellerRequest(t, SandBxoScopeId)
		list := make([]*productad.CreateProductAdReq, 0)
		list = append(list, &productad.CreateProductAdReq{
			CampaignId: 109238869710315,
			AdGroupId:  213079305122647,
			Sku:        "XT-J179-5Q5T",
			State:      "enabled",
		})
		createProductAd := productad.NewCreateProductAd().SetBody(list)
		res := make([]productad.ProductAdResponse, 0)
		resp, err := req.Execute(context.Background(), createProductAd, &res)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
		t.Logf("%v", res)
		t.Logf("%v", resp)
	})
}
