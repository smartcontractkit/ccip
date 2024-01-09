package pricegetter

import (
	"encoding/json"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/rpclib"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

func TestDynamicPriceGetter(t *testing.T) {
	ctx := testutils.Context(t)

	tk1 := utils.RandomAddress()
	tk2 := utils.RandomAddress()
	tk3 := utils.RandomAddress()

	cfg := DynamicPriceGetterConfig{
		AggregatorPrices: map[common.Address]AggregatorPriceConfig{
			tk1: {
				ChainID:         1,
				ContractAddress: common.HexToAddress("0xA550011"), // aggregator contract
			},
			tk2: {
				ChainID:         2,
				ContractAddress: common.HexToAddress("0xA550022"), // aggregator contract
			},
		},
		StaticPrices: map[common.Address]StaticPriceConfig{
			tk3: {
				ChainID: 3,
				Price:   1_234_000,
			},
		},
	}

	b, _ := json.MarshalIndent(cfg, "", " ")
	t.Logf("%s", b)

	pg := NewDynamicPriceGetter(cfg, map[int64]rpclib.EvmBatchCaller{})
	prices, err := pg.TokenPricesUSD(ctx, []common.Address{tk1, tk2, tk3})
	assert.NoError(t, err)
	assert.Len(t, prices, 3)
	assert.Equal(t, cfg.StaticPrices[tk3].Price, prices[tk3].Uint64())
}
