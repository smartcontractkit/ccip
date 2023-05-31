package ccip

import (
	"context"
	"sync"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	mocklp "github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller/mocks"
	evmmocks "github.com/smartcontractkit/chainlink/v2/core/chains/evm/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/commit_store"
	mock_contracts "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/services/job"
	ccipconfig "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config"
)

func TestGetCommitPluginFilterNamesFromSpec(t *testing.T) {
	testCases := []struct {
		description   string
		spec          *job.OCR2OracleSpec
		expectedNames []string
		expectingErr  bool
	}{
		{
			description:   "should not panic with nil spec",
			spec:          nil,
			expectedNames: nil,
			expectingErr:  true,
		},
		{
			description: "invalid config",
			spec: &job.OCR2OracleSpec{
				PluginConfig: map[string]interface{}{},
			},
			expectingErr: true,
		},
		{
			description: "invalid contract id",
			spec: &job.OCR2OracleSpec{
				ContractID: "whatever...",
			},
			expectingErr: true,
		},
	}

	for _, tc := range testCases {
		chainSet := &evmmocks.ChainSet{}
		t.Run(tc.description, func(t *testing.T) {
			names, err := GetCommitPluginFilterNamesFromSpec(context.Background(), tc.spec, chainSet)
			assert.Equal(t, tc.expectedNames, names)
			if tc.expectingErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}

}

func TestGetCommitPluginFilterNames(t *testing.T) {
	onRampAddr := common.HexToAddress("0xdafea492d9c6733ae3d56b7ed1adb60692c98bc2")
	priceRegAddr := common.HexToAddress("0xdafea492d9c6733ae3d56b7ed1adb60692c98bc3")
	mockCommitStore := mock_contracts.NewCommitStoreInterface(t)
	mockCommitStore.On("GetStaticConfig", mock.Anything).Return(commit_store.CommitStoreStaticConfig{
		OnRamp: onRampAddr,
	}, nil)
	mockCommitStore.On("GetDynamicConfig", mock.Anything).Return(commit_store.CommitStoreDynamicConfig{
		PriceRegistry: priceRegAddr,
	}, nil)

	filterNames, err := getCommitPluginFilterNames(context.Background(), mockCommitStore)
	assert.NoError(t, err)
	assert.Equal(t, []string{
		"Commit ccip sends - 0xdafea492D9c6733aE3d56B7ED1aDb60692C98bc2",
		"Commit price updates - 0xdafEa492d9C6733aE3D56b7eD1aDb60692c98bc3",
	}, filterNames)
}

func Test_updateLogPollerFilters(t *testing.T) {
	srcLP := &mocklp.LogPoller{}
	dstLP := &mocklp.LogPoller{}

	onRampAddr := common.HexToAddress("0xdafea492d9c6733ae3d56b7ed1adb60692c98bc2")
	priceRegAddr := common.HexToAddress("0xdafea492d9c6733ae3d56b7ed1adb60692c98bc3")

	rf := &CommitReportingPluginFactory{
		config: CommitPluginConfig{
			sourceLP:      srcLP,
			destLP:        dstLP,
			onRampAddress: onRampAddr,
		},
		dstChainFilters: []logpoller.Filter{
			{Name: "a"}, {Name: "b"},
		},
		srcChainFilters: []logpoller.Filter{
			{Name: "c"}, {Name: "d"},
		},
		filtersMu: &sync.Mutex{},
	}

	// make sure existing filters get unregistered
	for _, f := range rf.dstChainFilters {
		dstLP.On("UnregisterFilter", f.Name, nil).Return(nil)
	}
	for _, f := range rf.srcChainFilters {
		srcLP.On("UnregisterFilter", f.Name, nil).Return(nil)
	}

	// make sure new filters are registered
	for _, f := range getCommitPluginDestLpFilters(priceRegAddr) {
		dstLP.On("RegisterFilter", f).Return(nil)
	}
	for _, f := range getCommitPluginSourceLpFilters(onRampAddr) {
		srcLP.On("RegisterFilter", f).Return(nil)
	}

	err := rf.updateLogPollerFilters(ccipconfig.CommitOnchainConfig{
		PriceRegistry: priceRegAddr,
	})
	assert.NoError(t, err)

	srcLP.AssertExpectations(t)
	dstLP.AssertExpectations(t)
}
