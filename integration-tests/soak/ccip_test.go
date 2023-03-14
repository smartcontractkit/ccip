package soak

import (
	"fmt"
	"math/big"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/integration-tests/actions"
)

var (
	getTestInterval = func() (time.Duration, error) {
		if intervalEnv := os.Getenv("CCIP_TEST_INTERVAL"); intervalEnv != "" {
			interval, err := time.ParseDuration(intervalEnv)
			if err != nil {
				return 0, err
			}
			return interval, nil
		} else {
			return 30 * time.Second, nil
		}
	}
	getTestDuration = func() (time.Duration, error) {
		if durationEnv := os.Getenv("CCIP_TEST_DURATION"); durationEnv != "" {
			duration, err := time.ParseDuration(durationEnv)
			if err != nil {
				return 0, err
			}
			return duration, nil
		} else {
			return 2 * time.Minute, nil
		}
	}
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
	var err error

	// if interval and duration is provided in env variable, use that
	interval, err := getTestInterval()
	require.NoError(t, err, "invalid interval provided")
	duration, err := getTestDuration()
	require.NoError(t, err, "invalid duration provided")

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
			if reqSuccessLaneA != totalReqLaneA || reqSuccessLaneB != totalReqLaneB {
				t.Fail()
			}
		}
	})

	transferAmounts := []*big.Int{big.NewInt(5e17), big.NewInt(5e17)}
	laneA, laneB, tearDown = actions.CCIPDefaultTestSetUp(t, "soak-ccip", map[string]interface{}{
		"replicas": "6",
	}, transferAmounts, 5, true, true, true)

	if laneA == nil {
		return
	}
	laneRuns := &sync.WaitGroup{}
	laneRuns.Add(1)
	go func() {
		defer laneRuns.Done()
		log.Info().
			Str("Test Duration", fmt.Sprintf("%s", duration)).
			Str("Request Triggering interval", fmt.Sprintf("%s", interval)).
			Str("Source", laneA.SourceNetworkName).
			Str("Destination", laneA.DestNetworkName).
			Msg("Starting lane A")
		totalReqLaneA, reqSuccessLaneA = laneA.SoakRun(interval, duration)
	}()
	if laneB != nil {
		laneRuns.Add(1)
		go func() {
			defer laneRuns.Done()
			log.Info().
				Str("Test Duration", fmt.Sprintf("%s", duration)).
				Str("Request Triggering interval", fmt.Sprintf("%s", interval)).
				Str("Source", laneB.SourceNetworkName).
				Str("Destination", laneB.DestNetworkName).
				Msg("Starting lane B")
			totalReqLaneB, reqSuccessLaneB = laneB.SoakRun(interval, duration)
		}()
	}
	laneRuns.Wait()
}

// TestCCIPSoakOnExistingDeployment assumes
// 1. contracts are already deployed on live networks
// 2. CL nodes are set up and configured with existing contracts
// TestCCIPSoakOnExistingDeployment reuses already deployed contracts from the addresses provided in ../contracts/ccip/laneconfig/contracts.json
// This test verifies that CCIP Lanes are working as expected by sending a series of requests and validating their successful delivery
func TestExistingDeploymentSoakCCIP(t *testing.T) {
	var (
		laneA, laneB    *actions.CCIPLane
		totalReqLaneA   = 0
		totalReqLaneB   = 0
		reqSuccessLaneA = 0
		reqSuccessLaneB = 0
	)

	// if interval and duration is provided in env variable, use that
	interval, err := getTestInterval()
	require.NoError(t, err, "invalid interval provided")
	duration, err := getTestDuration()
	require.NoError(t, err, "invalid duration provided")

	t.Cleanup(func() {
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
		if reqSuccessLaneA != totalReqLaneA || reqSuccessLaneB != totalReqLaneB {
			t.Fail()
		}
	})

	transferAmounts := []*big.Int{big.NewInt(1)}
	laneA, laneB = actions.CCIPLaneOnExistingDeployment(
		t, transferAmounts, true,
	)

	if laneA == nil {
		return
	}
	laneRuns := &sync.WaitGroup{}
	laneRuns.Add(1)
	go func() {
		defer laneRuns.Done()
		log.Info().
			Str("Test Duration", fmt.Sprintf("%s", duration)).
			Str("Request Triggering interval", fmt.Sprintf("%s", interval)).
			Str("Source", laneA.SourceNetworkName).
			Str("Destination", laneA.DestNetworkName).
			Msg("Starting lane A")
		laneA.ValidationTimeout = 5 * time.Minute
		totalReqLaneA, reqSuccessLaneA = laneA.SoakRun(interval, duration)
	}()
	if laneB != nil {
		laneRuns.Add(1)
		go func() {
			defer laneRuns.Done()
			log.Info().
				Str("Test Duration", fmt.Sprintf("%s", duration)).
				Str("Request Triggering interval", fmt.Sprintf("%s", interval)).
				Str("Source", laneB.SourceNetworkName).
				Str("Destination", laneB.DestNetworkName).
				Msg("Starting lane B")
			laneB.ValidationTimeout = 5 * time.Minute
			totalReqLaneB, reqSuccessLaneB = laneB.SoakRun(interval, duration)
		}()
	}
	laneRuns.Wait()
}
