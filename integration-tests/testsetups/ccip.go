package testsetups

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/smartcontractkit/chainlink-env/environment"
	"github.com/smartcontractkit/chainlink-testing-framework/blockchain"
	"github.com/smartcontractkit/chainlink-testing-framework/utils"
	"github.com/stretchr/testify/require"
	"go.uber.org/multierr"
	"go.uber.org/zap/zapcore"
	"golang.org/x/sync/errgroup"

	"github.com/smartcontractkit/chainlink/integration-tests/actions"
	"github.com/smartcontractkit/chainlink/integration-tests/testreporters"
)

const (
	TokenTransfer       string = "WithToken"
	DataOnlyTransfer    string = "WithoutToken"
	Load                string = "Load"
	Soak                string = "Soak"
	Chaos               string = "Chaos"
	Smoke               string = "Smoke"
	DefaultLoadRPS      int64  = 3
	DefaultLoadTimeOut         = 30 * time.Minute
	DefaultPhaseTimeout        = 10 * time.Minute
	DefaultSoakInterval        = 45 * time.Second
	DefaultTestDuration        = 5 * time.Minute
)

type CCIPTestConfig struct {
	Test               *testing.T
	EnvTTL             time.Duration
	MsgType            string
	PhaseTimeout       time.Duration
	TestDuration       time.Duration
	ExistingDeployment bool
	NodeFunding        *big.Float
	Load               *CCIPLoadInput
	Soak               *CCIPSoakInput
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
	if interval, err := utils.GetEnv("CCIP_SOAK_TEST_REQ_INTERVAL"); err != nil {
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
		p.Load.LoadTimeOut = p.PhaseTimeout * 3
	}
	if allError != nil {
		p.Test.Fatal(allError)
	}
}

// NewCCIPTestConfig collects all test related CCIPTestConfig from environment variables
func NewCCIPTestConfig(t *testing.T, tType string) *CCIPTestConfig {
	var allError error
	p := &CCIPTestConfig{}
	p.Test = t

	p.EnvTTL = 20 * time.Minute
	if ttlDuration, err := utils.GetEnv("KEEP_ENV_TTL"); err != nil {
		if ttlDuration != "" {
			keepEnvFor, err := time.ParseDuration(ttlDuration)
			if err != nil {
				allError = multierr.Append(allError, fmt.Errorf("invalid KEEP_ENV_TTL %s", ttlDuration))
			} else {
				if keepEnvFor.Minutes() < 10 {
					allError = multierr.Append(allError, fmt.Errorf("invalid timeout %s - must be greater than 10m", keepEnvFor))
				} else {
					p.EnvTTL = keepEnvFor
				}
			}
		}
	} else {
		allError = multierr.Append(allError, err)
	}

	p.PhaseTimeout = DefaultPhaseTimeout
	if phaseTimeOut, err := utils.GetEnv("CCIP_PHASE_VALIDATION_TIMEOUT"); err != nil {
		allError = multierr.Append(allError, err)
	} else {
		if phaseTimeOut != "" {
			timeout, err := time.ParseDuration(phaseTimeOut)
			if err != nil {
				allError = multierr.Append(allError, fmt.Errorf("invalid PHASE_VALIDATION_TIMEOUT %s", phaseTimeOut))
			} else {
				if timeout.Minutes() < 1 || timeout.Minutes() > 20 {
					allError = multierr.Append(allError, fmt.Errorf("invalid timeout %s - must be between 1m and 20m", timeout))
				} else {
					p.PhaseTimeout = timeout
				}
			}
		}
	}

	p.TestDuration = DefaultTestDuration
	if inputDuration, err := utils.GetEnv("CCIP_TEST_DURATION"); err != nil {
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

	p.MsgType = TokenTransfer
	inputMsgType, err := utils.GetEnv("MSG_TYPE")
	if err != nil {
		allError = multierr.Append(allError, err)
	} else {
		if inputMsgType != "" {
			if inputMsgType != DataOnlyTransfer || inputMsgType != TokenTransfer {
				allError = multierr.Append(allError, fmt.Errorf("invalid msg type %s", inputMsgType))
			} else {
				p.MsgType = inputMsgType
			}
		}
	}

	p.NodeFunding = big.NewFloat(1)
	if fundingAmountStr, err := utils.GetEnv("CCIP_CHAINLINK_NODE_FUNDING"); err != nil {
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

	p.ExistingDeployment = false
	if existing, err := utils.GetEnv("CCIP_TESTS_ON_EXISTING_DEPLOYMENT"); err != nil {
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

	switch tType {
	case Load:
		p.setLoadInputs()
	case Soak:
		p.setSoakInputs()
	}
	if allError != nil {
		t.Fatal(allError)
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
	Cfg      *CCIPTestConfig
	Lanes    []*BiDirectionalLaneConfig
	Reporter *testreporters.CCIPTestReporter
	TearDown func()
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
	envName string,
	clProps map[string]interface{},
	transferAmounts []*big.Int,
	numOfCommitNodes int, commitAndExecOnSameDON, bidirectional bool,
	inputs *CCIPTestConfig,
) *CCIPTestSetUpOutputs {
	setUpArgs := &CCIPTestSetUpOutputs{
		Cfg:      inputs,
		Reporter: testreporters.NewCCIPTestReporter(t),
	}
	bidirectionalLane := &BiDirectionalLaneConfig{
		NetworkA: actions.NetworkA,
		NetworkB: actions.NetworkB,
	}
	setUpArgs.Lanes = append(setUpArgs.Lanes, bidirectionalLane)

	parent, cancel := context.WithCancel(context.Background())
	defer cancel()
	var allErrors, err error
	setUpFuncs, ctx := errgroup.WithContext(parent)
	var ccipEnv *actions.CCIPTestEnv
	var k8Env *environment.Environment
	configureCLNode := !inputs.ExistingDeployment
	namespace := strings.ToLower(fmt.Sprintf("%s-%s-%s", envName, actions.NetworkAName, actions.NetworkBName))
	if configureCLNode {
		// deploy the env if configureCLNode is true
		ccipEnv = actions.DeployEnvironments(
			t,
			&environment.Config{
				TTL:             inputs.EnvTTL,
				NamespacePrefix: namespace,
				Test:            t,
			}, clProps)
		k8Env = ccipEnv.K8Env
		if ccipEnv.K8Env.WillUseRemoteRunner() {
			return setUpArgs
		}
	} else {
		// if configureCLNode is false, use a placeholder env to create remote runner
		k8Env = environment.New(
			&environment.Config{
				TTL:             inputs.EnvTTL,
				NamespacePrefix: namespace,
				Test:            t,
			})
		err = k8Env.Run()
		require.NoErrorf(t, err, "error creating environment remote runner %s", k8Env.Cfg.Namespace)
		if k8Env.WillUseRemoteRunner() {
			return setUpArgs
		}
	}

	sourceChainClientA2B, err := blockchain.NewEVMClient(actions.NetworkA, k8Env)
	require.NoError(t, err, "Connecting to blockchain nodes shouldn't fail")

	destChainClientA2B, err := blockchain.NewEVMClient(actions.NetworkB, k8Env)
	require.NoError(t, err, "Connecting to blockchain nodes shouldn't fail")

	// For the reverse lane another set of clients(sourceChainClient,destChainClient)
	// are required with new header subscriptions(otherwise transactions
	// on one lane will keep on waiting for transactions on other lane for the same network)
	// Currently for simulated network clients(from same network) created with NewEVMClient does not sync nonce
	// ConcurrentEVMClient is a work-around for that.
	sourceChainClientB2A, err := blockchain.ConcurrentEVMClient(actions.NetworkB, k8Env, destChainClientA2B)
	require.NoError(t, err, "Connecting to blockchain nodes shouldn't fail")

	destChainClientB2A, err := blockchain.ConcurrentEVMClient(actions.NetworkA, k8Env, sourceChainClientA2B)
	require.NoError(t, err, "Connecting to blockchain nodes shouldn't fail")

	setUpFuncs.Go(func() error {
		if !configureCLNode {
			return nil
		}

		err = ccipEnv.SetUpNodesAndKeys(ctx, inputs.NodeFunding, sourceChainClientA2B, destChainClientA2B)
		if err != nil {
			allErrors = multierr.Append(allErrors, fmt.Errorf("setting up nodes and keys shouldn't fail; err -  %+v", err))
		} else {
			// creates the channel to denote job creation for the lanes can be started
			ccipEnv.CLNodeWithKeyReady = make(chan struct{})
			// sends two values to the channel if bidirectional is true
			ccipEnv.CLNodeWithKeyReady <- struct{}{}
			if bidirectional {
				ccipEnv.CLNodeWithKeyReady <- struct{}{}
			}
		}
		return err
	})

	ccipLaneA2B := &actions.CCIPLane{
		Test:              t,
		TestEnv:           ccipEnv,
		SourceChain:       sourceChainClientA2B,
		DestChain:         destChainClientA2B,
		SourceNetworkName: actions.NetworkAName,
		DestNetworkName:   actions.NetworkBName,
		ValidationTimeout: inputs.PhaseTimeout,
		SentReqs:          make(map[int64]actions.CCIPRequest),
		TotalFee:          big.NewInt(0),
		SourceBalances:    make(map[string]*big.Int),
		DestBalances:      make(map[string]*big.Int),
		Context:           ctx,
		CommonContractsWg: &sync.WaitGroup{},
		Reports:           setUpArgs.Reporter.AddNewLane(fmt.Sprintf("%d To %d", actions.NetworkA.ChainID, actions.NetworkB.ChainID)),
	}
	bidirectionalLane.ForwardLane = ccipLaneA2B

	var ccipLaneB2A *actions.CCIPLane

	if bidirectional {
		ccipLaneB2A = &actions.CCIPLane{
			Test:              t,
			TestEnv:           ccipEnv,
			SourceNetworkName: actions.NetworkBName,
			DestNetworkName:   actions.NetworkAName,
			SourceChain:       sourceChainClientB2A,
			DestChain:         destChainClientB2A,
			ValidationTimeout: inputs.PhaseTimeout,
			SourceBalances:    make(map[string]*big.Int),
			DestBalances:      make(map[string]*big.Int),
			SentReqs:          make(map[int64]actions.CCIPRequest),
			TotalFee:          big.NewInt(0),
			Context:           ctx,
			CommonContractsWg: &sync.WaitGroup{},
			Reports:           setUpArgs.Reporter.AddNewLane(fmt.Sprintf("%d To %d", actions.NetworkB.ChainID, actions.NetworkA.ChainID)),
		}
		bidirectionalLane.ReverseLane = ccipLaneB2A
	}

	ccipLaneA2B.CommonContractsWg.Add(1)
	setUpFuncs.Go(func() error {
		log.Info().Msg("Setting up lane A to B")
		err := ccipLaneA2B.DeployNewCCIPLane(numOfCommitNodes, commitAndExecOnSameDON, nil, nil,
			transferAmounts, true, configureCLNode)
		if err != nil {
			allErrors = multierr.Append(allErrors, fmt.Errorf("deploying lane A to B; err - %+v", err))
		}
		return err
	})

	if ccipLaneB2A != nil {
		ccipLaneB2A.CommonContractsWg.Add(1)
	}

	setUpFuncs.Go(func() error {
		if bidirectional {
			ccipLaneA2B.CommonContractsWg.Wait()
			srcCommon := ccipLaneA2B.Dest.Common.CopyAddresses(ccipLaneB2A.Context, ccipLaneB2A.SourceChain)
			destCommon := ccipLaneA2B.Source.Common.CopyAddresses(ccipLaneB2A.Context, ccipLaneB2A.DestChain)
			log.Info().Msg("Setting up lane B to A")
			err := ccipLaneB2A.DeployNewCCIPLane(numOfCommitNodes, commitAndExecOnSameDON, srcCommon, destCommon,
				transferAmounts, false, configureCLNode)
			if err != nil {
				allErrors = multierr.Append(allErrors, fmt.Errorf("deploying lane B to A; err -  %+v", err))
			}
			return err
		}
		return nil
	})

	setUpArgs.TearDown = func() {
		if configureCLNode {
			err := actions.TeardownSuite(t, ccipEnv.K8Env, utils.ProjectRoot, ccipEnv.CLNodes, setUpArgs.Reporter,
				zapcore.ErrorLevel, ccipLaneA2B.SourceChain, ccipLaneA2B.DestChain)
			require.NoError(t, err, "Environment teardown shouldn't fail")
		}
	}

	errs := make(chan error, 1)
	go func() {
		errs <- setUpFuncs.Wait()
	}()

	// wait for either context to get cancelled or all the error-groups to finish execution
	for {
		select {
		case <-ctx.Done():
			// if the context is cancelled it means there has been some error
			// allErrors helps to print a stacktrace
			require.NoError(t, allErrors)
			return setUpArgs
		case err := <-errs:
			// check if there has been any error while waiting for the error groups
			// to finish execution
			require.NoError(t, err)
			return setUpArgs
		}
	}
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
	transferAmounts []*big.Int,
	bidirectional bool,
	input *CCIPTestConfig,
) *CCIPTestSetUpOutputs {
	if actions.NetworkA.Simulated || actions.NetworkB.Simulated {
		t.Fatalf("test cannot run in simulated env")
	}
	return CCIPDefaultTestSetUp(t, "runner", nil, transferAmounts,
		0, false, bidirectional, input)
}
