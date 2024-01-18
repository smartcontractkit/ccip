package pricegetter

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/v2/core/internal/gethwrappers2/generated/offchainaggregator"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/rpclib"
)

// AggregatorPriceConfig specifies a price retrieved from an aggregator contract.
type AggregatorPriceConfig struct {
	ChainID         uint64         `json:"chainID"`
	ContractAddress common.Address `json:"contractAddress"`
}

// StaticPriceConfig specifies a price defined statically.
type StaticPriceConfig struct {
	ChainID uint64 `json:"chainID"`
	Price   uint64 `json:"price,string"`
}

type DynamicPriceGetterConfig struct {
	AggregatorPrices map[common.Address]AggregatorPriceConfig `json:"aggregatorPrices"`
	StaticPrices     map[common.Address]StaticPriceConfig     `json:"staticPrices"`
}

type DynamicPriceGetter struct {
	cfg           DynamicPriceGetterConfig
	evmClients    map[uint64]rpclib.EvmBatchCaller
	aggregatorAbi abi.ABI
}

// NewDynamicPriceGetter build a DynamicPriceGetter from a configuration and a map of chain ID to batch callers.
// A batch caller should be provided for all retrieved prices.
func NewDynamicPriceGetter(cfg DynamicPriceGetterConfig, evmClients map[uint64]rpclib.EvmBatchCaller) (*DynamicPriceGetter, error) {
	aggregatorAbi, err := abi.JSON(strings.NewReader(offchainaggregator.OffchainAggregatorABI))
	if err != nil {
		return nil, err
	}
	return &DynamicPriceGetter{cfg, evmClients, aggregatorAbi}, nil
}

// TokenPricesUSD implements the PriceGetter interface.
func (d *DynamicPriceGetter) TokenPricesUSD(ctx context.Context, tokens []common.Address) (map[common.Address]*big.Int, error) {
	prices := make(map[common.Address]*big.Int, len(tokens))

	batchCallsPerChain := make(map[uint64][]rpclib.EvmCall)
	batchCallsTokensOrder := make(map[uint64][]common.Address)

	fmt.Printf("=> querying token prices for %d tokens: %s\n", len(tokens), tokens)

	for _, tk := range tokens {
		// group aggregator-based tokens to make batch call (one per chain)
		if aggCfg, exists := d.cfg.AggregatorPrices[tk]; exists {
			batchCallsPerChain[aggCfg.ChainID] = append(batchCallsPerChain[aggCfg.ChainID], rpclib.NewEvmCall(
				d.aggregatorAbi,
				"latestRoundData",
				aggCfg.ContractAddress,
			))
			batchCallsTokensOrder[aggCfg.ChainID] = append(batchCallsTokensOrder[aggCfg.ChainID], tk)
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
		resultsPerChain, err := evmCaller.BatchCall(ctx, 50, batchCalls)
		if err != nil {
			return nil, fmt.Errorf("batch call: %w", err)
		}

		// latestRoundData returns an array of integers (not a proper struct), therefore we get the answer at position 1.
		latestRounds, err := rpclib.ParseOutputs[*big.Int](resultsPerChain, func(d rpclib.DataAndErr) (*big.Int, error) {
			return rpclib.ParseOutput[*big.Int](d, 1)
		})
		if err != nil {
			return nil, fmt.Errorf("parse outputs: %w", err)
		}

		for i := range tokensOrder {
			// Convert prices to wei (10e18) -> already in wei when coming from aggregator.
			//prices[tokensOrder[i]] = big.NewInt(0).Mul(latestRounds[i].Answer, big.NewInt(10_000_000_000))
			prices[tokensOrder[i]] = latestRounds[i]
		}
	}

	return prices, nil
}
