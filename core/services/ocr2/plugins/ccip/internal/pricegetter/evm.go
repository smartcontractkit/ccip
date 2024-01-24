package pricegetter

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/internal/gethwrappers2/generated/offchainaggregator"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/rpclib"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
)

// AggregatorPriceConfig specifies a price retrieved from an aggregator contract.
type AggregatorPriceConfig struct {
	ChainID                   uint64         `json:"chainID"`
	AggregatorContractAddress common.Address `json:"contractAddress"`
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

type DynamicPriceGetterClient struct {
	BatchCaller rpclib.EvmBatchCaller
	LP          logpoller.LogPoller
}

type DynamicPriceGetter struct {
	cfg           DynamicPriceGetterConfig
	evmClients    map[uint64]DynamicPriceGetterClient
	aggregatorAbi abi.ABI
}

// NewDynamicPriceGetter build a DynamicPriceGetter from a configuration and a map of chain ID to batch callers.
// A batch caller should be provided for all retrieved prices.
func NewDynamicPriceGetter(cfg DynamicPriceGetterConfig, evmClients map[uint64]DynamicPriceGetterClient) (*DynamicPriceGetter, error) {
	aggregatorAbi, err := abi.JSON(strings.NewReader(offchainaggregator.OffchainAggregatorABI))
	if err != nil {
		return nil, err
	}
	priceGetter := DynamicPriceGetter{cfg, evmClients, aggregatorAbi}
	if err = priceGetter.Validate(); err != nil {
		return nil, err
	}
	return &priceGetter, nil
}

func (d *DynamicPriceGetter) Validate() error {
	// Ensure no duplication in token price resolution rules.
	if d.cfg.AggregatorPrices != nil && d.cfg.StaticPrices != nil {
		for tk := range d.cfg.AggregatorPrices {
			if _, exists := d.cfg.StaticPrices[tk]; exists {
				return fmt.Errorf("token %s defined in both aggregator and static price rules", tk.Hex())
			}
		}
	}
	return nil
}

// TokenPricesUSD implements the PriceGetter interface.
func (d *DynamicPriceGetter) TokenPricesUSD(ctx context.Context, tokens []common.Address) (map[common.Address]*big.Int, error) {
	prices := make(map[common.Address]*big.Int, len(tokens))

	batchCallsPerChain := make(map[uint64][]rpclib.EvmCall)
	batchCallsTokensOrder := make(map[uint64][]common.Address)

	for _, tk := range tokens {
		// group aggregator-based tokens to make batch call (one per chain)
		if aggCfg, exists := d.cfg.AggregatorPrices[tk]; exists {
			batchCallsPerChain[aggCfg.ChainID] = append(batchCallsPerChain[aggCfg.ChainID], rpclib.NewEvmCall(
				d.aggregatorAbi,
				"latestRoundData",
				aggCfg.AggregatorContractAddress,
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
		client, exists := d.evmClients[chainID]
		evmCaller := client.BatchCaller
		lp := client.LP
		if !exists {
			return nil, fmt.Errorf("evm caller for chain %d not found", chainID)
		}

		tokensOrder := batchCallsTokensOrder[chainID]
		latestBlock, err := lp.LatestBlock(pg.WithParentCtx(ctx))
		if err != nil {
			return nil, fmt.Errorf("get latest block: %w", err)
		}

		resultsPerChain, err := evmCaller.BatchCall(ctx, uint64(latestBlock.BlockNumber), batchCalls)
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
			// Prices are already in wei (10e18) when coming from aggregator, no conversion needed.
			prices[tokensOrder[i]] = latestRounds[i]
		}
	}

	return prices, nil
}
