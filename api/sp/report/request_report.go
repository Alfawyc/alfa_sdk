package report

import (
	"alfa_sdk/api"
	"net/http"
	"net/url"
	"strings"
)

var _ api.IParameter = (*RequestReport)(nil)

type RequestReport struct {
	Method     string
	Path       string
	PathParams map[string]string
	Body       *RequestBodyReq
	Query      url.Values
}

func (r *RequestReport) Validate() error {
	return nil
}

func (r *RequestReport) Build() *api.Params {
	return &api.Params{
		Method:     r.Method,
		Path:       r.Path,
		PathParams: r.PathParams,
		Body:       r.Body,
		Query:      r.Query,
	}
}

func NewRequestReport() *RequestReport {
	return &RequestReport{
		Method:     http.MethodPost,
		Path:       RequestReportURL,
		PathParams: map[string]string{},
		Body:       &RequestBodyReq{},
		Query:      url.Values{},
	}
}

// SetBody ...
func (r *RequestReport) SetBody(in *RequestBodyReq, reportType, formatType string) *RequestReport {
	metric := ReportMetrics[formatType]

	r.Body.Metrics = strings.Join(metric, ",")
	r.Body.ReportDate = in.ReportDate
	if reportType == "asins" {
		r.Body.CampaignType = "sponsoredProducts"
	}

	if reportType == "keywords_query" {
		r.Body.Segment = "query"
	}

	if reportType == "campaigns_placement" {
		r.Body.Segment = "placement"
	}

	if in.StateFilter != "" {
		r.Body.StateFilter = in.StateFilter
	}

	return r
}

// SetReportType ...
func (r *RequestReport) SetReportType(reportType string) *RequestReport {
	r.PathParams["recordType"] = reportType

	return r
}
