package alfa_sdk

import (
	"alfa_sdk/api"
	"alfa_sdk/sdkerr"
	"alfa_sdk/token"
	"alfa_sdk/types"
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

type Request struct {
	option       *Options
	token        *token.Client
	endpoint     string
	region       string
	sellerId     string
	refreshToken string
	scopeId      string
}

//GetAccessToken ...
func (c *Request) GetAccessToken(ctx context.Context) (*types.Token, error) {

	return c.token.GetAccessToken(ctx, c.refreshToken)
}

//retryCondition ...
func (c *Request) retryCondition(ctx context.Context) resty.RetryConditionFunc {
	return func(r *resty.Response, err error) bool {
		if c.option.AutoRequestRetry {
			//network error, auto retry
			if err != nil {
				if _, ok := err.(*url.Error); ok {
					return true
				}
				if _, ok := err.(*net.OpError); ok {
					return true
				}
			}
			if r == nil {
				return false
			}
			if r.StatusCode() == http.StatusForbidden {
				if doesMatch, _ := regexp.MatchString("access token.*expired", r.String()); doesMatch {
					_, err := c.token.GetAccessTokenSkipCache(ctx, c.refreshToken)
					if err != nil {
						log.Printf("[DEBUG] refresh token -> %s", c.refreshToken)
						return false
					}
				}
			}
		}
		return false
	}
}

func (c *Request) Execute(ctx context.Context, parameter api.IParameter, res interface{}) (*resty.Response, error) {
	if parameter == nil {
		return nil, sdkerr.NewCustomError(sdkerr.NoApiParameterProvided, "please provider api params interface")
	}

	err := parameter.Validate()
	if err != nil {
		return nil, err
	}

	params := parameter.Build()
	client := resty.New()
	client.SetError(&Error{})

	//set timeout
	client.SetTimeout(c.option.Timeout)
	//set retry
	client.SetRetryCount(c.option.AutoRetryCount)
	client.SetRetryWaitTime(c.option.RetryWaitTime)
	client.SetRetryMaxWaitTime(c.option.RetryMaxWaitTime)

	client.AddRetryCondition(c.retryCondition(ctx))

	accessToken, err := c.GetAccessToken(ctx)
	if err != nil {
		return nil, sdkerr.NewCustomError(sdkerr.RequestTokenError, "get accessToken error ,"+err.Error())
	}
	//set request Header
	headers := map[string]string{"Amazon-Advertising-API-ClientId": c.option.Credentials.ClientId, "Authorization": "bearer " + accessToken.AccessToken, "Content-Type": "application/json"}
	if c.scopeId != "" {
		headers["Amazon-Advertising-API-Scope"] = c.scopeId
	}

	req := client.R().SetHeaders(headers).SetContext(ctx).SetPathParams(params.PathParams).SetQueryParamsFromValues(params.Query).SetBody(params.Body)

	if res != nil {
		req.SetResult(res)
	}
	apiPath := fmt.Sprintf("%s%s", c.endpoint, params.Path)
	resp, err := req.Execute(params.Method, apiPath)
	if err != nil {

		return nil, err
	}

	if resp.IsError() {
		if errorResponse, ok := resp.Error().(*Error); ok {
			return nil, errorResponse.ToSdkError()
		}
	}

	return resp, nil
}

func getRegionEndpoint(region string, userSandBox bool) string {
	if userSandBox {
		return "https://advertising-api-test.amazon.com"
	}
	region = strings.ToLower(region)
	var endpoint string
	switch region {
	case "na":
		endpoint = "https://advertising-api.amazon.com"
		break
	case "eu":
		endpoint = "https://advertising-api-eu.amazon.com"
		break
	case "fe":
		endpoint = "https://advertising-api-fe.amazon.com"
		break
	}
	return endpoint
}
