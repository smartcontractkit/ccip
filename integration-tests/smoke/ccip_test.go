package smoke

//revive:disable:dot-imports
import (
	"math/big"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/integration-tests/actions"
)

var bidirectional = true

func TestSmokeCCIPForBidirectionalLane(t *testing.T) {
	var (
		tearDown     func()
		laneA, laneB *actions.CCIPLane
	)

	t.Cleanup(func() {
		log.Info().Msg("Tearing down the environment")
		tearDown()
	})

	transferAmounts := []*big.Int{big.NewInt(5e17), big.NewInt(5e17)}
	laneA, laneB, tearDown = actions.CCIPDefaultTestSetUp(t, "smoke-ccip", map[string]interface{}{
		"replicas": "6",
		"toml":     actions.DefaultCCIPCLNodeEnv(t),
		"env": map[string]interface{}{
			"CL_DEV": "true",
		},
	}, transferAmounts, 5, true, bidirectional)

	t.Run("CCIP message transfer from laneA to laneB", func(t *testing.T) {
		// initiate transfer with GE and verify
		log.Info().Msg("Multiple Token transfer with GE from laneA to laneB")
		require.NoError(t, laneA.IsLaneDeployed())
		laneA.SendGERequests(1)
		laneA.ValidateGERequests()

		// initiate transfer with toll and verify
		log.Info().Msg("Multiple Token transfer with toll from laneA to laneB")
		laneA.SendTollRequests(1)
		laneA.ValidateTollRequests()
	})

	t.Run("CCIP message transfer from laneB to laneA", func(t *testing.T) {
		if !bidirectional {
			t.Skip("Skipping validating reverse lane")
		}

		// initiate transfer with GE and verify
		log.Info().Msg("Multiple Token transfer with GE from laneB to laneA")
		require.NoError(t, laneB.IsLaneDeployed())
		laneB.SendGERequests(1)
		laneB.ValidateGERequests()

		// initiate transfer with toll and verify
		log.Info().Msg("Multiple Token transfer with toll from laneB to laneA")
		laneB.SendTollRequests(1)
		laneB.ValidateTollRequests()
	})
}
