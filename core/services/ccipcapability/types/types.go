package types

import (
	"context"

	kcr "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/keystone/generated/keystone_capability_registry"
	p2ptypes "github.com/smartcontractkit/chainlink/v2/core/services/p2p/types"
)

type DonID = uint32
type HashedCapabilityID = [32]byte

type RegistryState struct {
	IDsToDONs         map[DonID]kcr.CapabilityRegistryDONInfo
	IDsToNodes        map[p2ptypes.PeerID]kcr.CapabilityRegistryNodeInfo
	IDsToCapabilities map[HashedCapabilityID]kcr.CapabilityRegistryCapability
}

type CapabilityRegistry interface {
	// LatestState returns the latest state of the on-chain capability registry.
	LatestState() (RegistryState, error)
}

type ChainConfig interface {
	Readers() [][32]byte
	FChain() uint8
}

type PluginType uint8

const (
	PluginTypeCCIPCommit PluginType = 0
	PluginTypeCCIPExec   PluginType = 1
)

//go:generate mockery --name OCRConfig --output ./mocks/ --case underscore
type OCRConfig interface {
	PluginType() PluginType
	ChainSelector() uint64
	F() uint8
	OffchainConfigVersion() uint64
	OfframpAddress() []byte
	BootstrapP2PIDs() [][32]byte
	P2PIDs() [][32]byte
	Signers() [][]byte
	Transmitters() [][]byte
	OffchainConfig() []byte
	String() string
}

// CCIPOracle represents either a CCIP commit or exec oracle or a bootstrap node.
//
//go:generate mockery --name CCIPOracle --output ./mocks/ --case underscore
type CCIPOracle interface {
	Shutdown() error
	Start() error
}

// OracleCreator is an interface for creating CCIP oracles.
// Whether the oracle uses a LOOPP or not is an implementation detail.
//
//go:generate mockery --name OracleCreator --output ./mocks/ --case underscore
type OracleCreator interface {
	// CreateCommitOracle creates a new oracle that will run the CCIP commit plugin.
	// The oracle must be returned unstarted.
	CreateCommitOracle(config OCRConfig) (CCIPOracle, error)

	// CreateExecOracle creates a new oracle that will run the CCIP exec plugin.
	// The oracle must be returned unstarted.
	CreateExecOracle(config OCRConfig) (CCIPOracle, error)

	// CreateBootstrapOracle creates a new bootstrap node with the given OCR config.
	// The oracle must be returned unstarted.
	CreateBootstrapOracle(config OCRConfig) (CCIPOracle, error)
}

// HomeChainReader is an interface for reading CCIP chain and OCR configurations from the home chain.
//
//go:generate mockery --name HomeChainReader --output ./mocks/ --case underscore
type HomeChainReader interface {
	// GetAllChainConfigs returns all chain configurations defined on the home chain.
	// The key is the chain selector.
	GetAllChainConfigs(ctx context.Context) (map[uint64]ChainConfig, error)

	// GetOCRConfigs returns all OCR configurations for a given DON ID and plugin type.
	GetOCRConfigs(ctx context.Context, donID uint32, pluginType PluginType) ([]OCRConfig, error)

	// IsHealthy returns true if the home chain reader is healthy.
	IsHealthy() bool
}
