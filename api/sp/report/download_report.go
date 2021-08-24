package report

import (
	"alfa_sdk/api"
	"net/http"
	"net/url"
)

type DownloadReport struct {
	Method     string
	Path       string
	PathParams map[string]string
	Body       interface{}
	Query      url.Values
}

func (r *DownloadReport) Validate() error {

	return nil
}

func (r *DownloadReport) Build() *api.Params {
	return &api.Params{
		Method:     r.Method,
		Path:       r.Path,
		PathParams: r.PathParams,
		Body:       r.Body,
		Query:      r.Query,
	}
}

func NewDownloadReport() *DownloadReport {
	return &DownloadReport{
		Method:     http.MethodGet,
		Path:       "/v1/reports/{reportId}/download",
		PathParams: map[string]string{},
		Body:       nil,
		Query:      url.Values{},
	}
}

// SetReportId ...
func (r *DownloadReport) SetReportId(reportId string) *DownloadReport {
	r.PathParams["reportId"] = reportId

	return r
}
