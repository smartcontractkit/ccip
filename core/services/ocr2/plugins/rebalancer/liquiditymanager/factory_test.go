package liquiditymanager

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils/evmtest"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/models"
)

func TestNewBaseLiquidityManagerFactory(t *testing.T) {
	lp1 := mocks.NewLogPoller(t)
	lp2 := mocks.NewLogPoller(t)
	ethClient1 := evmtest.NewEthClientMock(t)
	ethClient2 := evmtest.NewEthClientMock(t)
	lmf := NewBaseLiquidityManagerFactory(
		WithEvmDep(models.NetworkID(1), lp1, ethClient1),
		WithEvmDep(models.NetworkID(2), lp2, ethClient2),
	)
	assert.Len(t, lmf.evmDeps, 2)
}
