package devenv

import (
	"fmt"

	"github.com/smartcontractkit/ccip/integration-tests/deployment"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
)

const (
	DevEnv = "devenv"
)

type EnvironmentConfig struct {
	Chains         []deployment.ChainConfig
	DON            deployment.DON
	RegistryConfig deployment.RegistryConfig
	JDConfig       deployment.JDConfig
}

func NewEnvironment(lggr logger.Logger, config EnvironmentConfig) (*deployment.Environment, error) {
	chains, err := deployment.NewChains(lggr, config.Chains)
	if err != nil {
		return nil, fmt.Errorf("failed to create chains: %w", err)
	}
	nodes := NewNodes(t, chains, config.Nodes, config.Bootstraps, config.RegistryConfig)
	var nodeIDs []string
	for id := range nodes {
		nodeIDs = append(nodeIDs, id)
	}
	return &deployment.Environment{
		Name:     DevEnv,
		Offchain: NewMemoryJobClient(nodes),
		NodeIDs:  nodeIDs,
		Chains:   chains,
		Logger:   lggr,
	}, nil
}
