package alfa_sdk

import (
	"alfa_sdk/cache"
	"alfa_sdk/sdkerr"
	"alfa_sdk/token"
	"context"
	"strings"
	"time"
)

type Client struct {
	option *Options
	token  *token.Client
	cache  cache.Cache
}

func NewClient(opts ...Option) (*Client, error) {
	o := Options{
		Debug:            true,
		Credentials:      nil,
		UserSandBox:      false,
		Timeout:          5 * time.Minute,
		AutoRequestRetry: true,
		AutoRetryCount:   3,
		RetryWaitTime:    100 * time.Microsecond,
		RetryMaxWaitTime: 5 * time.Second,
	}
	for _, opt := range opts {
		if err := opt(&o); err != nil {
			return nil, err
		}
	}
	c := &Client{
		option: &o,
	}

	return c, nil
}

func (c *Client) SetCache(cache cache.Cache) {
	c.cache = cache
}

func (c *Client) preNewRequest(seller *Seller) {
	if c.cache == nil {
		c.cache = cache.NewMemoryCache()
	}

	if c.token == nil {
		c.token = token.NewClient(
			&token.Config{
				ClientID:     c.option.Credentials.ClientId,
				ClientSecret: c.option.Credentials.ClientSecret,
				Region:       seller.Region,
				Debug:        c.option.Debug,
			}, c.cache,
		)
	}
}

func (c *Client) newRequest(ctx context.Context, seller *Seller) (*Request, error) {
	seller.Region = strings.ToLower(seller.Region)

	if seller.RefreshToken == "" {
		return nil, sdkerr.NewCustomError(sdkerr.NoRefreshTokenProvided, `please provide a refresh token create client`)
	}
	r := &Request{
		option:       c.option,
		token:        c.token,
		endpoint:     getRegionEndpoint(seller.Region, c.option.UserSandBox),
		region:       seller.Region,
		refreshToken: seller.RefreshToken,
		scopeId:      seller.ScopeId,
	}
	if c.option.UserSandBox {
		r.endpoint = "https://advertising-api-test.amazon.com"
	}
	_, err := r.token.GetAccessToken(ctx, r.refreshToken)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (c *Client) NewSellerRequest(ctx context.Context, seller *Seller) (*Request, error) {
	c.preNewRequest(seller)

	return c.newRequest(ctx, seller)
}
