package ccipdeployment

import (
	"context"
	"testing"
	"time"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"

	jobv1 "github.com/smartcontractkit/chainlink/integration-tests/deployment/jd/job/v1"

	"github.com/smartcontractkit/chainlink/integration-tests/deployment"
	"github.com/smartcontractkit/chainlink/integration-tests/deployment/memory"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
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

type DeployedTestEnvironment struct {
	Ab           deployment.AddressBook
	Env          deployment.Environment
	HomeChainSel uint64
	Nodes        map[string]memory.Node
}

// NewDeployedEnvironment creates a new CCIP environment
// with capreg and nodes set up.
func NewEnvironmentWithCR(t *testing.T, lggr logger.Logger, numChains int) DeployedTestEnvironment {
	ctx := Context(t)
	chains := memory.NewMemoryChains(t, numChains)
	homeChainSel := uint64(0)
	homeChainEVM := uint64(0)
	// Say first chain is home chain.
	for chainSel := range chains {
		homeChainEVM, _ = chainsel.ChainIdFromSelector(chainSel)
		homeChainSel = chainSel
		break
	}
	ab, capReg, err := DeployCapReg(lggr, chains, homeChainSel)
	require.NoError(t, err)

	nodes := memory.NewNodes(t, zapcore.InfoLevel, chains, 4, 1, memory.RegistryConfig{
		EVMChainID: homeChainEVM,
		Contract:   capReg,
	})
	for _, node := range nodes {
		require.NoError(t, node.App.Start(ctx))
	}

	e := memory.NewMemoryEnvironmentFromChainsNodes(t, lggr, chains, nodes)
	return DeployedTestEnvironment{
		Ab:           ab,
		Env:          e,
		HomeChainSel: homeChainSel,
		Nodes:        nodes,
	}
}

func NewEnvironmentWithCRAndJobs(t *testing.T, lggr logger.Logger, numChains int) DeployedTestEnvironment {
	ctx := Context(t)
	e := NewEnvironmentWithCR(t, lggr, numChains)
	jbs, err := NewCCIPJobSpecs(e.Env.NodeIDs, e.Env.Offchain)
	require.NoError(t, err)
	for nodeID, jobs := range jbs {
		for _, job := range jobs {
			// Note these auto-accept
			_, err := e.Env.Offchain.ProposeJob(ctx,
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
	require.NoError(t, ReplayAllLogs(e.Nodes, e.Env.Chains))
	return e
}

func ReplayAllLogs(nodes map[string]memory.Node, chains map[uint64]deployment.Chain) error {
	for _, node := range nodes {
		for sel := range chains {
			if err := node.ReplayLogs(map[uint64]uint64{sel: 1}); err != nil {
				return err
			}
		}
	}
	return nil
}
