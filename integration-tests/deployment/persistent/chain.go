package persistent

import (
	"context"
	"fmt"
	"github.com/avast/retry-go/v4"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/chainlink-testing-framework/docker"
	"github.com/smartcontractkit/chainlink-testing-framework/logging"
	"path/filepath"
	"strings"

	chainselectors "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/chainlink-testing-framework/blockchain"
	ctf_config "github.com/smartcontractkit/chainlink-testing-framework/config"
	ctf_test_env "github.com/smartcontractkit/chainlink-testing-framework/docker/test_env"
	"github.com/smartcontractkit/chainlink-testing-framework/networks"
	"github.com/smartcontractkit/chainlink-testing-framework/seth"
	"github.com/smartcontractkit/chainlink-testing-framework/utils/osutil"
	seth_utils "github.com/smartcontractkit/chainlink-testing-framework/utils/seth"
	"github.com/smartcontractkit/chainlink/integration-tests/deployment"
	"github.com/smartcontractkit/chainlink/v2/core/logger"

	ccipconfig "github.com/smartcontractkit/ccip/integration-tests/ccip-tests/testconfig"
)

type RegistryConfig struct {
	EVMChainID uint64
	Contract   common.Address
}

type ChainConfig struct {
	// ExistingEVMChains are Chains that are already running in a separate process or machine.
	ExistingEVMChains []ExistingEVMChainConfig
	// NewEVMChains are Chains that will be started by the test environment.
	NewEVMChains []NewEVMChainConfig
}

// EVMChainConfigFromTestConfig creates a ChainConfig from a test config.
// TODO in the future Seth config should be part of the test config
func EVMChainConfigFromTestConfig(testCfg ccipconfig.Config, sethConfig seth.Config) (ChainConfig, error) {
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
			chainCfg, err := getSimulatedNetworkFromTestConfig(testCfg, uint64(network.ChainID))
			if err != nil {
				return evmChainConfig, err
			}
			chainCfg.DockerNetworkNames = []string{dockerNetwork.Name}
			evmChainConfig.NewEVMChains = append(evmChainConfig.NewEVMChains, CreateNewPrivateEVMChainConfig(chainCfg, sethConfig))
		} else {
			evmChainConfig.ExistingEVMChains = append(evmChainConfig.ExistingEVMChains, CreateExistingEVMChainConfigWithSeth(network, sethConfig))
		}
	}

	return evmChainConfig, nil
}

type NewEVMChainConfig interface {
	ctf_config.PrivateEthereumNetworkConfig
	SethConfig() seth.Config
	DockerNetworks() []string
}

type NewEVMChainConfigWithSeth struct {
	ctf_config.EthereumNetworkConfig
	sethConfig seth.Config
}

func (n *NewEVMChainConfigWithSeth) SethConfig() seth.Config {
	return n.sethConfig
}

func (n *NewEVMChainConfigWithSeth) DockerNetworks() []string {
	var dockerNetworks []string
	for _, network := range n.GetDockerNetworkNames() {
		contains := false
		for _, dockerNetwork := range dockerNetworks {
			if strings.EqualFold(dockerNetwork, network) {
				contains = true
				break
			}
		}
		if !contains {
			dockerNetworks = append(dockerNetworks, network)
		}
	}
	return dockerNetworks
}

func CreateNewPrivateEVMChainConfig(config ctf_config.EthereumNetworkConfig, sethConfig seth.Config) NewEVMChainConfig {
	return &NewEVMChainConfigWithSeth{
		EthereumNetworkConfig: config,
		sethConfig:            sethConfig,
	}
}

// ExistingEVMChainConfig is a configuration for an existing chain, i.e. chain that is already running in a separate process or machine.
type ExistingEVMChainConfig interface {
	GetEVMNetwork() blockchain.EVMNetwork
	GetSethConfig() seth.Config
}

// ExistingEVMChainConfigWithSeth is a configuration for an existing chain, i.e. chain that is already running in a separate process or machine and has a Seth client.
type ExistingEVMChainConfigWithSeth struct {
	EVMNetwork blockchain.EVMNetwork
	SethConfig seth.Config
}

func (e *ExistingEVMChainConfigWithSeth) GetEVMNetwork() blockchain.EVMNetwork {
	return e.EVMNetwork
}

func (e *ExistingEVMChainConfigWithSeth) GetSethConfig() seth.Config {
	return e.SethConfig
}

func CreateExistingEVMChainConfigWithSeth(evmNetwork blockchain.EVMNetwork, sethConfig seth.Config) ExistingEVMChainConfig {
	return &ExistingEVMChainConfigWithSeth{
		EVMNetwork: evmNetwork,
		SethConfig: sethConfig,
	}
}

// NewChains creates Chains based on the provided configuration. It returns a map of chain id to chain.
// You can mix existing and new Chains in the configuration, meaning that you can have Chains that are already running and Chains that will be started by the test environment.
func NewChains(lggr logger.Logger, config ChainConfig) (map[uint64]deployment.Chain, error) {
	lggr.Info("Creating devenv Chains")
	existingChains, err := newExistingChains(config.ExistingEVMChains)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create existing Chains")
	}
	createdChains, err := newChains(config.NewEVMChains)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create new Chains")
	}
	chains := make(map[uint64]deployment.Chain)
	for k, v := range existingChains {
		if _, ok := chains[k]; ok {
			return nil, errors.Wrapf(err, "duplicate chain id %d used by new and existing Chains", k)
		}
		chains[k] = v
	}
	for k, v := range createdChains {
		chains[k] = v
	}
	return chains, nil
}

func newExistingChains(configs []ExistingEVMChainConfig) (map[uint64]deployment.Chain, error) {
	chains := make(map[uint64]deployment.Chain)
	for _, chainCfg := range configs {
		contractsRootFolder, err := findGethWrappersFolderRoot(5)
		if err != nil {
			return nil, fmt.Errorf("failed to find contracts root folder: %w", err)
		}

		evmNetwork := chainCfg.GetEVMNetwork()
		c, err := seth_utils.MergeSethAndEvmNetworkConfigs(evmNetwork, chainCfg.GetSethConfig())
		if err != nil {
			return nil, fmt.Errorf("failed to merge seth and evm network configs: %w", err)
		}

		sethClient, err := seth.NewClientBuilderWithConfig(&c).
			// we want to set it dynamically, because the path depends on the location of the file in the project
			WithGethWrappersFolders([]string{fmt.Sprintf("%s/ccip", contractsRootFolder)}).
			Build()
		if err != nil {
			return nil, fmt.Errorf("failed to create seth client: %w", err)
		}

		rpcProvider := ctf_test_env.NewRPCProvider(evmNetwork.HTTPURLs, evmNetwork.URLs, evmNetwork.HTTPURLs, evmNetwork.URLs)

		chain, err := buildChain(sethClient, evmNetwork, rpcProvider)
		if err != nil {
			return make(map[uint64]deployment.Chain), err
		}
		chains[uint64(evmNetwork.ChainID)] = chain
	}
	return chains, nil
}

func newChains(configs []NewEVMChainConfig) (map[uint64]deployment.Chain, error) {
	chains := make(map[uint64]deployment.Chain)

	contractsRootFolder, err := findGethWrappersFolderRoot(5)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to find contracts root folder")
	}

	for _, config := range configs {
		if config.GetEthereumVersion() == nil {
			return nil, fmt.Errorf("ethereum version is required")
		}

		if config.GetExecutionLayer() == nil {
			return nil, fmt.Errorf("execution layer is required")
		}

		ethBuilder := ctf_test_env.NewEthereumNetworkBuilder()
		network, err := ethBuilder.
			WithEthereumVersion(*config.GetEthereumVersion()).
			WithExecutionLayer(*config.GetExecutionLayer()).
			WithEthereumChainConfig(config.GetChainConfig()).
			WithDockerNetworks(config.GetDockerNetworkNames()).
			WithCustomDockerImages(config.GetCustomDockerImages()).
			Build()

		if err != nil {
			return chains, err
		}

		evmNetwork, rpcProvider, err := network.Start()
		if err != nil {
			return nil, err
		}

		evmNetwork.Name = fmt.Sprintf("%s-%d", *config.GetExecutionLayer(), evmNetwork.ChainID)

		sethConfig := config.SethConfig()
		sethClient, err := seth.NewClientBuilderWithConfig(&sethConfig).
			// we want to set it dynamically, because the path depends on the location of the file in the project
			WithGethWrappersFolders([]string{fmt.Sprintf("%s/ccip", contractsRootFolder)}).
			WithRpcUrl(evmNetwork.URLs[0]).
			WithPrivateKeys(evmNetwork.PrivateKeys).
			Build()

		if err != nil {
			return nil, errors.Wrapf(err, "failed to create seth client")
		}

		chain, err := buildChain(sethClient, evmNetwork, rpcProvider)
		if err != nil {
			return make(map[uint64]deployment.Chain), err
		}
		chains[uint64(evmNetwork.ChainID)] = chain
	}

	return chains, nil
}

func buildChain(sethClient *seth.Client, evmNetwork blockchain.EVMNetwork, rpcProvider ctf_test_env.RpcProvider) (deployment.Chain, error) {
	shouldRetryOnErrFn := func(err error) bool {
		// some retry logic here
		return true
	}

	prepareReplacementTransactionFn := func(sethClient *seth.Client, tx *types.Transaction) (*types.Transaction, error) {
		// TODO some replacement tx creation logic could go here
		// TODO for example: adjusting base fee aggressively if it's too low for transaction to even be included in the block
		return tx, nil
	}

	sel, err := chainselectors.SelectorFromChainId(uint64(evmNetwork.ChainID))
	if err != nil {
		return deployment.Chain{}, err
	}

	return deployment.Chain{
		Selector: sel,
		Client:   sethClient.Client,
		DeployerKey: func() *bind.TransactOpts {
			// this will use the first private key from the seth client
			// if you want to use N private key you can use sethClient.NewTXKeyOpts(N)
			// we set the nonce to nil, because we want go-ethereum to use pending nonce it gets from the node
			opts := sethClient.NewTXOpts(seth.WithNonce(nil))
			return opts
		}(),
		DeployerKeys: func() []*bind.TransactOpts {
			var keys []*bind.TransactOpts
			// use all private keys set for network, in case we want to use them for concurrent transactions
			for i := range sethClient.Cfg.Network.PrivateKeys {
				// we set the nonce to nil, because we want go-ethereum to use pending nonce it gets from the node
				opts := sethClient.NewTXKeyOpts(i, seth.WithNonce(nil))
				keys = append(keys, opts)
			}

			return keys
		}(),
		EVMNetwork: &deployment.EVMNetworkWithEndpoints{
			EVMNetwork:  evmNetwork,
			RpcProvider: rpcProvider,
		},
		Confirm: func(txHash common.Hash) error {
			ctx, cancelFn := context.WithTimeout(context.Background(), sethClient.Cfg.Network.TxnTimeout.Duration())
			tx, _, err := sethClient.Client.TransactionByHash(ctx, txHash)
			cancelFn()
			if err != nil {
				return err
			}
			_, revertErr := sethClient.DecodeTx(tx)
			return revertErr
		},
		RetrySubmit: func(tx *types.Transaction, err error) (*types.Transaction, error) {
			if err == nil {
				return tx, nil
			}

			retryErr := retry.Do(
				func() error {
					ctx, cancel := context.WithTimeout(context.Background(), sethClient.Cfg.Network.TxnTimeout.Duration())
					defer cancel()

					return sethClient.Client.SendTransaction(ctx, tx)
				}, retry.OnRetry(func(i uint, retryErr error) {
					replacementTx, replacementErr := prepareReplacementTransactionFn(sethClient, tx)
					if replacementErr != nil {
						return
					}
					tx = replacementTx
				}),
				retry.DelayType(retry.FixedDelay),
				retry.Attempts(10),
				retry.RetryIf(shouldRetryOnErrFn),
			)

			return tx, sethClient.DecodeSendErr(retryErr)
		},
	}, nil
}

// findGethWrappersFolderRoot finds the root folder of the geth wrappers. It looks for a file named ".geth_wrappers_root" or ".repo_root" in the current directory and its `folderLimit` parents.
func findGethWrappersFolderRoot(folderLimit int) (string, error) {
	contractsRootFile, err := osutil.FindFile(".geth_wrappers_root", ".repo_root", folderLimit)
	if err != nil {
		return "", fmt.Errorf("failed to find contracts root folder: %w", err)
	}
	return filepath.Dir(contractsRootFile), nil
}
