package db

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
	ccipmocks "github.com/smartcontractkit/chainlink/v2/core/services/ccip/mocks"
)

func Test_priceCleanup(t *testing.T) {
	lggr := logger.TestLogger(t)
	destChainSelector := uint64(12345)

	testCases := []struct {
		name            string
		gasPriceError   bool
		tokenPriceError bool
		expectedErr     bool
	}{
		{
			name:            "ORM called successfully",
			gasPriceError:   false,
			tokenPriceError: false,
			expectedErr:     false,
		},
		{
			name:            "gasPrice clear failed",
			gasPriceError:   true,
			tokenPriceError: false,
			expectedErr:     true,
		},
		{
			name:            "tokenPrice clear failed",
			gasPriceError:   false,
			tokenPriceError: true,
			expectedErr:     true,
		},
		{
			name:            "both ORM calls failed",
			gasPriceError:   true,
			tokenPriceError: true,
			expectedErr:     true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := tests.Context(t)

			var gasPricesError error
			var tokenPricesError error
			if tc.gasPriceError {
				gasPricesError = fmt.Errorf("gas prices error")
			}
			if tc.tokenPriceError {
				tokenPricesError = fmt.Errorf("token prices error")
			}

			mockOrm := ccipmocks.NewORM(t)
			mockOrm.On("ClearGasPricesByDestChain", ctx, destChainSelector, priceExpireSec).Return(gasPricesError).Once()
			mockOrm.On("ClearTokenPricesByDestChain", ctx, destChainSelector, priceExpireSec).Return(tokenPricesError).Once()

			priceCleanup := NewPriceCleanup(lggr, mockOrm, destChainSelector).(*priceCleanup)
			err := priceCleanup.clean(ctx)
			if tc.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_priceCleanupInBackground(t *testing.T) {
	lggr := logger.TestLogger(t)
	destChainSelector := uint64(12345)

	expectedPriceExpireSec := 1
	expectedCleanupInterval := 1 * time.Second

	mockOrm := ccipmocks.NewORM(t)

	priceCleanup := NewPriceCleanup(lggr, mockOrm, destChainSelector).(*priceCleanup)
	priceCleanup.priceExpireSec = expectedPriceExpireSec
	priceCleanup.cleanupInterval = expectedCleanupInterval

	mockOrm.On("ClearGasPricesByDestChain", priceCleanup.backgroundCtx, destChainSelector, expectedPriceExpireSec).Return(nil).Once()
	mockOrm.On("ClearTokenPricesByDestChain", priceCleanup.backgroundCtx, destChainSelector, expectedPriceExpireSec).Return(nil).Once()
	assert.NoError(t, priceCleanup.Start(tests.Context(t)))

	for i := 0; i < 10; i++ {
		mockOrm.On("ClearGasPricesByDestChain", priceCleanup.backgroundCtx, destChainSelector, expectedPriceExpireSec).Return(nil).Once()
		mockOrm.On("ClearTokenPricesByDestChain", priceCleanup.backgroundCtx, destChainSelector, expectedPriceExpireSec).Return(nil).Once()
		time.Sleep(expectedCleanupInterval)
	}

	assert.NoError(t, priceCleanup.Close())
	// No more calls after closing
	time.Sleep(expectedCleanupInterval)
}
