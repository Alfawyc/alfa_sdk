package api

import "net/url"

type Params struct {
	Method     string
	Path       string
	PathParams map[string]string
	Body       interface{}
	Query      url.Values
}

type IParameter interface {
	Validate() error
	Build() *Params
}
