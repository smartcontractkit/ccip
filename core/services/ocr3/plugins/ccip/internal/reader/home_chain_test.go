package reader

import (
	"context"
	"fmt"
	"testing"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/ccipocr3/internal/mocks"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/libocr/commontypes"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var (
	chainA       = cciptypes.ChainSelector(1)
	chainB       = cciptypes.ChainSelector(2)
	chainC       = cciptypes.ChainSelector(3)
	oracleAId    = commontypes.OracleID(1)
	p2pOracleAId = libocrtypes.PeerID{byte(oracleAId)}
	oracleBId    = commontypes.OracleID(2)
	p2pOracleBId = libocrtypes.PeerID{byte(oracleBId)}
	oracleCId    = commontypes.OracleID(3)
	p2pOracleCId = libocrtypes.PeerID{byte(oracleCId)}
)

func TestHomeChainConfigPoller_HealthReport(t *testing.T) {
	homeChainReader := mocks.NewContractReaderMock()
	homeChainReader.On(
		"GetLatestValue",
		mock.Anything,
		"CCIPCapabilityConfiguration",
		"getAllChainConfigs",
		mock.Anything,
		mock.Anything).Return(fmt.Errorf("error"))

	var (
		tickTime       = 1 * time.Millisecond
		totalSleepTime = 11 * tickTime // to allow it to fail 10 times at least
	)
	configPoller := NewHomeChainConfigPoller(
		homeChainReader,
		logger.Test(t),
		tickTime,
	)
	_ = configPoller.Start(context.Background())
	// Initially it's healthy
	healthy := configPoller.HealthReport()
	require.Equal(t, map[string]error{configPoller.Name(): error(nil)}, healthy)
	// After one second it will try polling 10 times and fail
	time.Sleep(totalSleepTime)
	errors := configPoller.HealthReport()
	require.Equal(t, 1, len(errors))
	require.Errorf(t, errors[configPoller.Name()], "polling failed %d times in a row", MaxFailedPolls)

	closGracefully(t, configPoller)
}

func Test_PollingWorking(t *testing.T) {
	onChainConfigs := []ChainConfigInfo{
		{
			ChainSelector: chainA,
			ChainConfig: HomeChainConfigMapper{
				FChain: 1,
				Readers: []libocrtypes.PeerID{
					p2pOracleAId,
					p2pOracleBId,
					p2pOracleCId,
				},
				Config: []byte{0},
			},
		},
		{
			ChainSelector: chainB,
			ChainConfig: HomeChainConfigMapper{
				FChain: 2,
				Readers: []libocrtypes.PeerID{
					p2pOracleAId,
					p2pOracleBId,
				},
				Config: []byte{0},
			},
		},
		{
			ChainSelector: chainC,
			ChainConfig: HomeChainConfigMapper{
				FChain: 3,
				Readers: []libocrtypes.PeerID{
					p2pOracleCId,
				},
				Config: []byte{0},
			},
		},
	}
	homeChainConfig := map[cciptypes.ChainSelector]ChainConfig{
		chainA: {
			FChain:         1,
			SupportedNodes: mapset.NewSet(p2pOracleAId, p2pOracleBId, p2pOracleCId),
		},
		chainB: {
			FChain:         2,
			SupportedNodes: mapset.NewSet(p2pOracleAId, p2pOracleBId),
		},
		chainC: {
			FChain:         3,
			SupportedNodes: mapset.NewSet(p2pOracleCId),
		},
	}

	homeChainReader := mocks.NewContractReaderMock()
	homeChainReader.On(
		"GetLatestValue", mock.Anything, "CCIPCapabilityConfiguration", "getAllChainConfigs", mock.Anything, mock.Anything).Run(
		func(args mock.Arguments) {
			arg := args.Get(4).(*[]ChainConfigInfo)
			*arg = onChainConfigs
		}).Return(nil)

	var (
		tickTime       = 2 * time.Millisecond
		totalSleepTime = (tickTime * 2) + (10 * time.Millisecond)
		expNumCalls    = int(totalSleepTime/tickTime) + 1 // +1 for the initial call
	)

	configPoller := NewHomeChainConfigPoller(
		homeChainReader,
		logger.Test(t),
		tickTime,
	)

	ctx := context.Background()
	err := configPoller.Start(ctx)
	require.NoError(t, err)
	time.Sleep(totalSleepTime)
	err = configPoller.Close()
	require.NoError(t, err)

	// called 3 times, once when it's started, and 2 times when it's polling
	homeChainReader.AssertNumberOfCalls(t, "GetLatestValue", expNumCalls)

	configs, err := configPoller.GetAllChainConfigs()
	require.NoError(t, err)
	require.Equal(t, homeChainConfig, configs)
}

func closGracefully(t *testing.T, homeChain HomeChain) {
	err := homeChain.Close()
	require.NoError(t, err)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		// Make sure it's closed gracefully, give it 2 seconds to do so or fail
		err := homeChain.Ready()
		if err != nil {
			return
		}
		select {
		case <-ticker.C:
			t.Fatal("HomeChainReader did not close gracefully")
		}
	}
}
