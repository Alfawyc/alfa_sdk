package example

import (
	"alfa_sdk/api/sp/report"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRequestReport(t *testing.T) {
	t.Run("RequestReport", func(t *testing.T) {
		req := getSellerRequest(t, SandBxoScopeId)
		var res report.RequestReportResponse
		requestReport := report.NewRequestReport().SetBody(&report.RequestBodyReq{
			ReportDate: "20210701",
		}, "campaigns", "campaigns").SetReportType("campaigns")
		resp, err := req.Execute(context.Background(), requestReport, &res)
		//json.Unmarshal(resp.Body(), &res)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
		t.Logf("%v", resp)
		t.Logf("%v", res)
	})
}
