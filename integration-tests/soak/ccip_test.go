package soak

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/smartcontractkit/chainlink-testing-framework/utils"

	"github.com/smartcontractkit/chainlink/integration-tests/actions"
	"github.com/smartcontractkit/chainlink/integration-tests/testsetups"
)

func TestSoakCCIPMultiChain(t *testing.T) {
	type subtestInput struct {
		testName string
		lane     *actions.CCIPLane
	}
	var (
		tearDown func()
		tcs      []subtestInput
	)
	l := utils.GetTestLogger(t)
	TestCfg := testsetups.NewCCIPTestConfig(t, l, testsetups.Soak)
	interval := TestCfg.Soak.SoakInterval
	duration := TestCfg.TestDuration

	transferAmounts := []*big.Int{big.NewInt(100)}
	var setUpOutput *testsetups.CCIPTestSetUpOutputs
	if !TestCfg.ExistingDeployment {
		setUpOutput = testsetups.CCIPDefaultTestSetUp(t, l, "soak-ccip", map[string]interface{}{
			"replicas":   "6",
			"prometheus": "true",
		}, transferAmounts, 5, true, true, TestCfg)
	} else {
		setUpOutput = testsetups.CCIPExistingDeploymentTestSetUp(t, l, transferAmounts, true, TestCfg)
	}

	if len(setUpOutput.Lanes) == 0 {
		return
	}

	setUpOutput.Reporter.SetDuration(duration)
	setUpOutput.Reporter.SetSoakRunInterval(interval)
	tearDown = setUpOutput.TearDown
	t.Cleanup(func() {
		tearDown()
	})
	for i := range setUpOutput.Lanes {
		tcs = append(tcs, subtestInput{
			testName: fmt.Sprintf("CCIP message transfer from network %s to network %s for %s",
				setUpOutput.Lanes[i].ForwardLane.SourceNetworkName, setUpOutput.Lanes[i].ForwardLane.DestNetworkName, duration),
			lane: setUpOutput.Lanes[i].ForwardLane,
		}, subtestInput{
			testName: fmt.Sprintf("CCIP message transfer from network %s to network %s for %s",
				setUpOutput.Lanes[i].ReverseLane.SourceNetworkName, setUpOutput.Lanes[i].ReverseLane.DestNetworkName, duration),
			lane: setUpOutput.Lanes[i].ReverseLane,
		})
		setUpOutput.Lanes[i].ForwardLane.SourceChain.ParallelTransactions(false)
		setUpOutput.Lanes[i].ReverseLane.SourceChain.ParallelTransactions(false)
	}

	for _, testcase := range tcs {
		tc := testcase
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()
			l.Info().
				Str("Test Duration", fmt.Sprintf("%s", duration)).
				Str("Request Triggering interval", fmt.Sprintf("%s", interval)).
				Str("Source", tc.lane.SourceNetworkName).
				Str("Destination", tc.lane.DestNetworkName).
				Msgf("Starting lane %s -> %s", tc.lane.SourceNetworkName, tc.lane.DestNetworkName)
			totalReqLane, reqSuccessLane := tc.lane.SoakRun(interval, duration)
			if totalReqLane != reqSuccessLane {
				t.Fatalf("Failed to deliver %d requests out of %d for lane %s -> %s",
					totalReqLane-reqSuccessLane, totalReqLane, tc.lane.SourceNetworkName, tc.lane.DestNetworkName)
			}
		})
	}
}
