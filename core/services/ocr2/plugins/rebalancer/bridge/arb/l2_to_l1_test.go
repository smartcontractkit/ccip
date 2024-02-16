package arb

import (
	"context"
	"errors"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/lib/pq"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/mock"
	"github.com/test-go/testify/require"

	clientmocks "github.com/smartcontractkit/chainlink/v2/core/chains/evm/client/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	lpmocks "github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
	utilsbig "github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils/big"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/arbitrum_l1_bridge_adapter"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/arbitrum_l2_bridge_adapter"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/arbitrum_rollup_core"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/mocks/mock_arbitrum_rollup_core"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/models"
)

const (
	arbSepolia     uint64 = 421614
	sepoliaChainID uint64 = 11155111
)

func Test_L2ToL1Bridge_New(t *testing.T) {
	rollupAddress, l1RebalancerAddress, l2RebalancerAddress, l2BridgeAdapterAddress, l1BridgeAdapterAddress := testutils.NewAddress(), testutils.NewAddress(), testutils.NewAddress(), testutils.NewAddress(), testutils.NewAddress()

	t.Run("happy path", func(t *testing.T) {
		l2LogPoller := lpmocks.NewLogPoller(t)
		l2LogPoller.On("RegisterFilter", mock.MatchedBy(func(f logpoller.Filter) bool {
			if len(f.EventSigs) != 1 {
				return false
			}
			if f.EventSigs[0] != L2ToL1ERC20SentTopic {
				return false
			}
			if f.Retention != DurationMonth {
				return false
			}
			if f.Addresses[0] != l2BridgeAdapterAddress {
				return false
			}
			return true
		})).Return(nil)
		defer l2LogPoller.AssertExpectations(t)

		l1LogPoller := lpmocks.NewLogPoller(t)
		l1LogPoller.On("RegisterFilter", mock.MatchedBy(func(f logpoller.Filter) bool {
			if len(f.EventSigs) != 3 {
				return false
			}
			if f.EventSigs[0] != L2toL1ERC20FinalizedTopic || f.EventSigs[1] != NodeConfirmedTopic || f.EventSigs[2] != LiquidityTransferredTopic {
				return false
			}
			if f.Retention != DurationMonth {
				return false
			}
			if f.Addresses[0] != l1BridgeAdapterAddress || f.Addresses[1] != rollupAddress || f.Addresses[2] != l1RebalancerAddress {
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
			l2RebalancerAddress,
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
			l2RebalancerAddress,
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
			l2RebalancerAddress,
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

func Test_L2ToL1Bridge_GetBridgePayloadAndFee(t *testing.T) {
	bridge := &l2ToL1Bridge{}
	payload, fee, err := bridge.GetBridgePayloadAndFee(testutils.Context(t), models.Transfer{})
	require.NoError(t, err)
	require.Empty(t, payload)
	require.Equal(t, big.NewInt(0), fee)
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
	t.Run("happy path", func(t *testing.T) {
		// data from a real transaction:
		// https://sepolia.arbiscan.io/tx/0x12012fdc48132435be2821b6f6fa0b7da5acb7872934620d4221efc60bfe0e4d#eventlog
		lgs := []logpoller.Log{
			{
				EvmChainId:  utilsbig.New(big.NewInt(int64(arbSepolia))),
				Address:     common.HexToAddress("0x66e4037a15b5c15620b3cfac064f53ffba24361d"),
				BlockHash:   common.HexToHash("0x20e3784f8431bc9f5d5d0f830fb781217e4a1c46621d9f8735448dd7a542c0a7"),
				TxHash:      common.HexToHash("0x12012fdc48132435be2821b6f6fa0b7da5acb7872934620d4221efc60bfe0e4d"),
				LogIndex:    hexutil.MustDecodeBig("0x1c").Int64(),
				BlockNumber: hexutil.MustDecodeBig("0xc38f48").Int64(),
				Data:        hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000051fe"),
				Topics: pq.ByteaArray{
					hexutil.MustDecode("0xffa51f05034aaeaa4c0a623d14cae59dd4a579d2c66ed37f38c4baef2d504c01"),
					hexutil.MustDecode("0x000000000000000000000000980b62da83eff3d4576c647993b0c1d7faf17c73"),
					hexutil.MustDecode("0x0000000000000000000000007b79995e5f793a07bc00c21412e50ecae098e7f9"),
					hexutil.MustDecode("0x000000000000000000000000d19e526da8b7e1cb970af3da92bfcef3e135f468"),
				},
				EventSig:  common.HexToHash("0xffa51f05034aaeaa4c0a623d14cae59dd4a579d2c66ed37f38c4baef2d504c01"),
				CreatedAt: time.Now(),
			},
		}
		// create the wrapper, doesn't need to connect to anything, just for parsing
		l2Adapter, err := arbitrum_l2_bridge_adapter.NewArbitrumL2BridgeAdapter(utils.ZeroAddress, nil)
		require.NoError(t, err)
		bridge := &l2ToL1Bridge{
			l2BridgeAdapter: l2Adapter,
		}
		parsed, parsedToLPLogs, err := bridge.parseL2ToL1Transfers(lgs)
		require.NoError(t, err)
		require.Len(t, parsed, 1)
		require.Len(t, parsedToLPLogs, 1)
		require.Equal(t, lgs[0], parsedToLPLogs[logKey{
			txHash:   lgs[0].TxHash,
			logIndex: lgs[0].LogIndex,
		}])
		require.Equal(t, common.HexToHash("0xffa51f05034aaeaa4c0a623d14cae59dd4a579d2c66ed37f38c4baef2d504c01"), parsed[0].Topic())
		require.Equal(t, common.HexToAddress("0x980b62da83eff3d4576c647993b0c1d7faf17c73"), parsed[0].LocalToken)
		require.Equal(t, common.HexToAddress("0x7b79995e5f793a07bc00c21412e50ecae098e7f9"), parsed[0].RemoteToken)
		require.Equal(t, common.HexToAddress("0xd19e526da8b7e1cb970af3da92bfcef3e135f468"), parsed[0].Recipient)
		require.Equal(t, big.NewInt(1), parsed[0].Amount)
		require.Equal(t, hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000051FE"), parsed[0].OutboundTransferResult)
	})

	t.Run("parse error", func(t *testing.T) {
		// bad log data
		lgs := []logpoller.Log{
			{
				EvmChainId:  utilsbig.New(big.NewInt(int64(arbSepolia))),
				Address:     common.HexToAddress("0x66e4037a15b5c15620b3cfac064f53ffba24361d"),
				BlockHash:   common.HexToHash("0x20e3784f8431bc9f5d5d0f830fb781217e4a1c46621d9f8735448dd7a542c0a7"),
				TxHash:      common.HexToHash("0x12012fdc48132435be2821b6f6fa0b7da5acb7872934620d4221efc60bfe0e4d"),
				LogIndex:    hexutil.MustDecodeBig("0x1c").Int64(),
				BlockNumber: hexutil.MustDecodeBig("0xc38f48").Int64(),
				Data:        hexutil.MustDecode("0x"),
				Topics:      pq.ByteaArray{},
			},
		}
		// create the wrapper, doesn't need to connect to anything, just for parsing
		l2Adapter, err := arbitrum_l2_bridge_adapter.NewArbitrumL2BridgeAdapter(utils.ZeroAddress, nil)
		require.NoError(t, err)
		bridge := &l2ToL1Bridge{
			l2BridgeAdapter: l2Adapter,
		}
		_, _, err = bridge.parseL2ToL1Transfers(lgs)
		require.Error(t, err)
	})
}

func Test_L2ToL1Bridge_parseL2ToL1Finalizations(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		lgs := []logpoller.Log{
			{
				EvmChainId:  utilsbig.New(big.NewInt(int64(sepoliaChainID))),
				Address:     common.HexToAddress("0x97dc5c3c16954a0305f0419ed1a527cca0650dab"),
				BlockHash:   common.HexToHash("0xc9e1989459d9c3640ac0b43d88a9845ed7a2cc836236c866e1f3be5464724849"),
				TxHash:      common.HexToHash("0xecc26b602811b33a146c0e2803f67b4c931d31c48db974c70b8acb137234183a"),
				EventSig:    common.HexToHash("0x72ecb2b30e3013e4dcbd23949793c0ad87766498cc85c69a637847c1a790b8f6"),
				Data:        hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000000000051fe000000000000000000000000cfb1f08a4852699a979909e22c30263ca249556d000000000000000000000000a8ad8d7e13cbf556ee75cb0324c13535d8100e1e0000000000000000000000000000000000000000000000000000000000c38f480000000000000000000000000000000000000000000000000000000000501e920000000000000000000000000000000000000000000000000000000065c5e97f00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000320000000000000000000000000000000000000000000000000000000000000000f6632b3f37ef2be542e585204f3ed7826c158f6414c7388d12579f35eaa10420970b8e564f4708b24f214d2dc9e360727341760abba72bec52f64cbf45d1573fd9703436db5204779a947305958d7b285b918ec9cc919562c6293e182509be1982d766d4a0f1155f58d749967c4da5d0a3d275be0fa29c2a354ab41b061cd59fde9634b8838c95f4cfd7099ec952dc0593371f13d145cdf1b0738c89d225854f5c1cbee2f11e51364bb7b838b4a7631f3f46ccf9ea3664014c0e0c9564ccf737d92422592c2d063a2a466a4bd53d29c163f45fc7f2bd845b389c0bd89f688bb28b3b7492663240be7e70cec690b22ad5d43dff03ff1e264dec3e9982f3ec36f286edba00d9c21e36a7b567b1d6e4f828613c4d46328eca6fef826f8014de82a19b33d2a5f7ba320d574bf4bbd933ab5f5b2a991ef38b7a35ef56e734e6ed4f1d0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004bde3fbfa26b03bfcbbc80884a5c229d987992ce2366bef5882742cac7cc7efe000000000000000000000000000000000000000000000000000000000000000037ef20c8140b820c15303931a9e839810ec66f4d9523c3fa8187434076a25b1300000000000000000000000000000000000000000000000000000000000001242e567b360000000000000000000000007b79995e5f793a07bc00c21412e50ecae098e7f900000000000000000000000066e4037a15b5c15620b3cfac064f53ffba24361d000000000000000000000000d19e526da8b7e1cb970af3da92bfcef3e135f468000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000090000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
				BlockNumber: hexutil.MustDecodeBig("0x5023e1").Int64(),
				LogIndex:    hexutil.MustDecodeBig("0xca").Int64(),
				Topics: pq.ByteaArray{
					hexutil.MustDecode("0x72ecb2b30e3013e4dcbd23949793c0ad87766498cc85c69a637847c1a790b8f6"),
					hexutil.MustDecode("0x000000000000000000000000cfb1f08a4852699a979909e22c30263ca249556d"),
					hexutil.MustDecode("0x000000000000000000000000a8ad8d7e13cbf556ee75cb0324c13535d8100e1e"),
				},
				CreatedAt: time.Now(),
			},
		}
		// create the wrapper, doesn't need to connect to anything, just for parsing
		l1Adapter, err := arbitrum_l1_bridge_adapter.NewArbitrumL1BridgeAdapter(utils.ZeroAddress, nil)
		require.NoError(t, err)
		bridge := &l2ToL1Bridge{
			l1BridgeAdapter: l1Adapter,
		}
		parsed, err := bridge.parseL2ToL1Finalizations(lgs)
		require.NoError(t, err)
		require.Len(t, parsed, 1)
		require.Equal(t, common.HexToHash("0x72ecb2b30e3013e4dcbd23949793c0ad87766498cc85c69a637847c1a790b8f6"), parsed[0].Topic())
		require.Equal(t, common.HexToAddress("0xCFB1f08A4852699a979909e22c30263ca249556D"), parsed[0].RemoteSender)
		require.Equal(t, common.HexToAddress("0xA8aD8d7e13cbf556eE75CB0324c13535d8100e1E"), parsed[0].LocalReceiver)
		require.Equal(t, big.NewInt(1), parsed[0].Amount)
		// TODO: check payload data?
	})

	t.Run("parse error", func(t *testing.T) {
		lgs := []logpoller.Log{
			{
				EvmChainId: utilsbig.New(big.NewInt(int64(sepoliaChainID))),
				Address:    common.HexToAddress("0x97dc5c3c16954a0305f0419ed1a527cca0650dab"),
				BlockHash:  common.HexToHash("0xc9e1989459d9c3640ac0b43d88a9845ed7a2cc836236c866e1f3be5464724849"),
				TxHash:     common.HexToHash("0xecc26b602811b33a146c0e2803f67b4c931d31c48db974c70b8acb137234183a"),
				EventSig:   common.HexToHash("0x72ecb2b30e3013e4dcbd23949793c0ad87766498cc85c69a637847c1a790b8f6"),
				Data:       hexutil.MustDecode("0x"), // bad data
			},
		}
		// create the wrapper, doesn't need to connect to anything, just for parsing
		l1Adapter, err := arbitrum_l1_bridge_adapter.NewArbitrumL1BridgeAdapter(utils.ZeroAddress, nil)
		require.NoError(t, err)
		bridge := &l2ToL1Bridge{
			l1BridgeAdapter: l1Adapter,
		}
		_, err = bridge.parseL2ToL1Finalizations(lgs)
		require.Error(t, err)
	})
}

func Test_L2ToL1Bridge_findMatchingFinalization(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		l2Adapter, err := arbitrum_l2_bridge_adapter.NewArbitrumL2BridgeAdapter(utils.ZeroAddress, nil)
		require.NoError(t, err)
		l1Adapter, err := arbitrum_l1_bridge_adapter.NewArbitrumL1BridgeAdapter(utils.ZeroAddress, nil)
		require.NoError(t, err)
		bridge := &l2ToL1Bridge{
			l2BridgeAdapter: l2Adapter,
			l1BridgeAdapter: l1Adapter,
			// need to override this because the function will check if the transfer is destined for the rebalancer
			// these sample txes are destined for a test EOA on testnet
			l1RebalancerAddress: common.HexToAddress("0xd19e526da8b7e1cb970af3da92bfcef3e135f468"),
			lggr:                logger.TestLogger(t),
		}
		// l2 -> l1 real withdrawal
		// https://sepolia.arbiscan.io/tx/0x12012fdc48132435be2821b6f6fa0b7da5acb7872934620d4221efc60bfe0e4d#eventlog
		l2ToL1TransferLogs := []logpoller.Log{
			{
				EvmChainId:  utilsbig.New(big.NewInt(int64(arbSepolia))),
				Address:     common.HexToAddress("0x66e4037a15b5c15620b3cfac064f53ffba24361d"),
				BlockHash:   common.HexToHash("0x20e3784f8431bc9f5d5d0f830fb781217e4a1c46621d9f8735448dd7a542c0a7"),
				TxHash:      common.HexToHash("0x12012fdc48132435be2821b6f6fa0b7da5acb7872934620d4221efc60bfe0e4d"),
				LogIndex:    hexutil.MustDecodeBig("0x1c").Int64(),
				BlockNumber: hexutil.MustDecodeBig("0xc38f48").Int64(),
				Data:        hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000051fe"),
				Topics: pq.ByteaArray{
					hexutil.MustDecode("0xffa51f05034aaeaa4c0a623d14cae59dd4a579d2c66ed37f38c4baef2d504c01"),
					hexutil.MustDecode("0x000000000000000000000000980b62da83eff3d4576c647993b0c1d7faf17c73"),
					hexutil.MustDecode("0x0000000000000000000000007b79995e5f793a07bc00c21412e50ecae098e7f9"),
					hexutil.MustDecode("0x000000000000000000000000d19e526da8b7e1cb970af3da92bfcef3e135f468"),
				},
				EventSig:  common.HexToHash("0xffa51f05034aaeaa4c0a623d14cae59dd4a579d2c66ed37f38c4baef2d504c01"),
				CreatedAt: time.Now(),
			},
		}
		// l2 -> l1 real finalizations
		// 1. https://sepolia.etherscan.io/tx/0xe9219865868eea6654f15f7915924b3484c650400f7b7a9ed6fe3c5abcb50092#eventlog <- incorrect one
		// 2. https://sepolia.etherscan.io/tx/0xecc26b602811b33a146c0e2803f67b4c931d31c48db974c70b8acb137234183a#eventlog <- correct one
		finalizationLogs := []logpoller.Log{
			{
				EvmChainId:  utilsbig.New(big.NewInt(int64(sepoliaChainID))),
				Address:     common.HexToAddress("0x97dc5c3c16954a0305f0419ed1a527cca0650dab"),
				BlockHash:   common.HexToHash("0xebb5704d693dd2df8809eaf5564c306f23db452f503d58184628fded51d4b3e2"),
				TxHash:      common.HexToHash("0xe9219865868eea6654f15f7915924b3484c650400f7b7a9ed6fe3c5abcb50092"),
				EventSig:    common.HexToHash("0x72ecb2b30e3013e4dcbd23949793c0ad87766498cc85c69a637847c1a790b8f6"),
				Data:        hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000000000051fd000000000000000000000000cfb1f08a4852699a979909e22c30263ca249556d000000000000000000000000a8ad8d7e13cbf556ee75cb0324c13535d8100e1e0000000000000000000000000000000000000000000000000000000000c38bd20000000000000000000000000000000000000000000000000000000000501e7d0000000000000000000000000000000000000000000000000000000065c5e86e00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000320000000000000000000000000000000000000000000000000000000000000000f8b51fa17ea4f1dff4db1151de00b9868b9d6ba80ca7077a08dc55203fb6f698a6f1638fc7d51bcde617e57cfd8c89f669f5b9e5603e8f2ec48d55bbc5d1f57f99703436db5204779a947305958d7b285b918ec9cc919562c6293e182509be1982d766d4a0f1155f58d749967c4da5d0a3d275be0fa29c2a354ab41b061cd59fde9634b8838c95f4cfd7099ec952dc0593371f13d145cdf1b0738c89d225854f5c1cbee2f11e51364bb7b838b4a7631f3f46ccf9ea3664014c0e0c9564ccf737d92422592c2d063a2a466a4bd53d29c163f45fc7f2bd845b389c0bd89f688bb28b3b7492663240be7e70cec690b22ad5d43dff03ff1e264dec3e9982f3ec36f286edba00d9c21e36a7b567b1d6e4f828613c4d46328eca6fef826f8014de82a195a9351bbbf7c8f3e8c757b46a3fb28a114376fe8c77012b2f99dc20f5d0e932e000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004bde3fbfa26b03bfcbbc80884a5c229d987992ce2366bef5882742cac7cc7efe000000000000000000000000000000000000000000000000000000000000000037ef20c8140b820c15303931a9e839810ec66f4d9523c3fa8187434076a25b1300000000000000000000000000000000000000000000000000000000000001242e567b360000000000000000000000007b79995e5f793a07bc00c21412e50ecae098e7f900000000000000000000000066e4037a15b5c15620b3cfac064f53ffba24361d000000000000000000000000d19e526da8b7e1cb970af3da92bfcef3e135f468000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
				BlockNumber: hexutil.MustDecodeBig("0x502835").Int64(),
				LogIndex:    hexutil.MustDecodeBig("0x7f").Int64(),
				Topics: pq.ByteaArray{
					hexutil.MustDecode("0x72ecb2b30e3013e4dcbd23949793c0ad87766498cc85c69a637847c1a790b8f6"),
					hexutil.MustDecode("0x000000000000000000000000cfb1f08a4852699a979909e22c30263ca249556d"),
					hexutil.MustDecode("0x000000000000000000000000a8ad8d7e13cbf556ee75cb0324c13535d8100e1e"),
				},
				CreatedAt: time.Now(),
			},
			{
				EvmChainId:  utilsbig.New(big.NewInt(int64(sepoliaChainID))),
				Address:     common.HexToAddress("0x97dc5c3c16954a0305f0419ed1a527cca0650dab"),
				BlockHash:   common.HexToHash("0xc9e1989459d9c3640ac0b43d88a9845ed7a2cc836236c866e1f3be5464724849"),
				TxHash:      common.HexToHash("0xecc26b602811b33a146c0e2803f67b4c931d31c48db974c70b8acb137234183a"),
				EventSig:    common.HexToHash("0x72ecb2b30e3013e4dcbd23949793c0ad87766498cc85c69a637847c1a790b8f6"),
				Data:        hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000000000051fe000000000000000000000000cfb1f08a4852699a979909e22c30263ca249556d000000000000000000000000a8ad8d7e13cbf556ee75cb0324c13535d8100e1e0000000000000000000000000000000000000000000000000000000000c38f480000000000000000000000000000000000000000000000000000000000501e920000000000000000000000000000000000000000000000000000000065c5e97f00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000320000000000000000000000000000000000000000000000000000000000000000f6632b3f37ef2be542e585204f3ed7826c158f6414c7388d12579f35eaa10420970b8e564f4708b24f214d2dc9e360727341760abba72bec52f64cbf45d1573fd9703436db5204779a947305958d7b285b918ec9cc919562c6293e182509be1982d766d4a0f1155f58d749967c4da5d0a3d275be0fa29c2a354ab41b061cd59fde9634b8838c95f4cfd7099ec952dc0593371f13d145cdf1b0738c89d225854f5c1cbee2f11e51364bb7b838b4a7631f3f46ccf9ea3664014c0e0c9564ccf737d92422592c2d063a2a466a4bd53d29c163f45fc7f2bd845b389c0bd89f688bb28b3b7492663240be7e70cec690b22ad5d43dff03ff1e264dec3e9982f3ec36f286edba00d9c21e36a7b567b1d6e4f828613c4d46328eca6fef826f8014de82a19b33d2a5f7ba320d574bf4bbd933ab5f5b2a991ef38b7a35ef56e734e6ed4f1d0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004bde3fbfa26b03bfcbbc80884a5c229d987992ce2366bef5882742cac7cc7efe000000000000000000000000000000000000000000000000000000000000000037ef20c8140b820c15303931a9e839810ec66f4d9523c3fa8187434076a25b1300000000000000000000000000000000000000000000000000000000000001242e567b360000000000000000000000007b79995e5f793a07bc00c21412e50ecae098e7f900000000000000000000000066e4037a15b5c15620b3cfac064f53ffba24361d000000000000000000000000d19e526da8b7e1cb970af3da92bfcef3e135f468000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000090000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
				BlockNumber: hexutil.MustDecodeBig("0x5023e1").Int64(),
				LogIndex:    hexutil.MustDecodeBig("0xca").Int64(),
				Topics: pq.ByteaArray{
					hexutil.MustDecode("0x72ecb2b30e3013e4dcbd23949793c0ad87766498cc85c69a637847c1a790b8f6"),
					hexutil.MustDecode("0x000000000000000000000000cfb1f08a4852699a979909e22c30263ca249556d"),
					hexutil.MustDecode("0x000000000000000000000000a8ad8d7e13cbf556ee75cb0324c13535d8100e1e"),
				},
				CreatedAt: time.Now(),
			},
		}
		parsedWithdrawals, _, err := bridge.parseL2ToL1Transfers(l2ToL1TransferLogs)
		require.NoError(t, err)
		require.Len(t, parsedWithdrawals, 1)
		parsedWithdrawal := parsedWithdrawals[0]
		parsedFinalizations, err := bridge.parseL2ToL1Finalizations(finalizationLogs)
		require.NoError(t, err)
		require.Len(t, parsedFinalizations, 2)
		found, idx, err := bridge.findMatchingFinalization(parsedWithdrawal, parsedFinalizations)
		require.NoError(t, err)
		require.True(t, found)
		require.Equal(t, 1, idx)
	})

	t.Run("no matching finalization", func(t *testing.T) {
		l2Adapter, err := arbitrum_l2_bridge_adapter.NewArbitrumL2BridgeAdapter(utils.ZeroAddress, nil)
		require.NoError(t, err)
		l1Adapter, err := arbitrum_l1_bridge_adapter.NewArbitrumL1BridgeAdapter(utils.ZeroAddress, nil)
		require.NoError(t, err)
		bridge := &l2ToL1Bridge{
			l2BridgeAdapter: l2Adapter,
			l1BridgeAdapter: l1Adapter,
			// need to override this because the function will check if the transfer is destined for the rebalancer
			// these sample txes are destined for a test EOA on testnet
			l1RebalancerAddress: common.HexToAddress("0xd19e526da8b7e1cb970af3da92bfcef3e135f468"),
			lggr:                logger.TestLogger(t),
		}
		// l2 -> l1 real withdrawal
		// https://sepolia.arbiscan.io/tx/0x12012fdc48132435be2821b6f6fa0b7da5acb7872934620d4221efc60bfe0e4d#eventlog
		l2ToL1TransferLogs := []logpoller.Log{
			{
				EvmChainId:  utilsbig.New(big.NewInt(int64(arbSepolia))),
				Address:     common.HexToAddress("0x66e4037a15b5c15620b3cfac064f53ffba24361d"),
				BlockHash:   common.HexToHash("0x20e3784f8431bc9f5d5d0f830fb781217e4a1c46621d9f8735448dd7a542c0a7"),
				TxHash:      common.HexToHash("0x12012fdc48132435be2821b6f6fa0b7da5acb7872934620d4221efc60bfe0e4d"),
				LogIndex:    hexutil.MustDecodeBig("0x1c").Int64(),
				BlockNumber: hexutil.MustDecodeBig("0xc38f48").Int64(),
				Data:        hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000051fe"),
				Topics: pq.ByteaArray{
					hexutil.MustDecode("0xffa51f05034aaeaa4c0a623d14cae59dd4a579d2c66ed37f38c4baef2d504c01"),
					hexutil.MustDecode("0x000000000000000000000000980b62da83eff3d4576c647993b0c1d7faf17c73"),
					hexutil.MustDecode("0x0000000000000000000000007b79995e5f793a07bc00c21412e50ecae098e7f9"),
					hexutil.MustDecode("0x000000000000000000000000d19e526da8b7e1cb970af3da92bfcef3e135f468"),
				},
				EventSig:  common.HexToHash("0xffa51f05034aaeaa4c0a623d14cae59dd4a579d2c66ed37f38c4baef2d504c01"),
				CreatedAt: time.Now(),
			},
		}
		// l2 -> l1 real finalizations
		// 1. https://sepolia.etherscan.io/tx/0xe9219865868eea6654f15f7915924b3484c650400f7b7a9ed6fe3c5abcb50092#eventlog <- incorrect one
		// 2. https://sepolia.etherscan.io/tx/0xecc26b602811b33a146c0e2803f67b4c931d31c48db974c70b8acb137234183a#eventlog <- correct one
		finalizationLogs := []logpoller.Log{
			{
				EvmChainId:  utilsbig.New(big.NewInt(int64(sepoliaChainID))),
				Address:     common.HexToAddress("0x97dc5c3c16954a0305f0419ed1a527cca0650dab"),
				BlockHash:   common.HexToHash("0xebb5704d693dd2df8809eaf5564c306f23db452f503d58184628fded51d4b3e2"),
				TxHash:      common.HexToHash("0xe9219865868eea6654f15f7915924b3484c650400f7b7a9ed6fe3c5abcb50092"),
				EventSig:    common.HexToHash("0x72ecb2b30e3013e4dcbd23949793c0ad87766498cc85c69a637847c1a790b8f6"),
				Data:        hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000000000051fd000000000000000000000000cfb1f08a4852699a979909e22c30263ca249556d000000000000000000000000a8ad8d7e13cbf556ee75cb0324c13535d8100e1e0000000000000000000000000000000000000000000000000000000000c38bd20000000000000000000000000000000000000000000000000000000000501e7d0000000000000000000000000000000000000000000000000000000065c5e86e00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000320000000000000000000000000000000000000000000000000000000000000000f8b51fa17ea4f1dff4db1151de00b9868b9d6ba80ca7077a08dc55203fb6f698a6f1638fc7d51bcde617e57cfd8c89f669f5b9e5603e8f2ec48d55bbc5d1f57f99703436db5204779a947305958d7b285b918ec9cc919562c6293e182509be1982d766d4a0f1155f58d749967c4da5d0a3d275be0fa29c2a354ab41b061cd59fde9634b8838c95f4cfd7099ec952dc0593371f13d145cdf1b0738c89d225854f5c1cbee2f11e51364bb7b838b4a7631f3f46ccf9ea3664014c0e0c9564ccf737d92422592c2d063a2a466a4bd53d29c163f45fc7f2bd845b389c0bd89f688bb28b3b7492663240be7e70cec690b22ad5d43dff03ff1e264dec3e9982f3ec36f286edba00d9c21e36a7b567b1d6e4f828613c4d46328eca6fef826f8014de82a195a9351bbbf7c8f3e8c757b46a3fb28a114376fe8c77012b2f99dc20f5d0e932e000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004bde3fbfa26b03bfcbbc80884a5c229d987992ce2366bef5882742cac7cc7efe000000000000000000000000000000000000000000000000000000000000000037ef20c8140b820c15303931a9e839810ec66f4d9523c3fa8187434076a25b1300000000000000000000000000000000000000000000000000000000000001242e567b360000000000000000000000007b79995e5f793a07bc00c21412e50ecae098e7f900000000000000000000000066e4037a15b5c15620b3cfac064f53ffba24361d000000000000000000000000d19e526da8b7e1cb970af3da92bfcef3e135f468000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
				BlockNumber: hexutil.MustDecodeBig("0x502835").Int64(),
				LogIndex:    hexutil.MustDecodeBig("0x7f").Int64(),
				Topics: pq.ByteaArray{
					hexutil.MustDecode("0x72ecb2b30e3013e4dcbd23949793c0ad87766498cc85c69a637847c1a790b8f6"),
					hexutil.MustDecode("0x000000000000000000000000cfb1f08a4852699a979909e22c30263ca249556d"),
					hexutil.MustDecode("0x000000000000000000000000a8ad8d7e13cbf556ee75cb0324c13535d8100e1e"),
				},
				CreatedAt: time.Now(),
			},
			// correct one below, commented out
			//{
			//	EvmChainId:  utilsbig.New(big.NewInt(int64(sepoliaChainID))),
			//	Address:     common.HexToAddress("0x97dc5c3c16954a0305f0419ed1a527cca0650dab"),
			//	BlockHash:   common.HexToHash("0xc9e1989459d9c3640ac0b43d88a9845ed7a2cc836236c866e1f3be5464724849"),
			//	TxHash:      common.HexToHash("0xecc26b602811b33a146c0e2803f67b4c931d31c48db974c70b8acb137234183a"),
			//	EventSig:    common.HexToHash("0x72ecb2b30e3013e4dcbd23949793c0ad87766498cc85c69a637847c1a790b8f6"),
			//	Data:        hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000000000051fe000000000000000000000000cfb1f08a4852699a979909e22c30263ca249556d000000000000000000000000a8ad8d7e13cbf556ee75cb0324c13535d8100e1e0000000000000000000000000000000000000000000000000000000000c38f480000000000000000000000000000000000000000000000000000000000501e920000000000000000000000000000000000000000000000000000000065c5e97f00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000320000000000000000000000000000000000000000000000000000000000000000f6632b3f37ef2be542e585204f3ed7826c158f6414c7388d12579f35eaa10420970b8e564f4708b24f214d2dc9e360727341760abba72bec52f64cbf45d1573fd9703436db5204779a947305958d7b285b918ec9cc919562c6293e182509be1982d766d4a0f1155f58d749967c4da5d0a3d275be0fa29c2a354ab41b061cd59fde9634b8838c95f4cfd7099ec952dc0593371f13d145cdf1b0738c89d225854f5c1cbee2f11e51364bb7b838b4a7631f3f46ccf9ea3664014c0e0c9564ccf737d92422592c2d063a2a466a4bd53d29c163f45fc7f2bd845b389c0bd89f688bb28b3b7492663240be7e70cec690b22ad5d43dff03ff1e264dec3e9982f3ec36f286edba00d9c21e36a7b567b1d6e4f828613c4d46328eca6fef826f8014de82a19b33d2a5f7ba320d574bf4bbd933ab5f5b2a991ef38b7a35ef56e734e6ed4f1d0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004bde3fbfa26b03bfcbbc80884a5c229d987992ce2366bef5882742cac7cc7efe000000000000000000000000000000000000000000000000000000000000000037ef20c8140b820c15303931a9e839810ec66f4d9523c3fa8187434076a25b1300000000000000000000000000000000000000000000000000000000000001242e567b360000000000000000000000007b79995e5f793a07bc00c21412e50ecae098e7f900000000000000000000000066e4037a15b5c15620b3cfac064f53ffba24361d000000000000000000000000d19e526da8b7e1cb970af3da92bfcef3e135f468000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000090000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
			//	BlockNumber: hexutil.MustDecodeBig("0x5023e1").Int64(),
			//	LogIndex:    hexutil.MustDecodeBig("0xca").Int64(),
			//	Topics: pq.ByteaArray{
			//		hexutil.MustDecode("0x72ecb2b30e3013e4dcbd23949793c0ad87766498cc85c69a637847c1a790b8f6"),
			//		hexutil.MustDecode("0x000000000000000000000000cfb1f08a4852699a979909e22c30263ca249556d"),
			//		hexutil.MustDecode("0x000000000000000000000000a8ad8d7e13cbf556ee75cb0324c13535d8100e1e"),
			//	},
			//	CreatedAt: time.Now(),
			//},
		}
		parsedWithdrawals, _, err := bridge.parseL2ToL1Transfers(l2ToL1TransferLogs)
		require.NoError(t, err)
		require.Len(t, parsedWithdrawals, 1)
		parsedWithdrawal := parsedWithdrawals[0]
		parsedFinalizations, err := bridge.parseL2ToL1Finalizations(finalizationLogs)
		require.NoError(t, err)
		require.Len(t, parsedFinalizations, 1)
		found, idx, err := bridge.findMatchingFinalization(parsedWithdrawal, parsedFinalizations)
		require.NoError(t, err)
		require.False(t, found)
		require.Equal(t, -1, idx)
	})
}

func Test_l2ToL1Bridge_filterOutFinalizedTransfers(t *testing.T) {
	type fields struct {
		l1RebalancerAddress common.Address
	}
	type args struct {
		l2ToL1Transfers     []*arbitrum_l2_bridge_adapter.ArbitrumL2BridgeAdapterArbitrumL2ToL1ERC20Sent
		l2ToL1Finalizations []*arbitrum_l1_bridge_adapter.ArbitrumL1BridgeAdapterArbitrumL2ToL1ERC20Finalized
	}
	var (
		localToken  = testutils.NewAddress()
		remoteToken = testutils.NewAddress()
		recipient   = testutils.NewAddress()
	)
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*arbitrum_l2_bridge_adapter.ArbitrumL2BridgeAdapterArbitrumL2ToL1ERC20Sent
		wantErr bool
	}{
		{
			"no unfinalized transfers",
			fields{
				l1RebalancerAddress: recipient,
			},
			args{
				l2ToL1Transfers: []*arbitrum_l2_bridge_adapter.ArbitrumL2BridgeAdapterArbitrumL2ToL1ERC20Sent{
					genL2ToL1SentLog(t, localToken, remoteToken, recipient, big.NewInt(100), big.NewInt(1)),
					genL2ToL1SentLog(t, localToken, remoteToken, recipient, big.NewInt(200), big.NewInt(2)),
					genL2ToL1SentLog(t, localToken, remoteToken, recipient, big.NewInt(300), big.NewInt(3)),
				},
				l2ToL1Finalizations: []*arbitrum_l1_bridge_adapter.ArbitrumL1BridgeAdapterArbitrumL2ToL1ERC20Finalized{
					genL2ToL1FinalizedLog(
						t,
						testutils.NewAddress(), // remoteSender
						recipient,              // localReceiver
						big.NewInt(100),        // amount
						finalizeInboundTransferParams{
							amount:     big.NewInt(100),
							l1Receiver: recipient,
						},
						big.NewInt(1), // tx id
					),
					genL2ToL1FinalizedLog(
						t,
						testutils.NewAddress(), // remoteSender
						recipient,              // localReceiver
						big.NewInt(200),        // amount
						finalizeInboundTransferParams{
							amount:     big.NewInt(200),
							l1Receiver: recipient,
						},
						big.NewInt(2), // tx id
					),
					genL2ToL1FinalizedLog(
						t,
						testutils.NewAddress(), // remoteSender
						recipient,              // localReceiver
						big.NewInt(300),        // amount
						finalizeInboundTransferParams{
							amount:     big.NewInt(300),
							l1Receiver: recipient,
						},
						big.NewInt(3), // tx id
					),
				},
			},
			[]*arbitrum_l2_bridge_adapter.ArbitrumL2BridgeAdapterArbitrumL2ToL1ERC20Sent(nil),
			false,
		},
		{
			"finalizations not for rebalancer",
			fields{
				l1RebalancerAddress: testutils.NewAddress(),
			},
			args{
				l2ToL1Transfers: []*arbitrum_l2_bridge_adapter.ArbitrumL2BridgeAdapterArbitrumL2ToL1ERC20Sent{
					genL2ToL1SentLog(t, localToken, remoteToken, recipient, big.NewInt(100), big.NewInt(1)),
					genL2ToL1SentLog(t, localToken, remoteToken, recipient, big.NewInt(200), big.NewInt(2)),
					genL2ToL1SentLog(t, localToken, remoteToken, recipient, big.NewInt(300), big.NewInt(3)),
				},
				l2ToL1Finalizations: []*arbitrum_l1_bridge_adapter.ArbitrumL1BridgeAdapterArbitrumL2ToL1ERC20Finalized{
					genL2ToL1FinalizedLog(
						t,
						testutils.NewAddress(), // remoteSender
						recipient,              // localReceiver
						big.NewInt(100),        // amount
						finalizeInboundTransferParams{
							amount:     big.NewInt(100),
							l1Receiver: recipient,
						},
						big.NewInt(1), // tx id
					),
					genL2ToL1FinalizedLog(
						t,
						testutils.NewAddress(), // remoteSender
						recipient,              // localReceiver
						big.NewInt(200),        // amount
						finalizeInboundTransferParams{
							amount:     big.NewInt(200),
							l1Receiver: recipient,
						},
						big.NewInt(2), // tx id
					),
					genL2ToL1FinalizedLog(
						t,
						testutils.NewAddress(), // remoteSender
						recipient,              // localReceiver
						big.NewInt(300),        // amount
						finalizeInboundTransferParams{
							amount:     big.NewInt(300),
							l1Receiver: recipient,
						},
						big.NewInt(3), // tx id
					),
				},
			},
			[]*arbitrum_l2_bridge_adapter.ArbitrumL2BridgeAdapterArbitrumL2ToL1ERC20Sent(nil),
			false,
		},
		{
			"no unfinalized transfers",
			fields{
				l1RebalancerAddress: recipient,
			},
			args{
				l2ToL1Transfers: []*arbitrum_l2_bridge_adapter.ArbitrumL2BridgeAdapterArbitrumL2ToL1ERC20Sent{
					genL2ToL1SentLog(t, localToken, remoteToken, recipient, big.NewInt(100), big.NewInt(1)),
					genL2ToL1SentLog(t, localToken, remoteToken, recipient, big.NewInt(200), big.NewInt(2)),
					genL2ToL1SentLog(t, localToken, remoteToken, recipient, big.NewInt(300), big.NewInt(3)),
				},
				l2ToL1Finalizations: []*arbitrum_l1_bridge_adapter.ArbitrumL1BridgeAdapterArbitrumL2ToL1ERC20Finalized{
					genL2ToL1FinalizedLog(
						t,
						testutils.NewAddress(), // remoteSender
						recipient,              // localReceiver
						big.NewInt(100),        // amount
						finalizeInboundTransferParams{
							amount:     big.NewInt(100),
							l1Receiver: recipient,
						},
						big.NewInt(1), // tx id
					),
					genL2ToL1FinalizedLog(
						t,
						testutils.NewAddress(), // remoteSender
						recipient,              // localReceiver
						big.NewInt(200),        // amount
						finalizeInboundTransferParams{
							amount:     big.NewInt(200),
							l1Receiver: recipient,
						},
						big.NewInt(2), // tx id
					),
				},
			},
			[]*arbitrum_l2_bridge_adapter.ArbitrumL2BridgeAdapterArbitrumL2ToL1ERC20Sent{
				genL2ToL1SentLog(t, localToken, remoteToken, recipient, big.NewInt(300), big.NewInt(3)),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &l2ToL1Bridge{
				l1RebalancerAddress: tt.fields.l1RebalancerAddress,
				lggr:                logger.TestLogger(t),
			}
			got, err := l.filterOutFinalizedTransfers(tt.args.l2ToL1Transfers, tt.args.l2ToL1Finalizations)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			}
		})
	}
}

func Test_validateFinalizeInboundTransferABI(t *testing.T) {
	type args struct {
		tokenGatewayABI abi.ABI
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"valid abi",
			args{
				tokenGatewayABI: arbitrumTokenGatewayABI,
			},
			false,
		},
		{
			"invalid abi",
			args{
				tokenGatewayABI: abi.ABI{},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateFinalizeInboundTransferABI(tt.args.tokenGatewayABI)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func mustGetChainByID(t *testing.T, id uint64) chainsel.Chain {
	chain, ok := chainsel.ChainByEvmChainID(id)
	require.True(t, ok)
	return chain
}

func genL2ToL1SentLog(
	t *testing.T,
	localToken, remoteToken, recipient common.Address,
	amount, l2ToL1TxId *big.Int,
) *arbitrum_l2_bridge_adapter.ArbitrumL2BridgeAdapterArbitrumL2ToL1ERC20Sent {
	packedOutboundTransferResult, err := packUint256(l2ToL1TxId)
	require.NoError(t, err)
	return &arbitrum_l2_bridge_adapter.ArbitrumL2BridgeAdapterArbitrumL2ToL1ERC20Sent{
		LocalToken:             localToken,
		RemoteToken:            remoteToken,
		Recipient:              recipient,
		Amount:                 amount,
		OutboundTransferResult: packedOutboundTransferResult,
	}
}

func genL2ToL1FinalizedLog(
	t *testing.T,
	remoteSender,
	localReceiver common.Address,
	amount *big.Int,
	params finalizeInboundTransferParams,
	txId *big.Int,
) *arbitrum_l1_bridge_adapter.ArbitrumL1BridgeAdapterArbitrumL2ToL1ERC20Finalized {
	packed, err := arbitrumTokenGatewayABI.Pack(
		"finalizeInboundTransfer",
		params.l1Token,
		params.l2Sender,
		params.l1Receiver,
		params.amount,
		params.data,
	)
	require.NoError(t, err)
	return &arbitrum_l1_bridge_adapter.ArbitrumL1BridgeAdapterArbitrumL2ToL1ERC20Finalized{
		RemoteSender:  remoteSender,
		LocalReceiver: localReceiver,
		Amount:        amount,
		Payload: arbitrum_l1_bridge_adapter.ArbitrumL1BridgeAdapterArbitrumFinalizationPayload{
			Data:  packed,
			Index: txId,
		},
	}
}

func Test_l2ToL1Bridge_getLatestNodeConfirmed(t *testing.T) {
	type fields struct {
		l1LogPoller *lpmocks.LogPoller
		rollupCore  *mock_arbitrum_rollup_core.ArbRollupCoreInterface
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		want       *arbitrum_rollup_core.ArbRollupCoreNodeConfirmed
		wantErr    bool
		before     func(*testing.T, fields, *arbitrum_rollup_core.ArbRollupCoreNodeConfirmed)
		assertions func(*testing.T, fields)
	}{
		{
			"log found",
			fields{
				l1LogPoller: lpmocks.NewLogPoller(t),
				rollupCore:  mock_arbitrum_rollup_core.NewArbRollupCoreInterface(t),
			},
			args{
				ctx: testutils.Context(t),
			},
			&arbitrum_rollup_core.ArbRollupCoreNodeConfirmed{
				NodeNum:   1,
				BlockHash: testutils.Random32Byte(),
				SendRoot:  testutils.Random32Byte(),
			},
			false,
			func(t *testing.T, f fields, want *arbitrum_rollup_core.ArbRollupCoreNodeConfirmed) {
				topic1 := common.HexToHash(hexutil.EncodeUint64(want.NodeNum))
				data, err := utils.ABIEncode(`[{"type": "bytes32"}, {"type": "bytes32"}]`, want.BlockHash, want.SendRoot)
				require.NoError(t, err)
				rollupAddress := testutils.NewAddress()
				f.l1LogPoller.On("LatestLogByEventSigWithConfs", NodeConfirmedTopic, rollupAddress, logpoller.Finalized, mock.Anything).
					Return(&logpoller.Log{
						Topics: [][]byte{
							NodeConfirmedTopic[:],
							topic1[:],
						},
						Data: data,
					}, nil)
				f.rollupCore.On("Address").Return(rollupAddress)
				f.rollupCore.On("ParseNodeConfirmed", mock.Anything).Return(want, nil)
			},
			func(t *testing.T, f fields) {
				f.l1LogPoller.AssertExpectations(t)
				f.rollupCore.AssertExpectations(t)
			},
		},
		{
			"log not found",
			fields{
				l1LogPoller: lpmocks.NewLogPoller(t),
				rollupCore:  mock_arbitrum_rollup_core.NewArbRollupCoreInterface(t),
			},
			args{
				ctx: testutils.Context(t),
			},
			nil,
			true,
			func(t *testing.T, f fields, want *arbitrum_rollup_core.ArbRollupCoreNodeConfirmed) {
				rollupAddress := testutils.NewAddress()
				f.l1LogPoller.On("LatestLogByEventSigWithConfs", NodeConfirmedTopic, rollupAddress, logpoller.Finalized, mock.Anything).
					Return(nil, errors.New("not found"))
				f.rollupCore.On("Address").Return(rollupAddress)
			},
			func(t *testing.T, f fields) {
				f.l1LogPoller.AssertExpectations(t)
				f.rollupCore.AssertExpectations(t)
			},
		},
		{
			"parse error",
			fields{
				l1LogPoller: lpmocks.NewLogPoller(t),
				rollupCore:  mock_arbitrum_rollup_core.NewArbRollupCoreInterface(t),
			},
			args{
				ctx: testutils.Context(t),
			},
			&arbitrum_rollup_core.ArbRollupCoreNodeConfirmed{
				NodeNum:   1,
				BlockHash: testutils.Random32Byte(),
				SendRoot:  testutils.Random32Byte(),
			},
			true,
			func(t *testing.T, f fields, want *arbitrum_rollup_core.ArbRollupCoreNodeConfirmed) {
				topic1 := common.HexToHash(hexutil.EncodeUint64(want.NodeNum))
				data, err := utils.ABIEncode(`[{"type": "bytes32"}, {"type": "bytes32"}]`, want.BlockHash, want.SendRoot)
				require.NoError(t, err)
				rollupAddress := testutils.NewAddress()
				f.l1LogPoller.On("LatestLogByEventSigWithConfs", NodeConfirmedTopic, rollupAddress, logpoller.Finalized, mock.Anything).
					Return(&logpoller.Log{
						Topics: [][]byte{
							NodeConfirmedTopic[:],
							topic1[:],
						},
						Data: data,
					}, nil)
				f.rollupCore.On("Address").Return(rollupAddress)
				f.rollupCore.On("ParseNodeConfirmed", mock.Anything).Return(nil, errors.New("parse error"))
			},
			func(t *testing.T, f fields) {
				f.l1LogPoller.AssertExpectations(t)
				f.rollupCore.AssertExpectations(t)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &l2ToL1Bridge{
				l1LogPoller: tt.fields.l1LogPoller,
				rollupCore:  tt.fields.rollupCore,
			}
			if tt.before != nil {
				tt.before(t, tt.fields, tt.want)
				defer tt.assertions(t, tt.fields)
			}
			got, err := l.getLatestNodeConfirmed(tt.args.ctx)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			}
		})
	}
}
