package ccip

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func Test_tokenPriceUpdatesCache(t *testing.T) {
	tk := common.HexToAddress("1")
	ts := time.Now().Truncate(time.Second)

	c := newTokenPriceUpdatesCache()
	assert.Equal(t, time.Time{}, c.mostRecentTs())

	c.updateIfMoreRecent(ts, tk, big.NewInt(100))
	assert.Equal(t, ts, c.mostRecentTs())
	v := c.get(time.Time{})
	assert.Equal(t, big.NewInt(100), v[tk].value)

	// should not get updated if ts is older
	c.updateIfMoreRecent(ts.Add(-1*time.Minute), tk, big.NewInt(101))
	v = c.get(time.Time{})
	assert.Equal(t, big.NewInt(100), v[tk].value)

	// should not get anything when the provided timestamp is recent
	v = c.get(time.Now())
	assert.Len(t, v, 0)
}
