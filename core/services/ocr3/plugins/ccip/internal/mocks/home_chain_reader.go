package mocks

import "C"
import (
	"context"
	"fmt"

	cciptypes "github.com/smartcontractkit/ccipocr3/ccipocr3-dont-merge"

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
	if val, ok := returnVal.(*[]cciptypes.OnChainCapabilityConfig); ok {
		*val = hr.configs
	} else {
		return fmt.Errorf("unexpected type for returnVal")
	}
	return nil
}

func (*HomeChainContractReaderMock) Bind(ctx context.Context, bindings []types.BoundContract) error {
	return nil
}

func (*HomeChainContractReaderMock) QueryKey(ctx context.Context, contractName string, filter query.KeyFilter, limitAndSort query.LimitAndSort, sequenceDataType any) ([]types.Sequence, error) {
	return nil, nil
}

func (hr *HomeChainContractReaderMock) Start(ctx context.Context) error {
	return nil
}

func (hr *HomeChainContractReaderMock) Close() error {
	return nil
}

func (hr *HomeChainContractReaderMock) Ready() error {
	return nil
}

func (hr *HomeChainContractReaderMock) HealthReport() map[string]error {
	return nil
}

func (hr *HomeChainContractReaderMock) Name() string {
	return "HomeChainContractReaderMock"
}
