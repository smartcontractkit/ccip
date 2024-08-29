package persistent

import (
	"fmt"

	ccipconfig "github.com/smartcontractkit/ccip/integration-tests/ccip-tests/testconfig"
	ctf_config "github.com/smartcontractkit/chainlink-testing-framework/config"
	"github.com/smartcontractkit/chainlink-testing-framework/docker"
	"github.com/smartcontractkit/chainlink-testing-framework/logging"
	"github.com/smartcontractkit/chainlink-testing-framework/networks"
	"github.com/smartcontractkit/chainlink-testing-framework/seth"
)

// TODO in the future Seth config should be part of the test config
func EVMChainConfigFromTestConfig(testCfg ccipconfig.Config, sethConfig *seth.Config) (ChainConfig, error) {
	evmChainConfig := ChainConfig{
		NewEVMChains:      make([]NewEVMChainConfig, 0),
		ExistingEVMChains: make([]ExistingEVMChainConfig, 0),
	}

	var getSimulatedNetworkFromTestConfig = func(testConfig ccipconfig.Config, chainId uint64) (ctf_config.EthereumNetworkConfig, error) {
		for _, chainCfg := range testConfig.CCIP.Env.PrivateEthereumNetworks {
			if uint64(chainCfg.EthereumChainConfig.ChainID) == chainId {
				return *chainCfg, nil
			}
		}

		return ctf_config.EthereumNetworkConfig{}, fmt.Errorf("chain id %d not found in test config", chainId)
	}

	dockerNetwork, err := docker.CreateNetwork(logging.GetLogger(nil, "CORE_DOCKER_ENV_LOG_LEVEL"))
	if err != nil {
		return evmChainConfig, err
	}

	for _, network := range networks.MustGetSelectedNetworkConfig(testCfg.CCIP.Env.Network) {
		if network.Simulated {
			privateNetworkCfg, err := getSimulatedNetworkFromTestConfig(testCfg, uint64(network.ChainID))
			if err != nil {
				return evmChainConfig, err
			}
			privateNetworkCfg.DockerNetworkNames = []string{dockerNetwork.Name}
			var chainConfig NewEVMChainConfig
			if sethConfig == nil {
				chainConfig = CreateNewEVMChainWithGeth(&privateNetworkCfg)
			} else {
				chainConfig, err = CreateNewEVMChainWithSeth(&privateNetworkCfg, *sethConfig)
				if err != nil {
					return evmChainConfig, err
				}
			}
			evmChainConfig.NewEVMChains = append(evmChainConfig.NewEVMChains, chainConfig)
		} else {
			var chainConfig ExistingEVMChainConfig
			if sethConfig == nil {
				chainConfig = CreateExistingEVMChainConfigWithGeth(network)
			} else {
				chainConfig, err = CreateExistingEVMChainWithSeth(network, *sethConfig)
				if err != nil {
					return evmChainConfig, err
				}
			}
			evmChainConfig.ExistingEVMChains = append(evmChainConfig.ExistingEVMChains, chainConfig)
		}
	}

	return evmChainConfig, nil
}
