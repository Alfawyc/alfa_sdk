package example

import (
	"alfa_sdk/api/sp/report"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetReport(t *testing.T) {
	t.Run("GetReport", func(t *testing.T) {
		req := getSellerRequest(t, SandBxoScopeId)
		getReport := report.NewGetReport().SetReportId("amzn1.clicksAPI.v1.m1.60F97B9E.bea5d044-551c-44fd-a7be-46ddf64914cb")

		resp, err := req.Execute(context.Background(), getReport, nil)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
		t.Logf("%v", resp)
	})
}
