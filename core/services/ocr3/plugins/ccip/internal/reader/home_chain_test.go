package reader

import (
	"context"
	"testing"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	cciptypes "github.com/smartcontractkit/ccipocr3/ccipocr3-dont-merge"
	"github.com/smartcontractkit/ccipocr3/internal/mocks"
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/stretchr/testify/assert"
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
					ChainSelector: uint64(chainA),
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
					ChainSelector: uint64(chainB),
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
					ChainSelector: uint64(chainC),
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
		configPoller := HomeChainConfigPoller{
			homeChainConfig: cciptypes.HomeChainConfig{},
			p2pIdToOracleId: map[cciptypes.Bytes32]commontypes.OracleID{
				p2pOracleAId: oracleAId,
				p2pOracleBId: oracleBId,
				p2pOracleCId: oracleCId,
			},
		}
		t.Run(tc.name, func(t *testing.T) {
			resultConfig, err := configPoller.ConvertOnChainConfigToHomeChainConfig(tc.onChainConfigs)
			assert.NoError(t, err)
			assert.Equal(t, tc.homeChainConfig, resultConfig)
		})
	}
}

func Test_PollingWorking(t *testing.T) {
	onChainConfigs := []cciptypes.OnChainCapabilityConfig{
		{
			ChainSelector: uint64(chainA),
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
			ChainSelector: uint64(chainB),
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
			ChainSelector: uint64(chainC),
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

	configPoller := HomeChainConfigPoller{
		homeChainReader: mocks.NewHomeChainContractReader(onChainConfigs),
		homeChainConfig: cciptypes.HomeChainConfig{},
		p2pIdToOracleId: map[cciptypes.Bytes32]commontypes.OracleID{
			p2pOracleAId: oracleAId,
			p2pOracleBId: oracleBId,
			p2pOracleCId: oracleCId,
		},
	}
	ctx := context.Background()
	go configPoller.StartPolling(ctx, 1*time.Second)
	// sleep for 2 seconds
	time.Sleep(2 * time.Second)

	assert.Equal(t, homeChainConfig, configPoller.GetConfig())
}
