package ccipdata

import (
	"context"
)

type DAConfigCacheReader interface {
	Get(ctx context.Context) (destDAOverheadGas, destGasPerDAByte, destDAMultiplierBps int64, err error)
}
