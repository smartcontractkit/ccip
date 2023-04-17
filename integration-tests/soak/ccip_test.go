package soak

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/smartcontractkit/chainlink-testing-framework/utils"

	"github.com/smartcontractkit/chainlink/integration-tests/actions"
	"github.com/smartcontractkit/chainlink/integration-tests/testsetups"
)

// TestCCIPSoak verifies that CCIP requests can be successfully delivered for mentioned duration triggered at a certain interval
// If run on live networks it can reuse already deployed contracts if the addresses are provided in ../contracts/ccip/laneconfig/contracts.json
// This test does a full environment set up along with deploying CL nodes in K8 cluster
func TestSoakCCIP(t *testing.T) {
	var (
		tearDown        func()
		laneA, laneB    *actions.CCIPLane
		totalReqLaneA   = 0
		totalReqLaneB   = 0
		reqSuccessLaneA = 0
		reqSuccessLaneB = 0
	)
	l := utils.GetTestLogger(t)
	TestCfg := testsetups.NewCCIPTestConfig(t, l, testsetups.Soak)
	interval := TestCfg.Soak.SoakInterval
	duration := TestCfg.TestDuration

	t.Cleanup(func() {
		if tearDown != nil {
			l.Info().Msg("Tearing down the environment")
			tearDown()
			if laneA != nil {
				l.Info().
					Str("total duration", fmt.Sprint(duration)).
					Str("req interval", fmt.Sprint(interval)).
					Int("Total Requests", totalReqLaneA).
					Int("Successful Requests", reqSuccessLaneA).
					Msgf("Soak Result for lane %s --> %s", laneA.SourceNetworkName, laneA.DestNetworkName)
			}
			if laneB != nil {
				l.Info().
					Str("total duration", fmt.Sprint(duration)).
					Str("req interval", fmt.Sprint(interval)).
					Int("Total Requests", totalReqLaneB).
					Int("Successful Requests", reqSuccessLaneB).
					Msgf("Soak Result for lane %s --> %s", laneB.SourceNetworkName, laneB.DestNetworkName)
			}
		}
	})

	transferAmounts := []*big.Int{big.NewInt(5e17)}
	var setUpOutput *testsetups.CCIPTestSetUpOutputs
	if !TestCfg.ExistingDeployment {
		setUpOutput = testsetups.CCIPDefaultTestSetUp(t, l, "soak-ccip", map[string]interface{}{
			"replicas": "6",
		}, transferAmounts, 5, true, true, TestCfg)
	} else {
		setUpOutput = testsetups.CCIPExistingDeploymentTestSetUp(t, l, transferAmounts, true, TestCfg)
	}

	if len(setUpOutput.Lanes) == 0 {
		return
	}
	laneA = setUpOutput.Lanes[0].ForwardLane
	laneB = setUpOutput.Lanes[0].ReverseLane

	setUpOutput.Reporter.SetDuration(duration)
	setUpOutput.Reporter.SetSoakRunInterval(interval)
	tearDown = setUpOutput.TearDown

	t.Run(fmt.Sprintf("CCIP message transfer from network %s to network %s for %s",
		laneA.SourceNetworkName, laneA.DestNetworkName, duration), func(t *testing.T) {
		t.Parallel()
		totalReqLaneA, reqSuccessLaneA = laneA.SoakRun(interval, duration)
		l.Info().
			Str("Test Duration", fmt.Sprintf("%s", duration)).
			Str("Request Triggering interval", fmt.Sprintf("%s", interval)).
			Str("Source", laneA.SourceNetworkName).
			Str("Destination", laneA.DestNetworkName).
			Msg("Starting lane A")
		if totalReqLaneA != reqSuccessLaneA {
			t.Fatalf("Failed to deliver %d requests out of %d", totalReqLaneA-reqSuccessLaneA, totalReqLaneA)
		}
	})
	if laneB != nil {
		t.Run(fmt.Sprintf("CCIP message transfer from network %s to network %s for %s", laneB.SourceNetworkName, laneB.DestNetworkName, duration), func(t *testing.T) {
			t.Parallel()
			l.Info().
				Str("Test Duration", fmt.Sprintf("%s", duration)).
				Str("Request Triggering interval", fmt.Sprintf("%s", interval)).
				Str("Source", laneB.SourceNetworkName).
				Str("Destination", laneB.DestNetworkName).
				Msg("Starting lane B")
			totalReqLaneB, reqSuccessLaneB = laneB.SoakRun(interval, duration)
			if totalReqLaneB != reqSuccessLaneB {
				t.Fatalf("Failed to deliver %d requests out of %d", totalReqLaneB-reqSuccessLaneB, totalReqLaneB)
			}
		})
	}
}

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

	t.Cleanup(func() {
		if tearDown != nil {
			l.Info().Msg("Tearing down the environment")
			tearDown()
		}
	})

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
