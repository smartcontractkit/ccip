package load

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/rs/zerolog"
	"github.com/smartcontractkit/wasp"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"

	"github.com/smartcontractkit/chainlink/integration-tests/actions"
	"github.com/smartcontractkit/chainlink/integration-tests/testsetups"
)

type loadArgs struct {
	t             *testing.T
	lggr          zerolog.Logger
	ctx           context.Context
	ccipLoad      []*CCIPE2ELoad
	loadRunner    []*wasp.Generator
	Wg            *errgroup.Group
	TestCfg       *testsetups.CCIPTestConfig
	TestSetupArgs *testsetups.CCIPTestSetUpOutputs
}

func (l *loadArgs) SetupStableLoad() {
	transferAmounts := []*big.Int{big.NewInt(1)}
	lggr := l.lggr
	var setUpArgs *testsetups.CCIPTestSetUpOutputs
	if !l.TestCfg.ExistingDeployment {
		setUpArgs = testsetups.CCIPDefaultTestSetUp(l.TestCfg.Test, lggr, "load-ccip",
			map[string]interface{}{
				"replicas":   "6",
				"prometheus": "true",
			}, transferAmounts, 5, true, true, l.TestCfg)
	} else {
		setUpArgs = testsetups.CCIPExistingDeploymentTestSetUp(l.TestCfg.Test, lggr, transferAmounts, true, l.TestCfg)
	}
	if len(setUpArgs.Lanes) == 0 {
		return
	}
	l.TestSetupArgs = setUpArgs
	var lanes []*actions.CCIPLane
	for i := range setUpArgs.Lanes {
		lanes = append(lanes, setUpArgs.Lanes[i].ForwardLane)
		lanes = append(lanes, setUpArgs.Lanes[i].ReverseLane)
	}
	for _, lane := range lanes {
		ccipLoad := NewCCIPLoad(l.TestCfg.Test, lane, l.TestCfg.PhaseTimeout, 100000, lane.Reports)
		ccipLoad.BeforeAllCall(testsetups.DataOnlyTransfer)
		loadRunner, err := wasp.NewGenerator(&wasp.Config{
			T:           l.TestCfg.Test,
			Schedule:    wasp.Plain(l.TestCfg.Load.LoadRPS, l.TestCfg.TestDuration),
			LoadType:    wasp.RPSScheduleType,
			CallTimeout: l.TestCfg.Load.LoadTimeOut,
			Gun:         ccipLoad,
			Logger:      zerolog.Logger{},
			SharedData:  l.TestCfg.MsgType,
			LokiConfig: wasp.NewDefaultLokiConfig(
				os.Getenv("LOKI_URL"),
				os.Getenv("LOKI_TOKEN")),
			Labels: map[string]string{
				"test_group":   "load",
				"cluster":      "sdlc",
				"namespace":    lane.TestEnv.K8Env.Cfg.Namespace,
				"test_id":      "ccip",
				"source_chain": lane.SourceNetworkName,
				"dest_chain":   lane.DestNetworkName,
			},
		})
		require.NoError(l.TestCfg.Test, err, "initiating loadgen")
		l.ccipLoad = append(l.ccipLoad, ccipLoad)
		l.loadRunner = append(l.loadRunner, loadRunner)
	}
	l.TestSetupArgs.Reporter.SetDuration(l.TestCfg.TestDuration)
	l.TestSetupArgs.Reporter.SetRPS(l.TestCfg.Load.LoadRPS)
}

func (l *loadArgs) Run() {
	for i := range l.loadRunner {
		l.loadRunner[i].Run(false)
		go func(index int) {
			// wait for load to finish
			l.Wg.Go(func() error {
				_, failed := l.loadRunner[index].Wait()
				if failed {
					return fmt.Errorf("load run is failed")
				}
				if len(l.loadRunner[index].Errors()) > 0 {
					return fmt.Errorf("error in load sequence call %v", l.loadRunner[index].Errors())
				}
				return nil
			})
		}(i)
	}

	err := l.Wg.Wait()
	require.NoError(l.t, err, "load run is failed")
}

func (l *loadArgs) TearDown() {
	if l.TestSetupArgs.TearDown != nil {
		for i := range l.ccipLoad {
			l.ccipLoad[i].ReportAcceptedLog()
		}
		l.TestSetupArgs.TearDown()
	}
}

func NewLoadArgs(t *testing.T, lggr zerolog.Logger, parent context.Context) *loadArgs {
	wg, ctx := errgroup.WithContext(parent)
	return &loadArgs{
		t:       t,
		lggr:    lggr,
		Wg:      wg,
		ctx:     ctx,
		TestCfg: testsetups.NewCCIPTestConfig(t, lggr, testsetups.Load),
	}
}
