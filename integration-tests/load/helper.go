package load

import (
	"math/big"
	"testing"

	"github.com/rs/zerolog"
	"github.com/smartcontractkit/chainlink-testing-framework/loadgen"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/integration-tests/actions"
	"github.com/smartcontractkit/chainlink/integration-tests/testsetups"
)

type loadArgs struct {
	t             *testing.T
	ccipLoad      *CCIPE2ELoad
	loadRunner    *loadgen.Generator
	TestCfg       *testsetups.CCIPTestConfig
	TestSetupArgs *testsetups.CCIPTestSetUpOutputs
}

func (l *loadArgs) Setup() {
	transferAmounts := []*big.Int{big.NewInt(10)}
	var forwardLane *actions.CCIPLane
	var setUpArgs *testsetups.CCIPTestSetUpOutputs
	if !l.TestCfg.ExistingDeployment {
		setUpArgs = testsetups.CCIPDefaultTestSetUp(l.TestCfg.Test, "load-ccip", map[string]interface{}{
			"replicas": "6",
			"env": map[string]interface{}{
				"CL_DEV": "true",
			},
		}, transferAmounts, 5, true, false, l.TestCfg)
	} else {
		setUpArgs = testsetups.CCIPExistingDeploymentTestSetUp(l.TestCfg.Test, transferAmounts, false, l.TestCfg)
	}
	require.Greater(l.TestCfg.Test, len(setUpArgs.Lanes), 0, "error in default set up")
	l.TestSetupArgs = setUpArgs
	forwardLane = setUpArgs.Lanes[0].ForwardLane
	if forwardLane == nil {
		return
	}
	source := forwardLane.Source
	dest := forwardLane.Dest

	ccipLoad := NewCCIPLoad(l.TestCfg.Test, source, dest, l.TestCfg.PhaseTimeout, 100000, forwardLane.Reports)
	ccipLoad.BeforeAllCall(l.TestCfg.MsgType)
	loadRunner, err := loadgen.NewLoadGenerator(&loadgen.Config{
		T:           nil,
		Schedule:    loadgen.Plain(l.TestCfg.Load.LoadRPS, l.TestCfg.TestDuration),
		LoadType:    loadgen.RPSScheduleType,
		CallTimeout: l.TestCfg.Load.LoadTimeOut,
		Gun:         ccipLoad,
		Logger:      zerolog.Logger{},
		SharedData:  l.TestCfg.MsgType,
	})
	require.NoError(l.TestCfg.Test, err, "initiating loadgen")
	l.ccipLoad = ccipLoad
	l.loadRunner = loadRunner
	l.TestSetupArgs.Reporter.SetDuration(l.TestCfg.TestDuration)
	l.TestSetupArgs.Reporter.SetRPS(l.TestCfg.Load.LoadRPS)
}

func (l *loadArgs) Run() {
	l.loadRunner.Run()
	_, failed := l.loadRunner.Wait()
	require.False(l.t, failed, "load run is failed")
	require.Empty(l.t, l.loadRunner.Errors(), "error in load sequence call")
}

func (l *loadArgs) TearDown() {
	if l.TestSetupArgs.TearDown != nil {
		l.ccipLoad.ReportAcceptedLog()
		l.TestSetupArgs.TearDown()
	}
}
