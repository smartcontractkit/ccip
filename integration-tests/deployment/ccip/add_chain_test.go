package ccipdeployment

import (
	"encoding/json"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/integration-tests/deployment/executable"

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

	executorClients := make(map[string]executable.ContractDeployBackend)
	for _, chain := range e.Env.Chains {
		chainselc, exists := chainsel.ChainBySelector(chain.Selector)
		require.True(t, exists)
		executorClients[chainselc.Name] = chain.Client
	}

	// Enable inbound to new 4th chain.
	proposals, ab, err := NewChainInbound(e.Env, e.Ab, e.HomeChainSel, newChain, initialDeploy)
	require.NoError(t, err)
	require.Equal(t, 3, len(proposals[0].ChainMetadata))
	// Sign this proposal with the deployer key.
	execProposal, err := proposals[0].ToExecutableMCMSProposal()
	require.NoError(t, err)
	require.Equal(t, 4, len(execProposal.ExecutableMCMSProposalBase.ChainMetadata))
	exec, err := execProposal.ToExecutor(executorClients)
	require.NoError(t, err)
	payload, err := exec.SigningHash()
	require.NoError(t, err)
	// Sign the payload
	sig, err := crypto.Sign(payload.Bytes(), TestXXXMCMSSigner)
	require.NoError(t, err)

	// Sign the payload
	unmarshalledSig := executable.Signature{}
	err = json.Unmarshal(sig, &unmarshalledSig)
	require.NoError(t, err)

	// Add signature to proposal
	proposals[0].AddSignature(unmarshalledSig)
	require.NoError(t, exec.Execute())

	t.Log(proposals, ab)
}
