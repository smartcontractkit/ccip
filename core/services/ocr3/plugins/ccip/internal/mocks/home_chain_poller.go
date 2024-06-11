package mocks

import (
	"context"

	cciptypes "github.com/smartcontractkit/ccipocr3/ccipocr3-dont-merge"

	"github.com/stretchr/testify/mock"
)

type HomeChainPollerMock struct {
	*mock.Mock
}

func NewHomeChainPollerMock() *HomeChainPollerMock {
	return &HomeChainPollerMock{
		Mock: &mock.Mock{},
	}
}

func (hr *HomeChainPollerMock) Start(ctx context.Context) error {
	args := hr.Called(ctx)
	return args.Error(0)
}

func (hr *HomeChainPollerMock) GetConfig() cciptypes.HomeChainConfig {
	args := hr.Called()
	return args.Get(0).(cciptypes.HomeChainConfig)
}

func (hr *HomeChainPollerMock) Close() error {
	args := hr.Called()
	return args.Error(0)
}
