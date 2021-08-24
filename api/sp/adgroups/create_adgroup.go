package adgroups

import (
	"alfa_sdk/api"
	"net/http"
	"net/url"
)

var _ api.IParameter = (*CreateAdGroup)(nil)

type CreateAdGroup struct {
	Method     string
	Path       string
	PathParams map[string]string
	Body       []*CreateAdGroupReq
	Query      url.Values
}

func (c *CreateAdGroup) Validate() error {
	return nil
}

func (c *CreateAdGroup) Build() *api.Params {
	return &api.Params{
		Method:     c.Method,
		Path:       c.Path,
		PathParams: c.PathParams,
		Body:       c.Body,
		Query:      c.Query,
	}
}

func NewCreateAdGroup() *CreateAdGroup {
	body := make([]*CreateAdGroupReq, 0)
	return &CreateAdGroup{
		Method:     http.MethodPost,
		Path:       AdGroupsURL,
		PathParams: map[string]string{},
		Body:       body,
		Query:      url.Values{},
	}
}

// SetBody set request body
func (c *CreateAdGroup) SetBody(in []*CreateAdGroupReq) *CreateAdGroup {
	for _, val := range in {
		c.Body = append(c.Body, val)
	}

	return c
}
