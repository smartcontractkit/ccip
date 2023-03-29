package config

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/v2/core/store/models"
)

// CommitPluginConfig contains the plugin specific variables for the ccip.CCIPCommit plugin.
// We use ID here to keep it as general as possible, e.g. abstracting for chains which don't have an address concept.
type CommitPluginConfig struct {
	SourceChainID                    uint64          `json:"sourceChainID"`
	SourceStartBlock, DestStartBlock int64           // Only for first time job add.
	OnRampID                         string          `json:"onRampID"`
	OffRampID                        string          `json:"offRampID"`
	PollPeriod                       models.Duration `json:"pollPeriod"`
	InflightCacheExpiry              models.Duration `json:"inflightCacheExpiry"`
	// TokenPricesUSDPipeline should contain a token price pipeline for the following tokens:
	//		The SOURCE chain wrapped native
	// 		The DESTINATION fee tokens as defined in the destination chain PriceRegistry.
	TokenPricesUSDPipeline string `json:"tokenPricesUSDPipeline"`
}

// ValidateCommitPluginConfig validates the arguments for the CCIP commit plugin.
// It will return an error if there is anything wrong with the provided config.
func (cp *CommitPluginConfig) ValidateCommitPluginConfig() error {
	if cp.SourceChainID == 0 {
		return errors.Errorf("%v is not a valid source chain Id", cp.SourceChainID)
	}
	// TODO: Validation based on chainID
	// for now, all EVM.
	if !common.IsHexAddress(cp.OnRampID) {
		return errors.Errorf("%v is not a valid onRamp EIP155 address", cp.OnRampID)
	}
	if !common.IsHexAddress(cp.OffRampID) {
		return errors.Errorf("%v is not a valid offRamp EIP155 address", cp.OffRampID)
	}

	return nil
}

// ExecutionPluginConfig contains the plugin specific variables for the ccip.CCIPExecution plugin.
type ExecutionPluginConfig struct {
	SourceChainID                    uint64          `json:"sourceChainID"`
	OnRampID                         string          `json:"onRampID"`
	CommitStoreID                    string          `json:"commitStoreID"`
	SourceStartBlock, DestStartBlock int64           // Only for first time job add.
	InflightCacheExpiry              models.Duration `json:"inflightCacheExpiry"`
	RootSnoozeTime                   models.Duration `json:"rootSnoozeTime"`
}

// ValidateExecutionPluginConfig validates the arguments for the CCIP Execution plugin.
// It will return an error if there is anything wrong with the provided config.
func (ep *ExecutionPluginConfig) ValidateExecutionPluginConfig() error {
	// TODO: Validation based on chainIDs, for now all EVM.
	if ep.SourceChainID == 0 {
		return errors.Errorf("%v is not a valid source chain Id", ep.SourceChainID)
	}
	if !common.IsHexAddress(ep.OnRampID) {
		return errors.Errorf("%v is not a valid EIP155 address", ep.OnRampID)
	}
	if !common.IsHexAddress(ep.CommitStoreID) {
		return errors.Errorf("%v is not a valid EIP155 address", ep.CommitStoreID)
	}
	return nil
}
