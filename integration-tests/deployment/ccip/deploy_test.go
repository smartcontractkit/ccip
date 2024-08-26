package ccipdeployment

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"

	"github.com/smartcontractkit/chainlink/integration-tests/deployment"
	"github.com/smartcontractkit/chainlink/integration-tests/deployment/memory"
	"github.com/smartcontractkit/chainlink/integration-tests/deployment/persistent"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
)

func TestDeployCCIPContractsInMemory(t *testing.T) {
	lggr := logger.TestLogger(t)
	e := memory.NewMemoryEnvironment(t, lggr, zapcore.InfoLevel, memory.MemoryEnvironmentConfig{
		Bootstraps: 1,
		Chains:     1,
		Nodes:      4,
	})
	testDeployCCIPContractsWithEnv(t, lggr, e)
}

func TestDeployCCIPContractsPersistent(t *testing.T) {
	lggr := logger.TestLogger(t)
	e := persistent.NewPersistentEnvironment(t, lggr)
	testDeployCCIPContractsWithEnv(t, lggr, e)
}

func testDeployCCIPContractsWithEnv(t *testing.T, lggr logger.Logger, e deployment.Environment) {
	var ab deployment.AddressBook
	// Deploy all the CCIP contracts.
	for _, chain := range e.AllChainSelectors() {
		capRegAddresses, _, err := DeployCapReg(lggr, e.Chains, chain)
		require.NoError(t, err)
		s, err := LoadOnchainState(e, capRegAddresses)
		require.NoError(t, err)
		newAb, err := DeployCCIPContracts(e, DeployCCIPContractConfig{
			HomeChainSel:     chain,
			CCIPOnChainState: s,
		})
		require.NoError(t, err)
		if ab == nil {
			ab = newAb
		} else {
			mergeErr := ab.Merge(newAb)
			require.NoError(t, mergeErr)
		}
	}

	state, err := LoadOnchainState(e, ab)
	require.NoError(t, err)
	snap, err := state.Snapshot(e.AllChainSelectors())
	require.NoError(t, err)

	// Assert expect every deployed address to be in the address book.
	// TODO (CCIP-3047): Add the rest of CCIPv2 representation
	b, err := json.MarshalIndent(snap, "", "	")
	require.NoError(t, err)
	fmt.Println(string(b))
}

func TestJobSpecGeneration(t *testing.T) {
	lggr := logger.TestLogger(t)
	e := memory.NewMemoryEnvironment(t, lggr, zapcore.InfoLevel, memory.MemoryEnvironmentConfig{
		Chains: 1,
		Nodes:  1,
	})
	js, err := NewCCIPJobSpecs(e.NodeIDs, e.Offchain)
	require.NoError(t, err)
	for node, jb := range js {
		fmt.Println(node, jb)
	}
	// TODO: Add job assertions
}
