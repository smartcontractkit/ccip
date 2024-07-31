package environment

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hashicorp/consul/sdk/freeport"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink/v2/core/environment/memory"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"
	"testing"
)

type AddressBook interface {
	Save(chainSelector uint64, address string) error
	Addresses() (map[uint64]map[string]struct{}, error)
}

type OnchainClient interface {
	// For EVM specifically we can use existing geth interface
	// to abstract chain clients.
	bind.ContractBackend
}

type OffchainClient interface {
	// The job distributor grpc interface can be used to abstract offchain read/writes
	ProposeJob(ctx context.Context, nodeId string, spec string) (string, error)
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
	Chains      map[uint64]Chain
	NodeIds     []string
	AddressBook AddressBook
	Offchain    OffchainClient
	Logger      logger.Logger
}

// To be used by tests and any kind of deployment logic.
func NewMemoryEnvironment(t *testing.T) Environment {
	mchains := memory.GenerateChains(t, 3)
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
	ports := freeport.GetN(t, 1)
	node := memory.NewNode(t, ports[0], mchains, zapcore.DebugLevel)
	lggr, err := logger.New()
	if err != nil {
		panic(err)
	}
	return Environment{
		Offchain:    memory.NewMemoryJobClient(node.App),
		Chains:      chains,
		NodeIds:     []string{node.Keys.PeerID.String()},
		AddressBook: memory.NewMemoryAddressBook(),
		Logger:      lggr,
	}
}
