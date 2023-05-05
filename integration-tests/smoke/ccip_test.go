package smoke

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/smartcontractkit/chainlink-testing-framework/utils"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/integration-tests/actions"
	"github.com/smartcontractkit/chainlink/integration-tests/testsetups"
)

func TestSmokeCCIPForBidirectionalLane(t *testing.T) {
	t.Parallel()
	type subtestInput struct {
		testName string
		lane     *actions.CCIPLane
	}
	l := utils.GetTestLogger(t)
	TestCfg := testsetups.NewCCIPTestConfig(t, l, testsetups.Smoke)
	transferAmounts := []*big.Int{big.NewInt(5e17), big.NewInt(5e17)}
	setUpOutput := testsetups.CCIPDefaultTestSetUp(t, l, "smoke-ccip", map[string]interface{}{
		"replicas": "6",
	}, transferAmounts, 5, true, true, TestCfg)
	var tcs []subtestInput
	if len(setUpOutput.Lanes) == 0 {
		return
	}

	t.Cleanup(func() {
		setUpOutput.TearDown()
	})
	for i := range setUpOutput.Lanes {
		tcs = append(tcs, subtestInput{
			testName: fmt.Sprintf("CCIP message transfer from network %s to network %s",
				setUpOutput.Lanes[i].ForwardLane.SourceNetworkName, setUpOutput.Lanes[i].ForwardLane.DestNetworkName),
			lane: setUpOutput.Lanes[i].ForwardLane,
		})
		if setUpOutput.Lanes[i].ReverseLane != nil {
			tcs = append(tcs, subtestInput{
				testName: fmt.Sprintf("CCIP message transfer from network %s to network %s",
					setUpOutput.Lanes[i].ReverseLane.SourceNetworkName, setUpOutput.Lanes[i].ReverseLane.DestNetworkName),
				lane: setUpOutput.Lanes[i].ReverseLane,
			})
		}
	}

	for _, testcase := range tcs {
		tc := testcase
		t.Run(tc.testName, func(t *testing.T) {
			l.Info().
				Str("Source", tc.lane.SourceNetworkName).
				Str("Destination", tc.lane.DestNetworkName).
				Msgf("Starting lane %s -> %s", tc.lane.SourceNetworkName, tc.lane.DestNetworkName)
			tc.lane.RecordStateBeforeTransfer(t)
			_, err := tc.lane.SendRequests(1)
			require.NoError(t, err)
			tc.lane.ValidateRequests()
		})
	}
}
