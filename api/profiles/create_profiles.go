package profiles

import (
	"alfa_sdk/api"
	"net/http"
	"net/url"
)

//Note that this operation is only used for SANDBOX test environment.

var _ api.IParameter = (*CreateProfiles)(nil)

type CreateProfiles struct {
	Method     string
	Path       string
	PathParams map[string]string
	Body       *CreateProfilesReq
	Query      url.Values
}

func (c *CreateProfiles) Validate() error {
	return nil
}

func (c *CreateProfiles) Build() *api.Params {
	return &api.Params{
		Method:     c.Method,
		Path:       c.Path,
		PathParams: c.PathParams,
		Body:       c.Body,
		Query:      c.Query,
	}
}

func NewCreateProfiles() *CreateProfiles {
	return &CreateProfiles{
		Method:     http.MethodPut,
		Path:       "/v2/profiles/register",
		PathParams: map[string]string{},
		Body:       &CreateProfilesReq{},
		Query:      url.Values{},
	}
}

// SetRequestBody set request body
func (c *CreateProfiles) SetRequestBody(in *CreateProfilesReq) *CreateProfiles {
	c.Body.CountryCode = in.CountryCode

	return c
}
