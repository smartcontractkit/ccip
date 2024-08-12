package deployment

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	jobv1 "github.com/smartcontractkit/chainlink/integration-tests/deployment/jd/job/v1"
	nodev1 "github.com/smartcontractkit/chainlink/integration-tests/deployment/jd/node/v1"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

type OnchainClient interface {
	// For EVM specifically we can use existing geth interface
	// to abstract chain clients.
	bind.ContractBackend
}

type OffchainClient interface {
	// The job distributor grpc interface can be used to abstract offchain read/writes
	jobv1.JobServiceClient
	nodev1.NodeServiceClient
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
	Offchain OffchainClient
	NodeIDs  []string
	Logger   logger.Logger
}

func (e Environment) AllChainSelectors() []uint64 {
	var selectors []uint64
	for sel := range e.Chains {
		selectors = append(selectors, sel)
	}
	return selectors
}
