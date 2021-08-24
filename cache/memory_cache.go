package cache

import (
	"alfa_sdk/sdkerr"
	"alfa_sdk/types"
	"context"
	"github.com/patrickmn/go-cache"
	"time"
)

type MemoryCache struct {
	cache *cache.Cache
}

func NewMemoryCache() Cache {
	return &MemoryCache{
		cache: cache.New(defaultExpiration, cleanupInterval),
	}
}

func (m *MemoryCache) SetToken(ctx context.Context, key string, token *types.Token, expiration time.Duration) error {
	if token.Expiry.Round(0).Add(-expiryDelta).Before(time.Now()) {
		return sdkerr.NewCustomError(sdkerr.TokenExpired, "token expired")
	}
	m.cache.Set(key, token, expiration-expiryDelta)

	return nil
}

func (m *MemoryCache) GetToken(ctx context.Context, key string) (*types.Token, bool) {
	if x, found := m.cache.Get(key); found {
		return x.(*types.Token), true
	}

	return nil, false
}
