package v1_2_0

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	mocks2 "github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/rpclib"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/rpclib/rpclibmocks"
)

func TestTokenPool(t *testing.T) {
	latestBlock := logpoller.LogPollerBlock{
		BlockNumber:          1231230,
		BlockTimestamp:       time.Now(),
		FinalizedBlockNumber: 1231231,
		CreatedAt:            time.Now(),
	}

	offRamp := utils.RandomAddress()
	lp := mocks2.NewLogPoller(t)
	lp.On("LatestBlock", mock.Anything).Return(latestBlock, nil)

	ctx := context.Background()
	batchCallerMock := rpclibmocks.NewEvmBatchCaller(t)

	poolTypes := []string{"BurnMint", "LockRelease"}

	for _, tokenPoolType := range poolTypes {
		poolAddress := utils.RandomAddress()
		pool := NewTokenPool(tokenPoolType, poolAddress, offRamp, lp, batchCallerMock)

		assert.Equal(t, tokenPoolType, pool.Type())
		assert.Equal(t, poolAddress, pool.Address())

		rateLimits := ccipdata.TokenBucketRateLimit{
			Tokens:      big.NewInt(333333),
			LastUpdated: 33,
			IsEnabled:   true,
			Capacity:    big.NewInt(666666),
			Rate:        big.NewInt(444444),
		}

		call, err := pool.GetInboundTokenPoolRateLimitCall()
		require.NoError(t, err)
		assert.Equal(t, "currentOffRampRateLimiterState", call.MethodName())

		batchCallerMock.On("BatchCall", mock.Anything, uint64(latestBlock.BlockNumber), mock.Anything).Return([]rpclib.DataAndErr{{
			Outputs: []any{rateLimits},
			Err:     nil,
		}}, nil).Once()

		gotRateLimits, err := pool.GetInboundTokenPoolRateLimits(ctx, []ccipdata.TokenPoolReader{pool})
		require.NoError(t, err)
		assert.Len(t, gotRateLimits, 1)
		assert.Equal(t, rateLimits, gotRateLimits[0])

		// 2 pools

		batchCallerMock.On("BatchCall", mock.Anything, uint64(latestBlock.BlockNumber), mock.Anything).Return([]rpclib.DataAndErr{{
			Outputs: []any{rateLimits},
			Err:     nil,
		}, {
			Outputs: []any{rateLimits},
			Err:     nil,
		}}, nil).Once()

		pool2 := NewTokenPool(tokenPoolType, poolAddress, offRamp, lp, batchCallerMock)

		gotRateLimits, err = pool.GetInboundTokenPoolRateLimits(ctx, []ccipdata.TokenPoolReader{pool, pool2})
		require.NoError(t, err)
		assert.Len(t, gotRateLimits, 2)
		assert.Equal(t, rateLimits, gotRateLimits[0])
		assert.Equal(t, rateLimits, gotRateLimits[1])
	}
}
