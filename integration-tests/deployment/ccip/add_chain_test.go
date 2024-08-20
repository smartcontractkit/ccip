package ccipdeployment

import (
	"testing"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
)

func TestAddNewChain(t *testing.T) {
	// Cap reg
	tenv := NewDeployedTestEnvironment(t, logger.TestLogger(t))
	err := AddChain(tenv.Env, tenv.Ab, homeChainSel)
	require.NoError(t, err)
}
