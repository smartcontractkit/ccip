package smoke

//revive:disable:dot-imports
import (
	"fmt"
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
		if tearDown != nil {
			log.Info().Msg("Tearing down the environment")
			tearDown()
		}
	})

	transferAmounts := []*big.Int{big.NewInt(5e17), big.NewInt(5e17)}
	laneA, laneB, tearDown = actions.CCIPDefaultTestSetUp(t, fmt.Sprintf("smoke-ccip-%s-%s", actions.NetworkA.Name, actions.NetworkB.Name),
		map[string]interface{}{
			"replicas": "6",
			"toml":     actions.DefaultCCIPCLNodeEnv(t),
		}, transferAmounts, 5, true, bidirectional)

	if laneA == nil {
		return
	}

	t.Run("CCIP message transfer in bi-directional lane", func(t *testing.T) {
		// initiate transfer with GE and verify
		log.Info().Msg("Multiple Token transfer with GE from laneA to laneB")
		require.NoError(t, laneA.IsLaneDeployed())
		laneA.SendGERequests(1)
		laneA.ValidateGERequests()

		// initiate transfer with toll and verify
		log.Info().Msg("Multiple Token transfer with toll from laneA to laneB")
		laneA.SendTollRequests(1)
		laneA.ValidateTollRequests()

		if !bidirectional {
			log.Info().Msg("Skipping validating reverse lane")
			return
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
