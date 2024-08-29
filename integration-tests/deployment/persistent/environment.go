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
	DONConfig
	RegistryConfig RegistryConfig
}

func NewEnvironment(lggr logger.Logger, config EnvironmentConfig) (*deployment.Environment, error) {
	chains, err := NewChains(lggr, config.ChainConfig)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create Chains")
	}

	// TODO add logstream
	// TODO add clean ups? although that should be related to test, not to environment

	// TODO we probably need to pass chains to existing env as well, at least as long as we need to configure them
	if config.DONConfig.NewDON != nil {
		config.DONConfig.NewDON.Chains = chains
	}

	don, err := NewNodes(config.DONConfig)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create nodes")
	}
	_ = don
	var nodeIDs []string
	for _, node := range don.Keys {
		for _, keys := range node {
			nodeIDs = append(nodeIDs, keys.PeerID)
		}
		// peer ids are the same for all nodes, so we can iterate only once
		break
	}
	return &deployment.Environment{
		Name: DevEnv,
		//Offchain: NewMemoryJobClient(nodes),
		NodeIDs: nodeIDs,
		Chains:  chains,
		Logger:  lggr,
	}, nil
}
