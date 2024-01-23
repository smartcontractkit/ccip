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

func TestDynamicPriceGetter(t *testing.T) {
	ctx := testutils.Context(t)

	tk1 := utils.RandomAddress()
	tk2 := utils.RandomAddress()
	tk3 := utils.RandomAddress()

	cfg := DynamicPriceGetterConfig{
		AggregatorPrices: map[common.Address]AggregatorPriceConfig{
			tk1: {
				ChainID:         101,
				ContractAddress: common.HexToAddress("0xABC0011"), // aggregator contract
			},
			tk2: {
				ChainID:         102,
				ContractAddress: common.HexToAddress("0xABC0022"), // aggregator contract
			},
		},
		StaticPrices: map[common.Address]StaticPriceConfig{
			tk3: {
				ChainID: 103,
				Price:   1_234_000,
			},
		},
	}

	//b, _ := json.MarshalIndent(cfg, "", " ")
	//t.Logf("%s", b)

	caller1 := rpclibmocks.NewEvmBatchCaller(t)
	caller2 := rpclibmocks.NewEvmBatchCaller(t)

	lp1 := lpmocks.NewLogPoller(t)
	lp2 := lpmocks.NewLogPoller(t)

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

	//round3 := aggregator_v3_interface.LatestRoundData{
	//	RoundId:         big.NewInt(3000),
	//	Answer:          big.NewInt(238879815124),
	//	StartedAt:       big.NewInt(1704897198),
	//	UpdatedAt:       big.NewInt(1704897198),
	//	AnsweredInRound: big.NewInt(3000),
	//}

	caller1.On("BatchCall", mock.Anything, uint64(1000), mock.Anything).Return(
		[]rpclib.DataAndErr{
			{
				Outputs: []any{round1.RoundId, round1.Answer, round1.StartedAt, round1.UpdatedAt, round1.AnsweredInRound},
			},
		},
		nil,
	).Maybe()

	caller2.On("BatchCall", mock.Anything, uint64(2000), mock.Anything).Return(
		[]rpclib.DataAndErr{
			{
				Outputs: []any{round2.RoundId, round2.Answer, round2.StartedAt, round2.UpdatedAt, round2.AnsweredInRound},
			},
		},
		nil,
	).Maybe()
	lp1.On("LatestBlock", mock.Anything).Return(
		logpoller.LogPollerBlock{
			BlockNumber: int64(1000),
		}, nil).Maybe()
	lp2.On("LatestBlock", mock.Anything).Return(logpoller.LogPollerBlock{
		BlockNumber: int64(2000),
	}, nil).Maybe()

	evmClients := map[uint64]DynamicPriceGetterClient{
		uint64(101): {
			BatchCaller: caller1,
			LP:          lp1,
		},
		uint64(102): {
			BatchCaller: caller2,
			LP:          lp2,
		},
	}

	pg, err := NewDynamicPriceGetter(cfg, evmClients)
	require.NoError(t, err)
	prices, err := pg.TokenPricesUSD(ctx, []common.Address{tk1, tk2, tk3})
	require.NoError(t, err)
	assert.Len(t, prices, 3)
	assert.Equal(t, big.NewInt(1396818990), prices[tk1])
	assert.Equal(t, big.NewInt(238879815123), prices[tk2])
	assert.Equal(t, cfg.StaticPrices[tk3].Price, prices[tk3].Uint64())
}
