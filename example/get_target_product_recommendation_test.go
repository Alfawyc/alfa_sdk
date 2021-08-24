package example

import (
	"alfa_sdk/api/sp/producttarget"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTargetProductRecommendation(t *testing.T) {
	t.Run("targetProductRecommendation", func(t *testing.T) {
		req := getSellerRequest(t, ScopeId)
		targetProductRecommendation := producttarget.NewGetTargetProductRecommendation().SetBody(&producttarget.GetTargetProductRecommendationReq{
			PageSize:   20,
			PageNumber: 0,
			Asins:      []string{"B093GNQKYB", "B092QXVCDP", "B093H7D7Q3", "B093B51W77", "B0932T9F2G", "B0939ZN38T", "B092VL4BM5", "B093H4Y352", "B0924L32CV", "B093H7519Y"},
		})
		result := new(producttarget.TargetProductRecommendation)
		resp, err := req.Execute(context.Background(), targetProductRecommendation, &result)
		assert.Nil(t, err)
		t.Logf("%v", result)
		t.Logf("%v", resp)
	})
}
