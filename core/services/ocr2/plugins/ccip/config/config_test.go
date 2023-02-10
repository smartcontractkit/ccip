package config

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/core/store/models"
)

func TestCommitConfig(t *testing.T) {
	exampleConfig := CommitPluginConfig{
		SourceChainID:       1337,
		SourceStartBlock:    222,
		DestStartBlock:      333,
		OnRampID:            "0xAb5801a7D398351b8bE11C439e05C5B3259aeC9B",
		PollPeriod:          models.MustMakeDuration(5 * time.Second),
		InflightCacheExpiry: models.MustMakeDuration(23456 * time.Second),
	}

	bts, err := json.Marshal(exampleConfig)
	require.NoError(t, err)

	parsedConfig := CommitPluginConfig{}
	require.NoError(t, json.Unmarshal(bts, &parsedConfig))

	require.Equal(t, exampleConfig, parsedConfig)
	require.NoError(t, parsedConfig.ValidateCommitPluginConfig())
}

func TestExecutionConfig(t *testing.T) {
	exampleConfig := ExecutionPluginConfig{
		SourceChainID:            1337,
		OnRampID:                 "0xAb5801a7D398351b8bE11C439e05C5B3259aeC9B",
		CommitStoreID:            "0xC79b96044906550A5652BCf20a6EA02f139B9Ae5",
		SourceStartBlock:         222,
		DestStartBlock:           333,
		TokensPerFeeCoinPipeline: `merge [type=merge left="{}" right="{\"0xC79b96044906550A5652BCf20a6EA02f139B9Ae5\":\"1000000000000000000\"}"];`,
		InflightCacheExpiry:      models.MustMakeDuration(64 * time.Second),
		RootSnoozeTime:           models.MustMakeDuration(128 * time.Minute),
	}

	bts, err := json.Marshal(exampleConfig)
	require.NoError(t, err)

	parsedConfig := ExecutionPluginConfig{}
	require.NoError(t, json.Unmarshal(bts, &parsedConfig))

	require.Equal(t, exampleConfig, parsedConfig)
	require.NoError(t, parsedConfig.ValidateExecutionPluginConfig())
}
