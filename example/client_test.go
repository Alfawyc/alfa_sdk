package example

import (
	amz_ad_sdk "alfa_sdk"
	"alfa_sdk/cache"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var _client *amz_ad_sdk.Client
var _sellerRequest *amz_ad_sdk.Request

func TestClient(t *testing.T) {
	t.Run("clientTest", func(t *testing.T) {
		c := getClient(t)
		assert.NotNil(t, c)
	})
	t.Run("sellerRequest", func(t *testing.T) {
		req := getSellerRequest(t, "")
		assert.NotNil(t, req)
	})
}

func getClient(t *testing.T) *amz_ad_sdk.Client {
	if _client != nil {
		return _client
	}
	c, err := amz_ad_sdk.NewClient(
		amz_ad_sdk.WithAutoRequestRetry(true),
		amz_ad_sdk.WithAutoRetryCount(1),
		amz_ad_sdk.WithDebug(true),
		amz_ad_sdk.WithTimeout(time.Minute*2),
		amz_ad_sdk.WithUseSandbox(true),
		amz_ad_sdk.WithCredentialsFromEnv(),
	)
	assert.Nil(t, err)
	assert.NotNil(t, c)
	_client = c
	_client.SetCache(cache.NewRedisCache("127.0.0.1:6379", 0))

	return c
}

func getSellerRequest(t *testing.T, scopeId string) *amz_ad_sdk.Request {
	if _sellerRequest != nil {
		return _sellerRequest
	}
	c := getClient(t)
	seller := amz_ad_sdk.Seller{
		Region:       Region,
		RefreshToken: RefreshToken,
		ScopeId:      "",
	}
	if scopeId != "" {
		seller.ScopeId = scopeId
	}
	req, err := c.NewSellerRequest(context.Background(), &seller)
	assert.Nil(t, err)
	assert.NotNil(t, req)

	return req
}
