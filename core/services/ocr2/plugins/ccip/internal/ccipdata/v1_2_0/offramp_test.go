package v1_2_0

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	ccipconfig "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/v2/core/store/models"
)

func modifyCopy[T any](c T, f func(c *T)) T {
	f(&c)
	return c
}

func TestExecOffchainConfig120_Encoding(t *testing.T) {
	validConfig := ExecOffchainConfig{
		SourceFinalityDepth:         3,
		DestOptimisticConfirmations: 6,
		DestFinalityDepth:           3,
		BatchGasLimit:               5_000_000,
		RelativeBoostPerWaitHour:    0.07,
		MaxGasPrice:                 200e9,
		InflightCacheExpiry:         models.MustMakeDuration(64 * time.Second),
		RootSnoozeTime:              models.MustMakeDuration(128 * time.Minute),
	}

	tests := map[string]struct {
		want       ExecOffchainConfig
		errPattern string
	}{
		"legacy offchain config format parses": {
			want: validConfig,
		},
		"can omit finality depth": {
			want: modifyCopy(validConfig, func(c *ExecOffchainConfig) {
				c.SourceFinalityDepth = 0
				c.DestFinalityDepth = 0
			}),
		},
		"can set the DestMaxGasPrice": {
			want: modifyCopy(validConfig, func(c *ExecOffchainConfig) {
				c.MaxGasPrice = 0
				c.DestMaxGasPrice = 200e9
			}),
		},
		"cannot set both MaxGasPrice and DestMaxGasPrice": {
			want: modifyCopy(validConfig, func(c *ExecOffchainConfig) {
				c.DestMaxGasPrice = c.MaxGasPrice
			}),
			errPattern: "MaxGasPrice and DestMaxGasPrice",
		},
		"must set BatchGasLimit": {
			want: modifyCopy(validConfig, func(c *ExecOffchainConfig) {
				c.BatchGasLimit = 0
			}),
			errPattern: "BatchGasLimit",
		},
		"must set DestOptimisticConfirmations": {
			want: modifyCopy(validConfig, func(c *ExecOffchainConfig) {
				c.DestOptimisticConfirmations = 0
			}),
			errPattern: "DestOptimisticConfirmations",
		},
		"must set RelativeBoostPerWaitHour": {
			want: modifyCopy(validConfig, func(c *ExecOffchainConfig) {
				c.RelativeBoostPerWaitHour = 0
			}),
			errPattern: "RelativeBoostPerWaitHour",
		},
		"must set InflightCacheExpiry": {
			want: modifyCopy(validConfig, func(c *ExecOffchainConfig) {
				c.InflightCacheExpiry = models.MustMakeDuration(0)
			}),
			errPattern: "InflightCacheExpiry",
		},
		"must set RootSnoozeTime": {
			want: modifyCopy(validConfig, func(c *ExecOffchainConfig) {
				c.RootSnoozeTime = models.MustMakeDuration(0)
			}),
			errPattern: "RootSnoozeTime",
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			exp := tc.want
			encode, err := ccipconfig.EncodeOffchainConfig(&exp)
			require.NoError(t, err)
			got, err := ccipconfig.DecodeOffchainConfig[ExecOffchainConfig](encode)

			if tc.errPattern != "" {
				require.ErrorContains(t, err, tc.errPattern)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.want, got)
			}
		})
	}
}

func TestExecOffchainConfig120_MaxGasPrice(t *testing.T) {
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
	require.NoError(t, config.Validate())
	require.Equal(t, uint64(200e9), config.ComputeDestMaxGasPrice())

	config.MaxGasPrice = 0
	config.DestMaxGasPrice = 250e9
	require.NoError(t, config.Validate())
	require.Equal(t, uint64(250e9), config.ComputeDestMaxGasPrice())
}

func TestExecOffchainConfig120_ParseRawJson(t *testing.T) {
	decoded, err := ccipconfig.DecodeOffchainConfig[ExecOffchainConfig]([]byte(`{
		"DestOptimisticConfirmations": 6,
		"BatchGasLimit": 5000000,
		"RelativeBoostPerWaitHour": 0.07,
		"MaxGasPrice": 200000000000,
		"InflightCacheExpiry": "64s",
		"RootSnoozeTime": "128m"
	}`))
	require.NoError(t, err)
	require.Equal(t, ExecOffchainConfig{
		DestOptimisticConfirmations: 6,
		BatchGasLimit:               5_000_000,
		RelativeBoostPerWaitHour:    0.07,
		MaxGasPrice:                 200e9,
		InflightCacheExpiry:         models.MustMakeDuration(64 * time.Second),
		RootSnoozeTime:              models.MustMakeDuration(128 * time.Minute),
	}, decoded)
}
