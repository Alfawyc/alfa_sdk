package example

import (
	"alfa_sdk/api/portfolios"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdatePortfolios(t *testing.T) {
	t.Run("UpdatePortfolios", func(t *testing.T) {
		req := getSellerRequest(t, SandBxoScopeId)
		portfoliosList := make([]*portfolios.UpdatePortfoliosReq, 0)
		portfoliosList = append(portfoliosList, &portfolios.UpdatePortfoliosReq{
			PortfolioId: 66575202143589,
			Name:        "20210722_2",
			Budget: portfolios.Budget{
				Amount:    1,
				Policy:    "dateRange",
				StartDate: "20210723",
				EndState:  "20210724",
			},
		})
		updatePortfolios := portfolios.NewUpdatePortfolios().SetBody(portfoliosList)
		res := make([]portfolios.PortfoliosSuccessResp, 0)
		resp, err := req.Execute(context.Background(), updatePortfolios, &res)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
		t.Logf("%v", res)
		t.Logf("%v", resp)
	})
}
