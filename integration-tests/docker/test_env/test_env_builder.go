package test_env

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/kurtosis-tech/kurtosis/api/golang/core/lib/starlark_run_config"
	"github.com/kurtosis-tech/kurtosis/api/golang/engine/lib/kurtosis_context"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/test-go/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"gopkg.in/yaml.v2"

	"github.com/smartcontractkit/chainlink-testing-framework/blockchain"
	"github.com/smartcontractkit/chainlink-testing-framework/docker/test_env"
	"github.com/smartcontractkit/chainlink-testing-framework/logging"
	"github.com/smartcontractkit/chainlink-testing-framework/logwatch"
	"github.com/smartcontractkit/chainlink-testing-framework/networks"

	"github.com/smartcontractkit/chainlink/v2/core/services/chainlink"

	"github.com/smartcontractkit/chainlink/integration-tests/contracts"
	"github.com/smartcontractkit/chainlink/integration-tests/types/config/node"
)

type CleanUpType string

const (
	CleanUpTypeNone     CleanUpType = "none"
	CleanUpTypeStandard CleanUpType = "standard"
	CleanUpTypeCustom   CleanUpType = "custom"
)

type CLTestEnvBuilder struct {
	hasLogWatch         bool
	hasGeth             bool
	hasKillgrave        bool
	hasForwarders       bool
	clNodeConfig        *chainlink.Config
	secretsConfig       string
	nonDevGethNetworks  []blockchain.EVMNetwork
	kurtosisConfigFiles []string
	clNodesCount        int
	customNodeCsaKeys   []string
	defaultNodeCsaKeys  []string
	l                   zerolog.Logger
	t                   *testing.T
	te                  *CLClusterTestEnv
	isNonEVM            bool

	cleanUpType     CleanUpType
	cleanUpCustomFn func()
	/* funding */
	ETHFunds *big.Float
}

func NewCLTestEnvBuilder() *CLTestEnvBuilder {
	return &CLTestEnvBuilder{
		l: log.Logger,
	}
}

// WithTestEnv sets the test environment to use for the test.
// If nil, a new test environment is created.
// If not nil, the test environment is used as-is.
// If TEST_ENV_CONFIG_PATH is set, the test environment is created with the config at that path.
func (b *CLTestEnvBuilder) WithTestEnv(te *CLClusterTestEnv) (*CLTestEnvBuilder, error) {
	envConfigPath, isSet := os.LookupEnv("TEST_ENV_CONFIG_PATH")
	var cfg *TestEnvConfig
	var err error
	if isSet {
		cfg, err = NewTestEnvConfigFromFile(envConfigPath)
		if err != nil {
			return nil, err
		}
	}

	if te != nil {
		b.te = te
	} else {
		b.te, err = NewTestEnv()
		if err != nil {
			return nil, err
		}
	}

	if cfg != nil {
		b.te = b.te.WithTestEnvConfig(cfg)
	}
	return b, nil
}

// WithTestLogger sets the test logger to use for the test.
// Useful for parallel tests so the logging will be separated correctly in the results views.
func (b *CLTestEnvBuilder) WithTestLogger(t *testing.T) *CLTestEnvBuilder {
	b.t = t
	b.l = logging.GetTestLogger(t)
	return b
}

func (b *CLTestEnvBuilder) WithLogWatcher() *CLTestEnvBuilder {
	b.hasLogWatch = true
	return b
}

func (b *CLTestEnvBuilder) WithCLNodes(clNodesCount int) *CLTestEnvBuilder {
	b.clNodesCount = clNodesCount
	return b
}

func (b *CLTestEnvBuilder) WithForwarders() *CLTestEnvBuilder {
	b.hasForwarders = true
	return b
}

func (b *CLTestEnvBuilder) WithFunding(eth *big.Float) *CLTestEnvBuilder {
	b.ETHFunds = eth
	return b
}

func (b *CLTestEnvBuilder) WithGeth() *CLTestEnvBuilder {
	b.hasGeth = true
	return b
}

func (b *CLTestEnvBuilder) WithPrivateGethChains(evmNetworks []blockchain.EVMNetwork) *CLTestEnvBuilder {
	b.nonDevGethNetworks = evmNetworks
	return b
}

func (b *CLTestEnvBuilder) WithKurtosis(kurtosisConfigFiles []string) *CLTestEnvBuilder {
	b.kurtosisConfigFiles = kurtosisConfigFiles
	return b
}

func (b *CLTestEnvBuilder) WithCLNodeConfig(cfg *chainlink.Config) *CLTestEnvBuilder {
	b.clNodeConfig = cfg
	return b
}

func (b *CLTestEnvBuilder) WithSecretsConfig(secrets string) *CLTestEnvBuilder {
	b.secretsConfig = secrets
	return b
}

func (b *CLTestEnvBuilder) WithMockAdapter() *CLTestEnvBuilder {
	b.hasKillgrave = true
	return b
}

// WithNonEVM sets the test environment to not use EVM when built.
func (b *CLTestEnvBuilder) WithNonEVM() *CLTestEnvBuilder {
	b.isNonEVM = true
	return b
}

func (b *CLTestEnvBuilder) WithStandardCleanup() *CLTestEnvBuilder {
	b.cleanUpType = CleanUpTypeStandard
	return b
}

func (b *CLTestEnvBuilder) WithoutCleanup() *CLTestEnvBuilder {
	b.cleanUpType = CleanUpTypeNone
	return b
}

func (b *CLTestEnvBuilder) WithCustomCleanup(customFn func()) *CLTestEnvBuilder {
	b.cleanUpType = CleanUpTypeCustom
	b.cleanUpCustomFn = customFn
	return b
}

func (b *CLTestEnvBuilder) Build() (*CLClusterTestEnv, error) {
	if b.te == nil {
		var err error
		b, err = b.WithTestEnv(nil)
		if err != nil {
			return nil, err
		}
	}
	b.l.Info().
		Bool("hasGeth", b.hasGeth).
		Bool("hasKillgrave", b.hasKillgrave).
		Int("clNodesCount", b.clNodesCount).
		Strs("customNodeCsaKeys", b.customNodeCsaKeys).
		Strs("defaultNodeCsaKeys", b.defaultNodeCsaKeys).
		Msg("Building CL cluster test environment..")

	var err error
	if b.t != nil {
		b.te.WithTestLogger(b.t)
	}

	if b.hasLogWatch {
		b.te.LogWatch, err = logwatch.NewLogWatch(nil, nil)
		if err != nil {
			return nil, err
		}
	}

	if b.hasKillgrave {
		err = b.te.StartMockAdapter()
		if err != nil {
			return nil, err
		}
	}

	switch b.cleanUpType {
	case CleanUpTypeStandard:
		b.t.Cleanup(func() {
			if err := b.te.Cleanup(); err != nil {
				b.l.Error().Err(err).Msg("Error cleaning up test environment")
			}
		})
	case CleanUpTypeCustom:
		b.t.Cleanup(b.cleanUpCustomFn)
	case CleanUpTypeNone:
		b.l.Warn().Msg("test environment won't be cleaned up")
	case "":
		return b.te, errors.WithMessage(errors.New("explicit cleanup type must be set when building test environment"), "test environment builder failed")
	}

	if b.nonDevGethNetworks != nil {
		b.te.WithPrivateChain(b.nonDevGethNetworks)
		err := b.te.StartPrivateChain()
		if err != nil {
			return b.te, err
		}
		var nonDevNetworks []blockchain.EVMNetwork
		for i, n := range b.te.PrivateChain {
			primaryNode := n.GetPrimaryNode()
			if primaryNode == nil {
				return b.te, errors.WithStack(fmt.Errorf("primary node is nil in PrivateChain interface"))
			}
			nonDevNetworks = append(nonDevNetworks, *n.GetNetworkConfig())
			nonDevNetworks[i].URLs = []string{primaryNode.GetInternalWsUrl()}
			nonDevNetworks[i].HTTPURLs = []string{primaryNode.GetInternalHttpUrl()}
		}
		for _, n := range b.nonDevGethNetworks {
			nonDevNetworks = append(nonDevNetworks, n)
		}
		if nonDevNetworks == nil {
			return nil, errors.New("cannot create nodes with custom config without nonDevNetworks")
		}

		err = b.te.StartClCluster(b.clNodeConfig, b.clNodesCount, b.secretsConfig)
		if err != nil {
			return nil, err
		}
		return b.te, nil
	}

	networkConfig := networks.SelectedNetwork
	var internalDockerUrls test_env.InternalDockerUrls
	if b.hasGeth && networkConfig.Simulated {
		networkConfig, internalDockerUrls, err = b.te.StartGeth()
		if err != nil {
			return nil, err
		}
	}

	if len(b.kurtosisConfigFiles) > 0 {
		privateNetworks := make([]blockchain.EVMNetwork, 0)

		type kurtosisNetwork struct {
			NetworkParams struct {
				NetworkId string `yaml:"network_id"`
			} `yaml:"network_params"`
		}

		var (
			rpcKey = "rpc"
			wsKey  string
			wsKeys = []string{"ws-rpc", "ws"}
		)

		readKurtosisNetworkConfig := func(filePath string) (kurtosisNetwork, error) {
			var network kurtosisNetwork

			// Open YAML file
			file, err := os.Open(filePath)
			if err != nil {
				return network, err
			}
			defer file.Close()

			if file != nil {
				decoder := yaml.NewDecoder(file)
				if err := decoder.Decode(&network); err != nil {
					return network, err
				}
			}

			return network, nil
		}

		validateSerializedArgs := func(serializedArgs string) error {
			var result interface{}
			var jsonError error
			if jsonError = json.Unmarshal([]byte(serializedArgs), &result); jsonError == nil {
				return nil
			}
			var yamlError error
			if yamlError = yaml.Unmarshal([]byte(serializedArgs), &result); yamlError == nil {
				return nil
			}
			return fmt.Errorf("invalid serialized args: %s", jsonError.Error()+yamlError.Error())
		}

		kurtosisCtx, err := kurtosis_context.NewKurtosisContextFromLocalEngine()
		if err != nil {
			return nil, err
		}
		ctx := context.Background()

		for _, file := range b.kurtosisConfigFiles {
			network, err := readKurtosisNetworkConfig(file)
			if err != nil {
				return nil, err
			}

			enclaveName := strings.Replace(file, ".yaml", "", 1)
			enclaveCtx, err := kurtosisCtx.CreateEnclave(ctx, enclaveName)
			if err != nil {
				return nil, err
			}

			b.t.Cleanup(func() {
				err := kurtosisCtx.DestroyEnclave(ctx, enclaveName)
				require.NoError(b.t, err, fmt.Sprintf("Error destroying enclave %s", enclaveName))
			})

			packageArgsFileBytes, err := os.ReadFile(file)
			if err != nil {
				return nil, err
			}

			packageArgs := string(packageArgsFileBytes)

			if err := validateSerializedArgs(packageArgs); err != nil {
				return nil, err
			}

			starlarkRunConfig := starlark_run_config.NewRunStarlarkConfig(starlark_run_config.WithSerializedParams(packageArgs))
			starlarkRunResult, err := enclaveCtx.RunStarlarkPackageBlocking(ctx, "/Users/btofel/Desktop/repos/ethereum-package/", starlarkRunConfig)
			if err != nil {
				return nil, err
			}
			fmt.Print(starlarkRunResult.RunOutput)

			services, err := enclaveCtx.GetExistingAndHistoricalServiceIdentifiers(ctx)
			if err != nil {
				return nil, err
			}
			rpcs := make(map[string]string)

			// hardcoded for now, we only want to get non-bootstrap EL nodes endpoints
			// in the future depending on the service count we should either get the first one
			// if count == 1 or all the other, which have the index > 1
			for _, serviceName := range services.GetOrderedListOfNames() {
				if strings.Contains(serviceName, "el-2") {
					serviceCtx, err := enclaveCtx.GetServiceContext(serviceName)
					if err != nil {
						return nil, err
					}

					// publicPorts := serviceCtx.GetPublicPorts()
					// fmt.Printf("public ports: %v\n", publicPorts)
					privatePorts := serviceCtx.GetPrivatePorts()
					fmt.Printf("private ports: %v\n", privatePorts)

					for _, key := range wsKeys {
						if _, ok := privatePorts[key]; ok {
							wsKey = key
							break
						}
					}

					if wsKey == "" {
						return nil, errors.New(fmt.Sprintf("failed to find any key representing ws private port in %v", privatePorts))
					}

					if v, ok := privatePorts[wsKey]; ok {
						rpcs[wsKey] = fmt.Sprintf("ws://%s-%s:%s", serviceName, string(serviceCtx.GetServiceUUID()), v.String())
					} else {
						return nil, errors.New("no public port for ws")
					}

					// if there's no rpc/http port specified it means it is the same as ws port
					if v, ok := privatePorts[rpcKey]; ok {
						rpcs[rpcKey] = fmt.Sprintf("http://%s-%s:%s", serviceName, string(serviceCtx.GetServiceUUID()), v.String())
					} else {
						if v, ok := privatePorts[wsKey]; ok {
							rpcs[wsKey] = fmt.Sprintf("http://%s-%s:%s", serviceName, string(serviceCtx.GetServiceUUID()), v.String())
						} else {
							return nil, errors.New("no public port for ws")
						}
					}
				}
			}

			chainId, err := strconv.Atoi(network.NetworkParams.NetworkId)
			if err != nil {
				return nil, err
			}

			privateNetworks = append(privateNetworks, blockchain.EVMNetwork{
				Name:            enclaveName,
				ChainID:         int64(chainId),
				URLs:            []string{rpcs[wsKey]},
				HTTPURLs:        []string{rpcs[rpcKey]},
				FinalityTag:     true,
				Simulated:       false,
				SupportsEIP1559: true,
				Timeout: blockchain.JSONStrDuration{
					Duration: 4 * time.Minute,
				},
				SimulationType: "kurtosis",
				PrivateKeys:    []string{"bcdf20249abf0ed6d944c0288fad489e33f66b3960d9e6229c1cd214ed3bbe31"},
			})

			b.te.WithPrivateChain(privateNetworks)

			//TODO small hack to just set the name
			b.te.Network = &testcontainers.DockerNetwork{
				Name: fmt.Sprintf("kt-%s", enclaveName),
			}
		}

		return b.te, nil
	}

	if !b.isNonEVM {
		bc, err := blockchain.NewEVMClientFromNetwork(networkConfig, b.l)
		if err != nil {
			return nil, err
		}

		b.te.EVMClient = bc
		cd, err := contracts.NewContractDeployer(bc, b.l)
		if err != nil {
			return nil, err
		}
		b.te.ContractDeployer = cd

		cl, err := contracts.NewContractLoader(bc, b.l)
		if err != nil {
			return nil, err
		}
		b.te.ContractLoader = cl
	}

	var nodeCsaKeys []string

	// Start Chainlink Nodes
	if b.clNodesCount > 0 {
		var cfg *chainlink.Config
		if b.clNodeConfig != nil {
			cfg = b.clNodeConfig
		} else {
			cfg = node.NewConfig(node.NewBaseConfig(),
				node.WithOCR1(),
				node.WithP2Pv1(),
			)
		}

		if !b.isNonEVM {
			var httpUrls []string
			var wsUrls []string
			if networkConfig.Simulated {
				httpUrls = []string{internalDockerUrls.HttpUrl}
				wsUrls = []string{internalDockerUrls.WsUrl}
			} else {
				httpUrls = networkConfig.HTTPURLs
				wsUrls = networkConfig.URLs
			}

			node.SetChainConfig(cfg, wsUrls, httpUrls, networkConfig, b.hasForwarders)
		}

		err := b.te.StartClCluster(cfg, b.clNodesCount, b.secretsConfig)
		if err != nil {
			return nil, err
		}

		nodeCsaKeys, err = b.te.ClCluster.NodeCSAKeys()
		if err != nil {
			return nil, err
		}
		b.defaultNodeCsaKeys = nodeCsaKeys
	}

	if b.hasGeth && b.clNodesCount > 0 && b.ETHFunds != nil {
		b.te.ParallelTransactions(true)
		defer b.te.ParallelTransactions(false)
		if err := b.te.FundChainlinkNodes(b.ETHFunds); err != nil {
			return nil, err
		}
	}

	return b.te, nil
}
