package ccipdata

import (
	"context"
	"math/big"
)

type FeeEstimatorConfigReader interface {
	GetDataAvailabilityConfig(ctx context.Context) (destDAOverheadGas, destGasPerDAByte, destDAMultiplierBps int64, err error)
	ModifyDAGasPrice(gasPrice *big.Int) *big.Int
}
