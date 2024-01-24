package pricegetter

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	lpmocks "github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/aggregator_v3_interface"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/rpclib"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/rpclib/rpclibmocks"
)

type testParameters struct {
	cfg                        DynamicPriceGetterConfig
	evmClients                 map[uint64]DynamicPriceGetterClient
	invalidConfigErrorExpected bool
	expectedTokenPrices        map[common.Address]big.Int
}

func TestDynamicPriceGetter(t *testing.T) {
	tests := []struct {
		name  string
		param testParameters
	}{
		{
			name:  "aggregator_only_valid",
			param: testParamAggregatorOnly(t),
		},
		{
			name:  "static_only_valid",
			param: testParamStaticOnly(),
		},
		{
			name:  "aggregator_and_static_valid",
			param: testParamAggregatorAndStaticValid(t),
		},
		{
			name:  "aggregator_and_static_token_collision",
			param: testParamAggregatorAndStaticTokenCollision(t),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			pg, err := NewDynamicPriceGetter(test.param.cfg, test.param.evmClients)
			if test.param.invalidConfigErrorExpected {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				ctx := testutils.Context(t)
				// Build list of tokens to query.
				tokens := make([]common.Address, 0, len(test.param.expectedTokenPrices))
				for tk := range test.param.expectedTokenPrices {
					tokens = append(tokens, tk)
				}
				prices, err := pg.TokenPricesUSD(ctx, tokens)
				require.NoError(t, err)
				// we expect prices for at least all queried tokens (it is possible that additional tokens are returned).
				assert.True(t, len(prices) >= len(test.param.expectedTokenPrices))
				// Check prices are matching expected result.
				for tk, expectedPrice := range test.param.expectedTokenPrices {
					assert.Equal(t, expectedPrice, *prices[tk])
				}
			}
		})
	}
}

func testParamAggregatorOnly(t *testing.T) testParameters {
	tk1 := utils.RandomAddress()
	tk2 := utils.RandomAddress()
	cfg := DynamicPriceGetterConfig{
		AggregatorPrices: map[common.Address]AggregatorPriceConfig{
			tk1: {
				ChainID:                   101,
				AggregatorContractAddress: utils.RandomAddress(), // aggregator contract
			},
			tk2: {
				ChainID:                   102,
				AggregatorContractAddress: utils.RandomAddress(), // aggregator contract
			},
		},
		StaticPrices: map[common.Address]StaticPriceConfig{},
	}
	// Real LINK/USD example from OP.
	round1 := aggregator_v3_interface.LatestRoundData{
		RoundId:         big.NewInt(1000),
		Answer:          big.NewInt(1396818990),
		StartedAt:       big.NewInt(1704896575),
		UpdatedAt:       big.NewInt(1704896575),
		AnsweredInRound: big.NewInt(1000),
	}
	// Real ETH/USD example from OP.
	round2 := aggregator_v3_interface.LatestRoundData{
		RoundId:         big.NewInt(2000),
		Answer:          big.NewInt(238879815123),
		StartedAt:       big.NewInt(1704897197),
		UpdatedAt:       big.NewInt(1704897197),
		AnsweredInRound: big.NewInt(2000),
	}
	evmClients := map[uint64]DynamicPriceGetterClient{
		uint64(101): mockClientFromRound(t, round1),
		uint64(102): mockClientFromRound(t, round2),
	}
	expectedTokenPrices := map[common.Address]big.Int{
		tk1: *round1.Answer,
		tk2: *round2.Answer,
	}
	return testParameters{
		cfg:                        cfg,
		evmClients:                 evmClients,
		invalidConfigErrorExpected: false,
		expectedTokenPrices:        expectedTokenPrices,
	}
}

func testParamStaticOnly() testParameters {
	tk1 := utils.RandomAddress()
	tk2 := utils.RandomAddress()
	tk3 := utils.RandomAddress()
	cfg := DynamicPriceGetterConfig{
		AggregatorPrices: map[common.Address]AggregatorPriceConfig{},
		StaticPrices: map[common.Address]StaticPriceConfig{
			tk1: {
				ChainID: 101,
				Price:   1_234_000,
			},
			tk2: {
				ChainID: 102,
				Price:   2_234_000,
			},
			tk3: {
				ChainID: 103,
				Price:   3_234_000,
			},
		},
	}
	// Real LINK/USD example from OP.
	evmClients := map[uint64]DynamicPriceGetterClient{}
	expectedTokenPrices := map[common.Address]big.Int{
		tk1: *big.NewInt(int64(cfg.StaticPrices[tk1].Price)),
		tk2: *big.NewInt(int64(cfg.StaticPrices[tk2].Price)),
		tk3: *big.NewInt(int64(cfg.StaticPrices[tk3].Price)),
	}
	return testParameters{
		cfg:                        cfg,
		evmClients:                 evmClients,
		invalidConfigErrorExpected: false,
		expectedTokenPrices:        expectedTokenPrices,
	}
}

func testParamAggregatorAndStaticValid(t *testing.T) testParameters {
	tk1 := utils.RandomAddress()
	tk2 := utils.RandomAddress()
	tk3 := utils.RandomAddress()
	cfg := DynamicPriceGetterConfig{
		AggregatorPrices: map[common.Address]AggregatorPriceConfig{
			tk1: {
				ChainID:                   101,
				AggregatorContractAddress: utils.RandomAddress(), // aggregator contract
			},
			tk2: {
				ChainID:                   102,
				AggregatorContractAddress: utils.RandomAddress(), // aggregator contract
			},
		},
		StaticPrices: map[common.Address]StaticPriceConfig{
			tk3: {
				ChainID: 103,
				Price:   1_234_000,
			},
		},
	}
	// Real LINK/USD example from OP.
	round1 := aggregator_v3_interface.LatestRoundData{
		RoundId:         big.NewInt(1000),
		Answer:          big.NewInt(1396818990),
		StartedAt:       big.NewInt(1704896575),
		UpdatedAt:       big.NewInt(1704896575),
		AnsweredInRound: big.NewInt(1000),
	}
	// Real ETH/USD example from OP.
	round2 := aggregator_v3_interface.LatestRoundData{
		RoundId:         big.NewInt(2000),
		Answer:          big.NewInt(238879815123),
		StartedAt:       big.NewInt(1704897197),
		UpdatedAt:       big.NewInt(1704897197),
		AnsweredInRound: big.NewInt(2000),
	}
	evmClients := map[uint64]DynamicPriceGetterClient{
		uint64(101): mockClientFromRound(t, round1),
		uint64(102): mockClientFromRound(t, round2),
	}
	expectedTokenPrices := map[common.Address]big.Int{
		tk1: *round1.Answer,
		tk2: *round2.Answer,
		tk3: *big.NewInt(int64(cfg.StaticPrices[tk3].Price)),
	}
	return testParameters{
		cfg:                        cfg,
		evmClients:                 evmClients,
		invalidConfigErrorExpected: false,
		expectedTokenPrices:        expectedTokenPrices,
	}
}

func testParamAggregatorAndStaticTokenCollision(t *testing.T) testParameters {
	tk1 := utils.RandomAddress()
	tk2 := utils.RandomAddress()
	tk3 := utils.RandomAddress()
	cfg := DynamicPriceGetterConfig{
		AggregatorPrices: map[common.Address]AggregatorPriceConfig{
			tk1: {
				ChainID:                   101,
				AggregatorContractAddress: utils.RandomAddress(), // aggregator contract
			},
			tk2: {
				ChainID:                   102,
				AggregatorContractAddress: utils.RandomAddress(), // aggregator contract
			},
			tk3: {
				ChainID:                   103,
				AggregatorContractAddress: utils.RandomAddress(), // aggregator contract
			},
		},
		StaticPrices: map[common.Address]StaticPriceConfig{
			tk3: {
				ChainID: 103,
				Price:   1_234_000,
			},
		},
	}
	// Real LINK/USD example from OP.
	round1 := aggregator_v3_interface.LatestRoundData{
		RoundId:         big.NewInt(1000),
		Answer:          big.NewInt(1396818990),
		StartedAt:       big.NewInt(1704896575),
		UpdatedAt:       big.NewInt(1704896575),
		AnsweredInRound: big.NewInt(1000),
	}
	// Real ETH/USD example from OP.
	round2 := aggregator_v3_interface.LatestRoundData{
		RoundId:         big.NewInt(2000),
		Answer:          big.NewInt(238879815123),
		StartedAt:       big.NewInt(1704897197),
		UpdatedAt:       big.NewInt(1704897197),
		AnsweredInRound: big.NewInt(2000),
	}
	round3 := aggregator_v3_interface.LatestRoundData{
		RoundId:         big.NewInt(3000),
		Answer:          big.NewInt(238879815124),
		StartedAt:       big.NewInt(1704897198),
		UpdatedAt:       big.NewInt(1704897198),
		AnsweredInRound: big.NewInt(3000),
	}
	evmClients := map[uint64]DynamicPriceGetterClient{
		uint64(101): mockClientFromRound(t, round1),
		uint64(102): mockClientFromRound(t, round2),
		uint64(103): mockClientFromRound(t, round3),
	}
	return testParameters{
		cfg:                        cfg,
		evmClients:                 evmClients,
		invalidConfigErrorExpected: true,
		expectedTokenPrices:        nil,
	}
}

func mockClientFromRound(t *testing.T, round aggregator_v3_interface.LatestRoundData) DynamicPriceGetterClient {
	return DynamicPriceGetterClient{
		BatchCaller: mockCallerFromRound(t, round),
		LP:          mockLPFromRound(t, round),
	}
}

func mockCallerFromRound(t *testing.T, round aggregator_v3_interface.LatestRoundData) *rpclibmocks.EvmBatchCaller {
	caller := rpclibmocks.NewEvmBatchCaller(t)
	caller.On("BatchCall", mock.Anything, round.RoundId.Uint64(), mock.Anything).Return(
		[]rpclib.DataAndErr{
			{
				Outputs: []any{round.RoundId, round.Answer, round.StartedAt, round.UpdatedAt, round.AnsweredInRound},
			},
		},
		nil,
	).Maybe()
	return caller
}

func mockLPFromRound(t *testing.T, round aggregator_v3_interface.LatestRoundData) *lpmocks.LogPoller {
	lp := lpmocks.NewLogPoller(t)
	lp.On("LatestBlock", mock.Anything).Return(
		logpoller.LogPollerBlock{
			BlockNumber: int64(round.RoundId.Uint64()),
		}, nil).Maybe()
	return lp
}
