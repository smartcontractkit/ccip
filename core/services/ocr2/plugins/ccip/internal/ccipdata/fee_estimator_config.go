package ccipdata

import (
	"context"
)

//go:generate mockery --quiet --name FeeEstimatorConfigReader --filename fee_estimator_config_reader_mock.go --case=underscore
type FeeEstimatorConfigReader interface {
	GetDataAvailabilityConfig(ctx context.Context) (destDAOverheadGas, destGasPerDAByte, destDAMultiplierBps int64, err error)
}
