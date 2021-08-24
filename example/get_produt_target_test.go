package example

import (
	"alfa_sdk/api/sp/producttarget"
	"context"
	"log"
	"testing"
)

func TestGetProductTarget(t *testing.T) {
	t.Run("getProductTarget", func(t *testing.T) {
		req := getSellerRequest(t, ScopeId)
		result := make([]producttarget.ProductTarget, 0)
		productTarget := producttarget.NewGetProductTargetExtend().SetStartIndexAndCount(1, 1).SetCampaignIdFilter("38896158885101")

		resp, err := req.Execute(context.Background(), productTarget, &result)
		//assert.Nil(t, err)
		//assert.NotNil(t, resp)
		if err != nil {
			log.Println(err.Error())
		}

		t.Logf("%v", resp)
		t.Logf("%v", result)
	})
}
