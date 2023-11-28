package ccipdata

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_offramp_1_0_0"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/rpclib"
)

func TestOffRampV1_0_0_getPoolsByDestTokens(t *testing.T) {
	rpcURL := ""
	offRampAddr := common.HexToAddress("")

	if rpcURL == "" {
		t.Skip("set rpc url to run the test")
	}

	ctx := testutils.Context(t)
	rpcClient, err := ethclient.Dial(rpcURL)
	assert.NoError(t, err)

	batchCaller := rpclib.NewDynamicLimitedBatchCaller(logger.TestLogger(t), rpcClient.Client(), 20, 2)

	blockNum, err := rpcClient.BlockNumber(ctx)
	assert.NoError(t, err)

	lp := mocks.NewLogPoller(t)
	lp.On("LatestBlock", mock.Anything).Return(logpoller.LogPollerBlock{FinalizedBlockNumber: int64(blockNum)}, nil)

	offRampClient, err := evm_2_evm_offramp_1_0_0.NewEVM2EVMOffRamp(offRampAddr, rpcClient)
	assert.NoError(t, err)

	o := OffRampV1_0_0{
		offRamp:        offRampClient,
		addr:           offRampAddr,
		evmBatchCaller: batchCaller,
		lp:             lp,
	}

	destTokens, err := o.offRamp.GetDestinationTokens(&bind.CallOpts{Context: ctx})
	assert.NoError(t, err)
	t.Log(">>>>> dest tokens:", destTokens)

	pools, err := o.getPoolsByDestTokens(ctx, destTokens)
	assert.NoError(t, err)
	assert.Equal(t, len(destTokens), len(pools))
	t.Log(">>>>> dest pools:", pools)
}
