package arb

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"go.uber.org/multierr"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/arb_node_interface"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/arbitrum_l1_bridge_adapter"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/arbitrum_l2_bridge_adapter"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/arbitrum_rollup_core"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/arbitrum_token_gateway"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/arbsys"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/l2_arbitrum_gateway"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/l2_arbitrum_messenger"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/bridge"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/models"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
)

var (
	// Event emitted by the L2 bridge adapter
	L2ToL1ERC20SentTopic = arbitrum_l2_bridge_adapter.ArbitrumL2BridgeAdapterArbitrumL2ToL1ERC20Sent{}.Topic()
	// Event emitted by the L1 bridge adapter
	L2toL1ERC20FinalizedTopic = arbitrum_l1_bridge_adapter.ArbitrumL1BridgeAdapterArbitrumL2ToL1ERC20Finalized{}.Topic()

	// Arbitrum events emitted on L1
	NodeConfirmedTopic = arbitrum_rollup_core.ArbRollupCoreNodeConfirmed{}.Topic()

	// Arbitrum events emitted on L2
	TxToL1Topic              = l2_arbitrum_messenger.L2ArbitrumMessengerTxToL1{}.Topic()
	WithdrawalInitiatedTopic = l2_arbitrum_gateway.L2ArbitrumGatewayWithdrawalInitiated{}.Topic()
	L2ToL1TxTopic            = arbsys.ArbSysL2ToL1Tx{}.Topic()

	// Important addresses on L2
	// These are precompiles so their addresses will never change
	NodeInterfaceAddress = common.HexToAddress("0x00000000000000000000000000000000000000c8")
	ArbSysAddress        = common.HexToAddress("0x0000000000000000000000000000000000000064")

	arbitrumTokenGatewayABI = abihelpers.MustParseABI(arbitrum_token_gateway.ArbitrumTokenGatewayMetaData.ABI)
	l1AdapterABI            = abihelpers.MustParseABI(arbitrum_l1_bridge_adapter.ArbitrumL1BridgeAdapterMetaData.ABI)

	// type assertion
	_ bridge.Bridge = &l2ToL1Bridge{}
)

func init() {
	// Check that finalizeInboundTransfer is on the token gateway ABI
	finalizeInboundTransferMethod, ok := arbitrumTokenGatewayABI.Methods["finalizeInboundTransfer"]
	if !ok {
		panic("finalizeInboundTransfer not found in ArbitrumTokenGateway ABI")
	}
	// Check that it has the expected signature
	if len(finalizeInboundTransferMethod.Inputs) != 5 {
		panic("finalizeInboundTransfer has unexpected number of inputs")
	}
	if finalizeInboundTransferMethod.Inputs[0].Type.String() != "address" ||
		finalizeInboundTransferMethod.Inputs[1].Type.String() != "address" ||
		finalizeInboundTransferMethod.Inputs[2].Type.String() != "address" ||
		finalizeInboundTransferMethod.Inputs[3].Type.String() != "uint256" ||
		finalizeInboundTransferMethod.Inputs[4].Type.String() != "bytes" {
		panic("finalizeInboundTransfer has unexpected input type")
	}
}

const (
	DurationMonth = 720 * time.Hour
)

type l2ToL1Bridge struct {
	localSelector       models.NetworkSelector
	remoteSelector      models.NetworkSelector
	localChainID        *big.Int
	remoteChainID       *big.Int
	l1RebalancerAddress common.Address
	l2BridgeAdapter     *arbitrum_l2_bridge_adapter.ArbitrumL2BridgeAdapter
	l1BridgeAdapter     *arbitrum_l1_bridge_adapter.ArbitrumL1BridgeAdapter
	l2LogPoller         logpoller.LogPoller
	l1LogPoller         logpoller.LogPoller
	l2FilterName        string
	l1FilterName        string
	lggr                logger.Logger
	l2Client            client.Client
	arbSys              *arbsys.ArbSys
	l2ArbGateway        *l2_arbitrum_gateway.L2ArbitrumGateway
	l2ArbMessenger      *l2_arbitrum_messenger.L2ArbitrumMessenger
	rollupCore          *arbitrum_rollup_core.ArbRollupCore
	nodeInterface       *arb_node_interface.NodeInterface
}

// GetBridgeSpecificPayload implements bridge.Bridge.
// Arbitrum L2 to L1 transfers require no bridge specific payload.
func (l *l2ToL1Bridge) GetBridgeSpecificPayload(ctx context.Context, transfer models.Transfer) ([]byte, error) {
	return []byte{}, nil
}

// Close implements bridge.Bridge.
func (l *l2ToL1Bridge) Close(ctx context.Context) error {
	// close log poller filters
	err := l.l2LogPoller.UnregisterFilter(l.l2FilterName)
	err2 := l.l1LogPoller.UnregisterFilter(l.l1FilterName)
	return multierr.Combine(err, err2)
}

// RemoteChainSelector implements bridge.Bridge.
func (l *l2ToL1Bridge) RemoteChainSelector() models.NetworkSelector {
	return l.remoteSelector
}

// GetTransfers implements bridge.Bridge.
func (l *l2ToL1Bridge) GetTransfers(ctx context.Context, l2Token models.Address, l1Token models.Address) ([]models.PendingTransfer, error) {
	// get all the L2 -> L1 transfers in the past 14 days for the given l2Token
	// that should be enough time to catch all the transfers
	// that were potentially not finalized.
	l2ToL1Transfers, err := l.l2LogPoller.IndexedLogsCreatedAfter(
		L2ToL1ERC20SentTopic,
		l.l2BridgeAdapter.Address(),
		1, // topic index
		[]common.Hash{
			common.HexToHash(l2Token.String()),
		},
		time.Now().Add(-DurationMonth/2),
		logpoller.Finalized,
		pg.WithParentCtx(ctx),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get L2 -> L1 transfers from log poller (on L2): %w", err)
	}

	// get all L2 -> L1 finalizations in the past 14 days
	// we can't filter on token since we don't have the token address in the onchain event
	l2ToL1Finalizations, err := l.l1LogPoller.LogsCreatedAfter(
		L2toL1ERC20FinalizedTopic,
		l.l1BridgeAdapter.Address(),
		time.Now().Add(-DurationMonth/2),
		logpoller.Finalized,
		pg.WithParentCtx(ctx),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get L2 -> L1 finalizations from log poller (on L1): %w", err)
	}

	parsedL2toL1Transfers, parsedToLP, err := l.parseL2ToL1Transfers(l2ToL1Transfers)
	if err != nil {
		return nil, fmt.Errorf("failed to parse L2 -> L1 transfers: %w", err)
	}

	parsedL2ToL1Finalizations, err := l.parseL2ToL1Finalizations(l2ToL1Finalizations)
	if err != nil {
		return nil, fmt.Errorf("failed to parse L2 -> L1 finalizations: %w", err)
	}

	// filter out the L2 -> L1 transfers that have been finalized onchain already
	// all the transfers in unfinalizedTransfers are either in the "not-ready" or "ready" state
	unfinalizedTransfers, err := l.filterOutFinalizedTransfers(parsedL2toL1Transfers, parsedL2ToL1Finalizations)
	if err != nil {
		return nil, fmt.Errorf("failed to filter finalized transfers: %w", err)
	}

	// for the remaining as-of-yet unfinalized transfers, determine if they
	// are ready to finalize
	ready, readyData, notReady, err := l.partitionReadyTransfers(ctx, unfinalizedTransfers)
	if err != nil {
		return nil, fmt.Errorf("failed to partition ready transfers: %w", err)
	}

	return l.toPendingTransfers(ready, readyData, notReady, parsedToLP)
}

func (l *l2ToL1Bridge) toPendingTransfers(
	ready []*arbitrum_l2_bridge_adapter.ArbitrumL2BridgeAdapterArbitrumL2ToL1ERC20Sent,
	readyData [][]byte,
	notReady []*arbitrum_l2_bridge_adapter.ArbitrumL2BridgeAdapterArbitrumL2ToL1ERC20Sent,
	parsedToLP map[logKey]logpoller.Log,
) ([]models.PendingTransfer, error) {
	if len(ready) != len(readyData) {
		return nil, fmt.Errorf("length of ready and readyData should be the same: len(ready) = %d, len(readyData) = %d",
			len(ready), len(readyData))
	}
	var transfers []models.PendingTransfer
	for i, transfer := range ready {
		transfers = append(transfers, models.PendingTransfer{
			Transfer: models.Transfer{
				From:   l.localSelector,
				To:     l.remoteSelector,
				Amount: transfer.Amount,
				Date: parsedToLP[logKey{
					txHash:   transfer.Raw.TxHash,
					logIndex: int64(transfer.Raw.Index),
				}].BlockTimestamp,
				BridgeData: readyData[i], // finalization data for withdrawals that are ready
			},
			Status: models.TransferStatusReady,
		})
	}
	for _, transfer := range notReady {
		transfers = append(transfers, models.PendingTransfer{
			Transfer: models.Transfer{
				From:   l.localSelector,
				To:     l.remoteSelector,
				Amount: transfer.Amount,
				Date: parsedToLP[logKey{
					txHash:   transfer.Raw.TxHash,
					logIndex: int64(transfer.Raw.Index),
				}].BlockTimestamp,
				BridgeData: []byte{}, // No data since its not ready
			},
			Status: models.TransferStatusNotReady,
		})
	}
	return transfers, nil
}

func (l *l2ToL1Bridge) partitionReadyTransfers(
	ctx context.Context,
	unfinalized []*arbitrum_l2_bridge_adapter.ArbitrumL2BridgeAdapterArbitrumL2ToL1ERC20Sent,
) (
	ready []*arbitrum_l2_bridge_adapter.ArbitrumL2BridgeAdapterArbitrumL2ToL1ERC20Sent,
	readyDatas [][]byte,
	notReady []*arbitrum_l2_bridge_adapter.ArbitrumL2BridgeAdapterArbitrumL2ToL1ERC20Sent,
	err error,
) {
	var errs error
	for _, transfer := range unfinalized {
		readyData, readyToFinalize, err := l.getFinalizationData(ctx, transfer)
		if err != nil {
			errs = multierr.Append(
				errs,
				fmt.Errorf("failed to get finalization data for transfer %s: %w", transfer.Raw.TxHash, err),
			)
			continue
		}
		if readyToFinalize {
			l.lggr.Infow("transfer is ready to finalize!",
				"transfer", transfer.Raw.TxHash,
				"readyData", hexutil.Encode(readyData),
			)
			ready = append(ready, transfer)
			readyDatas = append(readyDatas, readyData)
		} else {
			notReady = append(notReady, transfer)
		}
	}
	return
}

func (l *l2ToL1Bridge) getFinalizationData(
	ctx context.Context,
	transfer *arbitrum_l2_bridge_adapter.ArbitrumL2BridgeAdapterArbitrumL2ToL1ERC20Sent,
) (
	[]byte,
	bool,
	error,
) {
	// function executeTransaction(
	//
	//	  bytes32[] calldata proof,
	//	  uint256 index,
	//	  address l2Sender,
	//	  address to,
	//	  uint256 l2Block,
	//	  uint256 l1Block,
	//	  uint256 l2Timestamp,
	//	  uint256 value,
	//	  bytes calldata data
	//	) external;
	//
	// Arg 0: proof. This takes multiple steps:
	// 1. Get the latest NodeConfirmed event on L1, which indicates the latest node that was confirmed by the rollup.
	// 2. Call eth_getBlockByHash on L2 specifying the L2 block hash in the NodeConfirmed event.
	// 3. Get the `sendCount` field from the response.
	// 4. Get the `l2ToL1Id` field from the `WithdrawalInitiated` log from the L2 withdrawal tx.
	// 5. Call `constructOutboxProof` on the L2 node interface contract with the `sendCount` as the first argument and `l2ToL1Id` as the second argument.
	// Arg 1: index. Fetch the index from the TxToL1 log in the L2 tx.
	// Arg 2: l2Sender. Fetch the source of the WithdrawalInitiated log in the L2 tx.
	// Arg 3: to. Fetch the `to` field of the WithdrawalInitiated log in the L2 tx.
	// Arg 4: l1Block. Fetch the `l1BlockNumber` field of the JSON-RPC response to eth_getTransactionReceipt
	// passing in the L2 tx hash as the param.
	// Arg 5: l2Block. This is the l2 block number in which the withdrawal tx was included.
	// Arg 6: l2Timestamp. Get the `timestamp` field from the L2ToL1Tx event emitted by ArbSys (0x64).
	// Arg 7: value. Fetch the `value` field from the WithdrawalInitiated log in the L2 tx.
	// Arg 8: data. Fetch the `data` field from the TxToL1 log in the L2 tx.
	receipt, err := l.l2Client.TransactionReceipt(ctx, transfer.Raw.TxHash)
	if err != nil {
		// should be a transient error
		return nil, false, fmt.Errorf("failed to get transaction receipt: %w", err)
	}
	var (
		l2ToL1TxLog, withdrawalInitiatedLog, txToL1Log *gethtypes.Log
	)
	for _, lg := range receipt.Logs {
		if lg.Topics[0] == L2ToL1TxTopic {
			l2ToL1TxLog = lg
		} else if lg.Topics[0] == WithdrawalInitiatedTopic {
			withdrawalInitiatedLog = lg
		} else if lg.Topics[0] == TxToL1Topic {
			txToL1Log = lg
		}
	}
	if l2ToL1TxLog == nil || withdrawalInitiatedLog == nil || txToL1Log == nil {
		return nil, false, fmt.Errorf("missing one or more logs: l2ToL1TxLog: %+v, withdrawalInitiatedLog: %+v, txToL1Log: %+v",
			l2ToL1TxLog, withdrawalInitiatedLog, txToL1Log)
	}
	l2ToL1Tx, err := l.arbSys.ParseL2ToL1Tx(*l2ToL1TxLog)
	if err != nil {
		return nil, false, fmt.Errorf("failed to parse L2ToL1Tx log in tx %s: %w", receipt.TxHash, err)
	}
	withdrawalInitiated, err := l.l2ArbGateway.ParseWithdrawalInitiated(*withdrawalInitiatedLog)
	if err != nil {
		return nil, false, fmt.Errorf("failed to parse WithdrawalInitiated log in tx %s: %w", receipt.TxHash, err)
	}
	txToL1, err := l.l2ArbMessenger.ParseTxToL1(*txToL1Log)
	if err != nil {
		return nil, false, fmt.Errorf("failed to parse TxToL1 log in tx %s: %w", receipt.TxHash, err)
	}
	// argument 0: proof
	arg0Proof, err := l.getProof(ctx, withdrawalInitiated.L2ToL1Id)
	if err != nil {
		return nil, false, fmt.Errorf("failed to get proof: %w, l2tol1id: %s",
			err, withdrawalInitiated.L2ToL1Id)
	}
	if arg0Proof == nil {
		// if there's no proof, it means the transfer is not yet ready to finalize
		return nil, false, nil
	}
	// argument 1: index
	arg1Index := withdrawalInitiated.L2ToL1Id
	// argument 2: l2Sender
	arg2L2Sender := withdrawalInitiatedLog.Address
	// argument 3: to
	arg3To := txToL1.To
	// argument 4: l1Block
	arg4L1Block, err := l.getL1BlockFromRPC(ctx, receipt.TxHash)
	if err != nil {
		return nil, false, fmt.Errorf("failed to get l1 block for tx (%s) from rpc: %w",
			receipt.TxHash, err)
	}
	// argument 5: l2Block
	arg5L2Block := receipt.BlockNumber
	// argument 6: l2Timestamp
	arg6L2Timestamp := l2ToL1Tx.Timestamp
	// argument 7: value
	arg7Value := withdrawalInitiated.Amount
	// argument 8: data
	arg8Data := txToL1.Data

	finalizationPayload, err := l1AdapterABI.Pack("exposeArbitrumFinalizationPayload", arbitrum_l1_bridge_adapter.ArbitrumL1BridgeAdapterArbitrumFinalizationPayload{
		Proof:       arg0Proof,
		Index:       arg1Index,
		L2Sender:    arg2L2Sender,
		To:          arg3To,
		L1Block:     arg4L1Block,
		L2Block:     arg5L2Block,
		L2Timestamp: arg6L2Timestamp,
		Value:       arg7Value,
		Data:        arg8Data,
	})
	if err != nil {
		return nil, false, fmt.Errorf("failed to pack finalization payload: %w", err)
	}
	// trim the first four bytes (function signature)
	finalizationPayload = finalizationPayload[4:]
	return finalizationPayload, true, nil
}

func (l *l2ToL1Bridge) getL1BlockFromRPC(ctx context.Context, txHash common.Hash) (*big.Int, error) {
	type Response struct {
		L1BlockNumber hexutil.Big `json:"l1BlockNumber"`
	}
	response := new(Response)
	err := l.l2Client.CallContext(ctx, response, "eth_getTransactionReceipt", txHash.Hex())
	if err != nil {
		return nil, fmt.Errorf("failed to call eth_getTransactionReceipt with tx hash %s: %w", txHash, err)
	}
	return response.L1BlockNumber.ToInt(), nil
}

func (l *l2ToL1Bridge) getProof(ctx context.Context, l2ToL1Id *big.Int) ([][32]byte, error) {
	// 1. Get the latest NodeConfirmed event on L1, which indicates the latest node that was confirmed by the rollup.
	latestNodeConfirmed, err := l.getLatestNodeConfirmed(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest node confirmed: %w", err)
	}
	// 2. Call eth_getBlockByHash on L2 specifying the L2 block hash in the NodeConfirmed event.
	sendCount, err := l.getSendCountForBlock(ctx, latestNodeConfirmed.BlockHash)
	if err != nil {
		return nil, fmt.Errorf("failed to get send count for block: %w", err)
	}
	// 5. Call `constructOutboxProof` on the L2 node interface contract with the `sendCount` as the first argument and `l2ToL1Id` as the second argument.
	outboxProof, err := l.nodeInterface.ConstructOutboxProof(&bind.CallOpts{
		Context: ctx,
	}, sendCount, l2ToL1Id.Uint64())
	if err != nil {
		// if there's an error calling constructOutboxProof it means that the
		// transfer is not yet ready to finalize.
		l.lggr.Infow("failed to construct outbox proof, transfer not ready to finalize",
			"l2ToL1Id", l2ToL1Id,
			"sendCount", sendCount,
			"err", err)
		return nil, nil
	}
	return outboxProof.Proof, nil
}

func (l *l2ToL1Bridge) getSendCountForBlock(ctx context.Context, blockHash [32]byte) (uint64, error) {
	type Response struct {
		SendCount hexutil.Big `json:"sendCount"`
	}
	response := new(Response)
	bhHex := hexutil.Encode(blockHash[:])
	err := l.l2Client.CallContext(ctx, response, "eth_getBlockByHash", bhHex, false)
	if err != nil {
		return 0, fmt.Errorf("failed to call eth_getBlockByHash with blockhash %s: %w", bhHex, err)
	}
	return response.SendCount.ToInt().Uint64(), nil
}

func (l *l2ToL1Bridge) getLatestNodeConfirmed(ctx context.Context) (*arbitrum_rollup_core.ArbRollupCoreNodeConfirmed, error) {
	lg, err := l.l1LogPoller.LatestLogByEventSigWithConfs(
		NodeConfirmedTopic,
		l.rollupCore.Address(),
		logpoller.Finalized,
		pg.WithParentCtx(ctx),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest node confirmed: %w, topic: %s, address: %s", err, NodeConfirmedTopic, l.rollupCore.Address())
	}

	parsed, err := l.rollupCore.ParseNodeConfirmed(lg.ToGethLog())
	if err != nil {
		return nil, fmt.Errorf("failed to parse node confirmed log: %w", err)
	}

	return parsed, nil
}

func (l *l2ToL1Bridge) filterOutFinalizedTransfers(
	l2ToL1Transfers []*arbitrum_l2_bridge_adapter.ArbitrumL2BridgeAdapterArbitrumL2ToL1ERC20Sent,
	l2ToL1Finalizations []*arbitrum_l1_bridge_adapter.ArbitrumL1BridgeAdapterArbitrumL2ToL1ERC20Finalized,
) (
	[]*arbitrum_l2_bridge_adapter.ArbitrumL2BridgeAdapterArbitrumL2ToL1ERC20Sent,
	error,
) {
	var unfinalized []*arbitrum_l2_bridge_adapter.ArbitrumL2BridgeAdapterArbitrumL2ToL1ERC20Sent
	for _, l2ToL1Transfer := range l2ToL1Transfers {
		// We only care about transfers where the recipient is the l1 rebalancer contract
		if l2ToL1Transfer.Recipient != l.l1RebalancerAddress {
			l.lggr.Debugw("Ignoring L2 -> L1 transfer not destined for rebalancer", "transfer", l2ToL1Transfer.Raw.TxHash)
			continue
		}
		foundFinalized, err := l.findMatchingFinalization(l2ToL1Transfer, l2ToL1Finalizations)
		if err != nil {
			return nil, fmt.Errorf("unable to find matching finalization (withdrawal tx: %s): %w", l2ToL1Transfer.Raw.TxHash, err)
		}
		if foundFinalized {
			l.lggr.Debugw("Ignoring L2 -> L1 transfer that has been finalized", "transfer", l2ToL1Transfer.Raw.TxHash)
			continue
		}
		unfinalized = append(unfinalized, l2ToL1Transfer)
	}

	return unfinalized, nil
}

func (l *l2ToL1Bridge) findMatchingFinalization(
	withdrawal *arbitrum_l2_bridge_adapter.ArbitrumL2BridgeAdapterArbitrumL2ToL1ERC20Sent,
	finalizations []*arbitrum_l1_bridge_adapter.ArbitrumL1BridgeAdapterArbitrumL2ToL1ERC20Finalized,
) (foundFinalized bool, err error) {
	for _, l2ToL1Finalization := range finalizations {
		// the destination address is in the payload data, need to unpack it to figure out
		// who it went to.
		finalizeInboundTransferData, err := unpackFinalizeInboundTransfer(l2ToL1Finalization.Payload.Data)
		if err != nil {
			// should never happen
			return false, fmt.Errorf("failed to unpack finalizeInboundTransfer: %w, raw: %s",
				err, hexutil.Encode(l2ToL1Finalization.Payload.Data))
		}
		// We only care about finalizations destined for the rebalancer contract
		if finalizeInboundTransferData.l1Receiver != l.l1RebalancerAddress {
			l.lggr.Debugw("Ignoring L2 -> L1 finalization not destined for rebalancer", "finalization", l2ToL1Finalization)
			continue
		}

		// decode the bridge specific data in the l2 to l1 event and extract the l2 to l1 ID
		// this ID is a unique identifier - it is emitted on L2 and passed into executeTransaction
		// when finalizing on L1. We can use it to match the withdrawal to the finalization.
		l2ToL1Id, err := unpackUint256(withdrawal.OutboundTransferResult)
		if err != nil {
			// should never happen
			return false, fmt.Errorf("failed to unpack l2 to l1 id from l2 -> l1 transfer event: %w, raw: %s",
				err, hexutil.Encode(withdrawal.OutboundTransferResult))
		}

		if l2ToL1Id.Cmp(l2ToL1Finalization.Payload.Index) != 0 {
			// This finalization is not for this transfer
			continue
		}

		// finalization is for this transfer
		foundFinalized = true
		break
	}
	return foundFinalized, nil
}

func (l *l2ToL1Bridge) parseL2ToL1Finalizations(logs []logpoller.Log) ([]*arbitrum_l1_bridge_adapter.ArbitrumL1BridgeAdapterArbitrumL2ToL1ERC20Finalized, error) {
	finalizations := make([]*arbitrum_l1_bridge_adapter.ArbitrumL1BridgeAdapterArbitrumL2ToL1ERC20Finalized, len(logs))
	for i, log := range logs {
		parsed, err := l.l1BridgeAdapter.ParseArbitrumL2ToL1ERC20Finalized(log.ToGethLog())
		if err != nil {
			// should never happen
			return nil, fmt.Errorf("failed to parse L2 -> L1 finalization log: %w", err)
		}
		finalizations[i] = parsed
	}
	return finalizations, nil
}

type logKey struct {
	txHash   common.Hash
	logIndex int64
}

func (l *l2ToL1Bridge) parseL2ToL1Transfers(
	logs []logpoller.Log,
) (
	[]*arbitrum_l2_bridge_adapter.ArbitrumL2BridgeAdapterArbitrumL2ToL1ERC20Sent,
	map[logKey]logpoller.Log,
	error,
) {
	transfers := make([]*arbitrum_l2_bridge_adapter.ArbitrumL2BridgeAdapterArbitrumL2ToL1ERC20Sent, len(logs))
	parsedToLPLog := make(map[logKey]logpoller.Log)
	for i, log := range logs {
		parsed, err := l.l2BridgeAdapter.ParseArbitrumL2ToL1ERC20Sent(log.ToGethLog())
		if err != nil {
			// should never happen
			return nil, nil, fmt.Errorf("failed to parse L2 -> L1 transfer log: %w", err)
		}
		transfers[i] = parsed
		parsedToLPLog[logKey{
			txHash:   log.TxHash,
			logIndex: log.LogIndex,
		}] = log
	}
	return transfers, parsedToLPLog, nil
}

// LocalChainSelector implements bridge.Bridge.
func (l *l2ToL1Bridge) LocalChainSelector() models.NetworkSelector {
	return l.localSelector
}

func NewL2ToL1Bridge(
	lggr logger.Logger,
	localSelector,
	remoteSelector models.NetworkSelector,
	l1RollupAddress,
	l1RebalancerAddress,
	l2BridgeAdapterAddress,
	l1BridgeAdapterAddress common.Address,
	l2LogPoller,
	l1LogPoller logpoller.LogPoller,
	l2Client,
	l1Client client.Client,
) (*l2ToL1Bridge, error) {
	localChain, ok := chainsel.ChainBySelector(uint64(localSelector))
	if !ok {
		return nil, fmt.Errorf("unknown chain selector for local chain: %d", localSelector)
	}
	remoteChain, ok := chainsel.ChainBySelector(uint64(remoteSelector))
	if !ok {
		return nil, fmt.Errorf("unknown chain selector for remote chain: %d", remoteSelector)
	}
	l2FilterName := fmt.Sprintf("ArbitrumL2ToL1Bridge-L2-%s-%s", localChain.Name, remoteChain.Name)
	err := l2LogPoller.RegisterFilter(logpoller.Filter{
		Name: l2FilterName,
		EventSigs: []common.Hash{
			L2ToL1ERC20SentTopic,
		},
		Addresses: []common.Address{l2BridgeAdapterAddress},
		Retention: DurationMonth,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to register filter for Arbitrum L2 to L1 bridge: %w", err)
	}

	l1FilterName := fmt.Sprintf("ArbitrumL2ToL1Bridge-L1-%s-%s", remoteChain.Name, localChain.Name)
	err = l1LogPoller.RegisterFilter(logpoller.Filter{
		Name: l1FilterName,
		EventSigs: []common.Hash{
			L2toL1ERC20FinalizedTopic, // emitted by l1 bridge adapter
			NodeConfirmedTopic,        // emitted by rollup
		},
		Addresses: []common.Address{
			l1BridgeAdapterAddress, // to get erc20 finalized logs
			l1RollupAddress,        // to get node confirmed logs
		},
		Retention: DurationMonth,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to register filter for Arbitrum L1 to L2 bridge: %w", err)
	}

	l2BridgeAdapter, err := arbitrum_l2_bridge_adapter.NewArbitrumL2BridgeAdapter(l2BridgeAdapterAddress, l2Client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate Arbitrum L2 bridge adapter: %w", err)
	}

	l1BridgeAdapter, err := arbitrum_l1_bridge_adapter.NewArbitrumL1BridgeAdapter(l1BridgeAdapterAddress, l1Client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate Arbitrum L1 bridge adapter: %w", err)
	}

	arbSys, err := arbsys.NewArbSys(ArbSysAddress, l2Client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate ArbSys contract: %w", err)
	}

	// Addresses provided to the below wrappers don't matter,
	// we're just using them to parse the needed logs easily.
	l2ArbGateway, err := l2_arbitrum_gateway.NewL2ArbitrumGateway(
		common.HexToAddress("0x0"),
		l2Client,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate L2ArbitrumGateway contract: %w", err)
	}

	l2ArbMessenger, err := l2_arbitrum_messenger.NewL2ArbitrumMessenger(
		common.HexToAddress("0x0"),
		l2Client,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate L2ArbitrumMessenger contract: %w", err)
	}

	// have to use the correct address here
	rollupCore, err := arbitrum_rollup_core.NewArbRollupCore(l1RollupAddress, l1Client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate ArbRollupCore contract: %w", err)
	}

	// and here
	nodeInterface, err := arb_node_interface.NewNodeInterface(NodeInterfaceAddress, l2Client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate NodeInterface contract: %w", err)
	}

	lggr = lggr.Named("ArbitrumL2ToL1Bridge").With(
		"localSelector", localSelector,
		"remoteSelector", remoteSelector,
		"localChainID", localChain.EvmChainID,
		"remoteChainID", remoteChain.EvmChainID,
		"l1BridgeAdapter", l1BridgeAdapterAddress,
		"l2BridgeAdapter", l2BridgeAdapterAddress,
		"l1RebalancerAddress", l1RebalancerAddress,
	)

	// TODO: replay log poller for any missed logs?
	return &l2ToL1Bridge{
		localSelector:       localSelector,
		remoteSelector:      remoteSelector,
		localChainID:        big.NewInt(int64(localChain.EvmChainID)),
		remoteChainID:       big.NewInt(int64(remoteChain.EvmChainID)),
		l2BridgeAdapter:     l2BridgeAdapter,
		l1BridgeAdapter:     l1BridgeAdapter,
		l2LogPoller:         l2LogPoller,
		l1LogPoller:         l1LogPoller,
		l2FilterName:        l2FilterName,
		l1FilterName:        l1FilterName,
		l1RebalancerAddress: l1RebalancerAddress,
		lggr:                lggr,
		l2Client:            l2Client,
		arbSys:              arbSys,
		l2ArbGateway:        l2ArbGateway,
		l2ArbMessenger:      l2ArbMessenger,
		rollupCore:          rollupCore,
		nodeInterface:       nodeInterface,
	}, nil
}

type finalizeInboundTransferParams struct {
	l1Token    common.Address
	l2Sender   common.Address
	l1Receiver common.Address
	amount     *big.Int
	data       []byte
}

func unpackFinalizeInboundTransfer(calldata []byte) (finalizeInboundTransferParams, error) {
	method, ok := arbitrumTokenGatewayABI.Methods["finalizeInboundTransfer"]
	if !ok {
		return finalizeInboundTransferParams{}, fmt.Errorf("finalizeInboundTransfer not found in ArbitrumTokenGateway ABI")
	}
	// trim first 4 bytes (function signature)
	ifaces, err := method.Inputs.Unpack(calldata[4:])
	if err != nil {
		return finalizeInboundTransferParams{}, fmt.Errorf("failed to unpack finalizeInboundTransfer: %w", err)
	}

	if len(ifaces) != 5 {
		return finalizeInboundTransferParams{}, fmt.Errorf("expected 5 arguments, got %d", len(ifaces))
	}

	var params finalizeInboundTransferParams
	params.l1Token = *abi.ConvertType(ifaces[0], new(common.Address)).(*common.Address)
	params.l2Sender = *abi.ConvertType(ifaces[1], new(common.Address)).(*common.Address)
	params.l1Receiver = *abi.ConvertType(ifaces[2], new(common.Address)).(*common.Address)
	params.amount = *abi.ConvertType(ifaces[3], new(*big.Int)).(**big.Int)
	params.data = *abi.ConvertType(ifaces[4], new([]byte)).(*[]byte)

	return params, nil
}

func unpackUint256(data []byte) (*big.Int, error) {
	iface, err := utils.ABIDecode(`[{"type": "uint256"}]`, data)
	if err != nil {
		return nil, fmt.Errorf("failed to decode uint256: %w", err)
	}
	ret := *abi.ConvertType(iface, new(*big.Int)).(**big.Int)
	return ret, nil
}
