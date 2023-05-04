// Package with set of configs that should be used only within tests suites

package testhelpers

import (
	"time"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/v2/core/store/models"
)

func createDefaultCommitOnchainConfig(c *CCIPContracts) []byte {
	config, err := ccip.EncodeAbiStruct(ccip.CommitOnchainConfig{
		PriceRegistry: c.Dest.PriceRegistry.Address(),
		Afn:           c.Dest.AFN.Address(),
	})
	require.NoError(c.t, err)
	return config
}

func createDefaultCommitOffchainConfig(c *CCIPContracts) []byte {
	return createCommitOffchainConfig(c, 10*time.Second, 5*time.Second)
}

func createCommitOffchainConfig(c *CCIPContracts, feeUpdateHearBeat time.Duration, inflightCacheExpiry time.Duration) []byte {
	config, err := ccip.EncodeOffchainConfig(ccip.CommitOffchainConfig{
		SourceIncomingConfirmations: 1,
		DestIncomingConfirmations:   1,
		FeeUpdateHeartBeat:          models.MustMakeDuration(feeUpdateHearBeat),
		FeeUpdateDeviationPPB:       1,
		MaxGasPrice:                 200e9,
		InflightCacheExpiry:         models.MustMakeDuration(inflightCacheExpiry),
	})
	require.NoError(c.t, err)
	return config
}

func createDefaultExecOnchainConfig(c *CCIPContracts) []byte {
	config, err := ccip.EncodeAbiStruct(ccip.ExecOnchainConfig{
		PermissionLessExecutionThresholdSeconds: 60,
		Router:                                  c.Dest.Router.Address(),
		Afn:                                     c.Dest.AFN.Address(),
		PriceRegistry:                           c.Dest.PriceRegistry.Address(),
		MaxDataSize:                             1e5,
		MaxTokensLength:                         5,
	})
	require.NoError(c.t, err)
	return config
}

func createDefaultExecOffchainConfig(c *CCIPContracts) []byte {
	return createExecOffchainConfig(c, 5*time.Second, 1*time.Second)
}

func createExecOffchainConfig(c *CCIPContracts, inflightCacheExpiry time.Duration, rootSnoozeTime time.Duration) []byte {
	config, err := ccip.EncodeOffchainConfig(ccip.ExecOffchainConfig{
		SourceIncomingConfirmations: 1,
		DestIncomingConfirmations:   1,
		BatchGasLimit:               5_000_000,
		RelativeBoostPerWaitHour:    0.07,
		MaxGasPrice:                 200e9,
		InflightCacheExpiry:         models.MustMakeDuration(inflightCacheExpiry),
		RootSnoozeTime:              models.MustMakeDuration(rootSnoozeTime),
	})
	require.NoError(c.t, err)
	return config
}
