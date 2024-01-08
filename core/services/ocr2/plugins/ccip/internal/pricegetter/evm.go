package pricegetter

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/rpclib"
)

type AggregatorPriceConfig struct {
	ChainID         int64          `json:"chainID"`
	ContractAddress common.Address `json:"contractAddress"`
}

type StaticPriceConfig struct {
	ChainID int64  `json:"chainID"`
	Price   uint64 `json:"price,string"`
}

type DynamicPriceGetterConfig struct {
	AggregatorPrices map[common.Address]AggregatorPriceConfig `json:"aggregatorPrices"`
	StaticPrices     map[common.Address]StaticPriceConfig     `json:"staticPrices"`
}

type DynamicPriceGetter struct {
	cfg        DynamicPriceGetterConfig
	evmClients map[int64]rpclib.EvmBatchCaller
}

func NewDynamicPriceGetter(cfg DynamicPriceGetterConfig, evmClients map[int64]rpclib.EvmBatchCaller) *DynamicPriceGetter {
	return &DynamicPriceGetter{
		cfg:        cfg,
		evmClients: evmClients,
	}
}

func (d *DynamicPriceGetter) TokenPricesUSD(ctx context.Context, tokens []common.Address) (map[common.Address]*big.Int, error) {
	prices := make(map[common.Address]*big.Int, len(tokens))

	batchCallsPerChain := make(map[int64][]rpclib.EvmCall)
	batchCallsTokensOrder := make(map[int64][]common.Address)
	for _, tk := range tokens {
		// group aggregator-based tokens to make batch call (one per chain)
		if dynCfg, exists := d.cfg.AggregatorPrices[tk]; exists {
			batchCallsPerChain[dynCfg.ChainID] = append(batchCallsPerChain[dynCfg.ChainID], rpclib.NewEvmCall(
				abi.ABI{}, // todo
				"latestRoundData",
				dynCfg.ContractAddress,
			))
			batchCallsTokensOrder[dynCfg.ChainID] = append(batchCallsTokensOrder[dynCfg.ChainID], tk)
			continue
		}

		// fill static prices
		if staticCfg, exists := d.cfg.StaticPrices[tk]; exists {
			prices[tk] = big.NewInt(0).SetUint64(staticCfg.Price)
			continue
		}
	}

	for chainID, batchCalls := range batchCallsPerChain {
		evmCaller, exists := d.evmClients[chainID]
		if !exists {
			return nil, fmt.Errorf("evm caller for chain %d not found", chainID)
		}

		tokensOrder := batchCallsTokensOrder[chainID]
		results, err := evmCaller.BatchCall(ctx, 0, batchCalls)
		if err != nil {
			return nil, fmt.Errorf("batch call: %w", err)
		}

		aggrPrices, err := rpclib.ParseOutputs[*big.Int](
			results,
			func(d rpclib.DataAndErr) (*big.Int, error) {
				return nil, fmt.Errorf("todo")
			},
		)
		if err != nil {
			return nil, fmt.Errorf("parse outputs: %w", err)
		}

		for i := range tokensOrder {
			prices[tokensOrder[i]] = aggrPrices[i]
		}
	}

	return prices, nil
}
