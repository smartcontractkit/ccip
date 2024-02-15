package testonlybridge

import (
	"context"
	"math/big"
	"reflect"
	"testing"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/mock_l1_bridge_adapter"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/models"
)

func TestNew(t *testing.T) {
	type args struct {
		sourceSelector   models.NetworkSelector
		destSelector     models.NetworkSelector
		sourceRebalancer models.Address
		destRebalancer   models.Address
		sourceAdapter    models.Address
		destAdapter      models.Address
		sourceLogPoller  logpoller.LogPoller
		destLogPoller    logpoller.LogPoller
		sourceClient     client.Client
		destClient       client.Client
		lggr             logger.Logger
	}
	tests := []struct {
		name    string
		args    args
		want    *testBridge
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.sourceSelector, tt.args.destSelector, tt.args.sourceRebalancer, tt.args.destRebalancer, tt.args.sourceAdapter, tt.args.destAdapter, tt.args.sourceLogPoller, tt.args.destLogPoller, tt.args.sourceClient, tt.args.destClient, tt.args.lggr)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPackFinalizeBridgePayload(t *testing.T) {
	type args struct {
		val1 *big.Int
		val2 *big.Int
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PackFinalizeBridgePayload(tt.args.val1, tt.args.val2)
			if (err != nil) != tt.wantErr {
				t.Errorf("PackFinalizeBridgePayload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PackFinalizeBridgePayload() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPackSendBridgePayload(t *testing.T) {
	type args struct {
		val *big.Int
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PackSendBridgePayload(tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("PackSendBridgePayload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PackSendBridgePayload() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnpackFinalizeBridgePayload(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Int
		want1   *big.Int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := UnpackFinalizeBridgePayload(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnpackFinalizeBridgePayload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnpackFinalizeBridgePayload() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("UnpackFinalizeBridgePayload() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestUnpackSendBridgePayload(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnpackSendBridgePayload(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnpackSendBridgePayload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnpackSendBridgePayload() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_testBridge_Close(t1 *testing.T) {
	type fields struct {
		sourceSelector   models.NetworkSelector
		destSelector     models.NetworkSelector
		sourceRebalancer models.Address
		destRebalancer   models.Address
		sourceAdapter    *mock_l1_bridge_adapter.MockL1BridgeAdapter
		destAdapter      *mock_l1_bridge_adapter.MockL1BridgeAdapter
		sourceLogPoller  logpoller.LogPoller
		destLogPoller    logpoller.LogPoller
		sourceClient     client.Client
		destClient       client.Client
		lggr             logger.Logger
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &testBridge{
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
			if err := t.Close(tt.args.ctx); (err != nil) != tt.wantErr {
				t1.Errorf("Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_testBridge_GetBridgePayloadAndFee(t1 *testing.T) {
	type fields struct {
		sourceSelector   models.NetworkSelector
		destSelector     models.NetworkSelector
		sourceRebalancer models.Address
		destRebalancer   models.Address
		sourceAdapter    *mock_l1_bridge_adapter.MockL1BridgeAdapter
		destAdapter      *mock_l1_bridge_adapter.MockL1BridgeAdapter
		sourceLogPoller  logpoller.LogPoller
		destLogPoller    logpoller.LogPoller
		sourceClient     client.Client
		destClient       client.Client
		lggr             logger.Logger
	}
	type args struct {
		ctx      context.Context
		transfer models.Transfer
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		want1   *big.Int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &testBridge{
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
			got, got1, err := t.GetBridgePayloadAndFee(tt.args.ctx, tt.args.transfer)
			if (err != nil) != tt.wantErr {
				t1.Errorf("GetBridgePayloadAndFee() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetBridgePayloadAndFee() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t1.Errorf("GetBridgePayloadAndFee() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_testBridge_GetTransfers(t1 *testing.T) {
	type fields struct {
		sourceSelector   models.NetworkSelector
		destSelector     models.NetworkSelector
		sourceRebalancer models.Address
		destRebalancer   models.Address
		sourceAdapter    *mock_l1_bridge_adapter.MockL1BridgeAdapter
		destAdapter      *mock_l1_bridge_adapter.MockL1BridgeAdapter
		sourceLogPoller  logpoller.LogPoller
		destLogPoller    logpoller.LogPoller
		sourceClient     client.Client
		destClient       client.Client
		lggr             logger.Logger
	}
	type args struct {
		ctx         context.Context
		localToken  models.Address
		remoteToken models.Address
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.PendingTransfer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &testBridge{
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
			got, err := t.GetTransfers(tt.args.ctx, tt.args.localToken, tt.args.remoteToken)
			if (err != nil) != tt.wantErr {
				t1.Errorf("GetTransfers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetTransfers() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_testBridge_QuorumizedBridgePayload(t1 *testing.T) {
	type fields struct {
		sourceSelector   models.NetworkSelector
		destSelector     models.NetworkSelector
		sourceRebalancer models.Address
		destRebalancer   models.Address
		sourceAdapter    *mock_l1_bridge_adapter.MockL1BridgeAdapter
		destAdapter      *mock_l1_bridge_adapter.MockL1BridgeAdapter
		sourceLogPoller  logpoller.LogPoller
		destLogPoller    logpoller.LogPoller
		sourceClient     client.Client
		destClient       client.Client
		lggr             logger.Logger
	}
	type args struct {
		payloads [][]byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &testBridge{
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
			got, err := t.QuorumizedBridgePayload(tt.args.payloads, 1)
			if (err != nil) != tt.wantErr {
				t1.Errorf("QuorumizedBridgePayload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("QuorumizedBridgePayload() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_testBridge_getReadyToFinalize(t1 *testing.T) {
	type fields struct {
		sourceSelector   models.NetworkSelector
		destSelector     models.NetworkSelector
		sourceRebalancer models.Address
		destRebalancer   models.Address
		sourceAdapter    *mock_l1_bridge_adapter.MockL1BridgeAdapter
		destAdapter      *mock_l1_bridge_adapter.MockL1BridgeAdapter
		sourceLogPoller  logpoller.LogPoller
		destLogPoller    logpoller.LogPoller
		sourceClient     client.Client
		destClient       client.Client
		lggr             logger.Logger
	}
	type args struct {
		sends     []*mock_l1_bridge_adapter.MockL1BridgeAdapterMockERC20Sent
		finalizes []*mock_l1_bridge_adapter.MockL1BridgeAdapterMockERC20Finalized
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*mock_l1_bridge_adapter.MockL1BridgeAdapterMockERC20Sent
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &testBridge{
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
			got, err := t.getReadyToFinalize(tt.args.sends, tt.args.finalizes)
			if (err != nil) != tt.wantErr {
				t1.Errorf("getReadyToFinalize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("getReadyToFinalize() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_testBridge_parseFinalizedLogs(t1 *testing.T) {
	type fields struct {
		sourceSelector   models.NetworkSelector
		destSelector     models.NetworkSelector
		sourceRebalancer models.Address
		destRebalancer   models.Address
		sourceAdapter    *mock_l1_bridge_adapter.MockL1BridgeAdapter
		destAdapter      *mock_l1_bridge_adapter.MockL1BridgeAdapter
		sourceLogPoller  logpoller.LogPoller
		destLogPoller    logpoller.LogPoller
		sourceClient     client.Client
		destClient       client.Client
		lggr             logger.Logger
	}
	type args struct {
		logs []logpoller.Log
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*mock_l1_bridge_adapter.MockL1BridgeAdapterMockERC20Finalized
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &testBridge{
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
			got, err := t.parseFinalizedLogs(tt.args.logs)
			if (err != nil) != tt.wantErr {
				t1.Errorf("parseFinalizedLogs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("parseFinalizedLogs() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_testBridge_parseSendLogs(t1 *testing.T) {
	type fields struct {
		sourceSelector   models.NetworkSelector
		destSelector     models.NetworkSelector
		sourceRebalancer models.Address
		destRebalancer   models.Address
		sourceAdapter    *mock_l1_bridge_adapter.MockL1BridgeAdapter
		destAdapter      *mock_l1_bridge_adapter.MockL1BridgeAdapter
		sourceLogPoller  logpoller.LogPoller
		destLogPoller    logpoller.LogPoller
		sourceClient     client.Client
		destClient       client.Client
		lggr             logger.Logger
	}
	type args struct {
		logs []logpoller.Log
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*mock_l1_bridge_adapter.MockL1BridgeAdapterMockERC20Sent
		want1   map[logKey]logpoller.Log
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &testBridge{
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
			got, got1, err := t.parseSendLogs(tt.args.logs)
			if (err != nil) != tt.wantErr {
				t1.Errorf("parseSendLogs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("parseSendLogs() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t1.Errorf("parseSendLogs() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_testBridge_toPendingTransfers(t1 *testing.T) {
	type fields struct {
		sourceSelector   models.NetworkSelector
		destSelector     models.NetworkSelector
		sourceRebalancer models.Address
		destRebalancer   models.Address
		sourceAdapter    *mock_l1_bridge_adapter.MockL1BridgeAdapter
		destAdapter      *mock_l1_bridge_adapter.MockL1BridgeAdapter
		sourceLogPoller  logpoller.LogPoller
		destLogPoller    logpoller.LogPoller
		sourceClient     client.Client
		destClient       client.Client
		lggr             logger.Logger
	}
	type args struct {
		ready      []*mock_l1_bridge_adapter.MockL1BridgeAdapterMockERC20Sent
		parsedToLP map[logKey]logpoller.Log
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
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &testBridge{
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
			if got := t.toPendingTransfers(tt.args.ready, tt.args.parsedToLP); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("toPendingTransfers() = %v, want %v", got, tt.want)
			}
		})
	}
}
