package v1_2_0

import (
	"math/big"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	ccipconfig "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
	"github.com/smartcontractkit/chainlink/v2/core/store/models"
)

func TestCommitReportEncoding(t *testing.T) {
	report := ccipdata.CommitStoreReport{
		TokenPrices: []ccipdata.TokenPrice{
			{
				Token: utils.RandomAddress(),
				Value: big.NewInt(9e18),
			},
			{
				Token: utils.RandomAddress(),
				Value: big.NewInt(1e18),
			},
		},
		GasPrices: []ccipdata.GasPrice{
			{
				DestChainSelector: rand.Uint64(),
				Value:             big.NewInt(2000e9),
			},
			{
				DestChainSelector: rand.Uint64(),
				Value:             big.NewInt(3000e9),
			},
		},
		MerkleRoot: [32]byte{123},
		Interval:   ccipdata.CommitStoreInterval{Min: 1, Max: 10},
	}

	c, err := NewCommitStore(logger.TestLogger(t), utils.RandomAddress(), nil, mocks.NewLogPoller(t), nil)
	assert.NoError(t, err)

	encodedReport, err := c.EncodeCommitReport(report)
	require.NoError(t, err)
	assert.Greater(t, len(encodedReport), 0)

	decodedReport, err := c.DecodeCommitReport(encodedReport)
	require.NoError(t, err)
	require.Equal(t, report, decodedReport)
}

func TestCommitStoreV120ffchainConfigEncoding(t *testing.T) {
	validConfig := CommitOffchainConfig{
		SourceFinalityDepth:      3,
		DestFinalityDepth:        4,
		MaxGasPrice:              200e9,
		GasPriceHeartBeat:        models.MustMakeDuration(1 * time.Minute),
		DAGasPriceDeviationPPB:   10,
		ExecGasPriceDeviationPPB: 11,
		TokenPriceHeartBeat:      models.MustMakeDuration(2 * time.Minute),
		TokenPriceDeviationPPB:   12,
		InflightCacheExpiry:      models.MustMakeDuration(3 * time.Minute),
	}

	require.NoError(t, validConfig.Validate())

	tests := map[string]struct {
		want       CommitOffchainConfig
		errPattern string
	}{
		"legacy offchain config format parses": {
			want: validConfig,
		},
		"can omit finality depth": {
			want: modifyCopy(validConfig, func(c *CommitOffchainConfig) {
				c.SourceFinalityDepth = 0
				c.DestFinalityDepth = 0
			}),
		},
		"can set the SourceMaxGasPrice": {
			want: modifyCopy(validConfig, func(c *CommitOffchainConfig) {
				c.MaxGasPrice = 0
				c.SourceMaxGasPrice = 200e9
			}),
		},
		"must set SourceMaxGasPrice": {
			want: modifyCopy(validConfig, func(c *CommitOffchainConfig) {
				c.MaxGasPrice = 0
				c.SourceMaxGasPrice = 0
			}),
			errPattern: "SourceMaxGasPrice",
		},
		"cannot set both MaxGasPrice and SourceMaxGasPrice": {
			want: modifyCopy(validConfig, func(c *CommitOffchainConfig) {
				c.SourceMaxGasPrice = c.MaxGasPrice
			}),
			errPattern: "MaxGasPrice and SourceMaxGasPrice",
		},
		"must set GasPriceHeartBeat": {
			want: modifyCopy(validConfig, func(c *CommitOffchainConfig) {
				c.GasPriceHeartBeat = models.MustMakeDuration(0)
			}),
			errPattern: "GasPriceHeartBeat",
		},
		"must set ExecGasPriceDeviationPPB": {
			want: modifyCopy(validConfig, func(c *CommitOffchainConfig) {
				c.ExecGasPriceDeviationPPB = 0
			}),
			errPattern: "ExecGasPriceDeviationPPB",
		},
		"must set TokenPriceHeartBeat": {
			want: modifyCopy(validConfig, func(c *CommitOffchainConfig) {
				c.TokenPriceHeartBeat = models.MustMakeDuration(0)
			}),
			errPattern: "TokenPriceHeartBeat",
		},
		"must set TokenPriceDeviationPPB": {
			want: modifyCopy(validConfig, func(c *CommitOffchainConfig) {
				c.TokenPriceDeviationPPB = 0
			}),
			errPattern: "TokenPriceDeviationPPB",
		},
		"must set InflightCacheExpiry": {
			want: modifyCopy(validConfig, func(c *CommitOffchainConfig) {
				c.InflightCacheExpiry = models.MustMakeDuration(0)
			}),
			errPattern: "InflightCacheExpiry",
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			exp := tc.want
			encode, err := ccipconfig.EncodeOffchainConfig(&exp)
			require.NoError(t, err)
			got, err := ccipconfig.DecodeOffchainConfig[CommitOffchainConfig](encode)

			if tc.errPattern != "" {
				require.ErrorContains(t, err, tc.errPattern)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.want, got)
			}
		})
	}
}

func TestCommitStoreV120ComputesGasPrice(t *testing.T) {
	validConfig := CommitOffchainConfig{
		SourceFinalityDepth:      3,
		DestFinalityDepth:        4,
		MaxGasPrice:              200e9,
		GasPriceHeartBeat:        models.MustMakeDuration(1 * time.Minute),
		DAGasPriceDeviationPPB:   10,
		ExecGasPriceDeviationPPB: 11,
		TokenPriceHeartBeat:      models.MustMakeDuration(2 * time.Minute),
		TokenPriceDeviationPPB:   12,
		InflightCacheExpiry:      models.MustMakeDuration(3 * time.Minute),
	}

	require.NoError(t, validConfig.Validate())
	require.Equal(t, uint64(200e9), validConfig.ComputeSourceMaxGasPrice())

	validConfig.MaxGasPrice = 0
	validConfig.SourceMaxGasPrice = 250e9
	require.NoError(t, validConfig.Validate())
	require.Equal(t, uint64(250e9), validConfig.ComputeSourceMaxGasPrice())
}
