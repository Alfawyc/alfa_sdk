package cache

import (
	"alfa_sdk/sdkerr"
	"alfa_sdk/types"
	"context"
	"encoding/json"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"

	"time"
)

type RedisCache struct {
	cache *cache.Cache
}

func NewRedisCache(addr string, db int, password ...string) Cache {
	opts := &redis.Options{
		Addr: addr,
		DB:   db,
	}
	if len(password) > 0 {
		opts.Password = password[0]
	}
	rdb := redis.NewClient(opts)
	return &RedisCache{
		cache: cache.New(&cache.Options{
			Redis:     rdb,
			Marshal:   json.Marshal,
			Unmarshal: json.Unmarshal,
		}),
	}
}

func (r *RedisCache) SetToken(ctx context.Context, key string, token *types.Token, expiration time.Duration) error {
	if token.Expiry.Round(0).Add(-expiryDelta).Before(time.Now()) {
		return sdkerr.NewCustomError(sdkerr.TokenExpired, "token expired")
	}
	err := r.cache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: token,
		TTL:   expiration,
	})
	if err != nil {
		return sdkerr.NewCustomErrorf(sdkerr.CacheSetError, "cache set error:%s", err)
	}
	return nil
}

func (r *RedisCache) GetToken(ctx context.Context, key string) (*types.Token, bool) {
	var t *types.Token
	err := r.cache.Get(ctx, key, &t)
	if err != nil {
		return nil, false
	}
	return t, true
}
