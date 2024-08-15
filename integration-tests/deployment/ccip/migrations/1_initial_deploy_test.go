package migrations

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/integration-tests/deployment"

	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	jobv1 "github.com/smartcontractkit/chainlink/integration-tests/deployment/jd/job/v1"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_multi_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/router"

	ccipdeployment "github.com/smartcontractkit/chainlink/integration-tests/deployment/ccip"
	"github.com/smartcontractkit/chainlink/integration-tests/deployment/memory"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
)

// Context returns a context with the test's deadline, if available.
func Context(tb testing.TB) context.Context {
	ctx := context.Background()
	var cancel func()
	switch t := tb.(type) {
	case *testing.T:
		if d, ok := t.Deadline(); ok {
			ctx, cancel = context.WithDeadline(ctx, d)
		}
	}
	if cancel == nil {
		ctx, cancel = context.WithCancel(ctx)
	}
	tb.Cleanup(cancel)
	return ctx
}

func Test0001_InitialDeploy(t *testing.T) {
	lggr := logger.TestLogger(t)
	ctx := Context(t)
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
	for _, node := range nodes {
		require.NoError(t, node.App.Start(ctx))
	}

	e := memory.NewMemoryEnvironmentFromChainsNodes(t, lggr, chains, nodes)
	state, err := ccipdeployment.GenerateOnchainState(e, ab)
	require.NoError(t, err)

	capabilities, err := state.CapabilityRegistry[homeChainSel].GetCapabilities(nil)
	require.NoError(t, err)
	require.Len(t, capabilities, 1)
	ccipCap, err := state.CapabilityRegistry[homeChainSel].GetHashedCapabilityId(nil,
		ccipdeployment.CapabilityLabelledName, ccipdeployment.CapabilityVersion)
	require.NoError(t, err)
	_, err = state.CapabilityRegistry[homeChainSel].GetCapability(nil, ccipCap)
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

	// Ensure capreg logs are up to date.
	ReplayAllLogs(nodes, chains)

	// Apply the jobs.
	for nodeID, jobs := range output.JobSpecs {
		for _, job := range jobs {
			// Note these auto-accept
			_, err := e.Offchain.ProposeJob(ctx,
				&jobv1.ProposeJobRequest{
					NodeId: nodeID,
					Spec:   job,
				})
			require.NoError(t, err)
		}
	}
	// Wait for plugins to register filters?
	// TODO: Investigate how to avoid.
	time.Sleep(30 * time.Second)

	// Ensure job related logs are up to date.
	ReplayAllLogs(nodes, chains)

	// Send a request from every router
	// Add all lanes
	for source := range e.Chains {
		for dest := range e.Chains {
			if source != dest {
				require.NoError(t, ccipdeployment.AddLane(e, state, source, dest))
			}
		}
	}
	for sel, chain := range e.Chains {
		dest := homeChainSel
		if sel == homeChainSel {
			continue
		}
		//require.NoError(t, ccipdeployment.AddLane(e, state, sel, dest))
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
		require.NoError(t, err, ccipdeployment.MaybeDataErr(err))
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
		tx, err = state.Routers[sel].CcipSend(e.Chains[sel].DeployerKey, homeChainSel, msg)
		require.NoError(t, err)
		require.NoError(t, chain.Confirm(tx.Hash()))
		waitForCommitWithInterval(t, chain, e.Chains[homeChainSel],
			state.EvmOffRampsV160[homeChainSel],
			ccipocr3.SeqNumRange{1, 1},
		)
	}
	// TODO: Apply the proposal.
}

func ReplayAllLogs(nodes map[string]memory.Node, chains map[uint64]deployment.Chain) {
	for _, node := range nodes {
		for sel := range chains {
			chainID, _ := chainsel.ChainIdFromSelector(sel)
			node.App.ReplayFromBlock(big.NewInt(int64(chainID)), 1, false)
		}
	}
}

func waitForCommitWithInterval(
	t *testing.T,
	src deployment.Chain,
	dest deployment.Chain,
	offRamp *evm_2_evm_multi_offramp.EVM2EVMMultiOffRamp,
	expectedSeqNumRange ccipocr3.SeqNumRange,
) {
	sink := make(chan *evm_2_evm_multi_offramp.EVM2EVMMultiOffRampCommitReportAccepted)
	subscription, err := offRamp.WatchCommitReportAccepted(&bind.WatchOpts{
		Context: context.Background(),
	}, sink)
	require.NoError(t, err)
	ticker := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-ticker.C:
			src.Client.(*backends.SimulatedBackend).Commit()
			dest.Client.(*backends.SimulatedBackend).Commit()
		case <-time.After(time.Minute):
			t.Logf("Waiting for commit report on chain selector %d from source selector %d expected seq nr range %s",
				dest.Selector, src.Selector, expectedSeqNumRange.String())
			t.Error("Timed out waiting for commit report")
			return
		case subErr := <-subscription.Err():
			t.Fatalf("Subscription error: %+v", subErr)
		case report := <-sink:
			if len(report.Report.MerkleRoots) > 0 {
				// Check the interval of sequence numbers and make sure it matches
				// the expected range.
				for _, mr := range report.Report.MerkleRoots {
					if mr.SourceChainSelector == src.Selector &&
						uint64(expectedSeqNumRange.Start()) == mr.Interval.Min &&
						uint64(expectedSeqNumRange.End()) == mr.Interval.Max {
						t.Logf("Received commit report on selector %d from source selector %d expected seq nr range %s",
							dest.Selector, src.Selector, expectedSeqNumRange.String())
						return
					}
				}
			}
		}
	}
}
