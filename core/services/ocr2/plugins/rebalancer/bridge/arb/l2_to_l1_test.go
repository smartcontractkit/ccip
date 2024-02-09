package arb

import (
	"errors"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/mock"

	clientmocks "github.com/smartcontractkit/chainlink/v2/core/chains/evm/client/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	lpmocks "github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/models"
	"github.com/test-go/testify/require"
)

const (
	arbSepolia     uint64 = 421614
	sepoliaChainID uint64 = 11155111
)

func Test_L2ToL1Bridge_New(t *testing.T) {
	rollupAddress, l1RebalancerAddress, l2BridgeAdapterAddress, l1BridgeAdapterAddress := testutils.NewAddress(), testutils.NewAddress(), testutils.NewAddress(), testutils.NewAddress()

	t.Run("happy path", func(t *testing.T) {
		l2LogPoller := lpmocks.NewLogPoller(t)
		l2LogPoller.On("RegisterFilter", mock.MatchedBy(func(f logpoller.Filter) bool {
			if len(f.EventSigs) != 1 {
				return false
			}
			if f.EventSigs[0] != L2ToL1ERC20SentTopic {
				return false
			}
			if len(f.Addresses) != 1 {
				return false
			}
			if f.Addresses[0] != l2BridgeAdapterAddress {
				return false
			}
			if f.Retention != DurationMonth {
				return false
			}
			return true
		})).Return(nil)
		defer l2LogPoller.AssertExpectations(t)

		l1LogPoller := lpmocks.NewLogPoller(t)
		l1LogPoller.On("RegisterFilter", mock.MatchedBy(func(f logpoller.Filter) bool {
			if len(f.EventSigs) != 2 {
				return false
			}
			if !(f.EventSigs[0] == L2toL1ERC20FinalizedTopic || f.EventSigs[1] == NodeConfirmedTopic) {
				return false
			}
			if !(f.EventSigs[1] == L2toL1ERC20FinalizedTopic || f.EventSigs[1] == NodeConfirmedTopic) {
				return false
			}
			if len(f.Addresses) != 2 {
				return false
			}
			if !(f.Addresses[0] == l1BridgeAdapterAddress || f.Addresses[1] == l1BridgeAdapterAddress) {
				return false
			}
			if !(f.Addresses[0] == rollupAddress || f.Addresses[1] == rollupAddress) {
				return false
			}
			if f.Retention != DurationMonth {
				return false
			}
			return true
		})).Return(nil)
		defer l1LogPoller.AssertExpectations(t)

		l2Client := clientmocks.NewClient(t)
		l1Client := clientmocks.NewClient(t)

		bridge, err := NewL2ToL1Bridge(
			logger.TestLogger(t),
			models.NetworkSelector(mustGetChainByID(t, arbSepolia).Selector),
			models.NetworkSelector(mustGetChainByID(t, sepoliaChainID).Selector),
			rollupAddress,
			l1RebalancerAddress,
			l2BridgeAdapterAddress,
			l1BridgeAdapterAddress,
			l2LogPoller,
			l1LogPoller,
			l2Client,
			l1Client,
		)
		require.NoError(t, err)

		require.Equal(t, rollupAddress, bridge.rollupCore.Address())
		require.Equal(t, l1RebalancerAddress, bridge.l1RebalancerAddress)
		require.Equal(t, l2BridgeAdapterAddress, bridge.l2BridgeAdapter.Address())
		require.Equal(t, l1BridgeAdapterAddress, bridge.l1BridgeAdapter.Address())
	})

	t.Run("l2 poller register filter error", func(t *testing.T) {
		l2LogPoller := lpmocks.NewLogPoller(t)
		l2LogPoller.On("RegisterFilter", mock.Anything).Return(errors.New("some error"), nil)
		defer l2LogPoller.AssertExpectations(t)

		l1LogPoller := lpmocks.NewLogPoller(t)
		l2Client := clientmocks.NewClient(t)
		l1Client := clientmocks.NewClient(t)

		_, err := NewL2ToL1Bridge(
			logger.TestLogger(t),
			models.NetworkSelector(mustGetChainByID(t, arbSepolia).Selector),
			models.NetworkSelector(mustGetChainByID(t, sepoliaChainID).Selector),
			rollupAddress,
			l1RebalancerAddress,
			l2BridgeAdapterAddress,
			l1BridgeAdapterAddress,
			l2LogPoller,
			l1LogPoller,
			l2Client,
			l1Client,
		)
		require.Error(t, err)
	})

	t.Run("l1 poller register filter error", func(t *testing.T) {
		l2LogPoller := lpmocks.NewLogPoller(t)
		l2LogPoller.On("RegisterFilter", mock.Anything).Return(nil)
		defer l2LogPoller.AssertExpectations(t)

		l1LogPoller := lpmocks.NewLogPoller(t)
		l1LogPoller.On("RegisterFilter", mock.Anything).Return(errors.New("some error"), nil)
		defer l1LogPoller.AssertExpectations(t)

		l2Client := clientmocks.NewClient(t)
		l1Client := clientmocks.NewClient(t)

		_, err := NewL2ToL1Bridge(
			logger.TestLogger(t),
			models.NetworkSelector(mustGetChainByID(t, arbSepolia).Selector),
			models.NetworkSelector(mustGetChainByID(t, sepoliaChainID).Selector),
			rollupAddress,
			l1RebalancerAddress,
			l2BridgeAdapterAddress,
			l1BridgeAdapterAddress,
			l2LogPoller,
			l1LogPoller,
			l2Client,
			l1Client,
		)
		require.Error(t, err)
	})
}

func Test_L2ToL1Bridge_GetBridgeSpecificPayload(t *testing.T) {
	bridge := &l2ToL1Bridge{}
	payload, err := bridge.GetBridgeSpecificPayload(testutils.Context(t), models.Transfer{})
	require.NoError(t, err)
	require.Empty(t, payload)
}

func Test_L2ToL1Bridge_RemoteChainSelector(t *testing.T) {
	sepoliaSelector := mustGetChainByID(t, sepoliaChainID).Selector
	bridge := &l2ToL1Bridge{
		remoteSelector: models.NetworkSelector(sepoliaSelector),
	}
	selector := bridge.RemoteChainSelector()
	require.Equal(t, models.NetworkSelector(sepoliaSelector), selector)
}

func Test_L2ToL1Bridge_LocalChainSelector(t *testing.T) {
	arbSelector := mustGetChainByID(t, arbSepolia).Selector
	bridge := &l2ToL1Bridge{
		localSelector: models.NetworkSelector(arbSelector),
	}
	selector := bridge.LocalChainSelector()
	require.Equal(t, models.NetworkSelector(arbSelector), selector)
}

func Test_L2ToL1Bridge_unpackFinalizeInboundTransfer(t *testing.T) {
	// Example from a real transaction:
	// https://sepolia.arbiscan.io/tx/0x12012fdc48132435be2821b6f6fa0b7da5acb7872934620d4221efc60bfe0e4d#eventlog
	finalizeInboundTransferCalldata := hexutil.MustDecode("0x2e567b360000000000000000000000007b79995e5f793a07bc00c21412e50ecae098e7f900000000000000000000000066e4037a15b5c15620b3cfac064f53ffba24361d000000000000000000000000d19e526da8b7e1cb970af3da92bfcef3e135f468000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000000000900000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000000")
	params, err := unpackFinalizeInboundTransfer(finalizeInboundTransferCalldata)
	require.NoError(t, err)
	expectedToken := common.HexToAddress("0x7b79995e5f793a07bc00c21412e50ecae098e7f9")
	expectedFrom := common.HexToAddress("0x66e4037a15b5c15620b3cfac064f53ffba24361d")
	expectedTo := common.HexToAddress("0xd19e526da8b7e1cb970af3da92bfcef3e135f468")
	expectedAmount := big.NewInt(1)
	expectedData := hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000000900000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000000")
	require.Equal(t, expectedToken, params.l1Token)
	require.Equal(t, expectedFrom, params.l2Sender)
	require.Equal(t, expectedTo, params.l1Receiver)
	require.Equal(t, expectedAmount, params.amount)
	require.Equal(t, expectedData, params.data)
}

func Test_L2ToL1Bridge_parseL2ToL1Transfers(t *testing.T) {

}

func Test_L2ToL1Bridge_parseL2ToL1Finalizations(t *testing.T) {

}

func Test_L2ToL1Bridge_findMatchingFinalization(t *testing.T) {

}

func Test_L2ToL1Bridge_filterOutFinalizedTransfers(t *testing.T) {

}

func Test_L2ToL1Bridge_getLatestNodeConfirmed(t *testing.T) {

}

func Test_L2ToL1Bridge_getSendCountForBlock(t *testing.T) {

}

func Test_L2ToL1Bridge_getProof(t *testing.T) {

}

func Test_L2ToL1Bridge_getL1BlockFromRPC(t *testing.T) {

}

func Test_L2ToL1Bridge_getFinalizationData(t *testing.T) {

}

func Test_L2ToL1Bridge_partitionReadyTransfers(t *testing.T) {

}

func Test_L2ToL1Bridge_toPendingTransfers(t *testing.T) {

}

func Test_L2ToL1Bridge_GetTransfers(t *testing.T) {

}

func Test_L2ToL1Bridge_unpackUint256(t *testing.T) {

}

func mustGetChainByID(t *testing.T, id uint64) chainsel.Chain {
	chain, ok := chainsel.ChainByEvmChainID(id)
	require.True(t, ok)
	return chain
}
