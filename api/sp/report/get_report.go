package report

import (
	"alfa_sdk/api"
	"net/http"
	"net/url"
)

var _ api.IParameter = (*GetReport)(nil)

type GetReport struct {
	Method     string
	Path       string
	PathParams map[string]string
	Body       interface{}
	Query      url.Values
}

func (c *GetReport) Validate() error {

	return nil
}

func (c *GetReport) Build() *api.Params {
	return &api.Params{
		Method:     c.Method,
		Path:       c.Path,
		PathParams: c.PathParams,
		Body:       c.Body,
		Query:      c.Query,
	}
}

func NewGetReport() *GetReport {
	return &GetReport{
		Method:     http.MethodGet,
		Path:       GetReportURL,
		PathParams: map[string]string{},
		Body:       nil,
		Query:      url.Values{},
	}
}

// SetReportId ...
func (c *GetReport) SetReportId(v string) *GetReport {
	c.PathParams["reportId"] = v

	return c
}
