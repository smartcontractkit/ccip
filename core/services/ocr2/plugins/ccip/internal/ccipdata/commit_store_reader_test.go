package ccipdata

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	lpmocks "github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
)

func assertFilterRegistration(t *testing.T, lp *lpmocks.LogPoller, buildCloser func(lp *lpmocks.LogPoller, addr common.Address) Closer, numFilter int) {
	// Expected filter properties for a closer:
	// - Should be the same filter set registered that is unregistered
	// - Should be registered to the address specified
	// - Number of events specific to this component should be registered
	addr := common.HexToAddress("0x1234")
	var filters []logpoller.Filter
	lp.On("RegisterFilter", mock.Anything).Run(func(args mock.Arguments) {
		f := args.Get(0).(logpoller.Filter)
		require.Equal(t, len(f.Addresses), 1)
		require.Equal(t, f.Addresses[0], addr)
		filters = append(filters, f)
	}).Return(nil).Times(numFilter)
	c := buildCloser(lp, addr)
	for _, filter := range filters {
		lp.On("UnregisterFilter", filter.Name).Return(nil)
	}
	require.NoError(t, c.Close())
	lp.AssertExpectations(t)
}

func TestCommitFilters(t *testing.T) {
	assertFilterRegistration(t, new(lpmocks.LogPoller), func(lp *lpmocks.LogPoller, addr common.Address) Closer {
		c, err := NewCommitStoreV1_0_0(logger.TestLogger(t), addr, new(mocks.Client), lp, nil)
		require.NoError(t, err)
		return c
	}, 1)
	assertFilterRegistration(t, new(lpmocks.LogPoller), func(lp *lpmocks.LogPoller, addr common.Address) Closer {
		c, err := NewCommitStoreV1_2_0(logger.TestLogger(t), addr, new(mocks.Client), lp, nil)
		require.NoError(t, err)
		return c
	}, 1)
}
