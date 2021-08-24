package cache

import (
	"alfa_sdk/types"
	"context"
	"time"
)

const (
	// expiryDelta determines how earlier a token should be considered
	// expired than its actual expiration time. It is used to avoid late
	// expirations due to client-server time mismatches.
	expiryDelta = 60 * time.Second

	// cleanupInterval cleanup interval
	cleanupInterval = 5 * time.Second

	// defaultExpiration default expiration duration
	defaultExpiration = 3600 * time.Second
)

type Cache interface {
	SetToken(ctx context.Context, key string, token *types.Token, expiration time.Duration) error
	GetToken(ctx context.Context, key string) (*types.Token, bool)
}
