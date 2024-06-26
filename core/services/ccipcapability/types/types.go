package types

import (
	"context"

	ccc "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/ccip_capability_configuration"
)

type PluginType uint8

const (
	PluginTypeCCIPCommit PluginType = 0
	PluginTypeCCIPExec   PluginType = 1
)

// CCIPOracle represents either a CCIP commit or exec oracle or a bootstrap node.
//
//go:generate mockery --name CCIPOracle --output ./mocks/ --case underscore
type CCIPOracle interface {
	Close() error
	Start() error
}

// OracleCreator is an interface for creating CCIP oracles.
// Whether the oracle uses a LOOPP or not is an implementation detail.
//
//go:generate mockery --name OracleCreator --output ./mocks/ --case underscore
type OracleCreator interface {
	// CreateCommitOracle creates a new oracle that will run the CCIP commit plugin.
	// The oracle must be returned unstarted.
	CreateCommitOracle(config ccc.CCIPCapabilityConfigurationOCR3ConfigWithMeta) (CCIPOracle, error)

	// CreateExecOracle creates a new oracle that will run the CCIP exec plugin.
	// The oracle must be returned unstarted.
	CreateExecOracle(config ccc.CCIPCapabilityConfigurationOCR3ConfigWithMeta) (CCIPOracle, error)

	// CreateBootstrapOracle creates a new bootstrap node with the given OCR config.
	// The oracle must be returned unstarted.
	CreateBootstrapOracle(config ccc.CCIPCapabilityConfigurationOCR3ConfigWithMeta) (CCIPOracle, error)
}

// HomeChainReader is an interface for reading CCIP chain and OCR configurations from the home chain.
//
//go:generate mockery --name HomeChainReader --output ./mocks/ --case underscore
type HomeChainReader interface {
	// GetAllChainConfigs returns all chain configurations defined on the home chain.
	// The key is the chain selector.
	GetAllChainConfigs(ctx context.Context) (map[uint64]ccc.CCIPCapabilityConfigurationChainConfigInfo, error)

	// GetOCRConfigs returns all OCR configurations for a given DON ID and plugin type.
	GetOCRConfigs(ctx context.Context, donID uint32, pluginType PluginType) ([]ccc.CCIPCapabilityConfigurationOCR3ConfigWithMeta, error)

	// IsHealthy returns true if the home chain reader is healthy.
	IsHealthy() bool
}
