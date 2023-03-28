package soak

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"

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

	TestCfg := testsetups.NewCCIPTestConfig(t, testsetups.Soak)
	interval := TestCfg.Soak.SoakInterval
	duration := TestCfg.TestDuration

	t.Cleanup(func() {
		if tearDown != nil {
			log.Info().Msg("Tearing down the environment")
			tearDown()
			if laneA != nil {
				log.Info().
					Str("total duration", fmt.Sprint(duration)).
					Str("req interval", fmt.Sprint(interval)).
					Int("Total Requests", totalReqLaneA).
					Int("Successful Requests", reqSuccessLaneA).
					Msgf("Soak Result for lane %s --> %s", laneA.SourceNetworkName, laneA.DestNetworkName)
			}
			if laneB != nil {
				log.Info().
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
		setUpOutput = testsetups.CCIPDefaultTestSetUp(t, "soak-ccip", map[string]interface{}{
			"replicas": "6",
		}, transferAmounts, 5, true, true, TestCfg)
	} else {
		setUpOutput = testsetups.CCIPExistingDeploymentTestSetUp(t, transferAmounts, true, TestCfg)
	}

	require.Greater(t, len(setUpOutput.Lanes), 0, "error in default set up")
	laneA = setUpOutput.Lanes[0].ForwardLane
	laneB = setUpOutput.Lanes[0].ReverseLane

	if laneA == nil {
		return
	}

	setUpOutput.Reporter.SetDuration(duration)
	setUpOutput.Reporter.SetSoakRunInterval(interval)
	tearDown = setUpOutput.TearDown

	t.Run(fmt.Sprintf("CCIP message transfer from network %s to network %s for %s",
		laneA.SourceNetworkName, laneA.DestNetworkName, duration), func(t *testing.T) {
		t.Parallel()
		totalReqLaneA, reqSuccessLaneA = laneA.SoakRun(interval, duration)
		log.Info().
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
			log.Info().
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
