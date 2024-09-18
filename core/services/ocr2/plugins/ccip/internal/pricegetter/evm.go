package pricegetter

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/aggregator_v3_interface"

	"go.uber.org/multierr"

	"github.com/smartcontractkit/chainlink-common/pkg/types"

	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/rpclib"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccip"

	"github.com/smartcontractkit/chainlink/v2/core/internal/gethwrappers2/generated/offchainaggregator"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipcalc"
)

const decimalsMethodName = "decimals"
const latestRoundDataMethodName = "latestRoundData"
const OFFCHAIN_AGGREGATOR = "OffchainAggregator"

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
	cfg             config.DynamicPriceGetterConfig
	contractReaders map[uint64]types.ContractReader
	aggregatorAbi   abi.ABI
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
func NewDynamicPriceGetter(cfg config.DynamicPriceGetterConfig, contractReaders map[uint64]types.ContractReader) (*DynamicPriceGetter, error) {
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("validating dynamic price getter config: %w", err)
	}
	aggregatorAbi, err := abi.JSON(strings.NewReader(offchainaggregator.OffchainAggregatorABI))
	if err != nil {
		return nil, fmt.Errorf("parsing offchainaggregator abi: %w", err)
	}
	priceGetter := DynamicPriceGetter{cfg, contractReaders, aggregatorAbi}
	return &priceGetter, nil
}

// FilterConfiguredTokens implements the PriceGetter interface.
// It filters a list of token addresses for only those that have a price resolution rule configured on the PriceGetterConfig
func (d *DynamicPriceGetter) FilterConfiguredTokens(_ context.Context, tokens []cciptypes.Address) (configured []cciptypes.Address, unconfigured []cciptypes.Address, err error) {
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

// GetJobSpecTokenPricesUSD returns the prices of all tokens defined in the price getter.
func (d *DynamicPriceGetter) GetJobSpecTokenPricesUSD(ctx context.Context) (map[cciptypes.Address]*big.Int, error) {
	return d.TokenPricesUSD(ctx, d.getAllTokensDefined())
}

// TokenPricesUSD implements the PriceGetter interface.
// It returns static prices stored in the price getter, and batch calls aggregators (one per chain) to retrieve aggregator-based prices.
func (d *DynamicPriceGetter) TokenPricesUSD(ctx context.Context, tokens []cciptypes.Address) (map[cciptypes.Address]*big.Int, error) {
	prices, batchCallsPerChain, err := d.preparePricesAndBatchCallsPerChain(tokens)
	if err != nil {
		return nil, err
	}
	if err = d.performBatchCalls(ctx, batchCallsPerChain, prices); err != nil {
		return nil, err
	}
	return prices, nil
}

func (d *DynamicPriceGetter) getAllTokensDefined() []cciptypes.Address {
	tokens := make([]cciptypes.Address, 0)

	for addr := range d.cfg.AggregatorPrices {
		tokens = append(tokens, ccipcalc.EvmAddrToGeneric(addr))
	}
	for addr := range d.cfg.StaticPrices {
		tokens = append(tokens, ccipcalc.EvmAddrToGeneric(addr))
	}
	return tokens
}

// performBatchCalls performs batch calls on all chains to retrieve token prices.
func (d *DynamicPriceGetter) performBatchCalls(ctx context.Context, batchCallsPerChain map[uint64]*batchCallsForChain, prices map[cciptypes.Address]*big.Int) error {
	for chainID, batchCalls := range batchCallsPerChain {
		if err := d.performBatchCall(ctx, chainID, batchCalls, prices); err != nil {
			return err
		}
	}
	return nil
}

// performBatchCall performs a batch call on a given chain to retrieve token prices.
func (d *DynamicPriceGetter) performBatchCall(ctx context.Context, chainID uint64, batchCalls *batchCallsForChain, prices map[cciptypes.Address]*big.Int) (err error) {
	nbDecimalCalls := len(batchCalls.decimalCalls)
	nbLatestRoundDataCalls := len(batchCalls.decimalCalls)

	// Retrieve contract reader for the chain
	contractReader := d.contractReaders[chainID]

	// Bind contract reader to the contract addresses necessary for the batch calls
	bindings := make([]types.BoundContract, 0)
	for i, call := range batchCalls.decimalCalls {
		bindings = append(bindings, types.BoundContract{
			Address: string(ccipcalc.EvmAddrToGeneric(call.ContractAddress())),
			Name:    fmt.Sprintf("%v_%v", OFFCHAIN_AGGREGATOR, i),
		})
	}

	err = contractReader.Bind(ctx, bindings)
	if err != nil {
		return fmt.Errorf("binding contracts failed: %w", err)
	}

	// Perform call
	var decimalsReq uint8
	var batchGetLatestValuesRequest types.BatchGetLatestValuesRequest
	batchGetLatestValuesRequest = make(map[string]types.ContractBatch)
	for i, call := range batchCalls.decimalCalls {
		contractName := fmt.Sprintf("%v_%v", OFFCHAIN_AGGREGATOR, i)
		batchGetLatestValuesRequest[contractName] = append(batchGetLatestValuesRequest[contractName], types.BatchRead{
			ReadName:  call.MethodName(),
			ReturnVal: &decimalsReq,
		})
	}

	for i, call := range batchCalls.latestRoundDataCalls {
		contractName := fmt.Sprintf("%v_%v", OFFCHAIN_AGGREGATOR, i)
		batchGetLatestValuesRequest[contractName] = append(batchGetLatestValuesRequest[contractName], types.BatchRead{
			ReadName:  call.MethodName(),
			ReturnVal: &aggregator_v3_interface.LatestRoundData{},
		})
	}

	result, err2 := contractReader.BatchGetLatestValues(ctx, batchGetLatestValuesRequest)
	if err2 != nil {
		return fmt.Errorf("BatchGetLatestValues failed %w", err2)
	}

	// Extract results
	// give result the method key and then you get slice of responses
	decimalsCR := make([]uint8, 0, nbDecimalCalls)
	latestRoundCR := make([]aggregator_v3_interface.LatestRoundData, 0, nbDecimalCalls)
	var respErr error
	for j := range nbDecimalCalls {
		contractName := fmt.Sprintf("%v_%v", OFFCHAIN_AGGREGATOR, j)
		offchainAggregatorRespSlice := result[contractName]

		for i, read := range offchainAggregatorRespSlice {
			val, readErr := read.GetResult()
			if readErr != nil {
				respErr = multierr.Append(respErr, fmt.Errorf("error with method call %v: %w", batchCalls.decimalCalls[i].MethodName(), readErr))
				continue
			}
			if read.ReadName == decimalsMethodName {
				decimal, ok := val.(*uint8)
				if !ok {
					return fmt.Errorf("expected type uint8 for method call %v on contract %v: %w", batchCalls.decimalCalls[i].MethodName(), batchCalls.decimalCalls[i].ContractAddress(), readErr)
				}

				decimalsCR = append(decimalsCR, *decimal)
			} else if read.ReadName == latestRoundDataMethodName {
				latestRoundDataRes, ok := val.(*aggregator_v3_interface.LatestRoundData)
				if !ok {
					return fmt.Errorf("expected type latestRoundDataConfig for method call %v on contract %v: %w", batchCalls.latestRoundDataCalls[i/2].MethodName(), batchCalls.latestRoundDataCalls[i/2].ContractAddress(), readErr)
				}

				latestRoundCR = append(latestRoundCR, *latestRoundDataRes)
			}
		}
	}
	if respErr != nil {
		return respErr
	}

	latestRoundAnswerCR := make([]*big.Int, 0, nbLatestRoundDataCalls)
	for i := range nbLatestRoundDataCalls {
		latestRoundAnswerCR = append(latestRoundAnswerCR, latestRoundCR[i].Answer)
	}

	// Normalize and store prices.
	for i := range batchCalls.tokenOrder {
		// Normalize to 1e18.
		if decimalsCR[i] < 18 {
			latestRoundAnswerCR[i].Mul(latestRoundAnswerCR[i], big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18-int64(decimalsCR[i])), nil))
		} else if decimalsCR[i] > 18 {
			latestRoundAnswerCR[i].Div(latestRoundAnswerCR[i], big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(decimalsCR[i])-18), nil))
		}
		prices[ccipcalc.EvmAddrToGeneric(batchCalls.tokenOrder[i])] = latestRoundAnswerCR[i]
	}
	return nil
}

// preparePricesAndBatchCallsPerChain uses this price getter to prepare for a list of tokens:
// - the map of token address to their prices (static prices)
// - the map of and batch calls per chain for the given tokens (dynamic prices)
func (d *DynamicPriceGetter) preparePricesAndBatchCallsPerChain(tokens []cciptypes.Address) (map[cciptypes.Address]*big.Int, map[uint64]*batchCallsForChain, error) {
	prices := make(map[cciptypes.Address]*big.Int, len(tokens))
	batchCallsPerChain := make(map[uint64]*batchCallsForChain)
	evmAddrs, err := ccipcalc.GenericAddrsToEvm(tokens...)
	if err != nil {
		return nil, nil, err
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
			return nil, nil, fmt.Errorf("no price resolution rule for token %s", tk.Hex())
		}
	}
	return prices, batchCallsPerChain, nil
}

// batchCallsForChain Defines the batch calls to perform on a given chain.
type batchCallsForChain struct {
	decimalCalls         []rpclib.EvmCall
	latestRoundDataCalls []rpclib.EvmCall
	tokenOrder           []common.Address // required to maintain the order of the batched rpc calls for mapping the results.
}

func (d *DynamicPriceGetter) Close() error {
	return nil
}
