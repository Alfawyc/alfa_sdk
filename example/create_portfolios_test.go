package example

import (
	"alfa_sdk/api/portfolios"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreatePortfolios(t *testing.T) {
	t.Run("CreatePortfolios", func(t *testing.T) {
		req := getSellerRequest(t, "2500755174273302")
		body := make([]*portfolios.CreatePortfoliosReq, 0)
		body = append(body, &portfolios.CreatePortfoliosReq{
			Name:  "20210722",
			State: "enabled",
		})
		createPortfolios := portfolios.NewCreatePortfolios().SetBody(body)
		res := make([]portfolios.PortfoliosSuccessResp, 0)
		resp, err := req.Execute(context.Background(), createPortfolios, &res)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
		t.Logf("%v", resp)
		t.Logf("%v", res)
	})
}
