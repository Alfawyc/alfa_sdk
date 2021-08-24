package example

import (
	"alfa_sdk/api/portfolios"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPortfolios(t *testing.T) {
	t.Run("getPortfolios", func(t *testing.T) {
		req := getSellerRequest(t, SandBxoScopeId)
		getPortfolios := portfolios.NewGetPortfolios()
		result := make([]portfolios.Portfolios, 0)
		resp, err := req.Execute(context.Background(), getPortfolios, &result)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
		t.Logf("%v", resp)
		t.Logf("%v", result)
	})
}
