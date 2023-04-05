package smoke

import (
	"math/big"
	"testing"

	"github.com/rs/zerolog/log"

	"github.com/smartcontractkit/chainlink/integration-tests/testsetups"
)

func TestSmokeCCIPForBidirectionalLane(t *testing.T) {
	t.Parallel()

	TestCfg := testsetups.NewCCIPTestConfig(t, testsetups.Smoke)
	transferAmounts := []*big.Int{big.NewInt(5e17), big.NewInt(5e17)}
	setUpOutput := testsetups.CCIPDefaultTestSetUp(t, "smoke-ccip", map[string]interface{}{
		"replicas": "6",
	}, transferAmounts, 5, true, true, TestCfg)

	if len(setUpOutput.Lanes) == 0 {
		return
	}
	laneA := setUpOutput.Lanes[0].ForwardLane
	laneB := setUpOutput.Lanes[0].ReverseLane
	t.Cleanup(func() {
		log.Info().Msg("Tearing down the environment")
		setUpOutput.TearDown()
	})

	// initiate transfer and verify
	log.Info().Msgf("Multiple Token transfer for lane %s --> %s", laneA.SourceNetworkName, laneA.DestNetworkName)
	laneA.RecordStateBeforeTransfer()
	laneA.SendRequests(1)
	laneA.ValidateRequests()

	if laneB == nil {
		return
	}

	log.Info().Msgf("Multiple Token transfer for lane %s --> %s", laneB.SourceNetworkName, laneB.DestNetworkName)
	laneB.RecordStateBeforeTransfer()
	laneB.SendRequests(1)
	laneB.ValidateRequests()
}
