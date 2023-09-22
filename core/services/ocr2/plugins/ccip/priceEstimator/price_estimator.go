package priceEstimator

import (
	"context"
	"math/big"

	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal"
)

const (
	FEE_BOOSTING_OVERHEAD_GAS   = 200_000
	EVM_MESSAGE_FIXED_BYTES     = 448 // Byte size of fixed-size fields in EVM2EVMMessage
	EVM_MESSAGE_BYTES_PER_TOKEN = 128 // Byte size of each token transfer, consisting of 1 EVMTokenAmount and 1 bytes, excl length of bytes
	DA_MULTIPLIER_BASE          = int64(10000)
)

type MsgCostConfig struct {
	daOverheadGas       int64
	gasPerDAByte        int64
	daMultiplier        int64
	feeBoostingOverhead int64
}

// GasPrice represents gas price as a single big.Int, same as gas price representation onchain.
// (multi-component gas prices are encoded into the int)
type GasPrice *big.Int

// Abstraction over multi-component gas prices
type GasPriceEstimator interface {
	GetGasPrice(context.Context) (GasPrice, error)
	DenoteInUSD(GasPrice, *big.Int) (GasPrice, error)
	Median([]GasPrice) (GasPrice, error)
	Deviates(GasPrice, GasPrice) (bool, error)
	EstimateMsgCostUSD(GasPrice, *big.Int, internal.EVM2EVMOnRampCCIPSendRequestedWithMeta, MsgCostConfig) (*big.Int, error)
	String(GasPrice) string
}
