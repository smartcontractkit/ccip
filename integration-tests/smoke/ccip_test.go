package smoke

//revive:disable:dot-imports
import (
	"math/big"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/integration-tests/actions"
)

func TestSmokeCCIPForBidirectionalLaneGE(t *testing.T) {
	t.Parallel()
	transferAmounts := []*big.Int{big.NewInt(5e17), big.NewInt(5e17)}
	laneA, laneB, tearDown := actions.CCIPDefaultTestSetUp(
		t, "smoke-ccip",
		map[string]interface{}{
			"replicas": "6",
			"toml":     actions.DefaultCCIPCLNodeEnv(t),
		}, transferAmounts, 5, true, true, actions.GE)
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

	// initiate transfer with GE and verify
	log.Info().Msgf("Multiple Token transfer with GE for lane %s --> %s", laneA.SourceNetworkName, laneA.DestNetworkName)
	laneA.RecordStateBeforeGETransfer()
	laneA.SendGERequests(1)
	laneA.ValidateGERequests()

	if laneB == nil {
		return
	}

	log.Info().Msgf("Multiple Token transfer with GE for lane %s --> %s", laneB.SourceNetworkName, laneB.DestNetworkName)
	laneB.RecordStateBeforeGETransfer()
	laneB.SendGERequests(1)
	laneB.ValidateGERequests()
}

func TestSmokeCCIPForBidirectionalLaneToll(t *testing.T) {
	t.Parallel()

	transferAmounts := []*big.Int{big.NewInt(5e17), big.NewInt(5e17)}
	laneA, laneB, tearDown := actions.CCIPDefaultTestSetUp(
		t, "smoke-ccip",
		map[string]interface{}{
			"replicas": "6",
			"toml":     actions.DefaultCCIPCLNodeEnv(t),
		}, transferAmounts, 5, true, true, actions.TOLL)
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
	// initiate transfer with toll and verify
	log.Info().Msgf("Multiple Token transfer with toll for lane %s --> %s", laneA.SourceNetworkName, laneA.DestNetworkName)
	laneA.RecordStateBeforeTollTransfer()
	laneA.SendTollRequests(1)
	laneA.ValidateTollRequests()

	if laneB == nil {
		return
	}

	log.Info().Msgf("Multiple Token transfer with toll for lane %s --> %s", laneB.SourceNetworkName, laneB.DestNetworkName)
	laneB.RecordStateBeforeTollTransfer()
	laneB.SendTollRequests(1)
	laneB.ValidateTollRequests()
}
