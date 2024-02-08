package batchreader

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	mocks2 "github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/rpclib"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/rpclib/rpclibmocks"
)

func TestTokenPoolFactory(t *testing.T) {
	latestBlock := logpoller.LogPollerBlock{
		BlockNumber:          1231230,
		BlockTimestamp:       time.Now(),
		FinalizedBlockNumber: 1231231,
		CreatedAt:            time.Now(),
	}

	lggr := logger.TestLogger(t)
	offRamp := utils.RandomAddress()
	lp := mocks2.NewLogPoller(t)
	lp.On("LatestBlock", mock.Anything).Return(latestBlock, nil)

	ctx := context.Background()
	remoteChainSelector := uint64(2000)
	batchCallerMock := rpclibmocks.NewEvmBatchCaller(t)

	tokenPoolBatchReader := NewTokenPoolBatchedReader(lggr, remoteChainSelector, offRamp, batchCallerMock, lp)

	poolTypes := []string{"BurnMint", "LockRelease"}

	rateLimits := ccipdata.TokenBucketRateLimit{
		Tokens:      big.NewInt(333333),
		LastUpdated: 33,
		IsEnabled:   true,
		Capacity:    big.NewInt(666666),
		Rate:        big.NewInt(444444),
	}

	for _, versionStr := range []string{ccipdata.V1_0_0, ccipdata.V1_1_0, ccipdata.V1_2_0, ccipdata.V1_4_0} {
		gotRateLimits, err := tokenPoolBatchReader.GetInboundTokenPoolRateLimits(ctx, []common.Address{})
		require.NoError(t, err)
		assert.Empty(t, gotRateLimits)

		var batchCallResult []rpclib.DataAndErr
		for _, poolType := range poolTypes {
			batchCallResult = append(batchCallResult, rpclib.DataAndErr{
				Outputs: []any{poolType + " " + versionStr},
				Err:     nil,
			})
		}

		// TypeAndVersion uses finalized block number
		batchCallerMock.On("BatchCall", mock.Anything, uint64(latestBlock.FinalizedBlockNumber), mock.Anything).Return(batchCallResult, nil)
		// Rate limits use latest block number
		batchCallerMock.On("BatchCall", mock.Anything, uint64(latestBlock.BlockNumber), mock.Anything).Return([]rpclib.DataAndErr{{
			Outputs: []any{rateLimits},
			Err:     nil,
		}, {
			Outputs: []any{rateLimits},
			Err:     nil,
		}}, nil).Once()

		var poolAddresses []common.Address

		for i := 0; i < len(poolTypes); i++ {
			poolAddresses = append(poolAddresses, utils.RandomAddress())
		}

		gotRateLimits, err = tokenPoolBatchReader.GetInboundTokenPoolRateLimits(ctx, poolAddresses)
		require.NoError(t, err)
		assert.Len(t, gotRateLimits, len(poolTypes))

		for _, gotRateLimit := range gotRateLimits {
			assert.Equal(t, rateLimits, gotRateLimit)
		}
	}
}
