package profiles

import (
	"alfa_sdk/api"
	"net/http"
	"net/url"
)

var _ api.IParameter = (*GetProfiles)(nil)

type GetProfiles struct {
	Method     string
	Path       string
	PathParams map[string]string
	Body       interface{}
	Query      url.Values
}

func (p *GetProfiles) Validate() error {
	return nil
}

func (p *GetProfiles) Build() *api.Params {
	return &api.Params{
		Method:     p.Method,
		Path:       p.Path,
		PathParams: p.PathParams,
		Body:       p.Body,
		Query:      p.Query,
	}
}

func NewGetProfiles() *GetProfiles {
	return &GetProfiles{
		Method:     http.MethodGet,
		Path:       GetProfilesURL,
		PathParams: map[string]string{},
		Body:       nil,
		Query:      url.Values{},
	}
}
