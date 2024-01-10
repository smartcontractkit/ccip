package pricegetter

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/aggregator_v2v3_interface"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/rpclib"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/rpclib/rpclibmocks"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

func TestDynamicPriceGetter(t *testing.T) {
	ctx := testutils.Context(t)

	tk1 := utils.RandomAddress()
	tk2 := utils.RandomAddress()
	tk3 := utils.RandomAddress()

	logger.TestLogger(t).Infof("tk1: %s, tk2: %s, tk3: %s", tk1, tk2, tk3)

	cfg := DynamicPriceGetterConfig{
		AggregatorPrices: map[common.Address]AggregatorPriceConfig{
			tk1: {
				ChainID:         101,
				ContractAddress: common.HexToAddress("0xA550011"), // aggregator contract
			},
			tk2: {
				ChainID:         102,
				ContractAddress: common.HexToAddress("0xA550022"), // aggregator contract
			},
		},
		StaticPrices: map[common.Address]StaticPriceConfig{
			tk3: {
				ChainID: 103,
				Price:   1_234_000,
			},
		},
	}

	b, _ := json.MarshalIndent(cfg, "", " ")
	t.Logf("%s", b)

	// TODO create fake batch callers (see example on http server in pipeline_test).
	//serverLinkEth := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
	//	_, err := w.Write([]byte(`{"JuelsPerETH": "200000000000000000000"}`))
	//	require.NoError(t, err)
	//}))
	//defer serverLinkEth.Close()

	//ec1 := evmclimocks.NewClient(t)
	//ec1.On("BatchCallContext", mock.Anything, mock.Anything).Return().Maybe()

	//ec1.On("BatchCallContext", mock.Anything, mock.MatchedBy(func(b []rpc.BatchElem) bool {
	//	return len(b) == 1 && b[0].Method == "latestRoundData"
	//})).Return(nil).Run(func(args mock.Arguments) {
	//	elems := args.Get(1).([]rpc.BatchElem)
	//	elems[0].Result = &evmtypes.Block{
	//		Number: 42,
	//		Hash:   utils.NewHash(),
	//	}
	//	elems[1].Result = &evmtypes.Block{
	//		Number: 41,
	//		Hash:   utils.NewHash(),
	//	}
	//}).Maybe()

	////chain, chainID, err := ccipconfig.GetChainByChainSelector(chainSet, chainSelector)
	//mockChainSet := mocks.NewLegacyChainContainer(t)
	////mockChainSet.On("Get", strconv.FormatInt(testChainID, 10)).Return(mockChain, nil).Maybe()
	//mockChain := mocks.NewChain(t)
	//mockChain.On("ID").Return(big.NewInt(3)).Maybe()
	//mockChainSet.On("Get", strconv.FormatInt(3, 10)).Return(mockChain, nil).Maybe()
	//chain, chainID, err := ccipconfig.GetChainByChainSelector(mockChainSet, 3)

	//caller1 := rpclib.NewDynamicLimitedBatchCaller(
	//	logger.TestLogger(t),
	//	ec1,
	//	rpclib.DefaultRpcBatchSizeLimit,
	//	rpclib.DefaultRpcBatchBackOffMultiplier,
	//)

	caller1 := rpclibmocks.NewEvmBatchCaller(t)
	caller2 := rpclibmocks.NewEvmBatchCaller(t)

	// Real LINK/USD example from OP.
	round1 := aggregator_v2v3_interface.LatestRoundData{
		RoundId:         big.NewInt(1000),
		Answer:          big.NewInt(1396818990),
		StartedAt:       big.NewInt(1704896575),
		UpdatedAt:       big.NewInt(1704896575),
		AnsweredInRound: big.NewInt(1000),
	}

	// Real ETH/USD example from OP.
	round2 := aggregator_v2v3_interface.LatestRoundData{
		RoundId:         big.NewInt(2000),
		Answer:          big.NewInt(238879815123),
		StartedAt:       big.NewInt(1704897197),
		UpdatedAt:       big.NewInt(1704897197),
		AnsweredInRound: big.NewInt(2000),
	}

	caller1.On("BatchCall", mock.Anything, uint64(0), mock.Anything).Return(
		[]rpclib.DataAndErr{
			{
				Outputs: []any{round1},
			},
		},
		nil,
	).Maybe()

	caller2.On("BatchCall", mock.Anything, uint64(0), mock.Anything).Return(
		[]rpclib.DataAndErr{
			{
				Outputs: []any{round2},
			},
		},
		nil,
	).Maybe()

	evmClients := map[uint64]rpclib.EvmBatchCaller{
		uint64(101): caller1,
		uint64(102): caller2,
	}

	pg := NewDynamicPriceGetter(cfg, evmClients)
	prices, err := pg.TokenPricesUSD(ctx, []common.Address{tk1, tk2, tk3})
	assert.NoError(t, err)
	assert.Len(t, prices, 3)
	assert.Equal(t, cfg.StaticPrices[tk3].Price, prices[tk3].Uint64())
	assert.Equal(t, big.NewInt(0).Mul(big.NewInt(1396818990), big.NewInt(10_000_000_000)), prices[tk1])
	assert.Equal(t, big.NewInt(0).Mul(big.NewInt(238879815123), big.NewInt(10_000_000_000)), prices[tk2])
}
