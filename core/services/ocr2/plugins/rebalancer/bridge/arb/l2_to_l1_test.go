package arb

import (
	"context"
	"errors"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/mock"
	"github.com/test-go/testify/require"

	clientmocks "github.com/smartcontractkit/chainlink/v2/core/chains/evm/client/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	lpmocks "github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
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
	rollupAddress, l1RebalancerAddress, l2RebalancerAddress := testutils.NewAddress(), testutils.NewAddress(), testutils.NewAddress()

	t.Run("happy path", func(t *testing.T) {
		l2LogPoller := lpmocks.NewLogPoller(t)
		l2LogPoller.On("RegisterFilter", mock.MatchedBy(func(f logpoller.Filter) bool {
			if len(f.EventSigs) != 1 {
				return false
			}
			if f.EventSigs[0] != LiquidityTransferredTopic {
				return false
			}
			if f.Retention != DurationMonth {
				return false
			}
			if f.Addresses[0] != l2RebalancerAddress {
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
			if f.EventSigs[0] != NodeConfirmedTopic || f.EventSigs[1] != LiquidityTransferredTopic {
				return false
			}
			if f.Retention != DurationMonth {
				return false
			}
			if f.Addresses[0] != rollupAddress || f.Addresses[1] != l1RebalancerAddress {
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
			l2LogPoller,
			l1LogPoller,
			l2Client,
			l1Client,
		)
		require.NoError(t, err)

		require.Equal(t, rollupAddress, bridge.rollupCore.Address())
		require.Equal(t, l1RebalancerAddress, bridge.l1Rebalancer.Address())
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

func mustGetChainByID(t *testing.T, id uint64) chainsel.Chain {
	chain, ok := chainsel.ChainByEvmChainID(id)
	require.True(t, ok)
	return chain
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
