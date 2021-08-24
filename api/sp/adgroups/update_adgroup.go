package adgroups

import (
	"alfa_sdk/api"
	"net/http"
	"net/url"
)

var _ api.IParameter = (*UpdateAdGroup)(nil)

type UpdateAdGroup struct {
	Method     string
	Path       string
	PathParams map[string]string
	Body       []*UpdateAdGroupReq
	Query      url.Values
}

func (a *UpdateAdGroup) Validate() error {
	return nil
}

func (a *UpdateAdGroup) Build() *api.Params {
	return &api.Params{
		Method:     a.Method,
		Path:       a.Path,
		PathParams: a.PathParams,
		Body:       a.Body,
		Query:      a.Query,
	}
}

func NewUpdateAdGroup() *UpdateAdGroup {
	body := make([]*UpdateAdGroupReq, 0)
	return &UpdateAdGroup{
		Method:     http.MethodPut,
		Path:       AdGroupsURL,
		PathParams: map[string]string{},
		Body:       body,
		Query:      url.Values{},
	}
}

// SetBody set request body
func (a *UpdateAdGroup) SetBody(in []*UpdateAdGroupReq) *UpdateAdGroup {
	for _, val := range in {
		a.Body = append(a.Body, val)
	}

	return a
}
