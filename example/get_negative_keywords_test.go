package example

import (
	"alfa_sdk/api/sp/negativekeywords"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetNegativeKeyword(t *testing.T) {
	t.Run("getNegativeKeyword", func(t *testing.T) {
		req := getSellerRequest(t, ScopeId)

		getNegativeKeyword := negativekeywords.NewGetNegativeKeyword().SetIndexAndCount(1, 1)

		result := make([]negativekeywords.NegativeKeyword, 0)
		resp, err := req.Execute(context.Background(), getNegativeKeyword, &result)

		assert.Nil(t, err)
		assert.NotNil(t, resp)

		t.Logf("%v", resp)
		t.Logf("%v", result)
	})
}
