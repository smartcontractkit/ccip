package config

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/core/services/pipeline"
	"github.com/smartcontractkit/chainlink/core/store/models"
)

// RelayPluginConfig contains the plugin specific variables for the ccip.CCIPRelay plugin.
// We use ID here to keep it as general as possible, e.g. abstracting for chains which don't have an address concept.
type RelayPluginConfig struct {
	SourceChainID                    uint64 `json:"sourceChainID"`
	DestChainID                      uint64 `json:"destChainID"`
	SourceStartBlock, DestStartBlock int64  // Only for first time job add.
	// We relay from multiple onramps from the same source chain. E.g. different message types.
	OnRampIDs           []string        `json:"onRampIDs"`
	PollPeriod          models.Duration `json:"pollPeriod"`
	InflightCacheExpiry models.Duration `json:"inflightCacheExpiry"`
}

// ValidateRelayPluginConfig validates the arguments for the CCIP Relay plugin.
// It will return an error if there is anything wrong with the provided config.
func (rp *RelayPluginConfig) ValidateRelayPluginConfig() error {
	// TODO: Validation based on chainID
	// for now, all EVM.
	for _, onRamp := range rp.OnRampIDs {
		if !common.IsHexAddress(onRamp) {
			return errors.Errorf("%v is not a valid EIP155 address", onRamp)
		}
	}

	return nil
}

// ExecutionPluginConfig contains the plugin specific variables for the ccip.CCIPExecution plugin.
type ExecutionPluginConfig struct {
	SourceChainID uint64 `json:"sourceChainID"`
	DestChainID   uint64 `json:"destChainID"`
	// We execute for a single on/offramp pair (lane) between a given source/dest chain. E.g. a single message types.
	OnRampID                         string          `json:"onRampID"`
	BlobVerifierID                   string          `json:"blobVerifierID"`
	SourceStartBlock, DestStartBlock int64           // Only for first time job add.
	TokensPerFeeCoinPipeline         string          `json:"tokensPerFeeCoinPipeline"`
	InflightCacheExpiry              models.Duration `json:"inflightCacheExpiry"`
	RootSnoozeTime                   models.Duration `json:"rootSnoozeTime"`
}

// ValidateExecutionPluginConfig validates the arguments for the CCIP Execution plugin.
// It will return an error if there is anything wrong with the provided config.
func (ep *ExecutionPluginConfig) ValidateExecutionPluginConfig() error {
	// TODO: Validation based on chainIDs, for now all EVM.
	if !common.IsHexAddress(ep.OnRampID) {
		return errors.Errorf("%v is not a valid EIP155 address", ep.OnRampID)
	}
	if !common.IsHexAddress(ep.BlobVerifierID) {
		return errors.Errorf("%v is not a valid EIP155 address", ep.BlobVerifierID)
	}
	_, err := pipeline.Parse(ep.TokensPerFeeCoinPipeline)
	return err
}
