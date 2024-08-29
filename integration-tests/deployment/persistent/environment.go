package persistent

import (
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/integration-tests/deployment"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
)

const (
	DevEnv = "devenv"
)

type EnvironmentConfig struct {
	ChainConfig
	//DON            DON
	RegistryConfig RegistryConfig
}

func NewEnvironment(lggr logger.Logger, config EnvironmentConfig) (*deployment.Environment, error) {
	chains, err := NewChains(lggr, config.ChainConfig)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create chains")
	}
	//nodes := NewNodes(t, chains, config.Nodes, config.Bootstraps, config.RegistryConfig)
	//var nodeIDs []string
	//for id := range nodes {
	//	nodeIDs = append(nodeIDs, id)
	//}
	return &deployment.Environment{
		Name: DevEnv,
		//Offchain: NewMemoryJobClient(nodes),
		//NodeIDs:  nodeIDs,
		Chains: chains,
		Logger: lggr,
	}, nil
}
