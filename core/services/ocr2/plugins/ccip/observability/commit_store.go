package observability

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/commit_store"
)

type ObservedCommitStore struct {
	commit_store.CommitStoreInterface
	metric metricDetails
}

func NewObservedCommitStore(address common.Address, pluginName string, backend bind.ContractBackend) (commit_store.CommitStoreInterface, error) {
	commitStore, err := commit_store.NewCommitStore(address, backend)
	if err != nil {
		return nil, err
	}
	return &ObservedCommitStore{
		CommitStoreInterface: commitStore,
		metric: metricDetails{
			histogram:  commitStoreHistogram,
			pluginName: pluginName,
		},
	}, nil
}

func (o *ObservedCommitStore) GetStaticConfig(opts *bind.CallOpts) (commit_store.CommitStoreStaticConfig, error) {
	return withObservedContract(o.metric, "GetStaticConfig", func() (commit_store.CommitStoreStaticConfig, error) {
		return o.CommitStoreInterface.GetStaticConfig(opts)
	})
}

func (o *ObservedCommitStore) GetExpectedNextSequenceNumber(opts *bind.CallOpts) (uint64, error) {
	return withObservedContract(o.metric, "GetExpectedNextSequenceNumber", func() (uint64, error) {
		return o.CommitStoreInterface.GetExpectedNextSequenceNumber(opts)
	})
}

func (o *ObservedCommitStore) Paused(opts *bind.CallOpts) (bool, error) {
	return withObservedContract(o.metric, "Paused", func() (bool, error) {
		return o.CommitStoreInterface.Paused(opts)
	})
}

func (o *ObservedCommitStore) IsAFNHealthy(opts *bind.CallOpts) (bool, error) {
	return withObservedContract(o.metric, "IsAFNHealthy", func() (bool, error) {
		return o.CommitStoreInterface.IsAFNHealthy(opts)
	})
}

func (o *ObservedCommitStore) IsBlessed(opts *bind.CallOpts, root [32]byte) (bool, error) {
	return withObservedContract(o.metric, "IsBlessed", func() (bool, error) {
		return o.CommitStoreInterface.IsBlessed(opts, root)
	})
}
