package mocks

import (
	"context"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query"
	"github.com/stretchr/testify/mock"
)

type HomeChainContractReaderMock struct {
	*mock.Mock
	configs []cciptypes.OnChainCapabilityConfig
}

func NewHomeChainContractReader(configs []cciptypes.OnChainCapabilityConfig) *HomeChainContractReaderMock {
	return &HomeChainContractReaderMock{
		Mock:    &mock.Mock{},
		configs: configs,
	}
}

// GetLatestValue Returns given configs at initialization
func (hr *HomeChainContractReaderMock) GetLatestValue(ctx context.Context, contractName, method string, params, returnVal any) error {
	args := hr.Called(ctx, contractName, method, params, returnVal)
	return args.Error(0)
}

func (hr *HomeChainContractReaderMock) Bind(ctx context.Context, bindings []types.BoundContract) error {
	args := hr.Called(ctx, bindings)
	return args.Error(0)
}

func (hr *HomeChainContractReaderMock) QueryKey(ctx context.Context, contractName string, filter query.KeyFilter, limitAndSort query.LimitAndSort, sequenceDataType any) ([]types.Sequence, error) {
	args := hr.Called(ctx, contractName, filter, limitAndSort, sequenceDataType)
	return args.Get(0).([]types.Sequence), args.Error(1)
}

func (hr *HomeChainContractReaderMock) Start(ctx context.Context) error {
	args := hr.Called(ctx)
	return args.Error(0)
}

func (hr *HomeChainContractReaderMock) Close() error {
	args := hr.Called()
	return args.Error(0)
}

func (hr *HomeChainContractReaderMock) Ready() error {
	args := hr.Called()
	return args.Error(0)
}

func (hr *HomeChainContractReaderMock) HealthReport() map[string]error {
	args := hr.Called()
	return args.Get(0).(map[string]error)
}

func (hr *HomeChainContractReaderMock) Name() string {
	return "HomeChainContractReaderMock"
}
