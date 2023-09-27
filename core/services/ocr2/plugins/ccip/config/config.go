package config

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

// CommitPluginJobSpecConfig contains the plugin specific variables for the ccip.CCIPCommit plugin.
// We use ID here to keep it as general as possible, e.g. abstracting for chains which don't have an address concept.
type CommitPluginJobSpecConfig struct {
	SourceStartBlock, DestStartBlock int64  // Only for first time job add.
	OffRamp                          string `json:"offRamp"`
	// TokenPricesUSDPipeline should contain a token price pipeline for the following tokens:
	//		The SOURCE chain wrapped native
	// 		The DESTINATION supported tokens (including fee tokens) as defined in destination OffRamp and PriceRegistry.
	TokenPricesUSDPipeline string `json:"tokenPricesUSDPipeline"`
}

// ExecutionPluginJobSpecConfig contains the plugin specific variables for the ccip.CCIPExecution plugin.
type ExecutionPluginJobSpecConfig struct {
	SourceStartBlock, DestStartBlock int64 // Only for first time job add.
	USDCConfig                       USDCConfig
}

type USDCConfig struct {
	SourceTokenAddress              common.Address
	SourceMessageTransmitterAddress common.Address
	AttestationAPI                  string
}

func (uc *USDCConfig) ValidateUSDCConfig() error {
	if uc.AttestationAPI == "" && uc.SourceTokenAddress == utils.ZeroAddress && uc.SourceMessageTransmitterAddress == utils.ZeroAddress {
		return nil
	}

	if uc.AttestationAPI == "" {
		return errors.New("AttestationAPI is required")
	}
	if uc.SourceTokenAddress == utils.ZeroAddress {
		return errors.New("SourceTokenAddress is required")
	}
	if uc.SourceMessageTransmitterAddress == utils.ZeroAddress {
		return errors.New("SourceMessageTransmitterAddress is required")
	}

	return nil
}
