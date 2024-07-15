package estimatorconfig

import (
	"context"
	"errors"

	"github.com/smartcontractkit/chainlink-common/pkg/types/ccip"
)

type FeeEstimatorConfigService struct {
	onRampReader ccip.OnRampReader
}

func NewFeeEstimatorConfigService() *FeeEstimatorConfigService {
	return &FeeEstimatorConfigService{}
}

// SetOnRampReader Sets the onRamp reader instance.
// must be called once for each instance.
func (c *FeeEstimatorConfigService) SetOnRampReader(reader ccip.OnRampReader) {
	c.onRampReader = reader
	return
}

// GetDataAvailabilityConfig Returns dynamic config data availability parameters.
// GetDynamicConfig should be cached in the onRamp reader to avoid unnecessary on-chain calls
func (c *FeeEstimatorConfigService) GetDataAvailabilityConfig(ctx context.Context) (destDataAvailabilityOverheadGas, destGasPerDataAvailabilityByte, destDataAvailabilityMultiplierBps int64, err error) {
	if c.onRampReader == nil {
		return 0, 0, 0, errors.New("no OnRampReader has been configured")
	}

	cfg, err := c.onRampReader.GetDynamicConfig(ctx)
	if err != nil {
		return 0, 0, 0, err
	}

	return int64(cfg.DestDataAvailabilityOverheadGas),
		int64(cfg.DestGasPerDataAvailabilityByte),
		int64(cfg.DestDataAvailabilityMultiplierBps),
		err
}
