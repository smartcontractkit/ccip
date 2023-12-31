package v1_0_0

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	ccipconfig "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/v2/core/store/models"
)

func TestExecOffchainConfig100_Encoding(t *testing.T) {
	tests := map[string]struct {
		want      ExecOffchainConfig
		expectErr bool
	}{
		"encodes and decodes config with all fields set": {
			want: ExecOffchainConfig{
				SourceFinalityDepth:         3,
				DestOptimisticConfirmations: 6,
				DestFinalityDepth:           3,
				BatchGasLimit:               5_000_000,
				RelativeBoostPerWaitHour:    0.07,
				MaxGasPrice:                 200e9,
				InflightCacheExpiry:         models.MustMakeDuration(64 * time.Second),
				RootSnoozeTime:              models.MustMakeDuration(128 * time.Minute),
			},
		},
		"fails decoding when all fields present but with 0 values": {
			want: ExecOffchainConfig{
				SourceFinalityDepth:         0,
				DestFinalityDepth:           0,
				DestOptimisticConfirmations: 0,
				BatchGasLimit:               0,
				RelativeBoostPerWaitHour:    0,
				MaxGasPrice:                 0,
				InflightCacheExpiry:         models.MustMakeDuration(0),
				RootSnoozeTime:              models.MustMakeDuration(0),
			},
			expectErr: true,
		},
		"fails decoding when all fields are missing": {
			want:      ExecOffchainConfig{},
			expectErr: true,
		},
		"fails decoding when some fields are missing": {
			want: ExecOffchainConfig{
				SourceFinalityDepth: 99999999,
				InflightCacheExpiry: models.MustMakeDuration(64 * time.Second),
			},
			expectErr: true,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			exp := tc.want
			encode, err := ccipconfig.EncodeOffchainConfig(&exp)
			require.NoError(t, err)
			got, err := ccipconfig.DecodeOffchainConfig[ExecOffchainConfig](encode)

			if tc.expectErr {
				require.ErrorContains(t, err, "must set")
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.want, got)
			}
		})
	}
}

func TestExecOffchainConfig100_AllFieldsRequired(t *testing.T) {
	config := ExecOffchainConfig{
		SourceFinalityDepth:         3,
		DestOptimisticConfirmations: 6,
		DestFinalityDepth:           3,
		BatchGasLimit:               5_000_000,
		RelativeBoostPerWaitHour:    0.07,
		MaxGasPrice:                 200e9,
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
		_, err = ccipconfig.DecodeOffchainConfig[ExecOffchainConfig](encodedPartialConfig)
		require.ErrorContains(t, err, keyToDelete)
	}
}
