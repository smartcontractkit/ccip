package load

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/rs/zerolog"
	ctfClient "github.com/smartcontractkit/chainlink-testing-framework/client"
	"github.com/smartcontractkit/chainlink-testing-framework/loadgen"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"

	"github.com/smartcontractkit/chainlink/integration-tests/actions"
	"github.com/smartcontractkit/chainlink/integration-tests/testsetups"
)

type loadArgs struct {
	t             *testing.T
	ctx           context.Context
	ccipLoad      []*CCIPE2ELoad
	loadRunner    []*loadgen.Generator
	Wg            *errgroup.Group
	TestCfg       *testsetups.CCIPTestConfig
	TestSetupArgs *testsetups.CCIPTestSetUpOutputs
}

func (l *loadArgs) SetupStableLoad() {
	transferAmounts := []*big.Int{big.NewInt(10)}
	var setUpArgs *testsetups.CCIPTestSetUpOutputs
	if !l.TestCfg.ExistingDeployment {
		setUpArgs = testsetups.CCIPDefaultTestSetUp(l.TestCfg.Test, "load-ccip",
			map[string]interface{}{
				"replicas":   "6",
				"prometheus": "true",
			}, transferAmounts, 5, true, true, l.TestCfg)
	} else {
		setUpArgs = testsetups.CCIPExistingDeploymentTestSetUp(l.TestCfg.Test, transferAmounts, true, l.TestCfg)
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
		source := lane.Source
		dest := lane.Dest

		ccipLoad := NewCCIPLoad(l.TestCfg.Test, source, dest, l.TestCfg.PhaseTimeout, 100000, lane.Reports)
		ccipLoad.BeforeAllCall(testsetups.DataOnlyTransfer)
		loadRunner, err := loadgen.NewLoadGenerator(&loadgen.Config{
			T:           l.TestCfg.Test,
			Schedule:    loadgen.Plain(l.TestCfg.Load.LoadRPS, l.TestCfg.TestDuration),
			LoadType:    loadgen.RPSScheduleType,
			CallTimeout: l.TestCfg.Load.LoadTimeOut,
			Gun:         ccipLoad,
			Logger:      zerolog.Logger{},
			SharedData:  l.TestCfg.MsgType,
			LokiConfig: ctfClient.NewDefaultLokiConfig(
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
		l.loadRunner[i].Run()
		// wait for load to finish
		l.Wg.Go(func() error {
			_, failed := l.loadRunner[i].Wait()
			if failed {
				return fmt.Errorf("load run is failed")
			}
			if len(l.loadRunner[i].Errors()) > 0 {
				return fmt.Errorf("error in load sequence call %v", l.loadRunner[i].Errors())
			}
			return nil
		})
	}
	err := l.Wg.Wait()
	require.NoError(l.t, err, "load run is failed")
}

func (l *loadArgs) TearDown() {
	if l.TestSetupArgs.TearDown != nil {
		for i, c := range l.ccipLoad {
			lggr := zerolog.New(zerolog.NewConsoleWriter(zerolog.ConsoleTestWriter(c.t))).
				With().Timestamp().Logger().With().
				Str("Lane",
					fmt.Sprintf("%d-->%d", c.Source.Common.ChainClient.GetChainID().Int64(),
						c.Destination.Common.ChainClient.GetChainID().Int64())).
				Logger()
			l.ccipLoad[i].ReportAcceptedLog(lggr)
		}
		l.TestSetupArgs.TearDown()
	}
}

func NewLoadArgs(t *testing.T, parent context.Context) *loadArgs {
	wg, ctx := errgroup.WithContext(parent)
	return &loadArgs{
		t:       t,
		Wg:      wg,
		ctx:     ctx,
		TestCfg: testsetups.NewCCIPTestConfig(t, testsetups.Load),
	}
}
