package v1_3_0_test

import (
	"testing"
	"time"

	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata/v1_3_0"
	"github.com/smartcontractkit/chainlink/v2/core/store/models"
)

func TestExecOffchainConfig130_AllFieldsRequired(t *testing.T) {
	ccipdata.AssertAllFieldsAreRequired(t, v1_3_0.ExecOffchainConfig{
		DestOptimisticConfirmations: 6,
		BatchGasLimit:               5_000_000,
		RelativeBoostPerWaitHour:    0.07,
		DestMaxGasPrice:             200e9,
		InflightCacheExpiry:         models.MustMakeDuration(64 * time.Second),
		RootSnoozeTime:              models.MustMakeDuration(128 * time.Minute),
	})
}
