package testsetups

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/smartcontractkit/chainlink-env/environment"
	"github.com/smartcontractkit/chainlink-testing-framework/blockchain"
	"github.com/smartcontractkit/chainlink-testing-framework/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/multierr"
	"go.uber.org/zap/zapcore"
	"golang.org/x/sync/errgroup"

	networks "github.com/smartcontractkit/chainlink/integration-tests"
	"github.com/smartcontractkit/chainlink/integration-tests/actions"
	"github.com/smartcontractkit/chainlink/integration-tests/contracts/ccip/laneconfig"
	"github.com/smartcontractkit/chainlink/integration-tests/testreporters"
)

const (
	TokenTransfer                   string = "WithToken"
	DataOnlyTransfer                string = "WithoutToken"
	Load                            string = "Load"
	Soak                            string = "Soak"
	Chaos                           string = "Chaos"
	Smoke                           string = "Smoke"
	DefaultTTL                             = 70 * time.Minute
	DefaultNoOfNetworks             int    = 2
	DefaultLoadRPS                  int64  = 2
	DefaultLoadTimeOut                     = 30 * time.Minute
	DefaultPhaseTimeoutForLongTests        = 20 * time.Minute
	DefaultPhaseTimeout                    = 5 * time.Minute
	DefaultSoakInterval                    = 45 * time.Second
	DefaultTestDuration                    = 15 * time.Minute
)

var (
	LongRunningTests = map[string]struct{}{
		Load: {},
		Soak: {},
	}
	GethResourceProfile = map[string]interface{}{
		"requests": map[string]interface{}{
			"cpu":    "4",
			"memory": "6Gi",
		},
		"limits": map[string]interface{}{
			"cpu":    "4",
			"memory": "6Gi",
		},
	}
	DONResourceProfile = map[string]interface{}{
		"requests": map[string]interface{}{
			"cpu":    "4",
			"memory": "8Gi",
		},
		"limits": map[string]interface{}{
			"cpu":    "4",
			"memory": "8Gi",
		},
	}
	DONDBResourceProfile = map[string]interface{}{
		"stateful": true,
		"capacity": "10Gi",
		"resources": map[string]interface{}{
			"requests": map[string]interface{}{
				"cpu":    "2",
				"memory": "4Gi",
			},
			"limits": map[string]interface{}{
				"cpu":    "2",
				"memory": "4Gi",
			},
		},
	}
	NodeFundingForLoad = big.NewFloat(50)
	DefaultNodeFunding = big.NewFloat(1)
)

type NetworkPair struct {
	NetworkA blockchain.EVMNetwork
	NetworkB blockchain.EVMNetwork
}

type CCIPTestConfig struct {
	Test                    *testing.T
	EnvTTL                  time.Duration
	MsgType                 string
	PhaseTimeout            time.Duration
	TestDuration            time.Duration
	ExistingDeployment      bool
	ReuseContracts          bool
	NodeFunding             *big.Float
	Load                    *CCIPLoadInput
	Soak                    *CCIPSoakInput
	AllNetworks             []blockchain.EVMNetwork
	NetworkPairs            []NetworkPair
	NoOfNetworks            int
	GethResourceProfile     map[string]interface{}
	CLNodeResourceProfile   map[string]interface{}
	CLNodeDBResourceProfile map[string]interface{}
}

type CCIPLoadInput struct {
	LoadRPS     int64
	LoadTimeOut time.Duration
}

type CCIPSoakInput struct {
	SoakInterval time.Duration
}

func (p *CCIPTestConfig) setSoakInputs() {
	var allError error
	p.Soak = &CCIPSoakInput{
		SoakInterval: DefaultSoakInterval,
	}
	if interval, err := utils.GetEnv("CCIP_SOAK_TEST_REQ_INTERVAL"); err == nil {
		if interval != "" {
			d, err := time.ParseDuration(interval)
			if err != nil {
				allError = multierr.Append(allError, err)
			} else {
				if d.Seconds() < 10 || d.Hours() > 2 {
					allError = multierr.Append(allError, fmt.Errorf("invalid interval %d - must be between 10s and 2h", d))
				} else {
					p.Soak.SoakInterval = d
				}
			}
		}
	} else {
		allError = multierr.Append(allError, err)
	}

	if allError != nil {
		p.Test.Fatal(allError)
	}
}

func (p *CCIPTestConfig) setLoadInputs() {
	var allError error
	p.Load = &CCIPLoadInput{
		LoadRPS:     DefaultLoadRPS,
		LoadTimeOut: DefaultLoadTimeOut,
	}

	if inputRps, err := utils.GetEnv("CCIP_LOAD_TEST_RPS"); err == nil {
		if inputRps != "" {
			rps, err := strconv.ParseInt(inputRps, 10, 64)
			if err != nil {
				allError = multierr.Append(allError, err)
			} else {
				maxRps := int64(16)
				if rps > maxRps {
					allError = multierr.Append(allError, fmt.Errorf("rps %d is too high - maximum value is %d", rps, maxRps))
				} else {
					p.Load.LoadRPS = rps
				}
			}
		}
	} else {
		allError = multierr.Append(allError, err)
	}

	if p.PhaseTimeout.Seconds() > 0 {
		p.Load.LoadTimeOut = time.Duration(p.PhaseTimeout.Minutes()+10) * time.Minute
	}
	if allError != nil {
		p.Test.Fatal(allError)
	}
}

func (p *CCIPTestConfig) FormNetworkPairs() {
	for i := 0; i < p.NoOfNetworks; i++ {
		for j := i + 1; j < p.NoOfNetworks; j++ {
			p.NetworkPairs = append(p.NetworkPairs, NetworkPair{
				NetworkA: p.AllNetworks[i],
				NetworkB: p.AllNetworks[j],
			})
		}
	}
}

// NewCCIPTestConfig collects all test related CCIPTestConfig from environment variables
func NewCCIPTestConfig(t *testing.T, lggr zerolog.Logger, tType string) *CCIPTestConfig {
	var allError error
	if len(networks.SelectedNetworks) < 3 {
		lggr.Fatal().
			Interface("SELECTED_NETWORKS", networks.SelectedNetworks).
			Msg("Set source and destination network in index 1 & 2 of env variable SELECTED_NETWORKS")
	}
	p := &CCIPTestConfig{
		Test:         t,
		MsgType:      TokenTransfer,
		PhaseTimeout: DefaultPhaseTimeout,
		TestDuration: DefaultTestDuration,
		NodeFunding:  DefaultNodeFunding,
		NoOfNetworks: DefaultNoOfNetworks,
		AllNetworks:  networks.SelectedNetworks[1:],
	}
	if _, ok := LongRunningTests[tType]; ok {
		p.EnvTTL = DefaultTTL
		p.GethResourceProfile = GethResourceProfile
		p.CLNodeResourceProfile = DONResourceProfile
		p.CLNodeDBResourceProfile = DONDBResourceProfile
		p.NodeFunding = NodeFundingForLoad
		p.PhaseTimeout = DefaultPhaseTimeoutForLongTests
	}

	inputNoOfNetworks, err := utils.GetEnv("CCIP_NO_OF_NETWORKS")
	if err != nil {
		allError = multierr.Append(allError, err)
	} else {
		if inputNoOfNetworks != "" {
			n, err := strconv.Atoi(inputNoOfNetworks)
			if err != nil {
				allError = multierr.Append(allError, err)
			} else {
				p.NoOfNetworks = n
			}
		}
	}
	// skip the first index as it is generally set to Simulated EVM in dev mode to be used by other tests
	simulated := p.AllNetworks[0].Simulated
	for i := 1; i < len(p.AllNetworks); i++ {
		if p.AllNetworks[i].Simulated != simulated {
			t.Fatal("networks must be of the same type either simulated or real")
		}
	}
	// if the networks are not simulated use the first p.NoOfNetworks networks from the selected networks
	if !simulated && len(p.AllNetworks) != p.NoOfNetworks {
		if len(p.AllNetworks) < p.NoOfNetworks {
			allError = multierr.Append(allError, fmt.Errorf("not enough networks provided"))
		} else {
			p.AllNetworks = p.AllNetworks[:p.NoOfNetworks]
		}
	}
	// If provided networks is lesser than the required number of networks
	// and the provided networks are simulated network, create replicas of the provided networks with
	// different chain ids
	if len(p.AllNetworks) < p.NoOfNetworks {
		if p.AllNetworks[0].Simulated {
			actualNoOfNetworks := len(p.AllNetworks)
			n := p.AllNetworks[0]
			for i := 0; i < p.NoOfNetworks-actualNoOfNetworks; i++ {
				chainID := networks.AdditionalSimulatedChainIds[i]
				p.AllNetworks = append(p.AllNetworks, blockchain.EVMNetwork{
					Name:                      fmt.Sprintf("simulated-non-dev%d", len(p.AllNetworks)+1),
					ChainID:                   chainID,
					Simulated:                 true,
					PrivateKeys:               []string{networks.AdditionalSimulatedPvtKeys[i]},
					ChainlinkTransactionLimit: n.ChainlinkTransactionLimit,
					Timeout:                   n.Timeout,
					MinimumConfirmations:      n.MinimumConfirmations,
					GasEstimationBuffer:       n.GasEstimationBuffer,
					ClientImplementation:      n.ClientImplementation,
				})
			}
		}
	}
	lggr.Info().Interface("Networks", p.AllNetworks).Msg("Running tests with networks")
	if p.NoOfNetworks > 2 {
		p.FormNetworkPairs()
	} else {
		p.NetworkPairs = []NetworkPair{
			{
				NetworkA: p.AllNetworks[0],
				NetworkB: p.AllNetworks[1],
			},
		}
	}

	for _, n := range p.NetworkPairs {
		lggr.Info().Str("NetworkA", n.NetworkA.Name).Str("NetworkB", n.NetworkB.Name).Msg("Network Pairs")
	}

	if ttlDuration, err := utils.GetEnv("CCIP_KEEP_ENV_TTL"); err == nil {
		if ttlDuration != "" {
			keepEnvFor, err := time.ParseDuration(ttlDuration)
			if err != nil {
				allError = multierr.Append(allError, fmt.Errorf("invalid KEEP_ENV_TTL %s", ttlDuration))
			} else {
				if keepEnvFor.Minutes() < 20 {
					allError = multierr.Append(allError, fmt.Errorf("invalid timeout %s - must be greater than 20m", keepEnvFor))
				} else {
					p.EnvTTL = keepEnvFor
				}
			}
		}
	} else {
		allError = multierr.Append(allError, err)
	}

	if phaseTimeOut, err := utils.GetEnv("CCIP_PHASE_VALIDATION_TIMEOUT"); err != nil {
		allError = multierr.Append(allError, err)
	} else {
		if phaseTimeOut != "" {
			timeout, err := time.ParseDuration(phaseTimeOut)
			if err != nil {
				allError = multierr.Append(allError, fmt.Errorf("invalid PHASE_VALIDATION_TIMEOUT %s", phaseTimeOut))
			} else {
				if timeout.Minutes() < 1 || timeout.Minutes() > 50 {
					allError = multierr.Append(allError, fmt.Errorf("invalid timeout %s - must be between 1m and 50m", timeout))
				} else {
					p.PhaseTimeout = timeout
				}
			}
		}
	}

	if inputDuration, err := utils.GetEnv("CCIP_TEST_DURATION"); err == nil {
		if inputDuration != "" {
			d, err := time.ParseDuration(inputDuration)
			if err != nil {
				allError = multierr.Append(allError, err)
			} else {
				if d.Minutes() < 1 || d.Hours() > 5 {
					allError = multierr.Append(allError, fmt.Errorf("invalid duration %d - must be between 1m and 5h", d))
				} else {
					p.TestDuration = d
				}
			}
		}
	} else {
		allError = multierr.Append(allError, err)
	}

	inputMsgType, err := utils.GetEnv("CCIP_MSG_TYPE")
	if err != nil {
		allError = multierr.Append(allError, err)
	} else {
		if inputMsgType != "" {
			if inputMsgType != DataOnlyTransfer && inputMsgType != TokenTransfer {
				allError = multierr.Append(allError, fmt.Errorf("invalid msg type %s", inputMsgType))
			} else {
				p.MsgType = inputMsgType
			}
		}
	}

	if fundingAmountStr, err := utils.GetEnv("CCIP_CHAINLINK_NODE_FUNDING"); err == nil {
		if fundingAmountStr != "" {
			fundingAmount, _ := big.NewFloat(0).SetString(fundingAmountStr)
			if fundingAmount == nil {
				allError = multierr.Append(allError, fmt.Errorf("invalid CCIP_CHAINLINK_NODE_FUNDING env variable value: %s", fundingAmountStr))
			} else {
				p.NodeFunding = fundingAmount
			}
		}
	} else {
		allError = multierr.Append(allError, err)
	}

	if existing, err := utils.GetEnv("CCIP_TESTS_ON_EXISTING_DEPLOYMENT"); err == nil {
		if existing != "" {
			e, err := strconv.ParseBool(existing)
			if err != nil {
				allError = multierr.Append(allError, err)
			} else {
				p.ExistingDeployment = e
			}
		}
	} else {
		allError = multierr.Append(allError, err)
	}

	if reuse, err := utils.GetEnv("CCIP_REUSE_CONTRACTS"); err == nil {
		if reuse != "" {
			e, err := strconv.ParseBool(reuse)
			if err != nil {
				allError = multierr.Append(allError, err)
			} else {
				p.ReuseContracts = e
			}
		}
	} else {
		allError = multierr.Append(allError, err)
	}

	if allError != nil {
		t.Fatal(allError)
	}
	switch tType {
	case Load:
		p.setLoadInputs()
	case Soak:
		p.setSoakInputs()
	}

	return p
}

type BiDirectionalLaneConfig struct {
	NetworkA    blockchain.EVMNetwork
	NetworkB    blockchain.EVMNetwork
	ForwardLane *actions.CCIPLane
	ReverseLane *actions.CCIPLane
}

type CCIPTestSetUpOutputs struct {
	Cfg            *CCIPTestConfig
	Lanes          []*BiDirectionalLaneConfig
	Reporter       *testreporters.CCIPTestReporter
	LaneConfigFile string
	LaneConfig     *laneconfig.Lanes
	TearDown       func()
}

func (o *CCIPTestSetUpOutputs) AddLanesForNetworkPair(
	lggr zerolog.Logger,
	networkA, networkB blockchain.EVMNetwork,
	chainClientA, chainClientB blockchain.EVMClient,
	transferAmounts []*big.Int,
	numOfCommitNodes int,
	commitAndExecOnSameDON, bidirectional bool,
	ccipEnv *actions.CCIPTestEnv,
	newBootstrap bool,
) error {
	var allErrors error
	t := o.Cfg.Test
	var k8Env *environment.Environment
	if ccipEnv != nil {
		k8Env = ccipEnv.K8Env
	}
	configureCLNode := !o.Cfg.ExistingDeployment
	setUpFuncs, ctx := errgroup.WithContext(context.Background())
	bidirectionalLane := &BiDirectionalLaneConfig{
		NetworkA: networkA,
		NetworkB: networkB,
	}

	// Use new set of clients(sourceChainClient,destChainClient)
	// with new header subscriptions(otherwise transactions
	// on one lane will keep on waiting for transactions on other lane for the same network)
	// Currently for simulated network clients(from same network) created with NewEVMClient does not sync nonce
	// ConcurrentEVMClient is a work-around for that.
	sourceChainClientA2B, err := blockchain.ConcurrentEVMClient(networkA, k8Env, chainClientA)
	require.NoError(t, err, "Connecting to blockchain nodes shouldn't fail")
	sourceChainClientA2B.ParallelTransactions(true)

	destChainClientA2B, err := blockchain.ConcurrentEVMClient(networkB, k8Env, chainClientB)
	require.NoError(t, err, "Connecting to blockchain nodes shouldn't fail")
	destChainClientA2B.ParallelTransactions(true)

	sourceChainClientB2A, err := blockchain.ConcurrentEVMClient(networkB, k8Env, destChainClientA2B)
	require.NoError(t, err, "Connecting to blockchain nodes shouldn't fail")
	sourceChainClientB2A.ParallelTransactions(true)

	destChainClientB2A, err := blockchain.ConcurrentEVMClient(networkA, k8Env, sourceChainClientA2B)
	require.NoError(t, err, "Connecting to blockchain nodes shouldn't fail")
	destChainClientB2A.ParallelTransactions(true)

	ccipLaneA2B := &actions.CCIPLane{
		Test:              t,
		Logger:            zerolog.New(zerolog.NewConsoleWriter(zerolog.ConsoleTestWriter(t))).With().Timestamp().Logger(),
		TestEnv:           ccipEnv,
		SourceChain:       sourceChainClientA2B,
		DestChain:         destChainClientA2B,
		SourceNetworkName: actions.NetworkName(networkA.Name),
		DestNetworkName:   actions.NetworkName(networkB.Name),
		ValidationTimeout: o.Cfg.PhaseTimeout,
		SentReqs:          make(map[int64]actions.CCIPRequest),
		TotalFee:          big.NewInt(0),
		SourceBalances:    make(map[string]*big.Int),
		DestBalances:      make(map[string]*big.Int),
		Context:           ctx,
		CommonContractsWg: &sync.WaitGroup{},
		Reports: o.Reporter.AddNewLane(fmt.Sprintf("%d To %d",
			networkA.ChainID, networkB.ChainID)),
	}
	ccipLaneA2B.SrcNetworkLaneCfg, err = o.LaneConfig.ReadLaneConfig(networkA.Name)
	require.NoError(t, err, "Reading lane config shouldn't fail")
	ccipLaneA2B.DstNetworkLaneCfg, err = o.LaneConfig.ReadLaneConfig(networkB.Name)
	require.NoError(t, err, "Reading lane config shouldn't fail")

	ccipLaneA2B.Logger = lggr.With().Str("Lane",
		fmt.Sprintf("%s-->%s", ccipLaneA2B.SourceNetworkName, ccipLaneA2B.DestNetworkName)).Logger()
	bidirectionalLane.ForwardLane = ccipLaneA2B
	var ccipLaneB2A *actions.CCIPLane

	if bidirectional {
		ccipLaneB2A = &actions.CCIPLane{
			Test:              t,
			TestEnv:           ccipEnv,
			SourceNetworkName: actions.NetworkName(networkB.Name),
			DestNetworkName:   actions.NetworkName(networkA.Name),
			SourceChain:       sourceChainClientB2A,
			DestChain:         destChainClientB2A,
			ValidationTimeout: o.Cfg.PhaseTimeout,
			SourceBalances:    make(map[string]*big.Int),
			DestBalances:      make(map[string]*big.Int),
			SentReqs:          make(map[int64]actions.CCIPRequest),
			TotalFee:          big.NewInt(0),
			Context:           ctx,
			CommonContractsWg: &sync.WaitGroup{},
			Reports:           o.Reporter.AddNewLane(fmt.Sprintf("%d To %d", networkB.ChainID, networkA.ChainID)),
			SrcNetworkLaneCfg: ccipLaneA2B.DstNetworkLaneCfg,
			DstNetworkLaneCfg: ccipLaneA2B.SrcNetworkLaneCfg,
		}
		ccipLaneB2A.Logger = lggr.With().Str("Lane",
			fmt.Sprintf("%s-->%s", ccipLaneB2A.SourceNetworkName, ccipLaneB2A.DestNetworkName)).Logger()
		bidirectionalLane.ReverseLane = ccipLaneB2A
	}
	o.Lanes = append(o.Lanes, bidirectionalLane)

	ccipLaneA2B.CommonContractsWg.Add(1)
	setUpFuncs.Go(func() error {
		lggr.Info().Msgf("Setting up lane %s to %s", networkA.Name, networkB.Name)
		err := ccipLaneA2B.DeployNewCCIPLane(numOfCommitNodes, commitAndExecOnSameDON, nil, nil,
			transferAmounts, newBootstrap, configureCLNode, o.Cfg.ExistingDeployment)
		if err != nil {
			allErrors = multierr.Append(allErrors, fmt.Errorf("deploying lane %s to %s; err - %+v", networkA.Name, networkB.Name, err))
		}
		return err
	})

	if ccipLaneB2A != nil {
		ccipLaneB2A.CommonContractsWg.Add(1)
	}

	setUpFuncs.Go(func() error {
		if bidirectional {
			ccipLaneA2B.CommonContractsWg.Wait()
			srcCommon := ccipLaneA2B.Dest.Common.CopyAddresses(ccipLaneB2A.Context, ccipLaneB2A.SourceChain, o.Cfg.ExistingDeployment)
			destCommon := ccipLaneA2B.Source.Common.CopyAddresses(ccipLaneB2A.Context, ccipLaneB2A.DestChain, o.Cfg.ExistingDeployment)
			lggr.Info().Msgf("Setting up lane %s to %s", networkB.Name, networkA.Name)
			err := ccipLaneB2A.DeployNewCCIPLane(numOfCommitNodes, commitAndExecOnSameDON, srcCommon, destCommon,
				transferAmounts, false, configureCLNode, o.Cfg.ExistingDeployment)
			if err != nil {
				allErrors = multierr.Append(allErrors, fmt.Errorf("deploying lane %s to %s; err -  %+v", networkB.Name, networkA.Name, err))
			}
			return err
		}
		return nil
	})

	errs := make(chan error, 1)
	go func() {
		errs <- setUpFuncs.Wait()
	}()

	// wait for either context to get cancelled or all the error-groups to finish execution
	for {
		select {
		case err := <-errs:
			// check if there has been any error while waiting for the error groups
			// to finish execution
			return err
		case <-ctx.Done():
			lggr.Print(ctx.Err())
			return allErrors
		}
	}
}

// CCIPDefaultTestSetUp sets up the environment for CCIP tests
// if configureCLNode is set as false, it assumes:
// 1. contracts are already deployed on live networks
// 2. CL nodes are set up and configured with existing contracts
// 3. No k8 env deployment is needed
// It reuses already deployed contracts from the addresses provided in ../contracts/ccip/laneconfig/contracts.json
//
// If bidirectional is true it sets up two-way lanes between NetworkA and NetworkB. Same CL nodes are used for both the lanes.
// If bidirectional is false only one way lane is set up.
//
// Returns -
// 1. CCIPLane for NetworkA --> NetworkB
// 2. If bidirectional is true, CCIPLane for NetworkB --> NetworkA
// 3. If configureCLNode is true, the tearDown func to call when environment needs to be destroyed
func CCIPDefaultTestSetUp(
	t *testing.T,
	lggr zerolog.Logger,
	envName string,
	clProps map[string]interface{},
	transferAmounts []*big.Int,
	numOfCommitNodes int, commitAndExecOnSameDON, bidirectional bool,
	inputs *CCIPTestConfig,
) *CCIPTestSetUpOutputs {
	var (
		ccipEnv *actions.CCIPTestEnv
		k8Env   *environment.Environment
		ctx     context.Context
		err     error
		chains  []blockchain.EVMClient
	)
	filename := fmt.Sprintf("./tmp_%s.json", strings.ReplaceAll(t.Name(), "/", "_"))
	setUpArgs := &CCIPTestSetUpOutputs{
		Cfg:            inputs,
		Reporter:       testreporters.NewCCIPTestReporter(t, lggr),
		LaneConfigFile: filename,
	}
	_, err = os.Stat(setUpArgs.LaneConfigFile)
	if err == nil {
		// remove the existing lane config file
		err = os.Remove(setUpArgs.LaneConfigFile)
		require.NoError(t, err, "error while removing existing lane config file - %s", setUpArgs.LaneConfigFile)
	}
	if inputs.ExistingDeployment || inputs.ReuseContracts {
		setUpArgs.LaneConfig, err = laneconfig.ReadLanesFromExistingDeployment()
		require.NoError(t, err)
	} else {
		setUpArgs.LaneConfig, err = laneconfig.CreateDeploymentJSON(setUpArgs.LaneConfigFile)
		require.NoError(t, err)
		if setUpArgs.LaneConfig == nil {
			setUpArgs.LaneConfig = &laneconfig.Lanes{LaneConfigs: make(map[string]*laneconfig.LaneConfig)}
		}
	}

	parent, cancel := context.WithCancel(context.Background())
	defer cancel()

	configureCLNode := !inputs.ExistingDeployment
	if configureCLNode {
		clProps["db"] = inputs.CLNodeDBResourceProfile
		clProps["chainlink"] = map[string]interface{}{
			"resources": inputs.CLNodeResourceProfile,
		}
		// deploy the env if configureCLNode is true
		ccipEnv = actions.DeployEnvironments(
			t,
			&environment.Config{
				TTL:             inputs.EnvTTL,
				NamespacePrefix: envName,
				Test:            t,
			}, clProps, inputs.GethResourceProfile, inputs.AllNetworks)
		k8Env = ccipEnv.K8Env
		ccipEnv.CLNodeWithKeyReady, ctx = errgroup.WithContext(parent)
		if ccipEnv.K8Env.WillUseRemoteRunner() {
			return setUpArgs
		}
	} else {
		// if configureCLNode is false, use a placeholder env to create remote runner
		k8Env = environment.New(
			&environment.Config{
				TTL:             inputs.EnvTTL,
				NamespacePrefix: envName,
				Test:            t,
			})
		err = k8Env.Run()
		require.NoErrorf(t, err, "error creating environment remote runner %s", k8Env.Cfg.Namespace)
		if k8Env.WillUseRemoteRunner() {
			return setUpArgs
		}
	}

	chainByChainID := make(map[int64]blockchain.EVMClient)
	for _, network := range inputs.AllNetworks {
		ec, err := blockchain.NewEVMClient(network, k8Env)
		require.NoError(t, err, "Connecting to blockchain nodes shouldn't fail")
		chains = append(chains, ec)
		chainByChainID[network.ChainID] = ec
	}
	if configureCLNode {
		ccipEnv.CLNodeWithKeyReady.Go(func() error {
			return ccipEnv.SetUpNodesAndKeys(ctx, inputs.NodeFunding, chains)
		})
	}

	for i, n := range inputs.NetworkPairs {
		newBootstrap := false
		if i == 0 {
			// create bootstrap job once
			newBootstrap = true
		}
		err = setUpArgs.AddLanesForNetworkPair(
			lggr, n.NetworkA, n.NetworkB,
			chainByChainID[n.NetworkA.ChainID], chainByChainID[n.NetworkB.ChainID],
			transferAmounts, numOfCommitNodes, commitAndExecOnSameDON,
			bidirectional, ccipEnv, newBootstrap)
		assert.NoError(t, err)
	}
	err = laneconfig.WriteLanesToJSON(setUpArgs.LaneConfigFile, setUpArgs.LaneConfig)
	assert.NoError(t, err)

	setUpArgs.TearDown = func() {
		for _, lanes := range setUpArgs.Lanes {
			assert.NoError(t, lanes.ForwardLane.CleanUp(), "error while cleaning up forward lane")
			if lanes.ReverseLane != nil {
				assert.NoError(t, lanes.ReverseLane.CleanUp(), "error while cleaning up reverse lane")
			}
		}
		if configureCLNode {
			require.NoError(t, err, "error while writing lane config to file - %s", setUpArgs.LaneConfigFile)
			err = actions.TeardownSuite(t, ccipEnv.K8Env, utils.ProjectRoot, ccipEnv.CLNodes, setUpArgs.Reporter,
				zapcore.ErrorLevel, chains...)
			require.NoError(t, err, "Environment teardown shouldn't fail")
		}
	}
	return setUpArgs
}

// CCIPExistingDeploymentTestSetUp is same as CCIPDefaultTestSetUp
// except it's called when
// 1. contracts are already deployed on live networks
// 2. CL nodes are set up and configured with existing contracts
// 3. No k8 env deployment is needed
// It reuses already deployed contracts from the addresses provided in ../contracts/ccip/laneconfig/contracts.json
// Returns -
// CCIPLane for NetworkA --> NetworkB
// CCIPLane for NetworkB --> NetworkA
func CCIPExistingDeploymentTestSetUp(
	t *testing.T,
	lggr zerolog.Logger,
	transferAmounts []*big.Int,
	bidirectional bool,
	input *CCIPTestConfig,
) *CCIPTestSetUpOutputs {
	return CCIPDefaultTestSetUp(t, lggr, "runner", nil, transferAmounts,
		0, false, bidirectional, input)
}
