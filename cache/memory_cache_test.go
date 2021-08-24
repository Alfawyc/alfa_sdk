package cache

import (
	"alfa_sdk/types"
	"context"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

func TestMemoryCache(t *testing.T) {
	ctx := context.Background()
	c := NewMemoryCache()

	t.Parallel()

	t.Run("TestMemoryCacheToken", func(t *testing.T) {
		for i := 0; i < 100; i++ {
			err := c.SetToken(ctx, "Token"+strconv.Itoa(i), &types.Token{
				AccessToken: "test",
				Expiry:      time.Now().Add(time.Hour),
			}, 5*time.Second)
			assert.NoError(t, err)
		}
		for i := 0; i < 100; i++ {
			x, found := c.GetToken(ctx, "Token"+strconv.Itoa(i))
			assert.Equal(t, found, true)
			assert.NotNil(t, x)
			assert.Equal(t, x.AccessToken, "test")
		}
	})
}
