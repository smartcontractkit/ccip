package testonlybridge

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/mock_l1_bridge_adapter"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/rebalancer"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/models"
	"github.com/stretchr/testify/require"
)

func Test_testBridge_toPendingTransfers(t *testing.T) {
	var (
		sourceSelector = models.NetworkSelector(1)
		destSelector   = models.NetworkSelector(2)
	)
	type fields struct {
		sourceSelector   models.NetworkSelector
		destSelector     models.NetworkSelector
		sourceRebalancer rebalancer.RebalancerInterface
		destRebalancer   rebalancer.RebalancerInterface
		sourceAdapter    *mock_l1_bridge_adapter.MockL1BridgeAdapter
		destAdapter      *mock_l1_bridge_adapter.MockL1BridgeAdapter
		sourceLogPoller  logpoller.LogPoller
		destLogPoller    logpoller.LogPoller
		sourceClient     client.Client
		destClient       client.Client
		lggr             logger.Logger
	}
	type args struct {
		localToken      models.Address
		remoteToken     models.Address
		readyToProve    []*rebalancer.RebalancerLiquidityTransferred
		readyToFinalize []*rebalancer.RebalancerLiquidityTransferred
		parsedToLP      map[logKey]logpoller.Log
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []models.PendingTransfer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &testBridge{
				sourceSelector:   tt.fields.sourceSelector,
				destSelector:     tt.fields.destSelector,
				sourceRebalancer: tt.fields.sourceRebalancer,
				destRebalancer:   tt.fields.destRebalancer,
				sourceAdapter:    tt.fields.sourceAdapter,
				destAdapter:      tt.fields.destAdapter,
				sourceLogPoller:  tt.fields.sourceLogPoller,
				destLogPoller:    tt.fields.destLogPoller,
				sourceClient:     tt.fields.sourceClient,
				destClient:       tt.fields.destClient,
				lggr:             tt.fields.lggr,
			}
			if got := tr.toPendingTransfers(tt.args.localToken, tt.args.remoteToken, tt.args.readyToProve, tt.args.readyToFinalize, tt.args.parsedToLP); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("testBridge.toPendingTransfers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_filterFinalized(t *testing.T) {
	type args struct {
		sends     []*rebalancer.RebalancerLiquidityTransferred
		finalizes []*rebalancer.RebalancerLiquidityTransferred
	}
	tests := []struct {
		name    string
		args    args
		want    []*rebalancer.RebalancerLiquidityTransferred
		wantErr bool
	}{
		{
			"no finalizes",
			args{
				[]*rebalancer.RebalancerLiquidityTransferred{
					{
						Amount:           big.NewInt(1),
						BridgeReturnData: mustPackSendReturnData(t, big.NewInt(1)),
					},
					{
						Amount:           big.NewInt(2),
						BridgeReturnData: mustPackSendReturnData(t, big.NewInt(2)),
					},
					{
						Amount:           big.NewInt(3),
						BridgeReturnData: mustPackSendReturnData(t, big.NewInt(3)),
					},
				},
				[]*rebalancer.RebalancerLiquidityTransferred{},
			},
			[]*rebalancer.RebalancerLiquidityTransferred{
				{
					Amount:           big.NewInt(1),
					BridgeReturnData: mustPackSendReturnData(t, big.NewInt(1)),
				},
				{
					Amount:           big.NewInt(2),
					BridgeReturnData: mustPackSendReturnData(t, big.NewInt(2)),
				},
				{
					Amount:           big.NewInt(3),
					BridgeReturnData: mustPackSendReturnData(t, big.NewInt(3)),
				},
			},
			false,
		},
		{
			"some finalizes",
			args{
				[]*rebalancer.RebalancerLiquidityTransferred{
					{
						Amount:           big.NewInt(1),
						BridgeReturnData: mustPackSendReturnData(t, big.NewInt(1)),
					},
					{
						Amount:           big.NewInt(2),
						BridgeReturnData: mustPackSendReturnData(t, big.NewInt(2)),
					},
					{
						Amount:           big.NewInt(3),
						BridgeReturnData: mustPackSendReturnData(t, big.NewInt(3)),
					},
				},
				[]*rebalancer.RebalancerLiquidityTransferred{
					{
						Amount:             big.NewInt(2),
						BridgeSpecificData: mustPackFinalizePayload(t, big.NewInt(2), big.NewInt(2)),
					},
				},
			},
			[]*rebalancer.RebalancerLiquidityTransferred{
				{
					Amount:           big.NewInt(1),
					BridgeReturnData: mustPackSendReturnData(t, big.NewInt(1)),
				},
				{
					Amount:           big.NewInt(3),
					BridgeReturnData: mustPackSendReturnData(t, big.NewInt(3)),
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := filterFinalized(tt.args.sends, tt.args.finalizes)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			}
		})
	}
}

func Test_groupByStage(t *testing.T) {
	type args struct {
		unfinalized    []*rebalancer.RebalancerLiquidityTransferred
		stepsCompleted []*rebalancer.RebalancerFinalizationStepCompleted
	}
	tests := []struct {
		name                string
		args                args
		wantReadyToProve    []*rebalancer.RebalancerLiquidityTransferred
		wantReadyToFinalize []*rebalancer.RebalancerLiquidityTransferred
		wantErr             bool
	}{
		{
			"all ready to prove",
			args{
				[]*rebalancer.RebalancerLiquidityTransferred{
					{
						Amount:           big.NewInt(1),
						BridgeReturnData: mustPackSendReturnData(t, big.NewInt(1)),
					},
					{
						Amount:           big.NewInt(2),
						BridgeReturnData: mustPackSendReturnData(t, big.NewInt(2)),
					},
					{
						Amount:           big.NewInt(3),
						BridgeReturnData: mustPackSendReturnData(t, big.NewInt(3)),
					},
				},
				[]*rebalancer.RebalancerFinalizationStepCompleted{}, // none proven
			},
			[]*rebalancer.RebalancerLiquidityTransferred{
				{
					Amount:           big.NewInt(1),
					BridgeReturnData: mustPackSendReturnData(t, big.NewInt(1)),
				},
				{
					Amount:           big.NewInt(2),
					BridgeReturnData: mustPackSendReturnData(t, big.NewInt(2)),
				},
				{
					Amount:           big.NewInt(3),
					BridgeReturnData: mustPackSendReturnData(t, big.NewInt(3)),
				},
			},
			nil,
			false,
		},
		{
			"all ready to finalize",
			args{
				[]*rebalancer.RebalancerLiquidityTransferred{
					{
						Amount:           big.NewInt(1),
						BridgeReturnData: mustPackSendReturnData(t, big.NewInt(1)),
					},
					{
						Amount:           big.NewInt(2),
						BridgeReturnData: mustPackSendReturnData(t, big.NewInt(2)),
					},
					{
						Amount:           big.NewInt(3),
						BridgeReturnData: mustPackSendReturnData(t, big.NewInt(3)),
					},
				},
				[]*rebalancer.RebalancerFinalizationStepCompleted{
					{
						BridgeSpecificData: mustPackProvePayload(t, big.NewInt(1)),
					},
					{
						BridgeSpecificData: mustPackProvePayload(t, big.NewInt(2)),
					},
					{
						BridgeSpecificData: mustPackProvePayload(t, big.NewInt(3)),
					},
				},
			},
			nil,
			[]*rebalancer.RebalancerLiquidityTransferred{
				{
					Amount:           big.NewInt(1),
					BridgeReturnData: mustPackSendReturnData(t, big.NewInt(1)),
				},
				{
					Amount:           big.NewInt(2),
					BridgeReturnData: mustPackSendReturnData(t, big.NewInt(2)),
				},
				{
					Amount:           big.NewInt(3),
					BridgeReturnData: mustPackSendReturnData(t, big.NewInt(3)),
				},
			},
			false,
		},
		{
			"mix of ready to prove and ready to finalize",
			args{
				[]*rebalancer.RebalancerLiquidityTransferred{
					{
						Amount:           big.NewInt(1),
						BridgeReturnData: mustPackSendReturnData(t, big.NewInt(1)),
					},
					{
						Amount:           big.NewInt(2),
						BridgeReturnData: mustPackSendReturnData(t, big.NewInt(2)),
					},
					{
						Amount:           big.NewInt(3),
						BridgeReturnData: mustPackSendReturnData(t, big.NewInt(3)),
					},
					{
						Amount:           big.NewInt(4),
						BridgeReturnData: mustPackSendReturnData(t, big.NewInt(4)),
					},
				},
				[]*rebalancer.RebalancerFinalizationStepCompleted{ // 1 and 3 already proven, ready to finalize
					{
						BridgeSpecificData: mustPackProvePayload(t, big.NewInt(1)),
					},
					{
						BridgeSpecificData: mustPackProvePayload(t, big.NewInt(3)),
					},
				},
			},
			[]*rebalancer.RebalancerLiquidityTransferred{
				{
					Amount:           big.NewInt(2),
					BridgeReturnData: mustPackSendReturnData(t, big.NewInt(2)),
				},
				{
					Amount:           big.NewInt(4),
					BridgeReturnData: mustPackSendReturnData(t, big.NewInt(4)),
				},
			},
			[]*rebalancer.RebalancerLiquidityTransferred{
				{
					Amount:           big.NewInt(1),
					BridgeReturnData: mustPackSendReturnData(t, big.NewInt(1)),
				},
				{
					Amount:           big.NewInt(3),
					BridgeReturnData: mustPackSendReturnData(t, big.NewInt(3)),
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotReadyToProve, gotReadyToFinalize, err := groupByStage(tt.args.unfinalized, tt.args.stepsCompleted)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.wantReadyToProve, gotReadyToProve)
				require.Equal(t, tt.wantReadyToFinalize, gotReadyToFinalize)
			}
		})
	}
}

func mustPackSendReturnData(t *testing.T, nonce *big.Int) []byte {
	packed, err := utils.ABIEncode(`[{"type": "uint256"}]`, nonce)
	require.NoError(t, err)
	return packed
}

func mustPackFinalizePayload(t *testing.T, nonce, amount *big.Int) []byte {
	packed, err := PackFinalizeBridgePayload(amount, nonce)
	require.NoError(t, err)
	return packed
}

func mustPackProvePayload(t *testing.T, nonce *big.Int) []byte {
	packed, err := PackProveBridgePayload(nonce)
	require.NoError(t, err)
	return packed
}
