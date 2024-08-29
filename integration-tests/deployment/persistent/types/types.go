package types

import (
	"github.com/smartcontractkit/ccip/integration-tests/deployment"
	"github.com/smartcontractkit/chainlink-testing-framework/blockchain"
)

type ChainConfig struct {
	// ExistingEVMChains are Chains that are already running in a separate process or machine.
	ExistingEVMChains []ExistingEVMChainConfig
	// NewEVMChains are Chains that will be started by the test environment.
	NewEVMChains []NewEVMChainConfig
}

type NewEVMChainConfig interface {
	Chain() (deployment.Chain, error)
}

type ExistingEVMChainConfig interface {
	Chain() (deployment.Chain, error)
	EVMNetwork() blockchain.EVMNetwork
}
