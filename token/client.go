package token

import (
	"alfa_sdk/cache"
	"alfa_sdk/sdkerr"
	"alfa_sdk/types"
	"context"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	tokenUrl     string
	clientId     string
	clientSecret string
	debug        bool
	cache        cache.Cache
}

func NewClient(c *Config, cache ...cache.Cache) *Client {

	client := &Client{
		tokenUrl:     "https://api.amazon.com/auth/o2/token",
		clientId:     c.ClientID,
		clientSecret: c.ClientSecret,
		debug:        c.Debug,
	}

	if len(cache) > 0 {
		client.cache = cache[0]
	}

	return client
}

func (c *Client) Exchange(ctx context.Context, authCode string, redirectUri string) (*types.Token, error) {

	if authCode == "" {
		return nil, sdkerr.NewCustomError(sdkerr.NoAuthCodeProvided, `please provide an authorization code received from a "getAuthorizationCode" operation to exchange it for a "refresh_token".`)
	}

	requestMeta := types.RequestMeta{
		GrantType:    "authorization_code",
		ClientId:     c.clientId,
		ClientSecret: c.clientSecret,
		Code:         authCode,
		RedirectUri:  redirectUri,
	}

	token, err := c.retrieveToken(ctx, requestMeta)
	if err != nil {
		return nil, err
	}

	if c.cache != nil {
		err := c.cache.SetToken(ctx, token.RefreshToken, token, token.ExpiryDuration())
		if err != nil {
			return nil, err
		}
	}

	return token, nil
}

// GetAccessToken get AccessToken from api
func (c *Client) GetAccessToken(ctx context.Context, refreshToken string) (*types.Token, error) {

	if refreshToken == "" {
		return nil, sdkerr.NewCustomError(sdkerr.NoRefreshTokenProvided, "please provide an authorization refresh_token")
	}

	if c.cache != nil {
		if x, ok := c.cache.GetToken(ctx, refreshToken); ok {
			return x, nil
		}
	}

	return c.GetAccessTokenSkipCache(ctx, refreshToken)
}

func (c *Client) GetAccessTokenSkipCache(ctx context.Context, refreshToken string) (*types.Token, error) {

	if refreshToken == "" {
		return nil, sdkerr.NewCustomError(sdkerr.NoRefreshTokenProvided, "please provide an authorization refresh_token")
	}

	requestMeta := types.RequestMeta{
		GrantType:    "refresh_token",
		ClientId:     c.clientId,
		ClientSecret: c.clientSecret,
		RefreshToken: refreshToken,
	}

	token, err := c.retrieveToken(ctx, requestMeta)
	if err != nil {
		return nil, err
	}

	if c.cache != nil {
		err = c.cache.SetToken(ctx, token.RefreshToken, token, token.ExpiryDuration())
		if err != nil {
			return nil, err
		}
	}

	return token, nil
}

// retrieveToken send request
func (c *Client) retrieveToken(ctx context.Context, requestMeta types.RequestMeta) (*types.Token, error) {

	token := &types.Token{}

	client := resty.New()
	client.SetDebug(c.debug)

	resp, err := client.R().
		SetContext(ctx).
		SetError(&types.ErrResp{}).
		SetBody(requestMeta).
		SetResult(&token).
		Post(c.tokenUrl)

	if err != nil {
		return nil, sdkerr.NewCustomErrorf(sdkerr.RequestTokenError, "%v", err)
	}

	if resp.IsError() {
		if errResponse, ok := resp.Error().(*types.ErrResp); ok {
			return nil, errResponse.ToSdkError()
		}
	}

	if token.AccessToken == "" {
		return nil, sdkerr.NewCustomError(sdkerr.UnknownRequestTokenError, resp.String())
	}

	if token.RefreshToken == "" {
		token.RefreshToken = requestMeta.RefreshToken
	}

	return token, nil
}
