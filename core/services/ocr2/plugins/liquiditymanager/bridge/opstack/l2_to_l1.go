package opstack

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"go.uber.org/multierr"

	evmtypes "github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"
	ubig "github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils/big"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/liquiditymanager/generated/optimism_l1_bridge_adapter_encoder"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/liquiditymanager/abiutils"
	bridgecommon "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/liquiditymanager/bridge/common"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/liquiditymanager/bridge/opstack/withdrawprover"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/liquiditymanager/generated/liquiditymanager"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/liquiditymanager/models"
)

type l2ToL1Bridge struct {
	localSelector      models.NetworkSelector
	remoteSelector     models.NetworkSelector
	l1LiquidityManager liquiditymanager.LiquidityManagerInterface
	l2LiquidityManager liquiditymanager.LiquidityManagerInterface
	l1Client           client.Client
	l2Client           client.Client
	l1LogPoller        logpoller.LogPoller
	l2LogPoller        logpoller.LogPoller
	l1FilterName       string
	l2FilterName       string
	l1Token, l2Token   common.Address
	lggr               logger.Logger
}

func NewL2ToL1Bridge(
	lggr logger.Logger,
	localSelector,
	remoteSelector models.NetworkSelector,
	l1LiquidityManagerAddress,
	l2LiquidityManagerAddress common.Address,
	l1Client,
	l2Client client.Client,
	l1LogPoller,
	l2LogPoller logpoller.LogPoller,
) (*l2ToL1Bridge, error) {
	localChain, ok := chainsel.ChainBySelector(uint64(localSelector))
	if !ok {
		return nil, fmt.Errorf("unknown chain selector for local chain: %d", localSelector)
	}
	remoteChain, ok := chainsel.ChainBySelector(uint64(remoteSelector))
	if !ok {
		return nil, fmt.Errorf("unknown chain selector for remote chain: %d", remoteSelector)
	}
	l2FilterName := fmt.Sprintf("OptimismL2ToL1Bridge-L2-LiquidityManager:%s-Local:%s-Remote:%s",
		l2LiquidityManagerAddress.Hex(), localChain.Name, remoteChain.Name)
	// TODO (ogtownsend): pass context from above
	ctx := context.Background()
	err := l2LogPoller.RegisterFilter(
		ctx,
		logpoller.Filter{
			Name: l2FilterName,
			EventSigs: []common.Hash{
				bridgecommon.LiquidityTransferredTopic,
			},
			Addresses: []common.Address{l2LiquidityManagerAddress},
			Retention: bridgecommon.DurationMonth,
		})
	if err != nil {
		return nil, fmt.Errorf("register L2 LM filter for Optimism L2 to L1 bridge: %w", err)
	}

	l1FilterName := fmt.Sprintf("OptimismL2ToL1Bridge-L1-LiquidityManager:%s-Local:%s-Remote:%s",
		l1LiquidityManagerAddress.Hex(), localChain.Name, remoteChain.Name)
	err = l1LogPoller.RegisterFilter(
		ctx,
		logpoller.Filter{
			Name: l1FilterName,
			EventSigs: []common.Hash{
				bridgecommon.FinalizationStepCompletedTopic, // emitted by LiquidityManager
				bridgecommon.LiquidityTransferredTopic,      // emitted by LiquidityManager
			},
			Addresses: []common.Address{
				l1LiquidityManagerAddress, // to get LiquidityTransferred and FinalizationStepCompleted logs
			},
			Retention: bridgecommon.DurationMonth,
		})
	if err != nil {
		return nil, fmt.Errorf("register L1 LM filter for Optimism L2 to L1 bridge: %w", err)
	}

	l1LiquidityManager, err := liquiditymanager.NewLiquidityManager(l1LiquidityManagerAddress, l1Client)
	if err != nil {
		return nil, fmt.Errorf("instantiate L1 LiquidityManager: %w", err)
	}

	l2LiquidityManager, err := liquiditymanager.NewLiquidityManager(l2LiquidityManagerAddress, l2Client)
	if err != nil {
		return nil, fmt.Errorf("instantiate L2 LiquidityManager: %w", err)
	}

	l2Token, err := l2LiquidityManager.ILocalToken(nil)
	if err != nil {
		return nil, fmt.Errorf("get L2 local token address: %w", err)
	}
	l1Token, err := l1LiquidityManager.ILocalToken(nil)
	if err != nil {
		return nil, fmt.Errorf("get L1 local token address: %w", err)
	}

	lggr = lggr.Named("OptimismL2ToL1Bridge").With(
		"localSelector", localSelector,
		"remoteSelector", remoteSelector,
		"localChainID", localChain.EvmChainID,
		"remoteChainID", remoteChain.EvmChainID,
		"localChainName", localChain.Name,
		"remoteChainName", remoteChain.Name,
		"l1LiquidityManager", l1LiquidityManagerAddress.Hex(),
		"l2LiquidityManager", l2LiquidityManagerAddress.Hex(),
		"l1Token", l1Token.Hex(),
		"l2Token", l2Token.Hex(),
	)
	lggr.Infow("Initialized Optimism L2 to L1 bridge")

	return &l2ToL1Bridge{
		localSelector:      localSelector,
		remoteSelector:     remoteSelector,
		l1LiquidityManager: l1LiquidityManager,
		l2LiquidityManager: l2LiquidityManager,
		l1Client:           l1Client,
		l2Client:           l2Client,
		l1LogPoller:        l1LogPoller,
		l2LogPoller:        l2LogPoller,
		l1FilterName:       l1FilterName,
		l2FilterName:       l2FilterName,
		l1Token:            l1Token,
		l2Token:            l2Token,
		lggr:               lggr,
	}, nil
}

func (l *l2ToL1Bridge) GetTransfers(
	ctx context.Context,
	localToken,
	remoteToken models.Address,
) ([]models.PendingTransfer, error) {
	lggr := l.lggr.With("l2Token", localToken, "l1Token", remoteToken)
	if l.l2Token.Cmp(common.Address(localToken)) != 0 {
		return nil, fmt.Errorf("local token mismatch: expected %s, got %s", l.l2Token, localToken)
	}
	if l.l1Token.Cmp(common.Address(remoteToken)) != 0 {
		return nil, fmt.Errorf("remote token mismatch: expected %s, got %s", l.l1Token, remoteToken)
	}

	sendLogs, proveFinalizationStepLogs, receivedLogs, err := l.getLogs(ctx)
	if err != nil {
		return nil, fmt.Errorf("get logs: %w", err)
	}

	lggr.Infow("Got L2 -> L1 transfer and finalization step logs",
		"sendLogs", len(sendLogs),
		"proveFinalizedLogs", len(proveFinalizationStepLogs),
		"receivedLogs", len(receivedLogs),
	)

	parsedSent, parsedToLp, err := bridgecommon.ParseLiquidityTransferred(l.l1LiquidityManager.ParseLiquidityTransferred, sendLogs)
	if err != nil {
		return nil, fmt.Errorf("parse L2 -> L1 transfer sent logs: %w", err)
	}

	parsedProveFinalizationSteps, err := bridgecommon.ParseFinalizationStepCompleted(l.l1LiquidityManager.ParseFinalizationStepCompleted, proveFinalizationStepLogs)
	if err != nil {
		return nil, fmt.Errorf("parse L2 -> L1 transfer prove finalization step logs: %w", err)
	}

	parsedReceived, _, err := bridgecommon.ParseLiquidityTransferred(l.l1LiquidityManager.ParseLiquidityTransferred, receivedLogs)
	if err != nil {
		return nil, fmt.Errorf("parse L2 -> L1 transfer received logs: %w", err)
	}

	lggr.Infow("parsed logs",
		"parsedSent", len(parsedSent),
		"parsedProveFinalizationSteps", len(parsedProveFinalizationSteps),
		"parsedReceived", len(parsedReceived),
	)

	needsToBeProven, needsToBeFinalized, missingSent, err := partitionWithdrawalTransfers(
		l.localSelector,
		l.l1LiquidityManager.Address(),
		parsedSent,
		parsedProveFinalizationSteps,
		parsedReceived,
	)
	if err != nil {
		return nil, fmt.Errorf("partition transfers: %w", err)
	}
	if len(missingSent) > 0 {
		l.lggr.Errorw("missing sent logs", "missingSent", missingSent)
	}

	return l.toPendingTransfers(ctx, localToken, remoteToken, needsToBeProven, needsToBeFinalized, parsedToLp)
}

/**
 * partitionWithdrawalTransfers matches and divides in-progress and completed transfers into three groups:
 * 1) needsToBeProven: transfers that have been started by the L2 LM but are not yet proven on L1
 * 2) needsToBeFinalized: transfers that have been proven on L1 but are not yet finalized (received) on L1
 * 3) missingSent: transfers that have a prove finalization step log but no matching sent log
 *
 * It does this by matching the transfer's unique nonce emitted in certain events' fields. These events and fields are:
 * - L2 LiquidityTransferred.bridgeReturnData: emitted by the L2 LM when a transfer is initiated
 * - L1 FinalizationStepCompleted.bridgeSpecificData: emitted by the L1 LM when a L2 to L1 withdrawal is proven
 * - L1 LiquidityTransferred.bridgeSpecificData: emitted by the L1 LM when a L2 to L1 withdrawal is finalized
 */
func partitionWithdrawalTransfers(
	localSelector models.NetworkSelector,
	l1LiquidityManagerAddress common.Address,
	sentLogs []*liquiditymanager.LiquidityManagerLiquidityTransferred,
	proveFinalizationStepLogs []*liquiditymanager.LiquidityManagerFinalizationStepCompleted,
	receivedLogs []*liquiditymanager.LiquidityManagerLiquidityTransferred,
) (
	needsToBeProven,
	needsToBeFinalized []*liquiditymanager.LiquidityManagerLiquidityTransferred,
	missingSent []*liquiditymanager.LiquidityManagerFinalizationStepCompleted,
	err error,
) {
	transferNonceToSentLogMap := make(map[string]*liquiditymanager.LiquidityManagerLiquidityTransferred)
	foundMatchingProveFinalizationStepMap := make(map[string]bool)
	for _, sentLog := range sentLogs {
		if sentLog.To != l1LiquidityManagerAddress {
			continue
		}
		if sentLog.FromChainSelector != uint64(localSelector) {
			continue
		}
		var transferNonce *big.Int
		transferNonce, err = abiutils.UnpackUint256(sentLog.BridgeReturnData)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("unpack transfer nonce from L2 LiquidityTransferred log. Log tx: %s. Err: %w, log bridgeReturnData: %s",
				sentLog.Raw.TxHash, err, hexutil.Encode(sentLog.BridgeReturnData))
		}
		transferNonceToSentLogMap[transferNonce.String()] = sentLog
		foundMatchingProveFinalizationStepMap[transferNonce.String()] = false
	}

	// For each proveFinalizationStep, check if it matches a sentLogs log
	for _, proveStep := range proveFinalizationStepLogs {
		// L1's prove finalization step log's remote chain selector should be L2
		if proveStep.RemoteChainSelector != uint64(localSelector) {
			continue
		}
		var transferNonce *big.Int
		transferNonce, err = getTransferNonceFromFinalizationStep(proveStep)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("get transfer nonce from L1 FinalizationStepCompleted log. Log tx: %s. Err: %w",
				proveStep.Raw.TxHash, err)
		}
		if sentLog, exists := transferNonceToSentLogMap[transferNonce.String()]; exists {
			// If a corresponding sentLog exists for this proveFinalizationStep, append to needsToBeFinalized and
			// mark it as found
			needsToBeFinalized = append(needsToBeFinalized, sentLog)
			foundMatchingProveFinalizationStepMap[transferNonce.String()] = true
		} else {
			// If no corresponding sentLog exists for this proveFinalizationStep, append to missingSent
			missingSent = append(missingSent, proveStep)
		}
	}

	// Any entries in foundMatchingProveFinalizationStepMap that are still false are transfers that need to be proven
	// TODO (ogtownsend / amirylm): is the plugin able to handle the case where we've already instructed the plugin to prove() a transfer, but
	//   the prove log hasn't been emitted or ingested by the log poller yet? We could potentially send two prove() txs
	for transferNonce, found := range foundMatchingProveFinalizationStepMap {
		if !found {
			if sentLog, exists := transferNonceToSentLogMap[transferNonce]; exists {
				needsToBeProven = append(needsToBeProven, sentLog)
			}
		}
	}

	// Filter out from needsToBeFinalized any entries that have already been receivedLogs by the L1 LM
	needsToBeFinalized, err = filterExecuted(needsToBeFinalized, receivedLogs)
	return
}

func getTransferNonceFromFinalizationStep(proveStep *liquiditymanager.LiquidityManagerFinalizationStepCompleted) (*big.Int, error) {
	encodedPayload := proveStep.BridgeSpecificData

	// Unpack outermost finalize withdraw erc20 payload
	unpackedFinalizeWithdrawERC20Payload, err := l1OPBridgeAdapterEncoderABI.Methods["encodeFinalizeWithdrawalERC20Payload"].Inputs.Unpack(encodedPayload)
	if err != nil {
		return nil, fmt.Errorf("unpack finalizeWithdrawalERC20Payload: %w", err)
	}
	outFinalizeWithdrawERC20Payload := *abi.ConvertType(unpackedFinalizeWithdrawERC20Payload[0], new(optimism_l1_bridge_adapter_encoder.OptimismL1BridgeAdapterFinalizeWithdrawERC20Payload)).(*optimism_l1_bridge_adapter_encoder.OptimismL1BridgeAdapterFinalizeWithdrawERC20Payload)

	// Unpack optimism prove withdrawal payload
	unpackedOptimismProveWithdrawalPayload, err := l1OPBridgeAdapterEncoderABI.Methods["encodeOptimismProveWithdrawalPayload"].Inputs.Unpack(outFinalizeWithdrawERC20Payload.Data)
	if err != nil {
		return nil, fmt.Errorf("unpack optimismProveWithdrawalPayload: %w", err)
	}
	outOptimismProveWithdrawalPayload := *abi.ConvertType(unpackedOptimismProveWithdrawalPayload[0], new(optimism_l1_bridge_adapter_encoder.OptimismL1BridgeAdapterOptimismProveWithdrawalPayload)).(*optimism_l1_bridge_adapter_encoder.OptimismL1BridgeAdapterOptimismProveWithdrawalPayload)

	// Unpack withdrawal transaction's data from relayMessage data. Trim the first 4 bytes since this was encoded with
	// the function selector: https://github.com/ethereum-optimism/optimism/blob/f707883038d527cbf1e9f8ea513fe33255deadbc/packages/contracts-bedrock/src/universal/CrossDomainMessenger.sol#L186
	decodedRelayMessage, err := opCrossDomainMessengerABI.Methods["relayMessage"].Inputs.Unpack(outOptimismProveWithdrawalPayload.WithdrawalTransaction.Data[4:])
	if err != nil {
		return nil, fmt.Errorf("unpack relayMessage data: %w", err)
	}

	// Unpack relay message Message field into StandardBridge's finalizeBridgeERC20 params. Trim the first 4 bytes since
	// this was encoded with the function selector. The nonce is the 6th parameter.
	unpackedFinalizeBridgeParams, err := opStandardBridgeABI.Methods["finalizeBridgeERC20"].Inputs.Unpack(decodedRelayMessage[5].([]byte)[4:])
	if err != nil {
		return nil, fmt.Errorf("unpack finalizeBridgeERC20 params: %w", err)
	}

	fmt.Println("Finalize bridge params decoded bridgeReturnData nonce:", hexutil.Encode(unpackedFinalizeBridgeParams[5].([]byte)))

	return abiutils.UnpackUint256(unpackedFinalizeBridgeParams[5].([]byte))
}

func (l *l2ToL1Bridge) getLogs(ctx context.Context) (sendLogs, proveFinalizationStepLogs, receivedLogs []logpoller.Log, err error) {
	// Get all L2 -> L1 transfers that have been sent from the L2 LM in the past 14 days
	sendLogs, err = l.l2LogPoller.IndexedLogsCreatedAfter(
		ctx,
		bridgecommon.LiquidityTransferredTopic,
		l.l2LiquidityManager.Address(),
		bridgecommon.LiquidityTransferredToChainSelectorTopicIndex,
		[]common.Hash{
			bridgecommon.NetworkSelectorToHash(l.remoteSelector),
		},
		time.Now().Add(-bridgecommon.DurationMonth/2),
		evmtypes.Finalized,
	)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get L2 -> L1 transfers from log poller on L2: %w", err)
	}

	// Get all L2 -> L1 transfers that have been proven/finalized in the past 14 days
	proveFinalizationStepLogs, err = l.l1LogPoller.IndexedLogsCreatedAfter(
		ctx,
		bridgecommon.FinalizationStepCompletedTopic,
		l.l1LiquidityManager.Address(),
		bridgecommon.FinalizationStepCompletedRemoteChainSelectorTopicIndex,
		[]common.Hash{
			bridgecommon.NetworkSelectorToHash(l.remoteSelector),
		},
		time.Now().Add(-bridgecommon.DurationMonth/2),
		evmtypes.Finalized, // TODO: confirm if we actually need this to be finalized
	)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get L1 -> L2 transfers from log poller on L1: %w", err)
	}

	receivedLogs, err = l.l1LogPoller.IndexedLogsCreatedAfter(
		ctx,
		bridgecommon.LiquidityTransferredTopic,
		l.l1LiquidityManager.Address(),
		bridgecommon.LiquidityTransferredFromChainSelectorTopicIndex,
		[]common.Hash{
			bridgecommon.NetworkSelectorToHash(l.localSelector),
		},
		time.Now().Add(-bridgecommon.DurationMonth/2),
		evmtypes.Finalized,
	)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get L1 -> L2 transfers from log poller on L1: %w", err)
	}

	return sendLogs, proveFinalizationStepLogs, receivedLogs, nil
}

func (l *l2ToL1Bridge) toPendingTransfers(
	ctx context.Context,
	localToken, remoteToken models.Address,
	needsToBeProven, needsToBeFinalized []*liquiditymanager.LiquidityManagerLiquidityTransferred,
	parsedToLP map[bridgecommon.LogKey]logpoller.Log,
) ([]models.PendingTransfer, error) {
	var transfers []models.PendingTransfer
	for _, transfer := range needsToBeProven {
		provePayload, err := l.generateTransferBridgeDataForProve(ctx, transfer)
		if err != nil {
			return nil, fmt.Errorf("generate transfer bridge data for prove: %w", err)
		}
		transfers = append(transfers, models.PendingTransfer{
			Transfer: models.Transfer{
				From:               l.localSelector,
				To:                 l.remoteSelector,
				Sender:             models.Address(l.l2LiquidityManager.Address()),
				Receiver:           models.Address(l.l1LiquidityManager.Address()),
				LocalTokenAddress:  localToken,
				RemoteTokenAddress: remoteToken,
				Amount:             ubig.New(transfer.Amount),
				Date: parsedToLP[bridgecommon.LogKey{
					TxHash:   transfer.Raw.TxHash,
					LogIndex: int64(transfer.Raw.Index),
				}].BlockTimestamp,
				BridgeData: provePayload,
				Stage:      bridgecommon.StageRebalanceConfirmed,
			},
			Status: models.TransferStatusNotReady, // Needs to be proved before it can be finalized
			ID:     fmt.Sprintf("%s-%d", transfer.Raw.TxHash.Hex(), transfer.Raw.Index),
		})
	}
	for _, transfer := range needsToBeFinalized {
		finalizePayload, err := l.generateTransferBridgeDataForFinalize(ctx, transfer)
		if err != nil {
			return nil, fmt.Errorf("generate transfer bridge data for finalize: %w", err)
		}
		transfers = append(transfers, models.PendingTransfer{
			Transfer: models.Transfer{
				From:               l.localSelector,
				To:                 l.remoteSelector,
				Sender:             models.Address(l.l2LiquidityManager.Address()),
				Receiver:           models.Address(l.l1LiquidityManager.Address()),
				LocalTokenAddress:  localToken,
				RemoteTokenAddress: remoteToken,
				Amount:             ubig.New(transfer.Amount),
				Date: parsedToLP[bridgecommon.LogKey{
					TxHash:   transfer.Raw.TxHash,
					LogIndex: int64(transfer.Raw.Index),
				}].BlockTimestamp,
				BridgeData: finalizePayload,
				Stage:      bridgecommon.StageFinalizeReady,
			},
			Status: models.TransferStatusReady, // Ready to be finalized
			ID:     fmt.Sprintf("%s-%d", transfer.Raw.TxHash.Hex(), transfer.Raw.Index),
		})
	}
	return transfers, nil
}

func (l *l2ToL1Bridge) generateTransferBridgeDataForProve(
	ctx context.Context,
	transfer *liquiditymanager.LiquidityManagerLiquidityTransferred,
) ([]byte, error) {
	prover, err := withdrawprover.New(
		l.l1Client,
		l.l2Client,
		OptimismContracts[uint64(l.localSelector)]["L2OutputOracle"],
		OptimismContracts[uint64(l.localSelector)]["OptimismPortal"],
	)
	if err != nil {
		return nil, fmt.Errorf("instantiate withdraw prover: %w", err)
	}

	messageProof, err := prover.Prove(ctx, transfer.Raw.TxHash)
	if err != nil {
		return nil, fmt.Errorf("prove message: %w", err)
	}
	fmt.Println("Calling proveWithdrawalTransaction on bridge adapter, nonce:", messageProof.LowLevelMessage.Nonce, "\n",
		"sender:", messageProof.LowLevelMessage.Sender.String(), "\n",
		"target:", messageProof.LowLevelMessage.Target.String(), "\n",
		"value:", messageProof.LowLevelMessage.Value.String(), "\n",
		"gasLimit:", messageProof.LowLevelMessage.GasLimit.String(), "\n",
		"data:", hexutil.Encode(messageProof.LowLevelMessage.Data), "\n",
		"l2OutputIndex:", messageProof.L2OutputIndex, "\n",
		"outputRootProof version:", hexutil.Encode(messageProof.OutputRootProof.Version[:]), "\n",
		"outputRootProof stateRoot:", hexutil.Encode(messageProof.OutputRootProof.StateRoot[:]), "\n",
		"outputRootProof messagePasserStorageRoot:", hexutil.Encode(messageProof.OutputRootProof.MessagePasserStorageRoot[:]), "\n",
		"outputRootProof latestBlockHash:", hexutil.Encode(messageProof.OutputRootProof.LatestBlockHash[:]), "\n",
		"withdrawalProof:", formatWithdrawalProof(messageProof.WithdrawalProof))

	encodedProveWithdrawal, err := l1OPBridgeAdapterEncoderABI.Methods["encodeOptimismProveWithdrawalPayload"].Inputs.Pack(
		optimism_l1_bridge_adapter_encoder.OptimismL1BridgeAdapterOptimismProveWithdrawalPayload{
			WithdrawalTransaction: optimism_l1_bridge_adapter_encoder.TypesWithdrawalTransaction{
				Nonce:    messageProof.LowLevelMessage.Nonce,
				Sender:   messageProof.LowLevelMessage.Sender,
				Target:   messageProof.LowLevelMessage.Target,
				Value:    messageProof.LowLevelMessage.Value,
				GasLimit: messageProof.LowLevelMessage.GasLimit,
				Data:     messageProof.LowLevelMessage.Data,
			},
			L2OutputIndex: messageProof.L2OutputIndex,
			OutputRootProof: optimism_l1_bridge_adapter_encoder.TypesOutputRootProof{
				Version:                  messageProof.OutputRootProof.Version,
				StateRoot:                messageProof.OutputRootProof.StateRoot,
				MessagePasserStorageRoot: messageProof.OutputRootProof.MessagePasserStorageRoot,
				LatestBlockhash:          messageProof.OutputRootProof.LatestBlockHash,
			},
			WithdrawalProof: messageProof.WithdrawalProof,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("encodeOptimismProveWithdrawalPayload: %w", err)
	}

	// Then encode the finalize withdraw ERC 20 payload
	encodedPayload, err := l1OPBridgeAdapterEncoderABI.Methods["encodeFinalizeWithdrawalERC20Payload"].Inputs.Pack(
		optimism_l1_bridge_adapter_encoder.OptimismL1BridgeAdapterFinalizeWithdrawERC20Payload{
			Action: FinalizationActionProveWithdrawal,
			Data:   encodedProveWithdrawal,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("encodeFinalizeWithdrawalERC20Payload: %w", err)
	}

	return encodedPayload, nil
}

func (l *l2ToL1Bridge) generateTransferBridgeDataForFinalize(
	ctx context.Context,
	transfer *liquiditymanager.LiquidityManagerLiquidityTransferred,
) ([]byte, error) {
	receipt, err := l.l2Client.TransactionReceipt(ctx, transfer.Raw.TxHash)
	if err != nil {
		return nil, fmt.Errorf("get transaction receipt: %w", err)
	}

	messagePassedLog := withdrawprover.GetMessagePassedLog(receipt.Logs)
	if messagePassedLog == nil {
		panic(fmt.Sprintf("No message passed log found in receipt %s", receipt.TxHash.String()))
	}

	messagePassed, err := withdrawprover.ParseMessagePassedLog(messagePassedLog)
	if err != nil {
		return nil, fmt.Errorf("parse message passed log: %w", err)
	}

	encodedFinalizeWithdrawal, err := l1OPBridgeAdapterEncoderABI.Methods["encodeOptimismFinalizationPayload"].Inputs.Pack(
		optimism_l1_bridge_adapter_encoder.OptimismL1BridgeAdapterOptimismFinalizationPayload{
			WithdrawalTransaction: optimism_l1_bridge_adapter_encoder.TypesWithdrawalTransaction{
				Nonce:    messagePassed.Nonce,
				Sender:   messagePassed.Sender,
				Target:   messagePassed.Target,
				Value:    messagePassed.Value,
				GasLimit: messagePassed.GasLimit,
				Data:     messagePassed.Data,
			},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("encodeOptimismFinalizationPayload: %w", err)
	}

	// then encode the finalize withdraw erc20 payload next.
	encodedPayload, err := l1OPBridgeAdapterEncoderABI.Methods["encodeFinalizeWithdrawalERC20Payload"].Inputs.Pack(
		optimism_l1_bridge_adapter_encoder.OptimismL1BridgeAdapterFinalizeWithdrawERC20Payload{
			Action: FinalizationActionFinalizeWithdrawal,
			Data:   encodedFinalizeWithdrawal,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("encodeFinalizeWithdrawalERC20Payload: %w", err)
	}

	return encodedPayload, nil
}

// GetBridgePayloadAndFee implements bridge.Bridge.
func (l *l2ToL1Bridge) GetBridgePayloadAndFee(
	_ context.Context,
	_ models.Transfer,
) ([]byte, *big.Int, error) {
	// Optimism L2 to L1 transfers require no bridge specific payload.
	return []byte{}, big.NewInt(0), nil
}

// QuorumizedBridgePayload implements bridge.Bridge.
func (l *l2ToL1Bridge) QuorumizedBridgePayload(_ [][]byte, _ int) ([]byte, error) {
	// Optimism L2 to L1 transfers require no bridge specific payload.
	return []byte{}, nil
}

// Close implements bridge.Bridge.
func (l *l2ToL1Bridge) Close(ctx context.Context) error {
	return multierr.Combine(
		l.l2LogPoller.UnregisterFilter(ctx, l.l2FilterName),
		l.l1LogPoller.UnregisterFilter(ctx, l.l1FilterName),
	)
}

func formatWithdrawalProof(proof [][]byte) string {
	var builder strings.Builder
	builder.WriteString("{")
	for i, p := range proof {
		builder.WriteString(hexutil.Encode(p))
		if i < len(proof)-1 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString("}")
	return builder.String()
}
