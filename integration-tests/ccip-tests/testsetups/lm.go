package testsetups

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/lib/pq"
	"github.com/smartcontractkit/ccip/integration-tests/client"
	"gopkg.in/guregu/null.v4"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/liquiditymanager/generated/liquiditymanager"
	"github.com/smartcontractkit/chainlink/v2/core/services/job"
	"github.com/smartcontractkit/chainlink/v2/core/store/models"

	"github.com/AlekSi/pointer"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	integrationactions "github.com/smartcontractkit/ccip/integration-tests/actions"
	chainselectors "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"
	"golang.org/x/sync/errgroup"

	"github.com/smartcontractkit/chainlink-testing-framework/blockchain"
	ctfClient "github.com/smartcontractkit/chainlink-testing-framework/client"
	"github.com/smartcontractkit/chainlink-testing-framework/k8s/config"
	"github.com/smartcontractkit/chainlink-testing-framework/k8s/environment"
	"github.com/smartcontractkit/chainlink/integration-tests/ccip-tests/actions"
	"github.com/smartcontractkit/chainlink/integration-tests/ccip-tests/contracts"
	"github.com/smartcontractkit/chainlink/integration-tests/docker/test_env"
)

type LMTestSetupOutputs struct {
	CCIPTestSetUpOutputs
	LMModules map[int64]*actions.LMCommon
}

func (o *LMTestSetupOutputs) CreateLMEnvironment(
	lggr zerolog.Logger,
	envName string,
	reportPath string,
) map[int64]blockchain.EVMClient {
	t := o.Cfg.Test
	testConfig := o.Cfg
	var (
		ccipEnv  *actions.CCIPTestEnv
		k8Env    *environment.Environment
		err      error
		chains   []blockchain.EVMClient
		local    *test_env.CLClusterTestEnv
		deployCL func() error
	)

	envConfig := createEnvironmentConfig(t, envName, testConfig, reportPath)

	configureCLNode := !testConfig.useExistingDeployment() || pointer.GetString(testConfig.EnvInput.EnvToConnect) != ""
	namespace := ""
	if testConfig.TestGroupInput.LoadProfile != nil {
		namespace = testConfig.TestGroupInput.LoadProfile.TestRunName
	}
	require.False(t, testConfig.localCluster() && testConfig.ExistingCLCluster(),
		"local cluster and existing cluster cannot be true at the same time")
	// if it's a new deployment, deploy the env
	// Or if EnvToConnect is given connect to that k8 environment
	if configureCLNode {
		if !testConfig.ExistingCLCluster() {
			// if it's a local cluster, deploy the local cluster in docker
			if testConfig.localCluster() {
				local, deployCL = DeployLocalCluster(t, testConfig)
				ccipEnv = &actions.CCIPTestEnv{
					LocalCluster: local,
				}
				namespace = "local-docker-deployment"
			} else {
				// Otherwise, deploy the k8s env
				lggr.Info().Msg("Deploying test environment")
				// deploy the env if configureCLNode is true
				k8Env = DeployEnvironments(t, envConfig, testConfig)
				ccipEnv = &actions.CCIPTestEnv{K8Env: k8Env}
				namespace = ccipEnv.K8Env.Cfg.Namespace
			}
		} else {
			// if there is already a cluster, use the existing cluster to connect to the nodes
			ccipEnv = &actions.CCIPTestEnv{}
			mockserverURL := pointer.GetString(testConfig.EnvInput.Mockserver)
			require.NotEmpty(t, mockserverURL, "mockserver URL cannot be nil")
			ccipEnv.MockServer = ctfClient.NewMockserverClient(&ctfClient.MockserverConfig{
				LocalURL:   mockserverURL,
				ClusterURL: mockserverURL,
			})
		}
		ccipEnv.CLNodeWithKeyReady, _ = errgroup.WithContext(o.SetUpContext)
		o.Env = ccipEnv
		if ccipEnv.K8Env != nil && ccipEnv.K8Env.WillUseRemoteRunner() {
			return nil
		}
	} else {
		// if configureCLNode is false it means we don't need to deploy any additional pods,
		// use a placeholder env to create just the remote runner in it.
		if value, set := os.LookupEnv(config.EnvVarJobImage); set && value != "" {
			k8Env = environment.New(envConfig)
			err = k8Env.Run()
			require.NoErrorf(t, err, "error creating environment remote runner")
			o.Env = &actions.CCIPTestEnv{K8Env: k8Env}
			if k8Env.WillUseRemoteRunner() {
				return nil
			}
		}
	}
	chainByChainID := make(map[int64]blockchain.EVMClient)
	if pointer.GetBool(testConfig.TestGroupInput.LocalCluster) {
		require.NotNil(t, ccipEnv.LocalCluster, "Local cluster shouldn't be nil")
		for _, n := range ccipEnv.LocalCluster.EVMNetworks {
			if evmClient, err := blockchain.NewEVMClientFromNetwork(*n, lggr); err == nil {
				chainByChainID[evmClient.GetChainID().Int64()] = evmClient
				chains = append(chains, evmClient)
			} else {
				lggr.Error().Err(err).Msgf("EVMClient for chainID %d not found", n.ChainID)
			}
		}
	} else {
		for _, n := range testConfig.SelectedNetworks {
			if _, ok := chainByChainID[n.ChainID]; ok {
				continue
			}
			var ec blockchain.EVMClient
			if k8Env == nil {
				ec, err = blockchain.ConnectEVMClient(n, lggr)
			} else {
				log.Info().Interface("urls", k8Env.URLs).Msg("URLs")
				ec, err = blockchain.NewEVMClient(n, k8Env, lggr)
			}
			require.NoError(t, err, "Connecting to blockchain nodes shouldn't fail")
			chains = append(chains, ec)
			chainByChainID[n.ChainID] = ec
		}
	}
	if configureCLNode {
		ccipEnv.CLNodeWithKeyReady.Go(func() error {
			var totalNodes int
			if !o.Cfg.ExistingCLCluster() {
				if ccipEnv.LocalCluster != nil {
					err = deployCL()
					if err != nil {
						return err
					}
				}
				err = ccipEnv.ConnectToDeployedNodes()
				if err != nil {
					return fmt.Errorf("error connecting to chainlink nodes: %w", err)
				}
				totalNodes = pointer.GetInt(testConfig.EnvInput.NewCLCluster.NoOfNodes)
			} else {
				totalNodes = pointer.GetInt(testConfig.EnvInput.ExistingCLCluster.NoOfNodes)
				err = ccipEnv.ConnectToExistingNodes(o.Cfg.EnvInput)
				if err != nil {
					return fmt.Errorf("error deploying and connecting to chainlink nodes: %w", err)
				}
			}
			err = ccipEnv.SetUpNodeKeysAndFund(lggr, big.NewFloat(testConfig.TestGroupInput.NodeFunding), chains)
			if err != nil {
				return fmt.Errorf("error setting up nodes and keys %w", err)
			}
			// first node is the bootstrapper
			ccipEnv.CommitNodeStartIndex = 1
			ccipEnv.ExecNodeStartIndex = 1
			ccipEnv.NumOfCommitNodes = testConfig.TestGroupInput.NoOfCommitNodes
			ccipEnv.NumOfExecNodes = ccipEnv.NumOfCommitNodes
			if !pointer.GetBool(testConfig.TestGroupInput.CommitAndExecuteOnSameDON) {
				if len(ccipEnv.CLNodesWithKeys) < 11 {
					return fmt.Errorf("not enough CL nodes for separate commit and execution nodes")
				}
				if testConfig.TestGroupInput.NoOfCommitNodes >= totalNodes {
					return fmt.Errorf("number of commit nodes can not be greater than total number of nodes in DON")
				}
				// first two nodes are reserved for bootstrap commit and bootstrap exec
				ccipEnv.CommitNodeStartIndex = 2
				ccipEnv.ExecNodeStartIndex = 2 + testConfig.TestGroupInput.NoOfCommitNodes
				ccipEnv.NumOfExecNodes = totalNodes - (2 + testConfig.TestGroupInput.NoOfCommitNodes)
				if ccipEnv.NumOfExecNodes < 4 {
					return fmt.Errorf("insufficient number of exec nodes")
				}
			}
			ccipEnv.NumOfAllowedFaultyExec = (ccipEnv.NumOfExecNodes - 1) / 3
			ccipEnv.NumOfAllowedFaultyCommit = (ccipEnv.NumOfCommitNodes - 1) / 3
			return nil
		})
	}

	t.Cleanup(func() {
		if configureCLNode {
			if ccipEnv.LocalCluster != nil {
				err := ccipEnv.LocalCluster.Terminate()
				require.NoError(t, err, "Local cluster termination shouldn't fail")
				//require.NoError(t, o.Reporter.SendReport(t, namespace, false), "Aggregating and sending report shouldn't fail")
				return
			}
			if pointer.GetBool(testConfig.TestGroupInput.KeepEnvAlive) || testConfig.ExistingCLCluster() {
				//require.NoError(t, o.Reporter.SendReport(t, namespace, true), "Aggregating and sending report shouldn't fail")
				return
			}
			lggr.Info().Msg("Tearing down the environment")
			err = integrationactions.TeardownSuite(t, ccipEnv.K8Env, ccipEnv.CLNodes, o.Reporter,
				zapcore.ErrorLevel, o.Cfg.EnvInput, chains...)
			require.NoError(t, err, "Environment teardown shouldn't fail")
		} else {
			//just send the report
			require.NoError(t, o.Reporter.SendReport(t, namespace, true), "Aggregating and sending report shouldn't fail")
		}
	})
	return chainByChainID
}

func (o *LMTestSetupOutputs) DeployLMChainContracts(
	lggr zerolog.Logger,
	chainClient blockchain.EVMClient,
	networkCfg blockchain.EVMNetwork,
	chainSelectors []uint64,
	lmCommon actions.LMCommon,
) error {
	var k8Env *environment.Environment
	ccipEnv := o.Env
	if ccipEnv != nil {
		k8Env = ccipEnv.K8Env
	}
	if k8Env != nil && chainClient.NetworkSimulated() {
		networkCfg.URLs = k8Env.URLs[chainClient.GetNetworkConfig().Name]
	}

	chain, err := blockchain.ConcurrentEVMClient(networkCfg, k8Env, chainClient, lggr)
	if err != nil {
		return errors.WithStack(fmt.Errorf("failed to create chain client for %s: %w", networkCfg.Name, err))
	}

	chain.ParallelTransactions(true)
	//defer chain.Close()

	cd, err := contracts.NewCCIPContractsDeployer(lggr, chain)
	if err != nil {
		return errors.WithStack(fmt.Errorf("failed to create contract deployer: %w", err))
	}

	// Deploy Mock ARM contract
	lggr.Info().Msg("Deploying Mock ARM contract")
	mockARMContract, err := cd.DeployMockARMContract()
	if err != nil {
		return errors.WithStack(fmt.Errorf("failed to deploy Mock ARM contract: %w", err))
	}
	lggr.Info().Str("Address", mockARMContract.String()).Msg("Deployed Mock ARM contract")
	lmCommon.MockArm = mockARMContract

	// Deploy ARM Proxy contract
	lggr.Info().Msg("Deploying ARM Proxy contract")
	armProxyContract, err := cd.DeployArmProxy(*mockARMContract)
	if err != nil {
		return errors.WithStack(fmt.Errorf("failed to deploy ARM Proxy contract: %w", err))
	}
	lggr.Info().Str("Address", armProxyContract.EthAddress.String()).Msg("Deployed ARM Proxy contract")
	lmCommon.ArmProxy = armProxyContract

	// Deploy Wrapped Native contract
	lggr.Info().Msg("Deploying Wrapped Native contract")
	wrapperNative, err := cd.DeployWrappedNative()
	if err != nil {
		return errors.WithStack(fmt.Errorf("failed to deploy Wrapped Native contract: %w", err))
	}
	lggr.Info().Str("Address", wrapperNative.String()).Msg("Deployed Wrapped Native contract")
	lmCommon.WrapperNative = wrapperNative

	// Deploy CCIP Router contract
	lggr.Info().Msg("Deploying CCIP Router contract")
	ccipRouterContract, err := cd.DeployRouter(common.Address{}, *lmCommon.ArmProxy.EthAddress)
	if err != nil {
		return errors.WithStack(fmt.Errorf("failed to deploy CCIP Router contract: %w", err))
	}
	lggr.Info().Str("Address", ccipRouterContract.EthAddress.String()).Msg("Deployed CCIP Router contract")
	lmCommon.CcipRouter = ccipRouterContract

	// Deploy Lock Release Token contract
	lggr.Info().Msg("Deploying Lock Release Token contract")
	lockReleaseTokenPool, err := cd.DeployLockReleaseTokenPoolContract(lmCommon.WrapperNative.String(), *lmCommon.MockArm, lmCommon.CcipRouter.EthAddress)
	if err != nil {
		return errors.WithStack(fmt.Errorf("failed to deploy Lock Release Token contract: %w", err))
	}
	lggr.Info().Str("Address", lockReleaseTokenPool.EthAddress.String()).Msg("Deployed Lock Release Token contract")
	lmCommon.TokenPool = lockReleaseTokenPool

	// Deploy Liquidity Manager contract
	lggr.Info().Msg("Deploying Liquidity Manager contract")
	liquidityManager, err := cd.DeployLiquidityManager(*lmCommon.WrapperNative, chainSelectors[0], lmCommon.TokenPool.EthAddress, lmCommon.MinimumLiquidity)
	if err != nil {
		return errors.WithStack(fmt.Errorf("failed to deploy Liquidity Manager contract: %w", err))
	}
	lggr.Info().Str("Address", liquidityManager.EthAddress.String()).Msg("Deployed Liquidity Manager contract")
	lmCommon.LM = liquidityManager

	// Set Liquidity Manager on Token Pool
	lggr.Info().Msg("Setting Liquidity Manager on Token Pool")
	err = lockReleaseTokenPool.SetRebalancer(*liquidityManager.EthAddress)
	if err != nil {
		return errors.WithStack(fmt.Errorf("failed to set Liquidity Manager on Token Pool: %w", err))
	}
	lggr.Info().Msg("Set Liquidity Manager on Token Pool")

	err = chain.WaitForEvents()
	if err != nil {
		return errors.WithStack(fmt.Errorf("failed to wait for events: %w", err))
	}

	// Verify on chain rebalancer from token pool matches deployed Liquidity Manager
	onchainRebalancer, err := lockReleaseTokenPool.GetRebalancer()
	if err != nil {
		return errors.WithStack(fmt.Errorf("failed to get rebalancer from Token Pool: %w", err))
	}
	if onchainRebalancer != *liquidityManager.EthAddress {
		return errors.WithStack(fmt.Errorf("onchainRebalancer doesn not match the deployed Liquidity Manager"))
	}

	// Deploy Bridge Adapter contracts
	if lmCommon.IsL2 {
		lggr.Info().Msg("Deploying Mock L2 Bridge Adapter contract")
		l2bridgeAdapter, err := cd.DeployMockL2BridgeAdapter()
		if err != nil {
			return errors.WithStack(fmt.Errorf("failed to deploy Mock L2 Bridge Adapter contract: %w", err))
		}
		lggr.Info().Str("Address", l2bridgeAdapter.EthAddress.String()).Msg("Deployed Mock L2 Bridge Adapter contract")
		lmCommon.BridgeAdapterAddr = l2bridgeAdapter.EthAddress
	} else {
		lggr.Info().Msg("Deploying Mock L1 Bridge Adapter contract")
		l1bridgeAdapter, err := cd.DeployMockL1BridgeAdapter(*lmCommon.WrapperNative, true)
		if err != nil {
			return errors.WithStack(fmt.Errorf("failed to deploy Mock L1 Bridge Adapter contract: %w", err))
		}
		lggr.Info().Str("Address", l1bridgeAdapter.EthAddress.String()).Msg("Deployed Mock L1 Bridge Adapter contract")
		lmCommon.BridgeAdapterAddr = l1bridgeAdapter.EthAddress
	}

	lggr.Debug().Interface("lmCommon", lmCommon).Msg("lmCommon")
	o.LMModules[chainClient.GetChainID().Int64()] = &lmCommon

	return nil
}

func LMDefaultTestSetup(
	t *testing.T,
	lggr zerolog.Logger,
	envName string,
	testConfig *CCIPTestConfig,
) *LMTestSetupOutputs {
	var (
		err error
	)
	reportPath := "tmp_laneconfig"
	parent, cancel := context.WithCancel(context.Background())
	defer cancel()
	lmModules := make(map[int64]*actions.LMCommon)
	setUpArgs := &LMTestSetupOutputs{
		CCIPTestSetUpOutputs{
			SetUpContext: parent,
			Cfg:          testConfig,
		},
		lmModules,
	}

	chainByChainID := setUpArgs.CreateLMEnvironment(lggr, envName, reportPath)

	chainAddGrp, _ := errgroup.WithContext(setUpArgs.SetUpContext)
	lggr.Info().Msg("Deploying common contracts")
	chainSelectors := make(map[int64]uint64)

	testConfig.SelectedNetworks, _, err = testConfig.EnvInput.EVMNetworks()
	require.NoError(t, err)

	testConfig.AllNetworks = make(map[string]blockchain.EVMNetwork)
	for _, net := range testConfig.SelectedNetworks {
		testConfig.AllNetworks[net.Name] = net
		if _, exists := chainSelectors[net.ChainID]; !exists {
			chainSelectors[net.ChainID], err = chainselectors.SelectorFromChainId(uint64(net.ChainID))
			require.NoError(t, err)
		}
	}

	i := 0
	for _, net := range testConfig.AllNetworks {
		chain := chainByChainID[net.ChainID]
		net := net
		net.HTTPURLs = chain.GetNetworkConfig().HTTPURLs
		net.URLs = chain.GetNetworkConfig().URLs
		var selectors []uint64
		for chainId, selector := range chainSelectors {
			if chainId != net.ChainID {
				selectors = append(selectors, selector)
			}
		}
		lmCommon, err := actions.DefaultLMModule(
			chain,
			big.NewInt(0),
			i == 1,
			selectors[0],
		)
		require.NoError(t, err)
		chainAddGrp.Go(func() error {
			return setUpArgs.DeployLMChainContracts(lggr, chain, net, selectors, *lmCommon)
		})
		i++
	}
	require.NoError(t, chainAddGrp.Wait(), "Deploying common contracts shouldn't fail")

	lggr.Debug().Interface("lmModules", lmModules).Msg("lmModules")

	l1ChainId := testConfig.SelectedNetworks[0].ChainID
	l2ChainId := testConfig.SelectedNetworks[1].ChainID

	//Set Cross Chain Rebalancer on L1 Rebalancer
	err = lmModules[l1ChainId].LM.SetCrossChainRebalancer(
		liquiditymanager.ILiquidityManagerCrossChainRebalancerArgs{
			RemoteRebalancer:    *lmModules[l2ChainId].LM.EthAddress,
			LocalBridge:         *lmModules[l1ChainId].BridgeAdapterAddr,
			RemoteToken:         *lmModules[l2ChainId].WrapperNative,
			RemoteChainSelector: lmModules[l2ChainId].ChainSelectror,
			Enabled:             true,
		})
	require.NoError(t, err, "Setting Cross Chain Rebalancer on L1 Rebalancer shouldn't fail")

	//Set Cross Chain Rebalancer on L2 Rebalancer
	err = lmModules[l2ChainId].LM.SetCrossChainRebalancer(
		liquiditymanager.ILiquidityManagerCrossChainRebalancerArgs{
			RemoteRebalancer:    *lmModules[l1ChainId].LM.EthAddress,
			LocalBridge:         *lmModules[l2ChainId].BridgeAdapterAddr,
			RemoteToken:         *lmModules[l1ChainId].WrapperNative,
			RemoteChainSelector: lmModules[l1ChainId].ChainSelectror,
			Enabled:             true,
		})
	require.NoError(t, err, "Setting Cross Chain Rebalancer on L1 Rebalancer shouldn't fail")

	err = lmModules[l1ChainId].ChainClient.WaitForEvents()
	require.NoError(t, err, "Waiting for events to confirm on L1 chain shouldn't fail")

	err = lmModules[l2ChainId].ChainClient.WaitForEvents()
	require.NoError(t, err, "Waiting for events to confirm on L2 chain shouldn't fail")

	onchainRebalancerL1, err := lmModules[l1ChainId].TokenPool.GetRebalancer()
	require.NoError(t, err, "Getting rebalancer from Token Pool shouldn't fail")

	onchainRebalancerL2, err := lmModules[l2ChainId].TokenPool.GetRebalancer()
	require.NoError(t, err, "Getting rebalancer from Token Pool shouldn't fail")

	if onchainRebalancerL1.String() != lmModules[l2ChainId].LM.EthAddress.String() ||
		onchainRebalancerL2.String() != lmModules[l1ChainId].LM.EthAddress.String() {
		lggr.Debug().
			Str("onchainRebalancerL1", onchainRebalancerL1.String()).
			Str("onchainRebalancerL2", onchainRebalancerL2.String()).
			Str("L2 LM", lmModules[l2ChainId].LM.EthAddress.String()).
			Str("L1 LM", lmModules[l1ChainId].LM.EthAddress.String()).
			Msg("Onchain rebalancer mismatch")
		t.Logf("Onchain rebalancer mismatch")
	}

	err = setUpArgs.Env.CLNodeWithKeyReady.Wait()
	require.NoError(t, err, "Waiting for CL nodes to be ready shouldn't fail")

	clNodesWithKeys := setUpArgs.Env.CLNodesWithKeys[strconv.FormatInt(l1ChainId, 10)]

	bootstrapNode := clNodesWithKeys[0]

	bootstrapSpec := &client.OCR2TaskJobSpec{
		Name:    "ocr2 bootstrap node " + lmModules[l1ChainId].LM.EthAddress.String(),
		JobType: "bootstrap",
		OCR2OracleSpec: job.OCR2OracleSpec{
			ContractID: lmModules[l1ChainId].LM.EthAddress.String(),
			Relay:      "evm",
			RelayConfig: map[string]interface{}{
				"chainID": int(l1ChainId),
			},
			ContractConfigTrackerPollInterval: *models.NewInterval(time.Second * 15),
		},
	}

	lggr.Info().Msg("Adding bootstrap job")
	j, err := bootstrapNode.Node.MustCreateJob(bootstrapSpec)
	require.NoError(t, err, "Adding bootstrap job shouldn't fail")
	lggr.Info().Str("jobId", j.Data.ID).Msg("Bootstrap job added")

	P2Pv2Bootstrapper := fmt.Sprintf("%s@%s:%d", bootstrapNode.KeysBundle.P2PKeys.Data[0].Attributes.PeerID, bootstrapNode.Node.InternalIP(), 6690)

	donNodes := clNodesWithKeys[1:]

	for _, node := range donNodes {
		lmPluginConf := job.JSONConfig{
			"closePluginTimeoutSec":   10,
			"liquidityManagerAddress": "\"" + lmModules[l1ChainId].LM.EthAddress.String() + "\"",
			"liquidityManagerNetwork": "\"" + strconv.FormatUint(lmModules[l1ChainId].ChainSelectror, 10) +
				"\"" + "\n[pluginConfig.rebalancerConfig]\n type = \"ping-pong\"\n",
		}

		lmJobSpec := &client.OCR2TaskJobSpec{
			Name:    "lm " + lmModules[l1ChainId].LM.EthAddress.String(),
			JobType: "offchainreporting2",
			OCR2OracleSpec: job.OCR2OracleSpec{
				PluginType: "liquiditymanager",
				Relay:      "evm",
				RelayConfig: map[string]interface{}{
					"chainID": int(l1ChainId),
				},
				PluginConfig: lmPluginConf,

				ContractConfigTrackerPollInterval: *models.NewInterval(time.Second * 15),
				ContractID:                        lmModules[l1ChainId].LM.EthAddress.String(),      // registryAddr
				OCRKeyBundleID:                    null.StringFrom(node.KeysBundle.OCR2Key.Data.ID), // get node ocr2config.ID
				TransmitterID:                     null.StringFrom(node.KeysBundle.EthAddress),      // node addr
				P2PV2Bootstrappers:                pq.StringArray{P2Pv2Bootstrapper},                // bootstrap node key and address <p2p-key>@bootstrap:8000
			},
		}
		lggr.Debug().Interface("lmJobSpec", lmJobSpec).Msg("lmJobSpec")
		lggr.Info().Str("Node URL", node.Node.URL()).Msg("Adding LM job")
		j, err := node.Node.MustCreateJob(lmJobSpec)
		require.NoError(t, err, "Adding LM job shouldn't fail")
		lggr.Info().Str("jobId", j.Data.ID).Msg("LM job added")

	}

	defer lmModules[l1ChainId].ChainClient.Close()
	defer lmModules[l2ChainId].ChainClient.Close()

	return setUpArgs
}
