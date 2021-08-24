package adgroups

import (
	"alfa_sdk/api"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

var _ api.IParameter = (*DeleteAdGroup)(nil)

type DeleteAdGroup struct {
	Method     string
	Path       string
	PathParams map[string]string
	Body       interface{}
	Query      url.Values
}

func (a *DeleteAdGroup) Validate() error {
	return nil
}

func (a *DeleteAdGroup) Build() *api.Params {
	return &api.Params{
		Method:     a.Method,
		Path:       a.Path,
		PathParams: a.PathParams,
		Body:       a.Body,
		Query:      a.Query,
	}
}

func NewDeleteAdGroup() *DeleteAdGroup {
	return &DeleteAdGroup{
		Method:     http.MethodDelete,
		Path:       fmt.Sprintf("%s/{adGroupId}", AdGroupsURL),
		PathParams: map[string]string{},
		Body:       nil,
		Query:      url.Values{},
	}
}

// SetAdGroupId ...
func (a *DeleteAdGroup) SetAdGroupId(v int64) *DeleteAdGroup {
	a.PathParams["adGroupId"] = strconv.Itoa(int(v))

	return a
}
