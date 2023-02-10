package config

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/core/services/pipeline"
	"github.com/smartcontractkit/chainlink/core/store/models"
)

// CommitPluginConfig contains the plugin specific variables for the ccip.CCIPCommit plugin.
// We use ID here to keep it as general as possible, e.g. abstracting for chains which don't have an address concept.
type CommitPluginConfig struct {
	SourceChainID                    uint64          `json:"sourceChainID"`
	SourceStartBlock, DestStartBlock int64           // Only for first time job add.
	OnRampID                         string          `json:"onRampID"`
	PollPeriod                       models.Duration `json:"pollPeriod"`
	InflightCacheExpiry              models.Duration `json:"inflightCacheExpiry"`
}

// ValidateCommitPluginConfig validates the arguments for the CCIP commit plugin.
// It will return an error if there is anything wrong with the provided config.
func (rp *CommitPluginConfig) ValidateCommitPluginConfig() error {
	// TODO: Validation based on chainID
	// for now, all EVM.
	if !common.IsHexAddress(rp.OnRampID) {
		return errors.Errorf("%v is not a valid EIP155 address", rp.OnRampID)
	}

	return nil
}

// ExecutionPluginConfig contains the plugin specific variables for the ccip.CCIPExecution plugin.
type ExecutionPluginConfig struct {
	SourceChainID                    uint64          `json:"sourceChainID"`
	OnRampID                         string          `json:"onRampID"`
	CommitStoreID                    string          `json:"commitStoreID"`
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
	if !common.IsHexAddress(ep.CommitStoreID) {
		return errors.Errorf("%v is not a valid EIP155 address", ep.CommitStoreID)
	}
	_, err := pipeline.Parse(ep.TokensPerFeeCoinPipeline)
	return err
}
