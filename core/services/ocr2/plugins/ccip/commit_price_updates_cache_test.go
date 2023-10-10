package ccip

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func Test_tokenPriceUpdatesCache(t *testing.T) {
	ctx := context.Background()

	tk := common.HexToAddress("1")
	ts := time.Now().Truncate(time.Second)

	t.Run("base", func(t *testing.T) {
		c := newTokenPriceUpdatesCache(ctx, time.Minute)
		assert.Equal(t, time.Time{}, c.mostRecentTs())

		c.updateIfMoreRecent(ts, tk, big.NewInt(100))
		assert.Equal(t, ts, c.mostRecentTs())
		v := c.get()
		assert.Equal(t, big.NewInt(100), v[tk].value)

		// should not get updated if ts is older
		c.updateIfMoreRecent(ts.Add(-1*time.Minute), tk, big.NewInt(101))
		v = c.get()
		assert.Equal(t, big.NewInt(100), v[tk].value)
	})

	t.Run("test expiration", func(t *testing.T) {
		c := newTokenPriceUpdatesCache(ctx, 200*time.Nanosecond) // every 1ns cache expires
		assert.Equal(t, time.Time{}, c.mostRecentTs())
		c.updateIfMoreRecent(ts, tk, big.NewInt(100))
		time.Sleep(5 * time.Millisecond)
		assert.Equal(t, time.Time{}, c.mostRecentTs()) // should have expired
		assert.Len(t, c.get(), 0)
	})

	t.Run("test expiration worker cancellation", func(t *testing.T) {
		ctx, cf := context.WithCancel(context.Background())
		c := newTokenPriceUpdatesCache(ctx, time.Nanosecond) // every 1ns cache expires
		cf()                                                 // stop the cancellation worker
		c.updateIfMoreRecent(ts, tk, big.NewInt(100))
		time.Sleep(10 * time.Nanosecond)
		assert.Equal(t, ts, c.mostRecentTs()) // should not have expired, since worker stopped
		assert.Len(t, c.get(), 1)
	})
}
