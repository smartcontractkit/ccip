package mocks

import (
	"context"
	"time"

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

func (hr *HomeChainPollerMock) StartPolling(ctx context.Context, interval time.Duration) {
	hr.Called(ctx, interval)
}

func (hr *HomeChainPollerMock) GetConfig() cciptypes.HomeChainConfig {
	args := hr.Called()
	return args.Get(0).(cciptypes.HomeChainConfig)
}

func (hr *HomeChainPollerMock) Close(ctx context.Context) error {
	args := hr.Called(ctx)
	return args.Error(0)
}
