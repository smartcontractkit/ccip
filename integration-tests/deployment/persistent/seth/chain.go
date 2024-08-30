package seth

import (
	"context"
	"fmt"
	"github.com/avast/retry-go/v4"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"

	chainselectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-testing-framework/blockchain"
	ctf_config "github.com/smartcontractkit/chainlink-testing-framework/config"
	ctf_test_env "github.com/smartcontractkit/chainlink-testing-framework/docker/test_env"
	"github.com/smartcontractkit/chainlink-testing-framework/seth"
	"github.com/smartcontractkit/chainlink-testing-framework/utils/osutil"
	seth_utils "github.com/smartcontractkit/chainlink-testing-framework/utils/seth"
	"github.com/smartcontractkit/chainlink/integration-tests/deployment"
	persistent_types "github.com/smartcontractkit/chainlink/integration-tests/deployment/persistent/types"
	"path/filepath"
	"strings"
)

func CreateNewEVMChainWithSeth(config ctf_config.PrivateEthereumNetworkConfig, sethConfig seth.Config) (persistent_types.NewEVMChainConfig, error) {
	contractsRootFolder, err := findGethWrappersFolderRoot(5)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to find contracts root folder")
	}

	return &NewEVMChainWithSeth{
		config:              config,
		sethConfig:          sethConfig,
		contractsRootFolder: contractsRootFolder,
	}, nil
}

type NewEVMChainWithSeth struct {
	SethChainBuilder
	config              ctf_config.PrivateEthereumNetworkConfig
	sethConfig          seth.Config
	contractsRootFolder string
}

func (n *NewEVMChainWithSeth) Chain() (deployment.Chain, error) {
	chain := deployment.Chain{}
	if n.config.GetEthereumVersion() == nil {
		return chain, fmt.Errorf("ethereum version is required")
	}

	if n.config.GetExecutionLayer() == nil {
		return chain, fmt.Errorf("execution layer is required")
	}

	ethBuilder := ctf_test_env.NewEthereumNetworkBuilder()
	network, err := ethBuilder.
		WithEthereumVersion(*n.config.GetEthereumVersion()).
		WithExecutionLayer(*n.config.GetExecutionLayer()).
		WithEthereumChainConfig(n.config.GetChainConfig()).
		WithDockerNetworks(n.config.GetDockerNetworkNames()).
		WithCustomDockerImages(n.config.GetCustomDockerImages()).
		Build()

	if err != nil {
		return chain, err
	}

	evmNetwork, rpcProvider, err := network.Start()
	if err != nil {
		return chain, err
	}

	evmNetwork.Name = fmt.Sprintf("%s-%d", *n.config.GetExecutionLayer(), evmNetwork.ChainID)
	sethClient, err := seth.NewClientBuilderWithConfig(&n.sethConfig).
		// we want to set it dynamically, because the path depends on the location of the file in the project
		WithGethWrappersFolders([]string{fmt.Sprintf("%s/ccip", n.contractsRootFolder)}).
		WithRpcUrl(evmNetwork.URLs[0]).
		WithPrivateKeys(evmNetwork.PrivateKeys).
		Build()

	if err != nil {
		return chain, errors.Wrapf(err, "failed to create seth client")
	}

	return n.Build(sethClient, evmNetwork, rpcProvider)
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

type ExistingEVMChainConfigWithSeth struct {
	SethChainBuilder
	evmNetwork          blockchain.EVMNetwork
	sethConfig          seth.Config
	contractsRootFolder string
}

func (e *ExistingEVMChainConfigWithSeth) EVMNetwork() blockchain.EVMNetwork {
	return e.evmNetwork
}

func CreateExistingEVMChainWithSeth(evmNetwork blockchain.EVMNetwork, sethConfig seth.Config) (persistent_types.ExistingEVMChainConfig, error) {
	contractsRootFolder, err := findGethWrappersFolderRoot(5)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to find contracts root folder")
	}

	return &ExistingEVMChainConfigWithSeth{
		evmNetwork:          evmNetwork,
		sethConfig:          sethConfig,
		contractsRootFolder: contractsRootFolder,
	}, nil
}

func (e *ExistingEVMChainConfigWithSeth) Chain() (deployment.Chain, error) {
	chain := deployment.Chain{}
	c, err := seth_utils.MergeSethAndEvmNetworkConfigs(e.evmNetwork, e.sethConfig)
	if err != nil {
		return chain, fmt.Errorf("failed to merge seth and evm network configs: %w", err)
	}

	sethClient, err := seth.NewClientBuilderWithConfig(&c).
		// we want to set it dynamically, because the path depends on the location of the file in the project
		WithGethWrappersFolders([]string{fmt.Sprintf("%s/ccip", e.contractsRootFolder)}).
		Build()
	if err != nil {
		return chain, fmt.Errorf("failed to create seth client: %w", err)
	}

	rpcProvider := ctf_test_env.NewRPCProvider(e.evmNetwork.HTTPURLs, e.evmNetwork.URLs, e.evmNetwork.HTTPURLs, e.evmNetwork.URLs)

	chain, err = e.Build(sethClient, e.evmNetwork, rpcProvider)
	if err != nil {
		return chain, err
	}

	return chain, nil
}

type SethChainBuilder struct{}

func (s *SethChainBuilder) Build(sethClient *seth.Client, evmNetwork blockchain.EVMNetwork, rpcProvider ctf_test_env.RpcProvider) (deployment.Chain, error) {
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
