package ccipexec

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink-common/pkg/hashutil"
	"github.com/smartcontractkit/chainlink-common/pkg/merklemulti"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccip"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/prices"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/tokendata"
	"github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/statuschecker"
)

type BatchContext struct {
	ctx                        context.Context
	report                     commitReportWithSendRequests
	lggr                       logger.Logger
	availableDataLen           int
	availableGas               uint64
	expectedNonces             map[cciptypes.Address]uint64
	sendersNonce               map[cciptypes.Address]uint64
	sourceTokenPricesUSD       map[cciptypes.Address]*big.Int
	destTokenPricesUSD         map[cciptypes.Address]*big.Int
	gasPrice                   *big.Int
	sourceToDestToken          map[cciptypes.Address]cciptypes.Address
	inflightAggregateValue     *big.Int
	aggregateTokenLimit        *big.Int
	tokenDataRemainingDuration time.Duration
	tokenDataWorker            tokendata.Worker
	gasPriceEstimator          prices.GasPriceEstimatorExec
	destWrappedNative          cciptypes.Address
	offchainConfig             cciptypes.ExecOffchainConfig
}

type BaseBatchingStrategy struct{}

type BatchingStrategy interface {
	BuildBatch(batchCtx *BatchContext) ([]ccip.ObservedMessage, []messageExecStatus)
}

type BestEffortBatchingStrategy struct {
	BaseBatchingStrategy
}

// BestEffortBatchingStrategy is a batching strategy that tries to batch as many messages as possible (up to certain limits).
func (s *BestEffortBatchingStrategy) BuildBatch(
	batchCtx *BatchContext,
) ([]ccip.ObservedMessage, []messageExecStatus) {
	batchBuilder := newBatchBuildContainer(len(batchCtx.report.sendRequestsWithMeta))
	for _, msg := range batchCtx.report.sendRequestsWithMeta {
		msgLggr := batchCtx.lggr.With("messageID", hexutil.Encode(msg.MessageID[:]), "seqNr", msg.SequenceNumber)
		shouldAdd, status, messageMaxGas, tokenData, msgValue, err := s.performCommonChecks(batchCtx, msg, msgLggr)

		if err != nil {
			return []ccip.ObservedMessage{}, []messageExecStatus{}
		}

		if !shouldAdd {
			batchBuilder.skip(msg, status)
			continue
		}

		batchCtx.availableGas -= messageMaxGas
		batchCtx.availableDataLen -= len(msg.Data)
		batchCtx.aggregateTokenLimit.Sub(batchCtx.aggregateTokenLimit, msgValue)
		if msg.Nonce > 0 {
			batchCtx.expectedNonces[msg.Sender] = msg.Nonce + 1
		}
		batchBuilder.addToBatch(msg, tokenData)

		msgLggr.Infow(
			"Message added to execution batch",
			"nonce", msg.Nonce,
			"sender", msg.Sender,
			"value", msgValue,
			"availableAggrTokenLimit", batchCtx.aggregateTokenLimit,
			"availableGas", batchCtx.availableGas,
			"availableDataLen", batchCtx.availableDataLen,
		)
	}
	return batchBuilder.batch, batchBuilder.statuses
}

type ZKOverflowBatchingStrategy struct {
	BaseBatchingStrategy
	statuschecker statuschecker.CCIPTransactionStatusChecker
}

// ZKOverflowBatchingStrategy is a batching strategy for ZK chains overflowing under certain conditions.
// It is a simple batching strategy that only allows one message to be added to the batch.
// TXM is used to perform the ZK check: if the message failed the check, it will be skipped.
func (bs *ZKOverflowBatchingStrategy) BuildBatch(
	batchCtx *BatchContext,
) ([]ccip.ObservedMessage, []messageExecStatus) {
	batchBuilder := newBatchBuildContainer(1)

	for _, msg := range batchCtx.report.sendRequestsWithMeta {
		msgLggr := batchCtx.lggr.With("messageID", hexutil.Encode(msg.MessageID[:]), "seqNr", msg.SequenceNumber)
		shouldAdd, status, messageMaxGas, tokenData, msgValue, err := bs.performCommonChecks(batchCtx, msg, msgLggr)

		if err != nil {
			return []ccip.ObservedMessage{}, []messageExecStatus{}
		}

		if !shouldAdd {
			batchBuilder.skip(msg, status)
			continue
		}

		// Check if the messsage is overflown using TXM
		statuses, _, err := bs.statuschecker.CheckMessageStatus(batchCtx.ctx, hexutil.Encode(msg.MessageID[:]))
		if err != nil {
			batchBuilder.skip(msg, TXMCheckError)
			continue
		}

		if len(statuses) == 0 {
			// No status found for message = first time we see it
			msgLggr.Infow("No status found for message, adding to batch")
		} else {
			// Status(es) found for message = check if any of them is final to decide if we should add it to the batch
			finalStatus := false
			for _, s := range statuses {
				if s == types.Fatal {
					msgLggr.Infow("Skipping message - ZK check failed (fatal status)")
					batchBuilder.skip(msg, TXMCheckFailed)
					finalStatus = true
					break
				}
				if s == types.Finalized {
					msgLggr.Infow("Skipping message - ZK check failed (final status)")
					batchBuilder.skip(msg, TXMCheckFailed)
					finalStatus = true
					break
				}
			}
			if finalStatus {
				continue
			}
			msgLggr.Infow("No final status found for message, adding to batch")
		}

		batchCtx.availableGas -= messageMaxGas
		batchCtx.availableDataLen -= len(msg.Data)
		batchCtx.aggregateTokenLimit.Sub(batchCtx.aggregateTokenLimit, msgValue)
		batchBuilder.addToBatch(msg, tokenData)

		msgLggr.Infow(
			"Message added to ZKOverflow execution batch",
			"nonce", msg.Nonce,
			"sender", msg.Sender,
			"value", msgValue,
			"availableAggrTokenLimit", batchCtx.aggregateTokenLimit,
			"availableGas", batchCtx.availableGas,
			"availableDataLen", batchCtx.availableDataLen,
		)

		// Batch size is limited to 1 for ZK Overflow chains
		break
	}
	return batchBuilder.batch, batchBuilder.statuses
}

func (bs *BaseBatchingStrategy) performCommonChecks(
	ctx *BatchContext,
	msg cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta,
	msgLggr logger.Logger,
) (bool, messageStatus, uint64, [][]byte, *big.Int, error) {
	if msg.Executed {
		msgLggr.Infow("Skipping message - already executed")
		return false, AlreadyExecuted, 0, nil, nil, nil
	}

	if len(msg.Data) > ctx.availableDataLen {
		msgLggr.Infow("Skipping message - insufficient remaining batch data length", "msgDataLen", len(msg.Data), "availableBatchDataLen", ctx.availableDataLen)
		return false, InsufficientRemainingBatchDataLength, 0, nil, nil, nil
	}

	messageMaxGas, err1 := calculateMessageMaxGas(
		msg.GasLimit,
		len(ctx.report.sendRequestsWithMeta),
		len(msg.Data),
		len(msg.TokenAmounts),
	)
	if err1 != nil {
		msgLggr.Errorw("Skipping message - message max gas calculation error", "err", err1)
		return false, MessageMaxGasCalcError, 0, nil, nil, nil
	}

	// Check sufficient gas in batch
	if ctx.availableGas < messageMaxGas {
		msgLggr.Infow("Skipping message - insufficient remaining batch gas limit", "availableGas", ctx.availableGas, "messageMaxGas", messageMaxGas)
		return false, InsufficientRemainingBatchGas, 0, nil, nil, nil
	}

	if _, ok := ctx.expectedNonces[msg.Sender]; !ok {
		nonce, ok1 := ctx.sendersNonce[msg.Sender]
		if !ok1 {
			msgLggr.Errorw("Skipping message - missing nonce", "sender", msg.Sender)
			return false, MissingNonce, 0, nil, nil, nil
		}
		ctx.expectedNonces[msg.Sender] = nonce + 1
	}

	// Check expected nonce is valid for sequenced messages.
	// Sequenced messages have non-zero nonces.
	if msg.Nonce > 0 && msg.Nonce != ctx.expectedNonces[msg.Sender] {
		msgLggr.Warnw("Skipping message - invalid nonce", "have", msg.Nonce, "want", ctx.expectedNonces[msg.Sender])
		return false, InvalidNonce, 0, nil, nil, nil
	}

	msgValue, err1 := aggregateTokenValue(ctx.lggr, ctx.destTokenPricesUSD, ctx.sourceToDestToken, msg.TokenAmounts)
	if err1 != nil {
		msgLggr.Errorw("Skipping message - aggregate token value compute error", "err", err1)
		return false, AggregateTokenValueComputeError, 0, nil, nil, nil
	}

	// if token limit is smaller than message value skip message
	if tokensLeft, hasCapacity := hasEnoughTokens(ctx.aggregateTokenLimit, msgValue, ctx.inflightAggregateValue); !hasCapacity {
		msgLggr.Warnw("Skipping message - aggregate token limit exceeded", "aggregateTokenLimit", tokensLeft.String(), "msgValue", msgValue.String())
		return false, AggregateTokenLimitExceeded, 0, nil, nil, nil
	}

	tokenData, elapsed, err1 := bs.getTokenDataWithTimeout(ctx.ctx, msg, ctx.tokenDataRemainingDuration, ctx.tokenDataWorker)
	ctx.tokenDataRemainingDuration -= elapsed
	if err1 != nil {
		if errors.Is(err1, tokendata.ErrNotReady) {
			msgLggr.Warnw("Skipping message - token data not ready", "err", err1)
			return false, TokenDataNotReady, 0, nil, nil, nil
		}
		msgLggr.Errorw("Skipping message - token data fetch error", "err", err1)
		return false, TokenDataFetchError, 0, nil, nil, nil
	}

	dstWrappedNativePrice, exists := ctx.destTokenPricesUSD[ctx.destWrappedNative]
	if !exists {
		msgLggr.Errorw("Skipping message - token not in destination token prices", "token", ctx.destWrappedNative)
		return false, TokenNotInDestTokenPrices, 0, nil, nil, nil
	}

	// calculating the source chain fee, dividing by 1e18 for denomination.
	// For example:
	// FeeToken=link; FeeTokenAmount=1e17 i.e. 0.1 link, price is 6e18 USD/link (1 USD = 1e18),
	// availableFee is 1e17*6e18/1e18 = 6e17 = 0.6 USD
	sourceFeeTokenPrice, exists := ctx.sourceTokenPricesUSD[msg.FeeToken]
	if !exists {
		msgLggr.Errorw("Skipping message - token not in source token prices", "token", msg.FeeToken)
		return false, TokenNotInSrcTokenPrices, 0, nil, nil, nil
	}

	// Fee boosting
	execCostUsd, err1 := ctx.gasPriceEstimator.EstimateMsgCostUSD(ctx.gasPrice, dstWrappedNativePrice, msg)
	if err1 != nil {
		msgLggr.Errorw("Failed to estimate message cost USD", "err", err1)
		return false, "", 0, nil, nil, errors.New("failed to estimate message cost USD")
	}

	availableFee := big.NewInt(0).Mul(msg.FeeTokenAmount, sourceFeeTokenPrice)
	availableFee = availableFee.Div(availableFee, big.NewInt(1e18))
	availableFeeUsd := waitBoostedFee(time.Since(msg.BlockTimestamp), availableFee, ctx.offchainConfig.RelativeBoostPerWaitHour)
	if availableFeeUsd.Cmp(execCostUsd) < 0 {
		msgLggr.Infow(
			"Skipping message - insufficient remaining fee",
			"availableFeeUsd", availableFeeUsd,
			"execCostUsd", execCostUsd,
			"sourceBlockTimestamp", msg.BlockTimestamp,
			"waitTime", time.Since(msg.BlockTimestamp),
			"boost", ctx.offchainConfig.RelativeBoostPerWaitHour,
		)
		return false, InsufficientRemainingFee, 0, nil, nil, nil
	}

	return true, "", messageMaxGas, tokenData, msgValue, nil
}

// getTokenDataWithCappedLatency gets the token data for the provided message.
// Stops and returns an error if more than allowedWaitingTime is passed.
func (bs *BaseBatchingStrategy) getTokenDataWithTimeout(
	ctx context.Context,
	msg cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta,
	timeout time.Duration,
	tokenDataWorker tokendata.Worker,
) ([][]byte, time.Duration, error) {
	if len(msg.TokenAmounts) == 0 {
		return nil, 0, nil
	}

	ctxTimeout, cf := context.WithTimeout(ctx, timeout)
	defer cf()
	tStart := time.Now()
	tokenData, err := tokenDataWorker.GetMsgTokenData(ctxTimeout, msg)
	tDur := time.Since(tStart)
	return tokenData, tDur, err
}

func getProofData(
	ctx context.Context,
	sourceReader ccipdata.OnRampReader,
	interval cciptypes.CommitStoreInterval,
) (sendReqsInRoot []cciptypes.EVM2EVMMessageWithTxMeta, leaves [][32]byte, tree *merklemulti.Tree[[32]byte], err error) {
	// We don't need to double-check if logs are finalized because we already checked that in the Commit phase.
	sendReqs, err := sourceReader.GetSendRequestsBetweenSeqNums(ctx, interval.Min, interval.Max, false)
	if err != nil {
		return nil, nil, nil, err
	}

	if err1 := validateSendRequests(sendReqs, interval); err1 != nil {
		return nil, nil, nil, err1
	}

	leaves = make([][32]byte, 0, len(sendReqs))
	for _, req := range sendReqs {
		leaves = append(leaves, req.Hash)
	}
	tree, err = merklemulti.NewTree(hashutil.NewKeccak(), leaves)
	if err != nil {
		return nil, nil, nil, err
	}
	return sendReqs, leaves, tree, nil
}

func validateSendRequests(sendReqs []cciptypes.EVM2EVMMessageWithTxMeta, interval cciptypes.CommitStoreInterval) error {
	if len(sendReqs) == 0 {
		return fmt.Errorf("could not find any requests in the provided interval %v", interval)
	}

	gotInterval := cciptypes.CommitStoreInterval{
		Min: sendReqs[0].SequenceNumber,
		Max: sendReqs[0].SequenceNumber,
	}

	for _, req := range sendReqs[1:] {
		if req.SequenceNumber < gotInterval.Min {
			gotInterval.Min = req.SequenceNumber
		}
		if req.SequenceNumber > gotInterval.Max {
			gotInterval.Max = req.SequenceNumber
		}
	}

	if (gotInterval.Min != interval.Min) || (gotInterval.Max != interval.Max) {
		return fmt.Errorf("interval %v is not the expected %v", gotInterval, interval)
	}
	return nil
}

func buildExecutionReportForMessages(
	msgsInRoot []cciptypes.EVM2EVMMessageWithTxMeta,
	tree *merklemulti.Tree[[32]byte],
	commitInterval cciptypes.CommitStoreInterval,
	observedMessages []ccip.ObservedMessage,
) (cciptypes.ExecReport, error) {
	innerIdxs := make([]int, 0, len(observedMessages))
	var messages []cciptypes.EVM2EVMMessage
	var offchainTokenData [][][]byte
	for _, observedMessage := range observedMessages {
		if observedMessage.SeqNr < commitInterval.Min || observedMessage.SeqNr > commitInterval.Max {
			// We only return messages from a single root (the root of the first message).
			continue
		}
		innerIdx := int(observedMessage.SeqNr - commitInterval.Min)
		if innerIdx >= len(msgsInRoot) || innerIdx < 0 {
			return cciptypes.ExecReport{}, fmt.Errorf("invalid inneridx SeqNr=%d IntervalMin=%d msgsInRoot=%d",
				observedMessage.SeqNr, commitInterval.Min, len(msgsInRoot))
		}
		messages = append(messages, msgsInRoot[innerIdx].EVM2EVMMessage)
		offchainTokenData = append(offchainTokenData, observedMessage.TokenData)
		innerIdxs = append(innerIdxs, innerIdx)
	}

	merkleProof, err := tree.Prove(innerIdxs)
	if err != nil {
		return cciptypes.ExecReport{}, err
	}

	// any capped proof will have length <= this one, so we reuse it to avoid proving inside loop, and update later if changed
	return cciptypes.ExecReport{
		Messages:          messages,
		Proofs:            merkleProof.Hashes,
		ProofFlagBits:     abihelpers.ProofFlagsToBits(merkleProof.SourceFlags),
		OffchainTokenData: offchainTokenData,
	}, nil
}

// Validates the given message observations do not exceed the committed sequence numbers
// in the commitStoreReader.
func validateSeqNumbers(serviceCtx context.Context, commitStore ccipdata.CommitStoreReader, observedMessages []ccip.ObservedMessage) error {
	nextMin, err := commitStore.GetExpectedNextSequenceNumber(serviceCtx)
	if err != nil {
		return err
	}
	// observedMessages are always sorted by SeqNr and never empty, so it's safe to take last element
	maxSeqNumInBatch := observedMessages[len(observedMessages)-1].SeqNr

	if maxSeqNumInBatch >= nextMin {
		return errors.Errorf("Cannot execute uncommitted seq num. nextMin %v, seqNums %v", nextMin, observedMessages)
	}
	return nil
}

// Gets the commit report from the saved logs for a given sequence number.
func getCommitReportForSeqNum(ctx context.Context, commitStoreReader ccipdata.CommitStoreReader, seqNum uint64) (cciptypes.CommitStoreReport, error) {
	acceptedReports, err := commitStoreReader.GetCommitReportMatchingSeqNum(ctx, seqNum, 0)
	if err != nil {
		return cciptypes.CommitStoreReport{}, err
	}

	if len(acceptedReports) == 0 {
		return cciptypes.CommitStoreReport{}, errors.Errorf("seq number not committed")
	}

	return acceptedReports[0].CommitStoreReport, nil
}

type messageStatus string

const (
	AlreadyExecuted                      messageStatus = "already_executed"
	SenderAlreadySkipped                 messageStatus = "sender_already_skipped"
	MessageMaxGasCalcError               messageStatus = "message_max_gas_calc_error"
	InsufficientRemainingBatchDataLength messageStatus = "insufficient_remaining_batch_data_length"
	InsufficientRemainingBatchGas        messageStatus = "insufficient_remaining_batch_gas"
	MissingNonce                         messageStatus = "missing_nonce"
	InvalidNonce                         messageStatus = "invalid_nonce"
	AggregateTokenValueComputeError      messageStatus = "aggregate_token_value_compute_error"
	AggregateTokenLimitExceeded          messageStatus = "aggregate_token_limit_exceeded"
	TokenDataNotReady                    messageStatus = "token_data_not_ready"
	TokenDataFetchError                  messageStatus = "token_data_fetch_error"
	TokenNotInDestTokenPrices            messageStatus = "token_not_in_dest_token_prices"
	TokenNotInSrcTokenPrices             messageStatus = "token_not_in_src_token_prices"
	InsufficientRemainingFee             messageStatus = "insufficient_remaining_fee"
	AddedToBatch                         messageStatus = "added_to_batch"
	TXMCheckError                        messageStatus = "txm_check_error"
	TXMCheckFailed                       messageStatus = "txm_check_failed"
)

type messageExecStatus struct {
	SeqNr     uint64
	MessageId string
	Status    messageStatus
}

func newMessageExecState(seqNr uint64, messageId cciptypes.Hash, status messageStatus) messageExecStatus {
	return messageExecStatus{
		SeqNr:     seqNr,
		MessageId: hexutil.Encode(messageId[:]),
		Status:    status,
	}
}

type batchBuildContainer struct {
	batch    []ccip.ObservedMessage
	statuses []messageExecStatus
}

func newBatchBuildContainer(capacity int) *batchBuildContainer {
	return &batchBuildContainer{
		batch:    make([]ccip.ObservedMessage, 0, capacity),
		statuses: make([]messageExecStatus, 0, capacity),
	}
}

func (m *batchBuildContainer) skip(msg cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta, status messageStatus) {
	m.addState(msg, status)
}

func (m *batchBuildContainer) addToBatch(msg cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta, tokenData [][]byte) {
	m.addState(msg, AddedToBatch)
	m.batch = append(m.batch, ccip.NewObservedMessage(msg.SequenceNumber, tokenData))
}

func (m *batchBuildContainer) addState(msg cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta, state messageStatus) {
	m.statuses = append(m.statuses, newMessageExecState(msg.SequenceNumber, msg.MessageID, state))
}
