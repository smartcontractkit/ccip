package reader

import (
	"context"
	"testing"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	cciptypes "github.com/smartcontractkit/ccipocr3/ccipocr3-dont-merge"
	"github.com/smartcontractkit/ccipocr3/internal/mocks"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	chainA       = cciptypes.ChainSelector(1)
	chainB       = cciptypes.ChainSelector(2)
	chainC       = cciptypes.ChainSelector(3)
	oracleAId    = commontypes.OracleID(1)
	p2pOracleAId = cciptypes.Bytes32{byte(oracleAId)}
	oracleBId    = commontypes.OracleID(2)
	p2pOracleBId = cciptypes.Bytes32{byte(oracleBId)}
	oracleCId    = commontypes.OracleID(3)
	p2pOracleCId = cciptypes.Bytes32{byte(oracleCId)}
)

func Test_ConvertOnChainConfigToHomeChainConfig(t *testing.T) {
	var tests = []struct {
		name            string
		onChainConfigs  []cciptypes.OnChainCapabilityConfig
		homeChainConfig cciptypes.HomeChainConfig
		expErr          string
	}{
		{
			name: "Convert",
			onChainConfigs: []cciptypes.OnChainCapabilityConfig{
				{
					ChainSelector: chainA,
					ChainConfig: cciptypes.OnChainConfig{
						FChain: 1,
						Readers: []cciptypes.Bytes32{
							p2pOracleAId,
							p2pOracleBId,
							p2pOracleCId,
						},
						Config: []byte{0},
					},
				},
				{
					ChainSelector: chainB,
					ChainConfig: cciptypes.OnChainConfig{
						FChain: 2,
						Readers: []cciptypes.Bytes32{
							p2pOracleAId,
							p2pOracleBId,
						},
						Config: []byte{0},
					},
				},
				{
					ChainSelector: chainC,
					ChainConfig: cciptypes.OnChainConfig{
						FChain: 3,
						Readers: []cciptypes.Bytes32{
							p2pOracleCId,
						},
						Config: []byte{0},
					},
				},
			},
			homeChainConfig: cciptypes.HomeChainConfig{
				FChain: map[cciptypes.ChainSelector]int{
					chainA: 1,
					chainB: 2,
					chainC: 3,
				},
				NodeSupportedChains: map[commontypes.OracleID]cciptypes.SupportedChains{
					oracleAId: {Supported: mapset.NewSet[cciptypes.ChainSelector](chainA, chainB)},
					oracleBId: {Supported: mapset.NewSet[cciptypes.ChainSelector](chainA, chainB)},
					oracleCId: {Supported: mapset.NewSet[cciptypes.ChainSelector](chainA, chainC)},
				},
			},
		},
	}
	for _, tc := range tests {
		configPoller := NewHomeChainConfigPoller(
			nil,
			logger.Test(t),
			map[cciptypes.Bytes32]commontypes.OracleID{
				p2pOracleAId: oracleAId,
				p2pOracleBId: oracleBId,
				p2pOracleCId: oracleCId,
			},
		)
		t.Run(tc.name, func(t *testing.T) {
			resultConfig, err := configPoller.convertOnChainConfigToHomeChainConfig(tc.onChainConfigs)
			assert.NoError(t, err)
			assert.Equal(t, tc.homeChainConfig, resultConfig)
		})
	}
}

func Test_PollingWorking(t *testing.T) {
	onChainConfigs := []cciptypes.OnChainCapabilityConfig{
		{
			ChainSelector: chainA,
			ChainConfig: cciptypes.OnChainConfig{
				FChain: 1,
				Readers: []cciptypes.Bytes32{
					p2pOracleAId,
					p2pOracleBId,
					p2pOracleCId,
				},
				Config: []byte{0},
			},
		},
		{
			ChainSelector: chainB,
			ChainConfig: cciptypes.OnChainConfig{
				FChain: 2,
				Readers: []cciptypes.Bytes32{
					p2pOracleAId,
					p2pOracleBId,
				},
				Config: []byte{0},
			},
		},
		{
			ChainSelector: chainC,
			ChainConfig: cciptypes.OnChainConfig{
				FChain: 3,
				Readers: []cciptypes.Bytes32{
					p2pOracleCId,
				},
				Config: []byte{0},
			},
		},
	}
	homeChainConfig := cciptypes.HomeChainConfig{
		FChain: map[cciptypes.ChainSelector]int{
			chainA: 1,
			chainB: 2,
			chainC: 3,
		},
		NodeSupportedChains: map[commontypes.OracleID]cciptypes.SupportedChains{
			oracleAId: {Supported: mapset.NewSet[cciptypes.ChainSelector](chainA, chainB)},
			oracleBId: {Supported: mapset.NewSet[cciptypes.ChainSelector](chainA, chainB)},
			oracleCId: {Supported: mapset.NewSet[cciptypes.ChainSelector](chainA, chainC)},
		},
	}

	homeChainReader := mocks.NewHomeChainContractReader(onChainConfigs)
	homeChainReader.On(
		"GetLatestValue", mock.Anything, "CCIPCapabilityConfiguration", "getAllChainConfigs", mock.Anything, mock.Anything).Run(
		func(args mock.Arguments) {
			arg := args.Get(4).(*[]cciptypes.OnChainCapabilityConfig)
			*arg = onChainConfigs
		}).Return(nil)

	configPoller := NewHomeChainConfigPoller(
		homeChainReader,
		logger.Test(t),
		map[cciptypes.Bytes32]commontypes.OracleID{
			p2pOracleAId: oracleAId,
			p2pOracleBId: oracleBId,
			p2pOracleCId: oracleCId,
		},
	)
	ctx := context.Background()
	configPoller.StartPolling(ctx, 1*time.Second)
	// sleep for 2 seconds
	time.Sleep(2 * time.Second)
	_ = configPoller.Close(ctx)

	assert.Equal(t, homeChainConfig, configPoller.GetConfig())
}
