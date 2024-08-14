package migrations

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"

	jobv1 "github.com/smartcontractkit/chainlink/integration-tests/deployment/jd/job/v1"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/router"

	ccipdeployment "github.com/smartcontractkit/chainlink/integration-tests/deployment/ccip"
	"github.com/smartcontractkit/chainlink/integration-tests/deployment/memory"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
)

func Test0001_InitialDeploy(t *testing.T) {
	lggr := logger.TestLogger(t)
	chains := memory.NewMemoryChains(t, 3)
	homeChainSel := uint64(0)
	homeChainEVM := uint64(0)
	// First chain is home chain.
	for chainSel := range chains {
		homeChainEVM, _ = chainsel.ChainIdFromSelector(chainSel)
		homeChainSel = chainSel
		break
	}
	ab, err := ccipdeployment.DeployCapReg(lggr, chains, homeChainSel)
	require.NoError(t, err)

	addrs, err := ab.AddressesForChain(homeChainSel)
	require.NoError(t, err)
	require.Len(t, addrs, 2)
	capReg := common.Address{}
	for addr := range addrs {
		capReg = common.HexToAddress(addr)
		break
	}
	nodes := memory.NewNodes(t, chains, 4, 1, memory.RegistryConfig{
		EVMChainID: homeChainEVM,
		Contract:   capReg,
	})

	e := memory.NewMemoryEnvironmentFromChainsNodes(t, lggr, chains, nodes)
	state, err := ccipdeployment.GenerateOnchainState(e, ab)
	require.NoError(t, err)
	// Apply migration
	output, err := Apply0001(e, ccipdeployment.DeployCCIPContractConfig{
		HomeChainSel: homeChainSel,
		// Capreg/config already exist.
		CCIPOnChainState: state,
	})
	require.NoError(t, err)
	// Get new state after migration.
	state, err = ccipdeployment.GenerateOnchainState(e, output.AddressBook)
	require.NoError(t, err)
	// Replay the log poller on all the chains so that the logs are in the db.
	// otherwise the plugins won't pick them up.
	for _, node := range nodes {
		for sel := range chains {
			chainID, _ := chainsel.ChainIdFromSelector(sel)
			t.Logf("Replaying logs for chain %d from block %d", chainID, 1)
			require.NoError(t, node.App.ReplayFromBlock(big.NewInt(int64(chainID)), 1, false), "failed to replay logs")
		}
	}

	// Apply the jobs.
	for nodeID, jobs := range output.JobSpecs {
		for _, job := range jobs {
			// Note these auto-accept
			_, err := e.Offchain.ProposeJob(context.Background(),
				&jobv1.ProposeJobRequest{
					NodeId: nodeID,
					Spec:   job,
				})
			require.NoError(t, err)
		}
	}
	// Replay the log poller on all the chains so that the logs are in the db.
	// otherwise the plugins won't pick them up.
	for _, node := range nodes {
		for sel := range chains {
			chainID, _ := chainsel.ChainIdFromSelector(sel)
			t.Logf("Replaying logs for chain %d from block %d", chainID, 1)
			require.NoError(t, node.App.ReplayFromBlock(big.NewInt(int64(chainID)), 1, false), "failed to replay logs")
		}
	}
	// Send a request from every router
	for sel, chain := range e.Chains {
		dest := homeChainSel
		if sel == homeChainSel {
			continue
		}
		require.NoError(t, ccipdeployment.AddLane(e, state, sel, dest))
		msg := router.ClientEVM2AnyMessage{
			Receiver:     common.LeftPadBytes(state.Receivers[dest].Address().Bytes(), 32),
			Data:         []byte("hello"),
			TokenAmounts: nil, // TODO: no tokens for now
			//FeeToken:     common.HexToAddress("0x0"),
			FeeToken:  state.Weth9s[sel].Address(),
			ExtraArgs: nil, // TODO: no extra args for now, falls back to default
		}
		fee, err := state.Routers[sel].GetFee(
			&bind.CallOpts{Context: context.Background()}, dest, msg)
		require.NoError(t, err)
		require.NoError(t, err, "%T", err)
		tx, err := state.Weth9s[sel].Deposit(&bind.TransactOpts{
			From:   e.Chains[sel].DeployerKey.From,
			Signer: e.Chains[sel].DeployerKey.Signer,
			Value:  fee,
		})
		require.NoError(t, err)
		require.NoError(t, chain.Confirm(tx.Hash()))
		tx, err = state.Weth9s[sel].Approve(e.Chains[sel].DeployerKey,
			state.Routers[sel].Address(), fee)
		require.NoError(t, err)
		require.NoError(t, chain.Confirm(tx.Hash()))

		t.Logf("Sending CCIP request from chain selector %d to chain selector %d",
			sel, dest)
		_, err = state.Routers[sel].CcipSend(e.Chains[sel].DeployerKey, homeChainSel, msg)
		require.NoError(t, err)
		break
	}
	// TODO: Apply the proposal.
}
