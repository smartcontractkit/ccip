package pricegetter

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccip"

	"github.com/smartcontractkit/chainlink/v2/core/internal/gethwrappers2/generated/offchainaggregator"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipcalc"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/rpclib"
)

const decimalsMethodName = "decimals"
const latestRoundDataMethodName = "latestRoundData"

func init() {
	// Ensure existence of latestRoundData method on the Aggregator contract.
	aggregatorABI, err := abi.JSON(strings.NewReader(offchainaggregator.OffchainAggregatorABI))
	if err != nil {
		panic(err)
	}
	ensureMethodOnContract(aggregatorABI, decimalsMethodName)
	ensureMethodOnContract(aggregatorABI, latestRoundDataMethodName)
}

func ensureMethodOnContract(abi abi.ABI, methodName string) {
	if _, ok := abi.Methods[methodName]; !ok {
		panic(fmt.Errorf("method %s not found on ABI: %+v", methodName, abi.Methods))
	}
}

type DynamicPriceGetterClient struct {
	BatchCaller rpclib.EvmBatchCaller
}

func NewDynamicPriceGetterClient(batchCaller rpclib.EvmBatchCaller) DynamicPriceGetterClient {
	return DynamicPriceGetterClient{
		BatchCaller: batchCaller,
	}
}

type DynamicPriceGetter struct {
	cfg           config.DynamicPriceGetterConfig
	evmClients    map[uint64]DynamicPriceGetterClient
	aggregatorAbi abi.ABI
}

func NewDynamicPriceGetterConfig(configJson string) (config.DynamicPriceGetterConfig, error) {
	priceGetterConfig := config.DynamicPriceGetterConfig{}
	err := json.Unmarshal([]byte(configJson), &priceGetterConfig)
	if err != nil {
		return config.DynamicPriceGetterConfig{}, fmt.Errorf("parsing dynamic price getter config: %w", err)
	}
	err = priceGetterConfig.Validate()
	if err != nil {
		return config.DynamicPriceGetterConfig{}, fmt.Errorf("validating price getter config: %w", err)
	}
	return priceGetterConfig, nil
}

// NewDynamicPriceGetter build a DynamicPriceGetter from a configuration and a map of chain ID to batch callers.
// A batch caller should be provided for all retrieved prices.
func NewDynamicPriceGetter(cfg config.DynamicPriceGetterConfig, evmClients map[uint64]DynamicPriceGetterClient) (*DynamicPriceGetter, error) {
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("validating dynamic price getter config: %w", err)
	}
	aggregatorAbi, err := abi.JSON(strings.NewReader(offchainaggregator.OffchainAggregatorABI))
	if err != nil {
		return nil, fmt.Errorf("parsing offchainaggregator abi: %w", err)
	}
	priceGetter := DynamicPriceGetter{cfg, evmClients, aggregatorAbi}
	return &priceGetter, nil
}

// FilterForConfiguredTokens implements the PriceGetter interface.
// It filters a list of token addresses for only those that have a price resolution rule configured on the PriceGetterConfig
func (d *DynamicPriceGetter) FilterConfiguredTokens(ctx context.Context, tokens []cciptypes.Address) (configured []cciptypes.Address, unconfigured []cciptypes.Address, err error) {
	configured = []cciptypes.Address{}
	unconfigured = []cciptypes.Address{}
	for _, tk := range tokens {
		evmAddr, err := ccipcalc.GenericAddrToEvm(tk)
		if err != nil {
			return nil, nil, err
		}

		if _, isAgg := d.cfg.AggregatorPrices[evmAddr]; isAgg {
			configured = append(configured, tk)
		} else if _, isStatic := d.cfg.StaticPrices[evmAddr]; isStatic {
			configured = append(configured, tk)
		} else {
			unconfigured = append(unconfigured, tk)
		}
	}
	return configured, unconfigured, nil
}

// BatchCallsForChain Defines the batch calls to perform on a given chain.
type batchCallsForChain struct {
	decimalCalls         []rpclib.EvmCall
	latestRoundDataCalls []rpclib.EvmCall
	tokenOrder           []common.Address // required to maintain the order of the batched rpc calls for mapping the results.
}

// TokenPricesUSD implements the PriceGetter interface.
// It returns static prices stored in the price getter, and batch calls to aggregators (on per chain) for aggregator-based prices.
func (d *DynamicPriceGetter) TokenPricesUSD(ctx context.Context, tokens []cciptypes.Address) (map[cciptypes.Address]*big.Int, error) {
	prices := make(map[cciptypes.Address]*big.Int, len(tokens))

	batchCallsPerChain := make(map[uint64]*batchCallsForChain)

	evmAddrs, err := ccipcalc.GenericAddrsToEvm(tokens...)
	if err != nil {
		return nil, err
	}
	for _, tk := range evmAddrs {
		if aggCfg, isAgg := d.cfg.AggregatorPrices[tk]; isAgg {
			// Batch calls for aggregator-based token prices (one per chain).
			if _, exists := batchCallsPerChain[aggCfg.ChainID]; !exists {
				batchCallsPerChain[aggCfg.ChainID] = &batchCallsForChain{
					decimalCalls:         []rpclib.EvmCall{},
					latestRoundDataCalls: []rpclib.EvmCall{},
					tokenOrder:           []common.Address{},
				}
			}
			chainCalls := batchCallsPerChain[aggCfg.ChainID]
			chainCalls.decimalCalls = append(chainCalls.decimalCalls, rpclib.NewEvmCall(
				d.aggregatorAbi,
				decimalsMethodName,
				aggCfg.AggregatorContractAddress,
			))
			chainCalls.latestRoundDataCalls = append(chainCalls.latestRoundDataCalls, rpclib.NewEvmCall(
				d.aggregatorAbi,
				latestRoundDataMethodName,
				aggCfg.AggregatorContractAddress,
			))
			chainCalls.tokenOrder = append(chainCalls.tokenOrder, tk)

		} else if staticCfg, isStatic := d.cfg.StaticPrices[tk]; isStatic {
			// Fill static prices.
			prices[ccipcalc.EvmAddrToGeneric(tk)] = staticCfg.Price
		} else {
			return nil, fmt.Errorf("no price resolution rule for token %s", tk.Hex())
		}
	}

	for chainID, batchCalls := range batchCallsPerChain {
		client, exists := d.evmClients[chainID]
		if !exists {
			return nil, fmt.Errorf("evm caller for chain %d not found", chainID)
		}

		evmCaller := client.BatchCaller

		resultsDecimals, err := evmCaller.BatchCall(ctx, 0, batchCalls.decimalCalls)
		if err != nil {
			return nil, fmt.Errorf("batch call: %w", err)
		}
		resultsLatestRoundData, err := evmCaller.BatchCall(ctx, 0, batchCalls.latestRoundDataCalls)
		if err != nil {
			return nil, fmt.Errorf("batch call: %w", err)
		}

		decimals, err := rpclib.ParseOutputs[uint8](resultsDecimals, func(d rpclib.DataAndErr) (uint8, error) {
			return rpclib.ParseOutput[uint8](d, 0)
		})
		if err != nil {
			return nil, fmt.Errorf("parse outputs: %w", err)
		}

		// latestRoundData function has multiple outputs (roundId,answer,startedAt,updatedAt,answeredInRound).
		// we want the second one (answer, idx=1).
		latestRounds, err := rpclib.ParseOutputs[*big.Int](resultsLatestRoundData, func(d rpclib.DataAndErr) (*big.Int, error) {
			return rpclib.ParseOutput[*big.Int](d, 1)
		})
		if err != nil {
			return nil, fmt.Errorf("parse outputs: %w", err)
		}

		for i := range batchCalls.tokenOrder {
			// Normalize to 1e18.
			if decimals[i] < 18 {
				latestRounds[i].Mul(latestRounds[i], big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18-int64(decimals[i])), nil))
			} else if decimals[i] > 18 {
				latestRounds[i].Div(latestRounds[i], big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(decimals[i])-18), nil))
			}
			prices[ccipcalc.EvmAddrToGeneric(batchCalls.tokenOrder[i])] = latestRounds[i]
		}
	}

	return prices, nil
}

func (d *DynamicPriceGetter) Close() error {
	return nil
}
