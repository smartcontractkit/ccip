package mocks

import (
	"context"

	mapset "github.com/deckarep/golang-set/v2"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/libocr/ragep2p/types"
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

func (hr *HomeChainPollerMock) GetChainConfig(chainSelector cciptypes.ChainSelector) (cciptypes.ChainConfig, error) {
	args := hr.Called(chainSelector)
	return args.Get(0).(cciptypes.ChainConfig), args.Error(0)
}

func (hr *HomeChainPollerMock) GetAllChainConfigs() (map[cciptypes.ChainSelector]cciptypes.ChainConfig, error) {
	args := hr.Called()
	return args.Get(0).(map[cciptypes.ChainSelector]cciptypes.ChainConfig), args.Error(0)
}

func (hr *HomeChainPollerMock) GetSupportedChains(id types.PeerID) mapset.Set[cciptypes.ChainSelector] {
	args := hr.Called()
	return args.Get(0).(mapset.Set[cciptypes.ChainSelector])
}

func (hr *HomeChainPollerMock) GetKnownChains() mapset.Set[cciptypes.ChainSelector] {
	args := hr.Called()
	return args.Get(0).(mapset.Set[cciptypes.ChainSelector])
}

func (hr *HomeChainPollerMock) GetFChain() map[cciptypes.ChainSelector]int {
	args := hr.Called()
	return args.Get(0).(map[cciptypes.ChainSelector]int)
}

func (hr *HomeChainPollerMock) Ready() error {
	args := hr.Called()
	return args.Error(0)
}

func (hr *HomeChainPollerMock) HealthReport() map[string]error {
	args := hr.Called()
	return args.Get(0).(map[string]error)
}

func (hr *HomeChainPollerMock) Name() string {
	args := hr.Called()
	return args.Get(0).(string)
}

func (hr *HomeChainPollerMock) Close() error {
	args := hr.Called()
	return args.Error(0)
}

// Interface compatibility check.
var _ cciptypes.HomeChainPoller = (*HomeChainPollerMock)(nil)
