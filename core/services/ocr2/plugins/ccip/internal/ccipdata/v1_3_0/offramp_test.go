package v1_3_0_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	ccipconfig "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata/v1_3_0"
	"github.com/smartcontractkit/chainlink/v2/core/store/models"
)

func TestExecOffchainConfig100_AllFieldsRequired(t *testing.T) {
	config := v1_3_0.ExecOffchainConfig{
		DestOptimisticConfirmations: 6,
		BatchGasLimit:               5_000_000,
		RelativeBoostPerWaitHour:    0.07,
		DestMaxGasPrice:             200e9,
		InflightCacheExpiry:         models.MustMakeDuration(64 * time.Second),
		RootSnoozeTime:              models.MustMakeDuration(128 * time.Minute),
	}
	encoded, err := ccipconfig.EncodeOffchainConfig(&config)
	require.NoError(t, err)

	var configAsMap map[string]any
	err = json.Unmarshal(encoded, &configAsMap)
	require.NoError(t, err)
	for keyToDelete := range configAsMap {
		partialConfig := make(map[string]any)
		for k, v := range configAsMap {
			if k != keyToDelete {
				partialConfig[k] = v
			}
		}
		encodedPartialConfig, err := json.Marshal(partialConfig)
		require.NoError(t, err)
		_, err = ccipconfig.DecodeOffchainConfig[v1_3_0.ExecOffchainConfig](encodedPartialConfig)
		require.ErrorContains(t, err, keyToDelete)
	}
}
