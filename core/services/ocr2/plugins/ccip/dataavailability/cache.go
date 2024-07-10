package dataavailability

import (
	"context"
	"errors"
	"github.com/smartcontractkit/chainlink-common/pkg/types/ccip"
	"sync"
)

type DAConfigCache struct {
	onRampReader ccip.OnRampReader
	once         sync.Once
}

func NewDAConfigCache() *DAConfigCache {
	return &DAConfigCache{}
}

func (c *DAConfigCache) SetOnRampReader(reader ccip.OnRampReader) {
	c.onRampReader = reader
}

func (c *DAConfigCache) Get(ctx context.Context) (destDataAvailabilityOverheadGas, destGasPerDataAvailabilityByte, destDataAvailabilityMultiplierBps int64, err error) {
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
