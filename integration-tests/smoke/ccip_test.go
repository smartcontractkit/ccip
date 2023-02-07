package smoke

//revive:disable:dot-imports
import (
	"math/big"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/integration-tests/actions"
)

func TestSmokeCCIPForBidirectionalLane(t *testing.T) {
	t.Parallel()
	transferAmounts := []*big.Int{big.NewInt(5e17), big.NewInt(5e17)}
	laneA, laneB, tearDown := actions.CCIPDefaultTestSetUp(
		t, "smoke-ccip",
		map[string]interface{}{
			"replicas": "6",
			"toml":     actions.DefaultCCIPCLNodeEnv(t),
		},
		transferAmounts,
		5,
		true,
		true,
	)
	if laneA == nil {
		return
	}
	t.Cleanup(func() {
		log.Info().Msg("Tearing down the environment")
		tearDown()
	})

	require.NoError(t, laneA.IsLaneDeployed())
	if laneB != nil {
		require.NoError(t, laneB.IsLaneDeployed())
	}

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
