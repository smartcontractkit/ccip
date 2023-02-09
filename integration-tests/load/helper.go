package load

import (
	"math/big"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/smartcontractkit/chainlink-testing-framework/client"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/integration-tests/actions"
)

type loadArgs struct {
	t           *testing.T
	rps         int
	duration    time.Duration
	ccipTimeout time.Duration
	loadTimeOut time.Duration
	msgType     string
	envTear     func()
	ccipLoad    *CCIPE2ELoad
	loadGen     *client.LoadGenerator
}

// PopulateAndValidate collects all loadArgs
func PopulateAndValidate(t *testing.T) *loadArgs {
	inputRps := os.Getenv("LOAD_TEST_TPS")
	inputTimeout := os.Getenv("LOAD_TEST_CALLTIMEOUT")
	inputDuration := os.Getenv("LOAD_TEST_DURATION")
	inputMsgType := os.Getenv("LOAD_TEST_MSG_TYPE")
	p := &loadArgs{
		t:           t,
		rps:         3,
		duration:    4 * time.Minute,
		ccipTimeout: 15 * time.Minute,
		loadTimeOut: 25 * time.Minute,
		msgType:     DataOnlyTransfer,
	}
	if inputRps != "" {
		rps, err := strconv.Atoi(inputRps)
		require.NoError(t, err)
		require.LessOrEqual(t, rps, 16, "rps too high")
		p.rps = rps
	}
	if inputTimeout != "" {
		d, err := strconv.Atoi(inputTimeout)
		require.NoError(t, err)
		require.Greater(t, d, 1, "invalid timeout")
		require.LessOrEqual(t, d, 20, "invalid timeout")
		p.ccipTimeout = time.Duration(d) * time.Minute
		p.loadTimeOut = time.Duration(d*3) * time.Minute
	}
	if inputDuration != "" {
		d, err := strconv.Atoi(inputDuration)
		require.NoError(t, err)
		require.LessOrEqual(t, d, 90, "invalid duration")
		p.duration = time.Duration(d) * time.Minute
	}
	if inputMsgType != "" {
		require.Containsf(t, []string{DataOnlyTransfer, TokenTransfer}, inputMsgType, "invalid msg type")
		p.msgType = inputMsgType
	}
	return p
}

func (loadArgs *loadArgs) Setup() {
	transferAmounts := []*big.Int{big.NewInt(5e17), big.NewInt(5e17)}
	forwardLane, _, tearDown := actions.CCIPDefaultTestSetUp(
		loadArgs.t,
		"load-ccip",
		map[string]interface{}{
			"replicas": "6",
			"toml":     actions.DefaultCCIPCLNodeEnv(loadArgs.t),
			"env": map[string]interface{}{
				"CL_DEV": "true",
			},
		},
		transferAmounts,
		5,
		true,
		false,
	)
	loadArgs.envTear = tearDown
	if forwardLane == nil {
		return
	}
	source := forwardLane.Source
	dest := forwardLane.Dest
	ccipLoad := NewCCIPLoad(loadArgs.t, source, dest, loadArgs.ccipTimeout, 100000)
	require.NoError(loadArgs.t, forwardLane.IsLaneDeployed())
	ccipLoad.BeforeAllCall()
	loadgen, err := client.NewLoadGenerator(&client.LoadGeneratorConfig{
		RPS:         loadArgs.rps,
		Gun:         ccipLoad,
		Duration:    loadArgs.duration,
		CallTimeout: loadArgs.loadTimeOut,
		SharedData:  loadArgs.msgType,
	})
	require.NoError(loadArgs.t, err, "initiating loadgen")
	loadArgs.ccipLoad = ccipLoad
	loadArgs.loadGen = loadgen
}

func (loadArgs *loadArgs) Run() {
	loadArgs.loadGen.Run()
	_, failed := loadArgs.loadGen.Wait()
	require.False(loadArgs.t, failed, "load run is failed")
	require.Empty(loadArgs.t, loadArgs.loadGen.Errors(), "error in load sequence call")
}

func (loadArgs *loadArgs) TearDown() {
	defer loadArgs.envTear()
	loadArgs.ccipLoad.PrintStats(loadArgs.rps, loadArgs.duration.Minutes())
}
