package mocks

import (
	"context"
	"time"

	cciptypes "github.com/smartcontractkit/ccipocr3/ccipocr3-dont-merge"

	"github.com/stretchr/testify/mock"
)

type HomeChainPollerMock struct {
	*mock.Mock
	config cciptypes.HomeChainConfig
}

func NewHomeChainPollerMock(config cciptypes.HomeChainConfig) *HomeChainPollerMock {
	return &HomeChainPollerMock{
		Mock:   &mock.Mock{},
		config: config,
	}
}

func (hr *HomeChainPollerMock) StartPolling(ctx context.Context, interval time.Duration) {
}

func (hr *HomeChainPollerMock) GetConfig() cciptypes.HomeChainConfig {
	return hr.config
}

func (hr *HomeChainPollerMock) Close(ctx context.Context) error {
	return nil
}
