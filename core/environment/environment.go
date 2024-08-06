package environment

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hashicorp/consul/sdk/freeport"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink/v2/core/environment/memory"
)

const (
	Memory = "memory"
)

type OnchainClient interface {
	// For EVM specifically we can use existing geth interface
	// to abstract chain clients.
	bind.ContractBackend
}

type OffchainClient interface {
	// The job distributor grpc interface can be used to abstract offchain read/writes
	ProposeJob(ctx context.Context, nodeId string, spec string) (int64, error)
	GetJob(ctx context.Context, nodeId string, jobID int64) (string, error)
}

type Chain struct {
	// Selectors used as canonical chain identifier.
	Selector uint64
	Client   OnchainClient
	// Note the Sign function can be abstract supporting a variety of key storage mechanisms (e.g. KMS etc).
	DeployerKey *bind.TransactOpts
	Confirm     func(tx common.Hash) error
}

type Environment struct {
	Name     string
	Chains   map[uint64]Chain
	NodeIds  []string
	Offchain OffchainClient
	Logger   logger.Logger
}

func (e Environment) AllChainSelectors() []uint64 {
	var selectors []uint64
	for sel := range e.Chains {
		selectors = append(selectors, sel)
	}
	return selectors
}

type MemoryEnvironmentConfig struct {
	Chains int
	Nodes  int
}

// To be used by tests and any kind of deployment logic.
func NewMemoryEnvironment(t *testing.T, config MemoryEnvironmentConfig) Environment {
	mchains := memory.GenerateChains(t, config.Chains)
	chains := make(map[uint64]Chain)
	for cid, chain := range mchains {
		sel, err := chainsel.SelectorFromChainId(cid)
		require.NoError(t, err)
		chains[sel] = Chain{
			Selector:    sel,
			Client:      chain.Backend,
			DeployerKey: chain.DeployerKey,
			Confirm: func(tx common.Hash) error {
				chain.Backend.Commit()
				return nil
			},
		}
	}

	nodesByPeerID := make(map[string]memory.Node)
	var keys []string
	ports := freeport.GetN(t, config.Nodes)
	for i := 0; i < config.Nodes; i++ {
		node := memory.NewNode(t, ports[i], mchains, zapcore.DebugLevel)
		nodesByPeerID[node.Keys.PeerID.String()] = *node
		keys = append(keys, node.Keys.PeerID.String())
	}
	lggr, err := logger.New()
	require.NoError(t, err)
	return Environment{
		Name:     Memory,
		Offchain: memory.NewMemoryJobClient(nodesByPeerID),
		Chains:   chains,
		NodeIds:  keys,
		Logger:   lggr,
	}
}
