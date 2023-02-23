package soak

//revive:disable:dot-imports
import (
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/integration-tests/actions"
)

var (
	interval = 10 * time.Second
	duration = 2 * time.Minute
)

// TestCCIPSoak verifies that CCIP requests can be successfully delivered for mentioned duration triggered at a certain interval
// If run on live networks it can reuse already deployed contracts if the addresses are provided in ../contracts/ccip/laneconfig/contracts.json
// This test does a full environment set up along with deploying CL nodes in K8 cluster
func TestCCIPSoak(t *testing.T) {
	var (
		tearDown        func()
		laneA, laneB    *actions.CCIPLane
		totalReqLaneA   = 0
		totalReqLaneB   = 0
		reqSuccessLaneA = 0
		reqSuccessLaneB = 0
	)

	t.Cleanup(func() {
		if tearDown != nil {
			log.Info().Msg("Tearing down the environment")
			tearDown()
			log.Info().
				Str("total duration", fmt.Sprint(duration)).
				Str("req interval", fmt.Sprint(interval)).
				Int("Total Requests", totalReqLaneA).
				Int("Successful Requests", reqSuccessLaneA).
				Msgf("Soak Result for lane %s --> %s", laneA.SourceNetworkName, laneA.DestNetworkName)
			log.Info().
				Str("total duration", fmt.Sprint(duration)).
				Str("req interval", fmt.Sprint(interval)).
				Int("Total Requests", totalReqLaneB).
				Int("Successful Requests", reqSuccessLaneB).
				Msgf("Soak Result for lane %s --> %s", laneB.SourceNetworkName, laneB.DestNetworkName)
		}
	})

	transferAmounts := []*big.Int{big.NewInt(5e17), big.NewInt(5e17)}
	laneA, laneB, tearDown = actions.CCIPDefaultTestSetUp(
		t, "soak-ccip",
		map[string]interface{}{
			"replicas": "6",
			"toml":     actions.DefaultCCIPCLNodeEnv(t),
		},
		transferAmounts, 5, true, true,
	)

	if laneA == nil {
		return
	}

	require.NoError(t, laneA.IsLaneDeployed())
	require.NoError(t, laneB.IsLaneDeployed())

	t.Run(fmt.Sprintf("CCIP message transfer from network %s to network %s for %s", laneA.SourceNetworkName, laneA.DestNetworkName, duration), func(t *testing.T) {
		t.Parallel()
		totalReqLaneA, reqSuccessLaneA = laneA.SoakRun(interval, duration)
	})

	t.Run(fmt.Sprintf("CCIP message transfer from network %s to network %s for %s", laneB.SourceNetworkName, laneB.DestNetworkName, duration), func(t *testing.T) {
		t.Parallel()
		totalReqLaneB, reqSuccessLaneB = laneB.SoakRun(interval, duration)
	})
}

// TestCCIPSoakOnExistingDeployment assumes
// 1. contracts are already deployed on live networks
// 2. CL nodes are set up and configured with existing contracts
// TestCCIPSoakOnExistingDeployment reuses already deployed contracts from the addresses provided in ../contracts/ccip/laneconfig/contracts.json
// This test verifies that CCIP Lanes are working as expected by sending a series of requests and validating their successful delivery
func TestCCIPSoakOnExistingDeployment(t *testing.T) {
	var (
		laneA, laneB    *actions.CCIPLane
		totalReqLaneA   = 0
		totalReqLaneB   = 0
		reqSuccessLaneA = 0
		reqSuccessLaneB = 0
		interval        = 30 * time.Second
		duration        = 1 * time.Minute
	)

	t.Cleanup(func() {
		log.Info().
			Str("total duration", fmt.Sprint(duration)).
			Str("req interval", fmt.Sprint(interval)).
			Int("Total Requests", totalReqLaneA).
			Int("Successful Requests", reqSuccessLaneA).
			Msgf("Soak Result for lane %s --> %s", laneA.SourceNetworkName, laneA.DestNetworkName)
		if laneB != nil {
			log.Info().
				Str("total duration", fmt.Sprint(duration)).
				Str("req interval", fmt.Sprint(interval)).
				Int("Total Requests", totalReqLaneB).
				Int("Successful Requests", reqSuccessLaneB).
				Msgf("Soak Result for lane %s --> %s", laneB.SourceNetworkName, laneB.DestNetworkName)
		}
	})

	transferAmounts := []*big.Int{big.NewInt(1)}
	laneA, laneB = actions.CCIPLaneOnExistingDeployment(
		t, transferAmounts, true,
	)

	if laneA == nil {
		return
	}

	require.NoError(t, laneA.IsLaneDeployed())
	if laneB != nil {
		require.NoError(t, laneB.IsLaneDeployed())
	}

	t.Run(fmt.Sprintf("CCIP message transfer from network %s to network %s for %s", laneA.SourceNetworkName, laneA.DestNetworkName, duration), func(t *testing.T) {
		t.Parallel()
		totalReqLaneA, reqSuccessLaneA = laneA.SoakRun(interval, duration)
	})

	if laneB != nil {
		t.Run(fmt.Sprintf("CCIP message transfer from network %s to network %s for %s", laneB.SourceNetworkName, laneB.DestNetworkName, duration), func(t *testing.T) {
			t.Parallel()
			totalReqLaneB, reqSuccessLaneB = laneB.SoakRun(interval, duration)
		})
	}
}
