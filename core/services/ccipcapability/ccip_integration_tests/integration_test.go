package ccip_integration_tests

import (
	"fmt"
	"github.com/smartcontractkit/chainlink/v2/core/environment"
	"github.com/smartcontractkit/chainlink/v2/core/services/ccipcapability/deployment"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCCIP(t *testing.T) {
	// Memory specific assertions against deployment configuration logic
	// We aim to ensure deployment.X (maybe better name) functions are bulletproof
	// for when they run against real environments.
	e := environment.NewMemoryEnvironment(t)
	require.NoError(t, deployment.DeployCCIPContracts(e))
	state, err := deployment.GenerateOnchainState(e)
	require.NoError(t, err)
	fmt.Println(state.TokenAdminRegistries)
}
