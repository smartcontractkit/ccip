package ccip

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"sort"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	gethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/pkg/errors"

	txmgrtypes "github.com/smartcontractkit/chainlink/v2/common/txmgr/types"
	"github.com/smartcontractkit/chainlink/v2/core/assets"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/gas"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	evmtypes "github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/price_registry"
	"github.com/smartcontractkit/chainlink/v2/core/logger"

	"github.com/smartcontractkit/libocr/offchainreporting2/types"
)

const (
	PERMISSIONLESS_EXECUTION_THRESHOLD = 7 * 24 * time.Hour
)

var (
	_ types.ReportingPluginFactory = &ExecutionReportingPluginFactory{}
	_ types.ReportingPlugin        = &ExecutionReportingPlugin{}
)

func MessagesFromExecutionReport(report types.Report) ([]uint64, [][]byte, error) {
	decodeExecutionReport, err := DecodeExecutionReport(report)
	if err != nil {
		return nil, nil, err
	}
	return decodeExecutionReport.SequenceNumbers, decodeExecutionReport.EncodedMessages, nil
}

func DecodeExecutionReport(report types.Report) (*evm_2_evm_offramp.InternalExecutionReport, error) {
	unpacked, err := makeExecutionReportArgs().Unpack(report)
	if err != nil {
		return nil, err
	}
	if len(unpacked) == 0 {
		return nil, errors.New("assumptionViolation: expected at least one element")
	}

	// Must be anonymous struct here
	erStruct, ok := unpacked[0].(struct {
		SequenceNumbers []uint64    `json:"sequenceNumbers"`
		EncodedMessages [][]byte    `json:"encodedMessages"`
		Proofs          [][32]uint8 `json:"proofs"`
		ProofFlagBits   *big.Int    `json:"proofFlagBits"`
	})
	if !ok {
		return nil, fmt.Errorf("got %T", unpacked[0])
	}
	var er evm_2_evm_offramp.InternalExecutionReport
	er.EncodedMessages = append(er.EncodedMessages, erStruct.EncodedMessages...)
	er.Proofs = append(er.Proofs, erStruct.Proofs...)
	er.SequenceNumbers = erStruct.SequenceNumbers
	// Unpack will populate with big.Int{false, <allocated empty nat>} for 0 values,
	// which is different from the expected big.NewInt(0). Rebuild to the expected value for this case.
	er.ProofFlagBits = big.NewInt(erStruct.ProofFlagBits.Int64())
	return &er, nil
}

func EncodeExecutionReport(seqNums []uint64,
	msgs [][]byte,
	proofs [][32]byte,
	proofSourceFlags []bool,
) (types.Report, error) {
	return makeExecutionReportArgs().PackValues([]interface{}{&evm_2_evm_offramp.InternalExecutionReport{
		SequenceNumbers: seqNums,
		EncodedMessages: msgs,
		Proofs:          proofs,
		ProofFlagBits:   ProofFlagsToBits(proofSourceFlags),
	}})
}

type ExecutionPluginConfig struct {
	onRamp                 *evm_2_evm_onramp.EVM2EVMOnRamp
	offRamp                *evm_2_evm_offramp.EVM2EVMOffRamp
	commitStore            *commit_store.CommitStore
	source, dest           logpoller.LogPoller
	eventSignatures        EventSignatures
	snoozeTime             time.Duration
	inflightCacheExpiry    time.Duration
	leafHasher             LeafHasherInterface[[32]byte]
	lggr                   logger.Logger
	srcPriceRegistry       *price_registry.PriceRegistry
	destPriceRegistry      *price_registry.PriceRegistry
	destGasEstimator       txmgrtypes.FeeEstimator[*evmtypes.Head, gas.EvmFee, *assets.Wei, common.Hash]
	srcWrappedNativeToken  common.Address
	destWrappedNativeToken common.Address
}

type ExecutionReportingPluginFactory struct {
	config ExecutionPluginConfig
}

func NewExecutionReportingPluginFactory(config ExecutionPluginConfig) types.ReportingPluginFactory {
	return &ExecutionReportingPluginFactory{config: config}
}

func (rf *ExecutionReportingPluginFactory) NewReportingPlugin(config types.ReportingPluginConfig) (types.ReportingPlugin, types.ReportingPluginInfo, error) {
	offchainConfig, err := Decode(config.OffchainConfig)
	if err != nil {
		return nil, types.ReportingPluginInfo{}, err
	}
	onRampABI, err := abi.JSON(strings.NewReader(evm_2_evm_onramp.EVM2EVMOnRampABI))
	if err != nil {
		return nil, types.ReportingPluginInfo{}, errors.Wrap(err, "failed to decode onRamp abi")
	}

	return &ExecutionReportingPlugin{
			lggr:            rf.config.lggr.Named("ExecutionReportingPlugin"),
			F:               config.F,
			offchainConfig:  offchainConfig,
			config:          rf.config,
			snoozedRoots:    make(map[[32]byte]time.Time),
			inflightReports: newInflightReportsContainer(rf.config.inflightCacheExpiry),
			onRampABI:       onRampABI,
		}, types.ReportingPluginInfo{
			Name:          "CCIPExecution",
			UniqueReports: true,
			Limits: types.ReportingPluginLimits{
				MaxObservationLength: MaxObservationLength,
				MaxReportLength:      MaxExecutionReportLength,
			},
		}, nil
}

type ExecutionReportingPlugin struct {
	lggr            logger.Logger
	F               int
	config          ExecutionPluginConfig
	inflightReports *inflightReportsContainer
	offchainConfig  OffchainConfig
	snoozedRoots    map[[32]byte]time.Time
	onRampABI       abi.ABI
}

func (r *ExecutionReportingPlugin) Query(context.Context, types.ReportTimestamp) (types.Query, error) {
	return types.Query{}, nil
}

func (r *ExecutionReportingPlugin) Observation(ctx context.Context, timestamp types.ReportTimestamp, query types.Query) (types.Observation, error) {
	lggr := r.lggr.Named("ExecutionObservation")
	if isCommitStoreDownNow(ctx, lggr, r.config.commitStore) {
		return nil, ErrCommitStoreIsDown
	}
	// Expire any inflight reports.
	r.inflightReports.expire(lggr)
	inFlight := r.inflightReports.getAll()

	batchBuilderStart := time.Now()
	// IMPORTANT: We build executable set based on the leaders token prices, ensuring consistency across followers.
	executableSequenceNumbers, err := r.getExecutableSeqNrs(ctx, inFlight)
	lggr.Infof("Batch building took %d ms", time.Since(batchBuilderStart).Milliseconds())
	if err != nil {
		return nil, err
	}
	lggr.Infof("executable seq nums %v %x", executableSequenceNumbers, r.config.eventSignatures.SendRequested)

	// Note can be empty
	return ExecutionObservation{SeqNrs: executableSequenceNumbers}.Marshal()
}

func (r *ExecutionReportingPlugin) getExecutedSeqNrsInRange(min, max uint64) (map[uint64]struct{}, error) {
	// Should be able to keep this log constant across msg types.
	executedLogs, err := r.config.dest.IndexedLogsTopicRange(r.config.eventSignatures.ExecutionStateChanged, r.config.offRamp.Address(), r.config.eventSignatures.ExecutionStateChangedSequenceNumberIndex, logpoller.EvmWord(min), logpoller.EvmWord(max), int(r.offchainConfig.DestIncomingConfirmations))
	if err != nil {
		return nil, err
	}
	executedMp := make(map[uint64]struct{})
	for _, executedLog := range executedLogs {
		exec, err := r.config.offRamp.ParseExecutionStateChanged(gethtypes.Log{Data: executedLog.Data, Topics: executedLog.GetTopics()})
		if err != nil {
			return nil, err
		}
		executedMp[exec.SequenceNumber] = struct{}{}
	}
	return executedMp, nil
}

func (r *ExecutionReportingPlugin) getExecutableSeqNrs(ctx context.Context, inflight []InflightInternalExecutionReport) ([]uint64, error) {
	unexpiredReports, err := getUnexpiredCommitReports(r.config.dest, r.config.commitStore)
	if err != nil {
		return nil, err
	}
	r.lggr.Infow("unexpired roots", "n", len(unexpiredReports))
	if len(unexpiredReports) == 0 {
		return []uint64{}, nil
	}

	// This could result in slightly different values on each call as
	// the function returns the allowed amount at the time of the last block.
	// Since this will only increase over time, the highest observed value will
	// always be the lower bound of what would be available on chain
	// since we already account for inflight txs.
	bucket, err := r.config.offRamp.CalculateCurrentTokenBucketState(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}
	allowedTokenAmount := bucket.Tokens

	// TODO don't build on every batch builder call but only change on changing configuration
	srcToDst := make(map[common.Address]common.Address)
	sourceTokens, err := r.config.offRamp.GetSupportedTokens(nil)
	if err != nil {
		return nil, err
	}

	for _, sourceToken := range sourceTokens {
		dst, err2 := r.config.offRamp.GetDestinationToken(&bind.CallOpts{Context: ctx}, sourceToken)
		if err2 != nil {
			return nil, err2
		}
		srcToDst[sourceToken] = dst
	}

	supportedDestTokens := make([]common.Address, 0, len(srcToDst))
	for _, destToken := range srcToDst {
		supportedDestTokens = append(supportedDestTokens, destToken)
	}

	destTokenPrices, err := r.config.offRamp.GetPricesForTokens(&bind.CallOpts{Context: ctx}, supportedDestTokens)
	if err != nil {
		return nil, err
	}

	pricePerDestToken := make(map[common.Address]*big.Int)
	for i, destToken := range supportedDestTokens {
		pricePerDestToken[destToken] = destTokenPrices[i]
	}

	srcFeeTokensPrices, err := getFeeTokensPrices(ctx, r.config.srcPriceRegistry, r.config.srcWrappedNativeToken)
	if err != nil {
		return nil, err
	}
	destFeeTokensPrices, err := getFeeTokensPrices(ctx, r.config.destPriceRegistry, r.config.destWrappedNativeToken)
	if err != nil {
		return nil, err
	}

	destGasPriceWei, _, err := r.config.destGasEstimator.GetFee(ctx, nil, BatchGasLimit, assets.NewWei(big.NewInt(MaxGasPrice)))
	if err != nil {
		return nil, errors.Wrap(err, "could not estimate destination gas price")
	}
	destGasPrice := destGasPriceWei.Legacy.ToInt()
	if destGasPriceWei.Dynamic != nil {
		destGasPrice = destGasPriceWei.Dynamic.FeeCap.ToInt()
	}

	r.lggr.Debugw("processing unexpired reports", "n", len(unexpiredReports))

	for _, unexpiredReport := range unexpiredReports {
		if ctx.Err() != nil {
			r.lggr.Warn("killed by context")
			break
		}
		snoozeUntil, haveSnoozed := r.snoozedRoots[unexpiredReport.MerkleRoot]
		if haveSnoozed && time.Now().Before(snoozeUntil) {
			incSkippedRequests(reasonSnoozed)
			continue
		}
		blessed, err := r.config.commitStore.IsBlessed(&bind.CallOpts{Context: ctx}, unexpiredReport.MerkleRoot)
		if err != nil {
			return nil, err
		}
		if !blessed {
			r.lggr.Infow("report is accepted but not blessed", "report", hexutil.Encode(unexpiredReport.MerkleRoot[:]))
			incSkippedRequests(reasonNotBlessed)
			continue
		}
		// Check this root for executable messages
		srcLogs, err := r.config.source.LogsDataWordRange(r.config.eventSignatures.SendRequested, r.config.onRamp.Address(), r.config.eventSignatures.SendRequestedSequenceNumberIndex, logpoller.EvmWord(unexpiredReport.Interval.Min), logpoller.EvmWord(unexpiredReport.Interval.Max), int(r.offchainConfig.SourceIncomingConfirmations))
		if err != nil {
			return nil, err
		}
		if len(srcLogs) != int(unexpiredReport.Interval.Max-unexpiredReport.Interval.Min+1) {
			return nil, errors.Errorf("unexpected missing msgs in committed root %x have %d want %d", unexpiredReport.MerkleRoot, len(srcLogs), int(unexpiredReport.Interval.Max-unexpiredReport.Interval.Min+1))
		}
		// TODO: Reorg risk here? I.e. 1 message in a batch, we see its executed so we snooze forever,
		// then it gets reorged out and we'll never retry.
		executedMp, err := r.getExecutedSeqNrsInRange(unexpiredReport.Interval.Min, unexpiredReport.Interval.Max)
		if err != nil {
			return nil, err
		}

		r.lggr.Debugw("building next batch", "executedMp", len(executedMp))

		batch, allMessagesExecuted := r.buildBatch(srcToDst, srcLogs, executedMp, inflight, allowedTokenAmount,
			pricePerDestToken, srcFeeTokensPrices, destFeeTokensPrices, destGasPrice)
		// If all messages are already executed, snooze the root for the PERMISSIONLESS_EXECUTION_THRESHOLD_SECONDS,
		// so it will never be considered again.
		if allMessagesExecuted {
			r.lggr.Infof("Snoozing root %s forever since there are no executable txs anymore %v", hex.EncodeToString(unexpiredReport.MerkleRoot[:]), executedMp)
			r.snoozedRoots[unexpiredReport.MerkleRoot] = time.Now().Add(PERMISSIONLESS_EXECUTION_THRESHOLD)
			incSkippedRequests(reasonAllExecuted)
			continue
		}
		if len(batch) != 0 {
			return batch, nil
		}
		r.snoozedRoots[unexpiredReport.MerkleRoot] = time.Now().Add(r.config.snoozeTime)
	}
	return []uint64{}, nil
}

func (r *ExecutionReportingPlugin) buildBatch(srcToDst map[common.Address]common.Address,
	srcLogs []logpoller.Log,
	executedSeq map[uint64]struct{},
	inflight []InflightInternalExecutionReport,
	aggregateTokenLimit *big.Int,
	tokenLimitPrices map[common.Address]*big.Int,
	srcFeeTokenPricesUSD map[common.Address]*big.Int,
	destFeeTokenPricesUSD map[common.Address]*big.Int,
	execGasPriceEstimate *big.Int,
) (executableSeqNrs []uint64, executedAllMessages bool) {
	inflightSeqNrs, inflightAggregateValue, maxInflightSenderNonces, err := r.inflight(inflight, tokenLimitPrices, srcToDst)
	if err != nil {
		r.lggr.Errorw("Unexpected error computing inflight values", "err", err)
		return []uint64{}, false
	}
	availableGas := uint64(BatchGasLimit)
	aggregateTokenLimit.Sub(aggregateTokenLimit, inflightAggregateValue)
	executedAllMessages = true
	expectedNonces := make(map[common.Address]uint64)
	for _, srcLog := range srcLogs {
		msg, err2 := parseCCIPSendRequestedLog(gethtypes.Log{
			Topics: srcLog.GetTopics(),
			Data:   srcLog.Data,
		}, r.onRampABI)
		if err2 != nil {
			r.lggr.Errorw("unable to parse message", "err", err2, "msg", msg)
			// Unable to parse so don't mark as executed
			executedAllMessages = false
			continue
		}
		lggr := r.lggr.With("messageID", hexutil.Encode(msg.Message.MessageId[:]))
		if _, executed := executedSeq[msg.Message.SequenceNumber]; executed {
			lggr.Infow("Skipping message already executed", "seqNr", msg.Message.SequenceNumber)
			continue
		}
		executedAllMessages = false
		if _, inflight := inflightSeqNrs[msg.Message.SequenceNumber]; inflight {
			lggr.Infow("Skipping message already inflight", "seqNr", msg.Message.SequenceNumber)
			continue
		}
		if _, ok := expectedNonces[msg.Message.Sender]; !ok {
			// First message in batch, need to populate expected nonce
			if maxInflight, ok := maxInflightSenderNonces[msg.Message.Sender]; ok {
				// Sender already has inflight nonce, populate from there
				expectedNonces[msg.Message.Sender] = maxInflight + 1
			} else {
				// Nothing inflight take from chain.
				// Chain holds existing nonce.
				nonce, err := r.config.offRamp.GetSenderNonce(nil, msg.Message.Sender)
				if err != nil {
					lggr.Errorw("unable to get sender nonce", "err", err)
					continue
				}
				expectedNonces[msg.Message.Sender] = nonce + 1
			}
		}
		// Check expected nonce is valid
		if msg.Message.Nonce != expectedNonces[msg.Message.Sender] {
			lggr.Warnw("Skipping message invalid nonce", "have", msg.Message.Nonce, "want", expectedNonces[msg.Message.Sender])
			continue
		}

		var tokens []common.Address
		var amounts []*big.Int
		for i := 0; i < len(msg.Message.TokenAmounts); i++ {
			tokens = append(tokens, msg.Message.TokenAmounts[i].Token)
			amounts = append(amounts, msg.Message.TokenAmounts[i].Amount)
		}
		msgValue, err := aggregateTokenValue(tokenLimitPrices, srcToDst, tokens, amounts)
		if err != nil {
			lggr.Errorw("Skipping message unable to compute aggregate value", "err", err)
			continue
		}
		// if token limit is smaller than message value skip message
		if aggregateTokenLimit.Cmp(msgValue) == -1 {
			lggr.Warnw("token limit is smaller than message value", "aggregateTokenLimit", aggregateTokenLimit.String(), "msgValue", msgValue.String())
			continue
		}
		// Fee boosting
		execCostUsd := computeExecCost(msg, execGasPriceEstimate, destFeeTokenPricesUSD[r.config.destWrappedNativeToken])
		// calculating the source chain fee, dividing by 1e18 for denomination.
		// For example:
		// FeeToken=link; FeeTokenAmount=1e17 i.e. 0.1 link, price is 6e18 USD/link (1 USD = 1e18),
		// availableFee is 1e17*6e18/1e18 = 6e17 = 0.6 USD
		availableFee := big.NewInt(0).Mul(msg.Message.FeeTokenAmount, srcFeeTokenPricesUSD[msg.Message.FeeToken])
		availableFee = availableFee.Div(availableFee, big.NewInt(1e18))
		availableFeeUsd := waitBoostedFee(srcLog.BlockTimestamp, availableFee)
		if availableFeeUsd.Cmp(execCostUsd) < 0 {
			lggr.Infow("Insufficient remaining fee", "availableFeeUsd", availableFeeUsd, "execCostUsd", execCostUsd)
			continue
		}

		messageMaxGas := msg.Message.GasLimit.Uint64() + maxGasOverHeadGas(len(srcLogs), msg.Message)
		// Check sufficient gas in batch
		if availableGas < messageMaxGas {
			lggr.Infow("Insufficient remaining gas in batch limit", "availableGas", availableGas, "messageMaxGas", messageMaxGas)
			continue
		}
		availableGas -= messageMaxGas
		aggregateTokenLimit.Sub(aggregateTokenLimit, msgValue)

		lggr.Infow("Adding msg to batch", "seqNum", msg.Message.SequenceNumber, "nonce", msg.Message.Nonce)
		executableSeqNrs = append(executableSeqNrs, msg.Message.SequenceNumber)
		expectedNonces[msg.Message.Sender] = msg.Message.Nonce + 1
	}
	return executableSeqNrs, executedAllMessages
}

func (r *ExecutionReportingPlugin) parseSeqNr(log logpoller.Log) (uint64, error) {
	s, err := r.config.onRamp.ParseCCIPSendRequested(log.ToGethLog())
	if err != nil {
		return 0, err
	}
	return s.Message.SequenceNumber, nil
}

// Assumes non-empty report. Messages to execute can span more than one report, but are assumed to be in order of increasing
// sequence number.
func (r *ExecutionReportingPlugin) buildReport(ctx context.Context, lggr logger.Logger, finalSeqNums []uint64) ([]byte, error) {
	me, err := buildExecution(
		ctx,
		lggr,
		r.config.source,
		r.config.dest,
		r.config.onRamp.Address(),
		finalSeqNums,
		r.config.commitStore,
		int(r.offchainConfig.SourceIncomingConfirmations),
		r.config.eventSignatures,
		r.parseSeqNr,
		r.config.leafHasher,
	)
	if err != nil {
		return nil, err
	}
	return EncodeExecutionReport(finalSeqNums,
		me.encMsgs,
		me.proofs,
		me.proofSourceFlags,
	)
}

func (r *ExecutionReportingPlugin) Report(ctx context.Context, timestamp types.ReportTimestamp, query types.Query, observations []types.AttributedObservation) (bool, types.Report, error) {
	lggr := r.lggr.Named("Report")
	if isCommitStoreDownNow(ctx, lggr, r.config.commitStore) {
		return false, nil, ErrCommitStoreIsDown
	}
	nonEmptyObservations := getNonEmptyObservations[ExecutionObservation](lggr, observations)
	// Need at least F+1 observations
	if len(nonEmptyObservations) <= r.F {
		lggr.Tracew("Non-empty observations <= F, need at least F+1 to continue")
		return false, nil, nil
	}

	finalSequenceNumbers := calculateSequenceNumberConsensus(nonEmptyObservations, r.F)
	if len(finalSequenceNumbers) == 0 {
		return false, nil, nil
	}

	report, err := r.buildReport(ctx, lggr, finalSequenceNumbers)
	if err != nil {
		return false, nil, err
	}
	lggr.Infow("Built report",
		"onRampAddr", r.config.onRamp.Address(),
		"finalSeqNums", finalSequenceNumbers)
	return true, report, nil
}

func calculateSequenceNumberConsensus(observations []ExecutionObservation, f int) []uint64 {
	tally := make(map[uint64]int)
	for _, obs := range observations {
		for _, seqNr := range obs.SeqNrs {
			tally[seqNr]++
		}
	}
	var finalSequenceNumbers []uint64
	for seqNr, count := range tally {
		// Note spec deviation - I think it's ok to rely on the batch builder for
		// capping the number of messages vs capping in two places/ways?
		if count > f {
			finalSequenceNumbers = append(finalSequenceNumbers, seqNr)
		}
	}
	// buildReport expects sorted sequence numbers (tally map is non-deterministic).
	sort.Slice(finalSequenceNumbers, func(i, j int) bool {
		return finalSequenceNumbers[i] < finalSequenceNumbers[j]
	})
	return finalSequenceNumbers
}

func (r *ExecutionReportingPlugin) ShouldAcceptFinalizedReport(ctx context.Context, timestamp types.ReportTimestamp, report types.Report) (bool, error) {
	lggr := r.lggr.Named("ShouldAcceptFinalizedReport")
	seqNrs, encMsgs, err := MessagesFromExecutionReport(report)
	if err != nil {
		lggr.Errorw("unable to decode report", "err", err)
		return false, nil
	}
	lggr.Infof("Seq nums %v", seqNrs)
	// If the first message is executed already, this execution report is stale, and we do not accept it.
	stale, err := r.isStaleReport(seqNrs)
	if err != nil {
		return false, err
	}
	if stale {
		return false, nil
	}
	// Else just assume in flight
	if err = r.inflightReports.add(lggr, seqNrs, encMsgs); err != nil {
		return false, err
	}
	return true, nil
}

func (r *ExecutionReportingPlugin) ShouldTransmitAcceptedReport(ctx context.Context, timestamp types.ReportTimestamp, report types.Report) (bool, error) {
	seqNrs, _, err := MessagesFromExecutionReport(report)
	if err != nil {
		return false, nil
	}
	// If report is not stale we transmit.
	// When the executeTransmitter enqueues the tx for tx manager,
	// we mark it as execution_sent, removing it from the set of inflight messages.
	stale, err := r.isStaleReport(seqNrs)
	return !stale, err
}

func (r *ExecutionReportingPlugin) isStaleReport(seqNrs []uint64) (bool, error) {
	// If the first message is executed already, this execution report is stale.
	msgState, err := r.config.offRamp.GetExecutionState(nil, seqNrs[0])
	if err != nil {
		// TODO: do we need to check for not present error?
		return true, err
	}
	if msgState == MessageStateFailure || msgState == MessageStateSuccess {
		return true, nil
	}

	return false, nil
}

func (r *ExecutionReportingPlugin) Close() error {
	return nil
}

func (r *ExecutionReportingPlugin) inflight(
	inflight []InflightInternalExecutionReport,
	tokenLimitPrices map[common.Address]*big.Int,
	srcToDst map[common.Address]common.Address,
) (map[uint64]struct{}, *big.Int, map[common.Address]uint64, error) {
	inflightSeqNrs := make(map[uint64]struct{})
	inflightAggregateValue := big.NewInt(0)
	maxInflightSenderNonces := make(map[common.Address]uint64)
	for _, rep := range inflight {
		for _, seqNr := range rep.seqNrs {
			inflightSeqNrs[seqNr] = struct{}{}
		}
		for _, encMsg := range rep.encMessages {
			msg, err := parseCCIPSendRequestedLog(gethtypes.Log{
				// Note this needs to change if we start indexing things.
				Topics: []common.Hash{r.config.eventSignatures.SendRequested},
				Data:   encMsg,
			}, r.onRampABI)
			if err != nil {
				return nil, nil, nil, err
			}
			var tokens []common.Address
			var amounts []*big.Int
			for i := 0; i < len(msg.Message.TokenAmounts); i++ {
				tokens = append(tokens, msg.Message.TokenAmounts[i].Token)
				amounts = append(amounts, msg.Message.TokenAmounts[i].Amount)
			}
			msgValue, err := aggregateTokenValue(tokenLimitPrices, srcToDst, tokens, amounts)
			if err != nil {
				return nil, nil, nil, err
			}
			inflightAggregateValue.Add(inflightAggregateValue, msgValue)
			maxInflightSenderNonce, ok := maxInflightSenderNonces[msg.Message.Sender]
			if !ok || msg.Message.Nonce > maxInflightSenderNonce {
				maxInflightSenderNonces[msg.Message.Sender] = msg.Message.Nonce
			}
		}
	}
	return inflightSeqNrs, inflightAggregateValue, maxInflightSenderNonces, nil
}

func parseCCIPSendRequestedLog(log gethtypes.Log, abi abi.ABI) (*evm_2_evm_onramp.EVM2EVMOnRampCCIPSendRequested, error) {
	event := new(evm_2_evm_onramp.EVM2EVMOnRampCCIPSendRequested)
	err := bind.NewBoundContract(common.Address{}, abi, nil, nil, nil).UnpackLog(event, "CCIPSendRequested", log)
	if err != nil {
		return nil, err
	}
	return event, nil
}

// getFeeTokensPrices returns fee token prices of the given price registry,
// price values are USD per full token, in base units 1e18 (e.g. 5$ = 5e18).
// this function is used for price reigstry of both source and destination chains.
func getFeeTokensPrices(ctx context.Context, priceRegistry *price_registry.PriceRegistry, wrappedNative common.Address) (map[common.Address]*big.Int, error) {
	prices := make(map[common.Address]*big.Int)

	srcFeeTokens, err := priceRegistry.GetFeeTokens(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, errors.Wrap(err, "could not get source fee tokens")
	}
	for _, feeToken := range srcFeeTokens {
		feeTokenPrice, err := priceRegistry.GetTokenPrice(&bind.CallOpts{Context: ctx}, feeToken)
		if err != nil {
			return nil, errors.Wrapf(err, "could not get token price of %s", feeToken.String())
		}
		prices[feeToken] = feeTokenPrice.Value
	}
	if _, ok := prices[wrappedNative]; !ok {
		srcTokenPrice, err := priceRegistry.GetTokenPrice(&bind.CallOpts{Context: ctx}, wrappedNative)
		if err != nil {
			return nil, errors.Wrap(err, "could not get token prices in USD")
		}
		prices[wrappedNative] = srcTokenPrice.Value
	}

	return prices, nil
}

func getUnexpiredCommitReports(dstLogPoller logpoller.LogPoller, commitStore *commit_store.CommitStore) ([]commit_store.CommitStoreCommitReport, error) {
	logs, err := dstLogPoller.LogsCreatedAfter(ReportAccepted, commitStore.Address(), time.Now().Add(-PERMISSIONLESS_EXECUTION_THRESHOLD))
	if err != nil {
		return nil, err
	}
	var reports []commit_store.CommitStoreCommitReport
	for _, log := range logs {
		reportAccepted, err := commitStore.ParseReportAccepted(gethtypes.Log{
			Topics: log.GetTopics(),
			Data:   log.Data,
		})
		if err != nil {
			return nil, err
		}
		reports = append(reports, reportAccepted.Report)
	}
	return reports, nil
}
