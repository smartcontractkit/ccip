package ccipdeployment

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
)

func TestAddChain(t *testing.T) {
	// 4 chains where the 4th is added after initial deployment.
	e := NewEnvironmentWithCRAndJobs(t, logger.TestLogger(t), 4)
	state, err := LoadOnchainState(e.Env, e.Ab)
	require.NoError(t, err)
	sels := e.Env.AllChainSelectors()
	initialDeploy := sels[0:3]
	newChain := sels[3]

	ab, err := DeployCCIPContracts(e.Env, DeployCCIPContractConfig{
		HomeChainSel:     e.HomeChainSel,
		ChainsToDeploy:   initialDeploy,
		CCIPOnChainState: state,
	})
	require.NoError(t, err)
	require.NoError(t, e.Ab.Merge(ab))
	state, err = LoadOnchainState(e.Env, e.Ab)
	require.NoError(t, err)

	// Contracts deployed and initial DONs set up.
	// Connect all the lanes
	for _, source := range initialDeploy {
		for _, dest := range initialDeploy {
			if source != dest {
				require.NoError(t, AddLane(e.Env, state, uint64(source), uint64(dest)))
			}
		}
	}

	// Enable inbound to new 4th chain.
	proposals, ab, err := NewChainInbound(e.Env, e.Ab, e.HomeChainSel, newChain, initialDeploy)
	require.NoError(t, err)
	t.Log(proposals, ab)
}
