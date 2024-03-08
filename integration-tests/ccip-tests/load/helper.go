package load

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"sync"
	"testing"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/rs/zerolog"
	"github.com/smartcontractkit/wasp"
	"github.com/stretchr/testify/require"
	"go.uber.org/atomic"
	"golang.org/x/sync/errgroup"

	"github.com/smartcontractkit/chainlink-testing-framework/k8s/chaos"
	"github.com/smartcontractkit/chainlink-testing-framework/utils/testcontext"

	"github.com/smartcontractkit/chainlink/integration-tests/ccip-tests/actions"
	"github.com/smartcontractkit/chainlink/integration-tests/ccip-tests/testconfig"
	"github.com/smartcontractkit/chainlink/integration-tests/ccip-tests/testsetups"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/router"
)

const (
	Source      = "Source"
	destination = "Destination"
)

type ChaosConfig struct {
	ChaosName        string
	ChaosFunc        chaos.ManifestFunc
	ChaosProps       *chaos.Props
	WaitBetweenChaos time.Duration
}

type LoadArgs struct {
	t                *testing.T
	lggr             zerolog.Logger
	schedules        []*wasp.Segment
	RunnerWg         *errgroup.Group // to wait on individual load generators run
	LoadStarterWg    *sync.WaitGroup // waits for all the runners to start
	TestCfg          *testsetups.CCIPTestConfig
	TestSetupArgs    *testsetups.CCIPTestSetUpOutputs
	ChaosExps        []ChaosConfig
	LoadgenTearDowns []func()
	Labels           map[string]string
	pauseLoad        *atomic.Bool
}

func (l *LoadArgs) SetReportParams() {
	var qParams []string
	for k, v := range l.Labels {
		qParams = append(qParams, fmt.Sprintf("var-%s=%s", k, v))
	}
	// add one of the source and destination network to the grafana query params
	if len(l.TestSetupArgs.Lanes) > 0 {
		qParams = append(qParams, fmt.Sprintf("var-source_chain=%s", l.TestSetupArgs.Lanes[0].ForwardLane.SourceNetworkName))
		qParams = append(qParams, fmt.Sprintf("var-dest_chain=%s", l.TestSetupArgs.Lanes[0].ForwardLane.DestNetworkName))
	}
	err := l.TestSetupArgs.Reporter.AddToGrafanaDashboardQueryParams(qParams...)
	require.NoError(l.t, err, "failed to set grafana query params")
}

func (l *LoadArgs) Setup() {
	lggr := l.lggr
	existing := pointer.GetBool(l.TestCfg.TestGroupInput.ExistingDeployment)
	envName := "load-ccip"
	if existing {
		envName = "ccip-runner"
	}
	l.TestSetupArgs = testsetups.CCIPDefaultTestSetUp(l.TestCfg.Test, lggr, envName, nil, l.TestCfg)
	namespace := l.TestCfg.TestGroupInput.TestRunName
	if l.TestSetupArgs.Env != nil && l.TestSetupArgs.Env.K8Env != nil && l.TestSetupArgs.Env.K8Env.Cfg != nil {
		namespace = l.TestSetupArgs.Env.K8Env.Cfg.Namespace
	}
	l.Labels = map[string]string{
		"test_group": "load",
		"cluster":    "sdlc",
		"test_id":    "ccip",
		"namespace":  namespace,
	}
	l.TestSetupArgs.Reporter.SetGrafanaURLProvider(l.TestCfg.EnvInput)
	l.SetReportParams()
}

func (l *LoadArgs) setSchedule() {
	var segments []*wasp.Segment
	var segmentDuration time.Duration
	require.Greater(l.t, len(l.TestCfg.TestGroupInput.RequestPerUnitTime), 0, "RequestPerUnitTime must be set")

	if len(l.TestCfg.TestGroupInput.RequestPerUnitTime) > 1 {
		for i, req := range l.TestCfg.TestGroupInput.RequestPerUnitTime {
			duration := l.TestCfg.TestGroupInput.StepDuration[i].Duration()
			segmentDuration += duration
			segments = append(segments, wasp.Plain(req, duration)...)
		}
		totalDuration := l.TestCfg.TestGroupInput.TestDuration.Duration()
		repeatTimes := totalDuration.Seconds() / segmentDuration.Seconds()
		l.schedules = wasp.CombineAndRepeat(int(math.Round(repeatTimes)), segments)
	} else {
		l.schedules = wasp.Plain(l.TestCfg.TestGroupInput.RequestPerUnitTime[0], l.TestCfg.TestGroupInput.TestDuration.Duration())
	}
}

func (l *LoadArgs) SanityCheck() {
	for _, lane := range l.TestSetupArgs.Lanes {
		lane.ForwardLane.RecordStateBeforeTransfer()
		err := lane.ForwardLane.SendRequests(1, l.TestCfg.TestGroupInput.MsgType, big.NewInt(600_000))
		require.NoError(l.t, err)
		lane.ForwardLane.ValidateRequests(true)
		lane.ReverseLane.RecordStateBeforeTransfer()
		err = lane.ReverseLane.SendRequests(1, l.TestCfg.TestGroupInput.MsgType, big.NewInt(600_000))
		require.NoError(l.t, err)
		lane.ReverseLane.ValidateRequests(true)
	}
}

func (l *LoadArgs) ValidateCurseFollowedByUncurse(srcorDest string) {
	if srcorDest != Source && srcorDest != destination {
		return
	}
	var lanes []*actions.CCIPLane
	for _, lane := range l.TestSetupArgs.Lanes {
		lanes = append(lanes, lane.ForwardLane)
		if lane.ReverseLane != nil {
			lanes = append(lanes, lane.ReverseLane)
		}
	}
	// before cursing set pause
	l.pauseLoad.Store(true)
	defer func() {
		l.pauseLoad.Store(false)
	}()
	// wait for some time for pause to be active
	l.lggr.Info().Msg("Waiting for 1 minute after applying pause on load")
	time.Sleep(1 * time.Minute)

	var curseTimeStamps []time.Time
	for _, lane := range lanes {
		if srcorDest == Source {
			require.NoError(l.t, lane.Source.Common.CurseARM(), "error in cursing arm")
		}
		if srcorDest == destination {
			require.NoError(l.t, lane.Dest.Common.CurseARM(), "error in cursing arm")
		}
		// try to send requests and teh request should revert
		failedTx, _, _, err := lane.Source.SendRequest(
			lane.Dest.ReceiverDapp.EthAddress,
			actions.TokenTransfer, "msg with token more than aggregated rate",
			big.NewInt(600_000), // gas limit
		)
		require.Error(l.t, err)
		receipt, err := lane.Source.Common.ChainClient.GetTxReceipt(failedTx)
		require.NoError(l.t, err)
		hdr, err := lane.Source.Common.ChainClient.HeaderByNumber(context.Background(), receipt.BlockNumber)
		require.NoError(l.t, err)
		curseTimeStamps = append(curseTimeStamps, hdr.Timestamp)
		errReason, v, err := lane.Source.Common.ChainClient.RevertReasonFromTx(failedTx, router.RouterABI)
		require.NoError(l.t, err)
		require.Equal(l.t, "BadARMSignal", errReason)
		lane.Logger.Info().
			Str("Revert Reason", errReason).
			Interface("Args", v).
			Str("FailedTx", failedTx.Hex()).
			Msg("Msg sent while source ARM is cursed")
	}
	// no execution event should be received after the curse is applied,
	l.lggr.Info().Msg("Curse is applied on all lanes. Waiting for 5 minutes")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	errGrp := &errgroup.Group{}
	for i, lane := range lanes {
		curseTimeStamp := curseTimeStamps[i]
		errGrp.Go(func() error {
			ticker := time.NewTicker(time.Second)
			for {
				select {
				case <-ticker.C:
					eventFoundAfterCursing := false
					// verify if executionstate changed is received, it's not generated much after lane is cursed
					lane.Dest.ExecStateChangedWatcher.Range(func(key, value any) bool {
						e, exists := value.(*evm_2_evm_offramp.EVM2EVMOffRampExecutionStateChanged)
						if exists {
							vLogs := e.Raw
							hdr, err := lane.Dest.Common.ChainClient.HeaderByNumber(ctx, big.NewInt(int64(vLogs.BlockNumber)))
							if err != nil {
								return true
							}
							if hdr.Timestamp.Add(30 * time.Second).After(curseTimeStamp) {
								eventFoundAfterCursing = true
								return false
							}
						}
						return true
					})
					if eventFoundAfterCursing {
						return fmt.Errorf("ExecutionStateChanged Event detected after curse")
					}
				case <-ctx.Done():
					return nil
				}
			}
		})
	}

	err := errGrp.Wait()
	require.NoError(l.t, err, "error received to validate no execution state changed is generated after lane is cursed")
	// now uncurse all
	for _, lane := range lanes {
		if srcorDest == Source {
			require.NoError(l.t, lane.Source.Common.UnvoteToCurseARM(), "error to unvote in cursing arm")
		}
		if srcorDest == destination {
			require.NoError(l.t, lane.Dest.Common.UnvoteToCurseARM(), "error to unvote in cursing arm")
		}
	}
}

func (l *LoadArgs) TriggerLoadByLane() {
	l.setSchedule()
	l.TestSetupArgs.Reporter.SetDuration(l.TestCfg.TestGroupInput.TestDuration.Duration())

	// start load for a lane
	startLoad := func(lane *actions.CCIPLane) {
		lane.Logger.Info().
			Str("Source Network", lane.SourceNetworkName).
			Str("Destination Network", lane.DestNetworkName).
			Msg("Starting load for lane")

		ccipLoad := NewCCIPLoad(l.TestCfg.Test, lane, l.TestCfg.TestGroupInput.PhaseTimeout.Duration(), 100000)
		ccipLoad.BeforeAllCall(l.TestCfg.TestGroupInput.MsgType, big.NewInt(*l.TestCfg.TestGroupInput.DestGasLimit))
		lokiConfig := l.TestCfg.EnvInput.Logging.Loki
		labels := make(map[string]string)
		for k, v := range l.Labels {
			labels[k] = v
		}
		labels["source_chain"] = lane.SourceNetworkName
		labels["dest_chain"] = lane.DestNetworkName
		loadRunner, err := wasp.NewGenerator(&wasp.Config{
			T:                     l.TestCfg.Test,
			GenName:               fmt.Sprintf("lane %s-> %s", lane.SourceNetworkName, lane.DestNetworkName),
			Schedule:              l.schedules,
			LoadType:              wasp.RPS,
			RateLimitUnitDuration: l.TestCfg.TestGroupInput.TimeUnit.Duration(),
			CallResultBufLen:      10, // we keep the last 10 call results for each generator, as the detailed report is generated at the end of the test
			CallTimeout:           (l.TestCfg.TestGroupInput.PhaseTimeout.Duration()) * 5,
			Gun:                   ccipLoad,
			Logger:                ccipLoad.Lane.Logger,
			SharedData:            l.TestCfg.TestGroupInput.MsgType,
			LokiConfig:            wasp.NewLokiConfig(lokiConfig.Endpoint, lokiConfig.TenantId, nil, nil),
			Labels:                labels,
			FailOnErr:             pointer.GetBool(l.TestCfg.TestGroupInput.FailOnFirstErrorInLoad),
		})
		require.NoError(l.TestCfg.Test, err, "initiating loadgen for lane %s --> %s",
			lane.SourceNetworkName, lane.DestNetworkName)
		loadRunner.Run(false)
		l.AddToRunnerGroup(loadRunner)
	}

	for _, lane := range l.TestSetupArgs.Lanes {
		lane := lane
		l.LoadStarterWg.Add(1)
		go func() {
			defer l.LoadStarterWg.Done()
			startLoad(lane.ForwardLane)
		}()
		if pointer.GetBool(l.TestSetupArgs.Cfg.TestGroupInput.BiDirectionalLane) {
			l.LoadStarterWg.Add(1)
			go func() {
				defer l.LoadStarterWg.Done()
				startLoad(lane.ReverseLane)
			}()
		}
	}
}

func (l *LoadArgs) AddToRunnerGroup(gen *wasp.Generator) {
	// watch for pause signal
	go func(gen *wasp.Generator) {
		ticker := time.NewTicker(time.Second)
		for {
			select {
			case <-ticker.C:
				if l.pauseLoad.Load() {
					gen.Pause()
					continue
				}
				gen.Resume()
			}
		}
	}(gen)
	l.RunnerWg.Go(func() error {
		_, failed := gen.Wait()
		if failed {
			return fmt.Errorf("load run is failed")
		}
		if len(gen.Errors()) > 0 {
			return fmt.Errorf("error in load sequence call %v", gen.Errors())
		}
		return nil
	})
}

func (l *LoadArgs) Wait() {
	l.lggr.Info().Msg("Waiting for load to start on all lanes")
	// wait for load runner to start
	l.LoadStarterWg.Wait()
	l.lggr.Info().Msg("Waiting for load to finish on all lanes")
	// wait for load runner to finish
	err := l.RunnerWg.Wait()
	require.NoError(l.t, err, "load run is failed")
	l.lggr.Info().Msg("Load finished on all lanes")
}

func (l *LoadArgs) ApplyChaos() {
	testEnv := l.TestSetupArgs.Env
	if testEnv == nil || testEnv.K8Env == nil {
		l.lggr.Warn().Msg("test environment is nil, skipping chaos")
		return
	}
	testEnv.ChaosLabelForCLNodes(l.TestCfg.Test)

	for _, exp := range l.ChaosExps {
		if exp.WaitBetweenChaos > 0 {
			l.lggr.Info().Msgf("sleeping for %s after chaos %s", exp.WaitBetweenChaos, exp.ChaosName)
			time.Sleep(exp.WaitBetweenChaos)
		}
		l.lggr.Info().Msgf("Starting to apply chaos %s at %s", exp.ChaosName, time.Now().UTC())
		// apply chaos
		chaosId, err := testEnv.K8Env.Chaos.Run(exp.ChaosFunc(testEnv.K8Env.Cfg.Namespace, exp.ChaosProps))
		require.NoError(l.t, err)
		if chaosId != "" {
			chaosDur, err := time.ParseDuration(exp.ChaosProps.DurationStr)
			require.NoError(l.t, err)
			err = testEnv.K8Env.Chaos.WaitForAllRecovered(chaosId, chaosDur+1*time.Minute)
			require.NoError(l.t, err)
			l.lggr.Info().Msgf("chaos %s is recovered at %s", exp.ChaosName, time.Now().UTC())
			err = testEnv.K8Env.Chaos.Stop(chaosId)
			require.NoError(l.t, err)
			l.lggr.Info().Msgf("stopped chaos %s at %s", exp.ChaosName, time.Now().UTC())
		}
	}
}

func (l *LoadArgs) TearDown() {
	for _, tearDn := range l.LoadgenTearDowns {
		tearDn()
	}
	if l.TestSetupArgs.TearDown != nil {
		require.NoError(l.t, l.TestSetupArgs.TearDown())
	}
}

func (l *LoadArgs) TriggerLoadBySource() {
	require.NotNil(l.t, l.TestCfg.TestGroupInput.TestDuration, "test duration input is nil")
	require.GreaterOrEqual(l.t, 1, len(l.TestCfg.TestGroupInput.RequestPerUnitTime), "time unit input must be specified")
	l.TestSetupArgs.Reporter.SetDuration(l.TestCfg.TestGroupInput.TestDuration.Duration())
	var laneBySource = make(map[string][]*actions.CCIPLane)
	for _, lane := range l.TestSetupArgs.Lanes {
		laneBySource[lane.ForwardLane.SourceNetworkName] = append(laneBySource[lane.ForwardLane.SourceNetworkName], lane.ForwardLane)
		if lane.ReverseLane != nil {
			laneBySource[lane.ReverseLane.SourceNetworkName] = append(laneBySource[lane.ReverseLane.SourceNetworkName], lane.ReverseLane)
		}
	}
	for source, lanes := range laneBySource {
		source := source
		lanes := lanes
		l.LoadStarterWg.Add(1)
		go func() {
			defer l.LoadStarterWg.Done()
			l.lggr.Info().
				Str("Source Network", source).
				Msg("Starting load for source")
			allLabels := make(map[string]string)
			for k, v := range l.Labels {
				allLabels[k] = v
			}
			allLabels["source_chain"] = source
			multiCallGen, err := NewMultiCallLoadGenerator(l.TestCfg, lanes, l.TestCfg.TestGroupInput.RequestPerUnitTime[0], allLabels)
			require.NoError(l.t, err)
			lokiConfig := l.TestCfg.EnvInput.Logging.Loki
			loadRunner, err := wasp.NewGenerator(&wasp.Config{
				T:                     l.TestCfg.Test,
				GenName:               fmt.Sprintf("Source %s", source),
				Schedule:              wasp.Plain(1, l.TestCfg.TestGroupInput.TestDuration.Duration()), // hardcoded request per unit time to 1 as we are using multiCallGen
				LoadType:              wasp.RPS,
				RateLimitUnitDuration: l.TestCfg.TestGroupInput.TimeUnit.Duration(),
				CallResultBufLen:      10, // we keep the last 10 call results for each generator, as the detailed report is generated at the end of the test
				CallTimeout:           (l.TestCfg.TestGroupInput.PhaseTimeout.Duration()) * 5,
				Gun:                   multiCallGen,
				Logger:                multiCallGen.logger,
				LokiConfig:            wasp.NewLokiConfig(lokiConfig.Endpoint, lokiConfig.TenantId, nil, nil),
				Labels:                allLabels,
				FailOnErr:             pointer.GetBool(l.TestCfg.TestGroupInput.FailOnFirstErrorInLoad),
			})
			require.NoError(l.TestCfg.Test, err, "initiating loadgen for source %s", source)
			loadRunner.Run(false)
			l.AddToRunnerGroup(loadRunner)
			l.LoadgenTearDowns = append(l.LoadgenTearDowns, func() {
				require.NoError(l.t, multiCallGen.Stop())
			})
		}()
	}
}

func NewLoadArgs(t *testing.T, lggr zerolog.Logger, chaosExps ...ChaosConfig) *LoadArgs {
	wg, _ := errgroup.WithContext(testcontext.Get(t))
	return &LoadArgs{
		t:             t,
		lggr:          lggr,
		RunnerWg:      wg,
		TestCfg:       testsetups.NewCCIPTestConfig(t, lggr, testconfig.Load),
		ChaosExps:     chaosExps,
		LoadStarterWg: &sync.WaitGroup{},
		pauseLoad:     atomic.NewBool(false),
	}
}
