package testonlybridge

import (
	"context"
	"fmt"
	"math/big"
	"slices"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/mock_l1_bridge_adapter"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/rebalancer"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/models"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
	"golang.org/x/exp/constraints"
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
	currSourceBlock  int64
	currDestBlock    int64
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
	payload, err := packUint256(transfer.Amount)
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

	if latestSourceBlock.BlockNumber <= t.currSourceBlock && latestDestBlock.BlockNumber <= t.currDestBlock {
		t.lggr.Debugw("No new blocks since last poll",
			"latestBlock", latestSourceBlock.BlockNumber,
			"currSourceBlock", t.currSourceBlock,
			"latestDestBlock", latestDestBlock.BlockNumber,
			"currDestBlock", t.currDestBlock)
		return nil, nil
	}

	sourceSendLogs, err := t.sourceLogPoller.LogsWithSigs(
		t.currSourceBlock,
		latestSourceBlock.BlockNumber,
		[]common.Hash{MockERC20SentTopic},
		t.sourceAdapter.Address(),
		pg.WithParentCtx(ctx),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get source MockERC20Sent logs: %w", err)
	}

	destFinalizeLogs, err := t.destLogPoller.LogsWithSigs(
		t.currDestBlock,
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

	t.currDestBlock = latestDestBlock.BlockNumber
	t.currSourceBlock = latestSourceBlock.BlockNumber

	return t.toPendingTransfers(ready, parsedToLP), nil
}

func (t *testBridge) toPendingTransfers(
	ready []*mock_l1_bridge_adapter.MockL1BridgeAdapterMockERC20Sent,
	parsedToLP map[logKey]logpoller.Log,
) []models.PendingTransfer {
	var transfers []models.PendingTransfer
	for _, send := range ready {
		lp := parsedToLP[logKey{txHash: send.Raw.TxHash, logIdx: int64(send.Raw.Index)}]
		bridgeData, err := packUint256(send.Amount)
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
				Amount:             send.Amount,
				LocalTokenAddress:  models.Address(send.LocalToken),
				RemoteTokenAddress: models.Address(send.RemoteToken),
				Date:               lp.BlockTimestamp,
				BridgeData:         bridgeData,
			},
			Status: models.TransferStatusReady,
		})
	}
	return transfers
}

func (t *testBridge) getReadyToFinalize(
	sends []*mock_l1_bridge_adapter.MockL1BridgeAdapterMockERC20Sent,
	finalizes []*mock_l1_bridge_adapter.MockL1BridgeAdapterMockERC20Finalized,
) ([]*mock_l1_bridge_adapter.MockL1BridgeAdapterMockERC20Sent, error) {
	// sort so that we can easily match the events to each other
	slices.SortFunc(sends, func(a, b *mock_l1_bridge_adapter.MockL1BridgeAdapterMockERC20Sent) int {
		return intComparator(a.Raw.BlockNumber, b.Raw.BlockNumber)
	})
	slices.SortFunc(finalizes, func(a, b *mock_l1_bridge_adapter.MockL1BridgeAdapterMockERC20Finalized) int {
		return intComparator(a.Raw.BlockNumber, b.Raw.BlockNumber)
	})

	// anything that has been mined but has not been already finalized is eligible
	var ready []*mock_l1_bridge_adapter.MockL1BridgeAdapterMockERC20Sent
	if len(sends) > len(finalizes) {
		t.lggr.Infow("New sends, marking ready to finalize", "sends", len(sends), "finalizes", len(finalizes))
		ready = sends[len(finalizes):]
	} else if len(sends) < len(finalizes) {
		// should be impossible
		t.lggr.Criticalw("more finalizes than sends, should be impossible", "sends", len(sends), "finalizes", len(finalizes))
		return nil, fmt.Errorf("more finalizes than sends")
	} else {
		// no new sends
		t.lggr.Debugw("No new sends", "sends", len(sends), "finalizes", len(finalizes))
		return nil, nil
	}

	return ready, nil
}

func (t *testBridge) parseFinalizedLogs(logs []logpoller.Log) ([]*mock_l1_bridge_adapter.MockL1BridgeAdapterMockERC20Finalized, error) {
	parsedFinalizeLogs := make([]*mock_l1_bridge_adapter.MockL1BridgeAdapterMockERC20Finalized, len(logs))
	for i, log := range logs {
		finalizeLog, err := t.destAdapter.ParseMockERC20Finalized(log.ToGethLog())
		if err != nil {
			return nil, fmt.Errorf("failed to parse finalize log: %w", err)
		}
		parsedFinalizeLogs[i] = finalizeLog
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
	parsedSendLogs := make([]*mock_l1_bridge_adapter.MockL1BridgeAdapterMockERC20Sent, len(logs))
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

// LocalChainSelector implements bridge.Bridge.
func (t *testBridge) LocalChainSelector() models.NetworkSelector {
	return t.sourceSelector
}

// RemoteChainSelector implements bridge.Bridge.
func (t *testBridge) RemoteChainSelector() models.NetworkSelector {
	return t.destSelector
}

func intComparator[T constraints.Integer](a, b T) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	} else {
		return 0
	}
}

func packUint256(val *big.Int) ([]byte, error) {
	return utils.ABIEncode(`[{"type": "uint256"}]`, val)
}
