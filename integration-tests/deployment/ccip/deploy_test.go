package ccipdeployment

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"

	ccipconfig "github.com/smartcontractkit/ccip/integration-tests/ccip-tests/testconfig"
	"github.com/smartcontractkit/chainlink-testing-framework/blockchain"
	ctf_config "github.com/smartcontractkit/chainlink-testing-framework/config"
	ctf_config_types "github.com/smartcontractkit/chainlink-testing-framework/config/types"
	"github.com/smartcontractkit/chainlink-testing-framework/seth"
	"github.com/smartcontractkit/chainlink/integration-tests/deployment"
	"github.com/smartcontractkit/chainlink/integration-tests/deployment/memory"
	"github.com/smartcontractkit/chainlink/integration-tests/deployment/persistent"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
)

func TestDeployCapReg_InMemory_Concurrent(t *testing.T) {
	lggr := logger.TestLogger(t)
	e := memory.NewMemoryEnvironment(t, lggr, zapcore.InfoLevel, memory.MemoryEnvironmentConfig{
		Bootstraps: 1,
		Chains:     4,
		Nodes:      4,
	})
	testDeployCapRegWithEnv_Concurrent(t, lggr, e)
}

func TestDeployCapReg_NewDevnet_Concurrent(t *testing.T) {
	lggr := logger.TestLogger(t)

	firstNetworkConfig := ctf_config.MustGetDefaultChainConfig()
	firstNetworkConfig.ChainID = 1337
	secondNetworkConfig := ctf_config.MustGetDefaultChainConfig()
	secondNetworkConfig.ChainID = 2337

	geth := ctf_config_types.ExecutionLayer_Geth
	eth1 := ctf_config_types.EthereumVersion_Eth1

	defaultSethConfig := seth.NewClientBuilder().WithGasPriceEstimations(false, 0, seth.Priority_Standard).BuildConfig()

	envConfig := persistent.EnvironmentConfig{
		ChainConfig: persistent.ChainConfig{
			NewEVMChains: []persistent.NewEVMChainConfig{
				persistent.CreateNewPrivateEVMChainConfig(ctf_config.EthereumNetworkConfig{
					ExecutionLayer:      &geth,
					EthereumVersion:     &eth1,
					EthereumChainConfig: &firstNetworkConfig,
				}, *defaultSethConfig),
				persistent.CreateNewPrivateEVMChainConfig(ctf_config.EthereumNetworkConfig{
					ExecutionLayer:      &geth,
					EthereumVersion:     &eth1,
					EthereumChainConfig: &secondNetworkConfig,
				}, *defaultSethConfig),
			},
		},
	}

	e, err := persistent.NewEnvironment(lggr, envConfig)
	require.NoError(t, err, "Error creating new persistent environment")
	testDeployCapRegWithEnv_Concurrent(t, lggr, *e)
}

func TestDeployCCIPContractsInMemory(t *testing.T) {
	lggr := logger.TestLogger(t)
	e := memory.NewMemoryEnvironment(t, lggr, zapcore.InfoLevel, memory.MemoryEnvironmentConfig{
		Bootstraps: 1,
		Chains:     1,
		Nodes:      4,
	})
	testDeployCCIPContractsWithEnv(t, lggr, e)
}

func TestDeployCCIPContractsNewDevnet(t *testing.T) {
	lggr := logger.TestLogger(t)

	firstNetworkConfig := ctf_config.MustGetDefaultChainConfig()
	firstNetworkConfig.ChainID = 1337
	secondNetworkConfig := ctf_config.MustGetDefaultChainConfig()
	secondNetworkConfig.ChainID = 2337

	geth := ctf_config_types.ExecutionLayer_Geth
	eth1 := ctf_config_types.EthereumVersion_Eth1

	defaultSethConfig := seth.NewClientBuilder().BuildConfig()

	envConfig := persistent.EnvironmentConfig{
		ChainConfig: persistent.ChainConfig{
			NewEVMChains: []persistent.NewEVMChainConfig{
				persistent.CreateNewPrivateEVMChainConfig(ctf_config.EthereumNetworkConfig{
					ExecutionLayer:      &geth,
					EthereumVersion:     &eth1,
					EthereumChainConfig: &firstNetworkConfig,
				}, *defaultSethConfig),
				persistent.CreateNewPrivateEVMChainConfig(ctf_config.EthereumNetworkConfig{
					ExecutionLayer:      &geth,
					EthereumVersion:     &eth1,
					EthereumChainConfig: &secondNetworkConfig,
				}, *defaultSethConfig),
			},
		},
	}

	e, err := persistent.NewEnvironment(lggr, envConfig)
	require.NoError(t, err, "Error creating new persistent environment")
	testDeployCCIPContractsWithEnv(t, lggr, *e)
}

func TestDeployCCIPContractsNewDevnet_FromTestConfig(t *testing.T) {
	lggr := logger.TestLogger(t)
	testCfg := ccipconfig.GlobalTestConfig()
	require.NoError(t, testCfg.Validate(), "Error validating test config")

	// here we are creating Seth config, but we should read it from the test config
	defaultSethConfig := seth.NewClientBuilder().BuildConfig()

	chainCfg, err := persistent.EVMChainConfigFromTestConfig(*testCfg, *defaultSethConfig)
	require.NoError(t, err, "Error creating chain config from test config")

	envConfig := persistent.EnvironmentConfig{
		ChainConfig: chainCfg,
	}

	e, err := persistent.NewEnvironment(lggr, envConfig)
	require.NoError(t, err, "Error creating new persistent environment")
	testDeployCCIPContractsWithEnv(t, lggr, *e)
}

// TODO: update urls before running
func TestDeployCCIPContractsExistingDevnet(t *testing.T) {
	lggr := logger.TestLogger(t)
	defaultSethConfig := seth.NewClientBuilder().BuildConfig()
	envConfig := persistent.EnvironmentConfig{
		ChainConfig: persistent.ChainConfig{
			ExistingEVMChains: []persistent.ExistingEVMChainConfig{
				persistent.CreateExistingEVMChainConfigWithSeth(
					blockchain.EVMNetwork{
						Name:        "SomeChain_1337",
						ChainID:     1337,
						URLs:        []string{"ws://127.0.0.1:57163"},
						HTTPURLs:    []string{"ws://127.0.0.1:57162"},
						PrivateKeys: []string{"ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"}, // default Geth PK
					},
					*defaultSethConfig,
				),
				persistent.CreateExistingEVMChainConfigWithSeth(
					blockchain.EVMNetwork{
						Name:        "SomeChain_2337",
						ChainID:     2337,
						URLs:        []string{"ws://127.0.0.1:57251"},
						HTTPURLs:    []string{"ws://127.0.0.1:57161"},
						PrivateKeys: []string{"ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"}, // default Geth PK
					},
					*defaultSethConfig,
				),
			},
		},
	}
	e, err := persistent.NewEnvironment(lggr, envConfig)
	require.NoError(t, err, "Error creating new persistent environment")
	testDeployCCIPContractsWithEnv(t, lggr, *e)
}

func testDeployCCIPContractsWithEnv(t *testing.T, lggr logger.Logger, e deployment.Environment) {
	var ab deployment.AddressBook
	// Deploy all the CCIP contracts.
	for _, chain := range e.AllChainSelectors() {
		capRegAddresses, _, err := DeployCapReg(lggr, e.Chains, chain)
		require.NoError(t, err)
		s, err := LoadOnchainState(e, capRegAddresses)
		require.NoError(t, err)
		newAb, err := DeployCCIPContracts(e, DeployCCIPContractConfig{
			HomeChainSel:     chain,
			CCIPOnChainState: s,
		})
		require.NoError(t, err)
		if ab == nil {
			ab = newAb
		} else {
			mergeErr := ab.Merge(newAb)
			require.NoError(t, mergeErr)
		}
	}

	state, err := LoadOnchainState(e, ab)
	require.NoError(t, err)
	snap, err := state.Snapshot(e.AllChainSelectors())
	require.NoError(t, err)

	// Assert expect every deployed address to be in the address book.
	// TODO (CCIP-3047): Add the rest of CCIPv2 representation
	b, err := json.MarshalIndent(snap, "", "	")
	require.NoError(t, err)
	fmt.Println(string(b))
}

func testDeployCapRegWithEnv_Concurrent(t *testing.T, lggr logger.Logger, e deployment.Environment) {
	var ab deployment.AddressBook
	// Deploy all the CCIP contracts.
	ab, _, err := DeployCapReg_Concurrent(lggr, e.Chains, e.AllChainSelectors())
	require.NoError(t, err)

	state, err := LoadOnchainState(e, ab)
	require.NoError(t, err)
	snap, err := state.Snapshot(e.AllChainSelectors())
	require.NoError(t, err)

	// Assert expect every deployed address to be in the address book.
	// TODO (CCIP-3047): Add the rest of CCIPv2 representation
	b, err := json.MarshalIndent(snap, "", "	")
	require.NoError(t, err)
	fmt.Println(string(b))
}

func TestJobSpecGeneration(t *testing.T) {
	lggr := logger.TestLogger(t)
	e := memory.NewMemoryEnvironment(t, lggr, zapcore.InfoLevel, memory.MemoryEnvironmentConfig{
		Chains: 1,
		Nodes:  1,
	})
	js, err := NewCCIPJobSpecs(e.NodeIDs, e.Offchain)
	require.NoError(t, err)
	for node, jb := range js {
		fmt.Println(node, jb)
	}
	// TODO: Add job assertions
}
