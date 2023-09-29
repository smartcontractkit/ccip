package prices

import (
	"context"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/gas"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal"
)

const (
	feeBoostingOverheadGas = 200_000
	// execGasPerToken is lower-bound estimation of ERC20 releaseOrMint gas cost (Mint with static minter).
	// Use this in per-token gas cost calc as heuristic to simplify estimation logic.
	execGasPerToken = 10_000
	// execGasPerPayloadByte is gas charged for passing each byte of `data` payload to CCIP receiver, ignores 4 gas per 0-byte rule.
	// This can be a constant as it is part of EVM spec. Changes should be rare.
	execGasPerPayloadByte = 16
	// evmMessageFixedBytes is byte size of fixed-size fields in EVM2EVMMessage
	// Updating EVM2EVMMessage involves an offchain upgrade, safe to keep this as constant in code.
	evmMessageFixedBytes     = 448
	evmMessageBytesPerToken  = 128          // Byte size of each token transfer, consisting of 1 EVMTokenAmount and 1 bytes, excl length of bytes
	daMultiplierBase         = int64(10000) // DA multiplier is in multiples of 0.0001, i.e. 1/daMultiplierBase
	daGasPriceEncodingLength = 112          // Each gas price takes up at most GasPriceEncodingLength number of bits
)

type GasPriceDeviationOptions struct {
	DADeviationPPB   int64
	ExecDeviationPPB int64
}

type MsgCostOptions struct {
	DAOverheadGas int64
	GasPerDAByte  int64
	DAMultiplier  int64
}

// GasPrice represents gas price as a single big.Int, same as gas price representation onchain.
// (multi-component gas prices are encoded into the int)
type GasPrice *big.Int

// GasPriceEstimator is abstraction over multi-component gas prices.
//
//go:generate mockery --quiet --name GasPriceEstimator --output . --filename gas_price_estimator_mock.go --inpackage --case=underscore
type GasPriceEstimator interface {
	// GetGasPrice fetches the current gas price.
	GetGasPrice(ctx context.Context) (GasPrice, error)
	// DenoteInUSD converts the gas price to be in units of USD. Input prices should not be nil.
	DenoteInUSD(p GasPrice, wrappedNativePrice *big.Int) (GasPrice, error)
	// Median finds the median gas price in slice. If gas price has multiple components, median of each individual component should be taken. Input prices should not contain nil.
	Median(gasPrices []GasPrice) (GasPrice, error)
	// Deviates checks if p1 gas price diffs from p2 by deviation options. Input prices should not be nil.
	Deviates(p1 GasPrice, p2 GasPrice, opts GasPriceDeviationOptions) (bool, error)
	// EstimateMsgCostUSD estimates the costs for msg execution, and converts to USD value scaled by 1e18 (e.g. 5$ = 5e18).
	EstimateMsgCostUSD(p GasPrice, wrappedNativePrice *big.Int, msg internal.EVM2EVMOnRampCCIPSendRequestedWithMeta, opts MsgCostOptions) (*big.Int, error)
	// String converts the gas price to string.
	String(p GasPrice) string
}

func NewGasPriceEstimator(
	commitStoreVersion semver.Version,
	estimator gas.EvmFeeEstimator,
	maxExecGasPrice *big.Int,
) (GasPriceEstimator, error) {
	switch commitStoreVersion.String() {
	case "1.0.0", "1.1.0":
		return ExecGasPriceEstimator{
			estimator:   estimator,
			maxGasPrice: maxExecGasPrice,
		}, nil
	case "1.2.0":
		return DAGasPriceEstimator{
			execEstimator: ExecGasPriceEstimator{
				estimator:   estimator,
				maxGasPrice: maxExecGasPrice,
			},
			l1Oracle:            estimator.L1Oracle(),
			priceEncodingLength: daGasPriceEncodingLength,
		}, nil
	default:
		return nil, errors.Errorf("Invalid commitStore version: %s", commitStoreVersion)
	}
}
