package example

import (
	"alfa_sdk/api/sp/producttarget"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetOneProductTargetTest(t *testing.T) {
	t.Run("getOneProductTargetTest", func(t *testing.T) {
		req := getSellerRequest(t, ScopeId)
		result := new(producttarget.ProductTarget)
		getOneProductTargetTest := producttarget.NewGetProductTargetSpecified().SetTargetId("175233455496634")
		resp, err := req.Execute(context.Background(), getOneProductTargetTest, &result)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
		t.Logf("%v", resp)
		t.Logf("%v", result)
	})
}
