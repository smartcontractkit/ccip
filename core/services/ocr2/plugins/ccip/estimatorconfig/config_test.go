package estimatorconfig_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink-common/pkg/types/ccip"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/estimatorconfig"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata/mocks"
)

func TestFeeEstimatorConfigService(t *testing.T) {
	svc := estimatorconfig.NewFeeEstimatorConfigService()
	ctx := context.Background()

	var expectedDestDataAvailabilityOverheadGas int64 = 1
	var expectedDestGasPerDataAvailabilityByte int64 = 2
	var expectedDestDataAvailabilityMultiplierBps int64 = 3

	onRampReader := mocks.NewOnRampReader(t)
	_, _, _, err := svc.GetDataAvailabilityConfig(ctx)
	assert.Error(t, err)
	assert.NoError(t, svc.SetOnRampReader(onRampReader))
	assert.Error(t, svc.SetOnRampReader(onRampReader))

	onRampReader.On("GetDynamicConfig", ctx).
		Return(ccip.OnRampDynamicConfig{
			DestDataAvailabilityOverheadGas:   uint32(expectedDestDataAvailabilityOverheadGas),
			DestGasPerDataAvailabilityByte:    uint16(expectedDestGasPerDataAvailabilityByte),
			DestDataAvailabilityMultiplierBps: uint16(expectedDestDataAvailabilityMultiplierBps),
		}, nil).Once()

	destDataAvailabilityOverheadGas, destGasPerDataAvailabilityByte, destDataAvailabilityMultiplierBps, err := svc.GetDataAvailabilityConfig(ctx)
	assert.NoError(t, err)
	assert.Equal(t, expectedDestDataAvailabilityOverheadGas, destDataAvailabilityOverheadGas)
	assert.Equal(t, expectedDestGasPerDataAvailabilityByte, destGasPerDataAvailabilityByte)
	assert.Equal(t, expectedDestDataAvailabilityMultiplierBps, destDataAvailabilityMultiplierBps)

	onRampReader.On("GetDynamicConfig", ctx).
		Return(ccip.OnRampDynamicConfig{}, errors.New("test")).Once()
	_, _, _, err = svc.GetDataAvailabilityConfig(ctx)
	assert.Error(t, err)
}
