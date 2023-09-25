package prices

import (
	"context"
	"math/big"

	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/gas"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal"
)

const (
	FeeBoostingOverheadGas   = 200_000
	ExecGasPerToken          = 25_000 // Reasonable estimation of ERC20 transfer cost (19_000-35_000 depending on receiver balances)
	ExecGasPerPayloadByte    = 16     // Gas charged for passing each byte of `data` payload to CCIP receiver. This can be a constant as it is part of EVM spec. Changes should be rare.
	EvmMessageFixedBytes     = 448    // Byte size of fixed-size fields in EVM2EVMMessage
	EvmMessageBytesPerToken  = 128    // Byte size of each token transfer, consisting of 1 EVMTokenAmount and 1 bytes, excl length of bytes
	DAMultiplierBase         = int64(10000)
	DAGasPriceEncodingLength = 112 // Each gas price takes up at most GasPriceEncodingLength number of bits
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

// GasPriceEstimator is abstraction over multi-component gas prices
//
//go:generate mockery --quiet --name GasPriceEstimator --output . --filename gas_price_estimator_mock.go --inpackage --case=underscore
type GasPriceEstimator interface {
	// GetGasPrice fetches the current gas price.
	GetGasPrice(ctx context.Context) (GasPrice, error)
	// DenoteInUSD converts the gas price to be in units of USD.
	DenoteInUSD(p GasPrice, wrappedNativePrice *big.Int) (GasPrice, error)
	// Median finds the median gas price in slice. If gas price has multiple components, median of each individual component should be taken.
	Median(gasPrices []GasPrice) (GasPrice, error)
	// Deviates checks if p1 diffs from p2 by deviation options.
	Deviates(p1 GasPrice, p2 GasPrice, opts GasPriceDeviationOptions) (bool, error)
	// EstimateMsgCostUSD estimates the costs for msg execution, and converts to USD value scaled by 1e18 (e.g. 5$ = 5e18).
	EstimateMsgCostUSD(p GasPrice, wrappedNativePrice *big.Int, msg internal.EVM2EVMOnRampCCIPSendRequestedWithMeta, opts MsgCostOptions) (*big.Int, error)
	// String converts the gas price to string.
	String(p GasPrice) string
}

func NewGasPriceEstimator(
	commitStoreVersion string,
	estimator gas.EvmFeeEstimator,
	maxExecGasPrice *big.Int,
) (GasPriceEstimator, error) {
	switch commitStoreVersion {
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
			priceEncodingLength: DAGasPriceEncodingLength,
		}, nil
	default:
		return nil, errors.Errorf("Invalid commitStore version: %s", commitStoreVersion)
	}
}
