package testonlybridge

import (
	"context"
	"fmt"
	"math/big"

	"golang.org/x/exp/constraints"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
	ubig "github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils/big"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/mock_l1_bridge_adapter"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/rebalancer"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/models"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
)

var (
	// Emitted on the source
	MockERC20SentTopic = mock_l1_bridge_adapter.MockL1BridgeAdapterMockERC20Sent{}.Topic()
	// Emitted on the destination
	MockERC20FinalizedTopic = mock_l1_bridge_adapter.MockL1BridgeAdapterMockERC20Finalized{}.Topic()
	// Emitted on both source and destination
	LiquidityTransferredTopic = rebalancer.RebalancerLiquidityTransferred{}.Topic()
)

type testBridge struct {
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

func New(
	sourceSelector, destSelector models.NetworkSelector,
	sourceRebalancer, destRebalancer, sourceAdapter, destAdapter models.Address,
	sourceLogPoller, destLogPoller logpoller.LogPoller,
	sourceClient, destClient client.Client,
	lggr logger.Logger,
) (*testBridge, error) {
	err := sourceLogPoller.RegisterFilter(logpoller.Filter{
		Name: logpoller.FilterName("MockERC20Sent", sourceSelector),
		EventSigs: []common.Hash{
			MockERC20SentTopic,
			LiquidityTransferredTopic,
		},
		Addresses: []common.Address{
			common.Address(sourceAdapter),
			common.Address(sourceRebalancer),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to register filter for source log poller: %w", err)
	}

	err = destLogPoller.RegisterFilter(logpoller.Filter{
		Name: logpoller.FilterName("MockERC20Finalized", destSelector),
		EventSigs: []common.Hash{
			MockERC20FinalizedTopic,
			LiquidityTransferredTopic,
		},
		Addresses: []common.Address{
			common.Address(destAdapter),
			common.Address(destRebalancer),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to register filter for dest log poller: %w", err)
	}

	lggr = lggr.Named("TestBridge").With(
		"sourceSelector", sourceSelector,
		"destSelector", destSelector,
		"sourceRebalancer", sourceRebalancer,
		"destRebalancer", destRebalancer,
		"sourceAdapter", sourceAdapter,
		"destAdapter", destAdapter,
	)
	lggr.Infow("TestBridge created")

	sourceAdapterWrapper, err := mock_l1_bridge_adapter.NewMockL1BridgeAdapter(common.Address(sourceAdapter), sourceClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create source adapter wrapper: %w", err)
	}

	destAdapterWrapper, err := mock_l1_bridge_adapter.NewMockL1BridgeAdapter(common.Address(destAdapter), destClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create dest adapter wrapper: %w", err)
	}

	return &testBridge{
		sourceSelector:   sourceSelector,
		destSelector:     destSelector,
		sourceRebalancer: sourceRebalancer,
		destRebalancer:   destRebalancer,
		sourceAdapter:    sourceAdapterWrapper,
		destAdapter:      destAdapterWrapper,
		sourceLogPoller:  sourceLogPoller,
		destLogPoller:    destLogPoller,
		sourceClient:     sourceClient,
		destClient:       destClient,
		lggr:             lggr,
	}, nil
}

// Close implements bridge.Bridge.
func (t *testBridge) Close(ctx context.Context) error {
	return nil
}

// QuorumizedBridgePayload implements bridge.Bridge.
func (t *testBridge) QuorumizedBridgePayload(payloads [][]byte) ([]byte, error) {
	// TODO: implement, should just return Amount and they should all be the same
	return payloads[0], nil
}

// GetBridgePayloadAndFee implements bridge.Bridge.
func (t *testBridge) GetBridgePayloadAndFee(ctx context.Context, transfer models.Transfer) ([]byte, *big.Int, error) {
	payload, err := PackSendBridgePayload(transfer.Amount.ToInt())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to pack bridge data: %w", err)
	}
	return payload, big.NewInt(0), nil
}

// GetTransfers implements bridge.Bridge.
func (t *testBridge) GetTransfers(ctx context.Context, localToken models.Address, remoteToken models.Address) ([]models.PendingTransfer, error) {
	latestSourceBlock, err := t.sourceLogPoller.LatestBlock(pg.WithParentCtx(ctx))
	if err != nil {
		return nil, fmt.Errorf("failed to get latest block: %w", err)
	}

	latestDestBlock, err := t.destLogPoller.LatestBlock(pg.WithParentCtx(ctx))
	if err != nil {
		return nil, fmt.Errorf("failed to get latest block: %w", err)
	}

	sourceSendLogs, err := t.sourceLogPoller.LogsWithSigs(
		1,
		latestSourceBlock.BlockNumber,
		[]common.Hash{MockERC20SentTopic},
		t.sourceAdapter.Address(),
		pg.WithParentCtx(ctx),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get source MockERC20Sent logs: %w", err)
	}

	destFinalizeLogs, err := t.destLogPoller.LogsWithSigs(
		1,
		latestDestBlock.BlockNumber,
		[]common.Hash{MockERC20FinalizedTopic},
		t.destAdapter.Address(),
		pg.WithParentCtx(ctx),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get dest MockERC20Finalized logs: %w", err)
	}

	parsedSendLogs, parsedToLP, err := t.parseSendLogs(sourceSendLogs)
	if err != nil {
		return nil, fmt.Errorf("failed to parse source send logs: %w", err)
	}

	parsedFinalizeLogs, err := t.parseFinalizedLogs(destFinalizeLogs)
	if err != nil {
		return nil, fmt.Errorf("failed to parse dest finalize logs: %w", err)
	}

	ready, err := t.getReadyToFinalize(parsedSendLogs, parsedFinalizeLogs)
	if err != nil {
		return nil, fmt.Errorf("failed to get ready to finalize: %w", err)
	}

	return t.toPendingTransfers(ready, parsedToLP), nil
}

func (t *testBridge) toPendingTransfers(
	ready []*mock_l1_bridge_adapter.MockL1BridgeAdapterMockERC20Sent,
	parsedToLP map[logKey]logpoller.Log,
) []models.PendingTransfer {
	var transfers []models.PendingTransfer
	for _, send := range ready {
		lp := parsedToLP[logKey{txHash: send.Raw.TxHash, logIdx: int64(send.Raw.Index)}]
		bridgeData, err := PackFinalizeBridgePayload(send.Amount, send.Nonce)
		if err != nil {
			t.lggr.Errorw("failed to pack bridge data", "err", err)
			continue
		}
		transfers = append(transfers, models.PendingTransfer{
			Transfer: models.Transfer{
				From:               t.sourceSelector,
				To:                 t.destSelector,
				Sender:             models.Address(t.sourceAdapter.Address()),
				Receiver:           t.destRebalancer,
				Amount:             ubig.New(send.Amount),
				LocalTokenAddress:  models.Address(send.LocalToken),
				RemoteTokenAddress: models.Address(send.RemoteToken),
				Date:               lp.BlockTimestamp,
				BridgeData:         bridgeData,
			},
			Status: models.TransferStatusReady,
			ID:     fmt.Sprintf("%s-%d", send.Raw.TxHash.Hex(), send.Raw.Index),
		})
	}

	if len(transfers) > 0 {
		t.lggr.Infow("produced pending transfers", "pendingTransfers", transfers)
	}

	return transfers
}

func (t *testBridge) getReadyToFinalize(
	sends []*mock_l1_bridge_adapter.MockL1BridgeAdapterMockERC20Sent,
	finalizes []*mock_l1_bridge_adapter.MockL1BridgeAdapterMockERC20Finalized,
) ([]*mock_l1_bridge_adapter.MockL1BridgeAdapterMockERC20Sent, error) {
	t.lggr.Debugw("Getting ready to finalize",
		"sendsLen", len(sends),
		"finalizesLen", len(finalizes),
		"sends", sends,
		"finalizes", finalizes)

	// find sent events that don't have a matching finalized event
	var ready []*mock_l1_bridge_adapter.MockL1BridgeAdapterMockERC20Sent
	for _, send := range sends {
		var finalized bool
		for _, finalize := range finalizes {
			if send.Nonce.Cmp(finalize.Nonce) == 0 {
				finalized = true
				break
			}
		}
		if !finalized {
			ready = append(ready, send)
		}
	}

	if len(ready) > 0 {
		t.lggr.Infow("found ready to finalize", "sendsLen", len(sends),
			"finalizesLen", len(finalizes),
			"sends", sends,
			"finalizes", finalizes,
			"ready", ready)
	} else {
		t.lggr.Debugw("no requests ready to finalize", "sendsLen", len(sends),
			"finalizesLen", len(finalizes),
			"sends", sends,
			"finalizes", finalizes)
	}

	return ready, nil
}

func (t *testBridge) parseFinalizedLogs(logs []logpoller.Log) ([]*mock_l1_bridge_adapter.MockL1BridgeAdapterMockERC20Finalized, error) {
	var parsedFinalizeLogs []*mock_l1_bridge_adapter.MockL1BridgeAdapterMockERC20Finalized
	for _, log := range logs {
		finalizeLog, err := t.destAdapter.ParseMockERC20Finalized(log.ToGethLog())
		if err != nil {
			return nil, fmt.Errorf("failed to parse finalize log: %w", err)
		}
		parsedFinalizeLogs = append(parsedFinalizeLogs, finalizeLog)
	}
	return parsedFinalizeLogs, nil
}

type logKey struct {
	txHash common.Hash
	logIdx int64
}

func (t *testBridge) parseSendLogs(logs []logpoller.Log) (
	[]*mock_l1_bridge_adapter.MockL1BridgeAdapterMockERC20Sent,
	map[logKey]logpoller.Log,
	error,
) {
	var parsedSendLogs []*mock_l1_bridge_adapter.MockL1BridgeAdapterMockERC20Sent
	parsedToLP := make(map[logKey]logpoller.Log)
	for _, log := range logs {
		sendLog, err := t.sourceAdapter.ParseMockERC20Sent(log.ToGethLog())
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse send log: %w", err)
		}
		parsedSendLogs = append(parsedSendLogs, sendLog)
		parsedToLP[logKey{txHash: log.TxHash, logIdx: log.LogIndex}] = log
	}
	return parsedSendLogs, parsedToLP, nil
}

func intComparator[T constraints.Integer](a, b T) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

func PackFinalizeBridgePayload(val1, val2 *big.Int) ([]byte, error) {
	return utils.ABIEncode(`[{"type": "uint256"}, {"type": "uint256"}]`, val1, val2)
}

func UnpackFinalizeBridgePayload(data []byte) (*big.Int, *big.Int, error) {
	ifaces, err := utils.ABIDecode(`[{"type": "uint256"}, {"type": "uint256"}]`, data)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to decode bridge data: %w", err)
	}
	if len(ifaces) != 2 {
		return nil, nil, fmt.Errorf("expected 2 arguments, got %d", len(ifaces))
	}
	val1 := *abi.ConvertType(ifaces[0], new(*big.Int)).(**big.Int)
	val2 := *abi.ConvertType(ifaces[1], new(*big.Int)).(**big.Int)
	return val1, val2, nil
}

func PackSendBridgePayload(val *big.Int) ([]byte, error) {
	return utils.ABIEncode(`[{"type": "uint256"}]`, val)
}

func UnpackSendBridgePayload(data []byte) (*big.Int, error) {
	ifaces, err := utils.ABIDecode(`[{"type": "uint256"}]`, data)
	if err != nil {
		return nil, fmt.Errorf("failed to decode bridge data: %w", err)
	}
	if len(ifaces) != 1 {
		return nil, fmt.Errorf("expected 1 argument, got %d", len(ifaces))
	}
	val := *abi.ConvertType(ifaces[0], new(*big.Int)).(**big.Int)
	return val, nil
}
