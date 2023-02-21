package ccip

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"sort"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/libocr/offchainreporting2/types"

	"github.com/smartcontractkit/chainlink/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink/core/logger"
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

var (
	_ types.ReportingPluginFactory = &ExecutionReportingPluginFactory{}
	_ types.ReportingPlugin        = &ExecutionReportingPlugin{}
)

type ExecutionPluginConfig struct {
	onRamp              *evm_2_evm_onramp.EVM2EVMOnRamp
	offRamp             *evm_2_evm_offramp.EVM2EVMOffRamp
	commitStore         *commit_store.CommitStore
	source, dest        logpoller.LogPoller
	eventSignatures     EventSignatures
	snoozeTime          time.Duration
	inflightCacheExpiry time.Duration
	builder             BatchBuilderInterface
	leafHasher          LeafHasherInterface[[32]byte]
	lggr                logger.Logger
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

	return &ExecutionReportingPlugin{
			lggr:           rf.config.lggr.Named("ExecutionReportingPlugin"),
			F:              config.F,
			offchainConfig: offchainConfig,
			config:         rf.config,
			snoozedRoots:   make(map[[32]byte]time.Time),
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
	lggr   logger.Logger
	F      int
	config ExecutionPluginConfig
	// We need to synchronize access to the inflight structure
	// as reporting plugin methods may be called from separate goroutines,
	// e.g. reporting vs transmission protocol.
	inFlightMu     sync.RWMutex
	inFlight       []InflightInternalExecutionReport
	offchainConfig OffchainConfig
	snoozedRoots   map[[32]byte]time.Time
}

type Query struct {
	TokenPrices map[common.Address]*big.Int `json:"tokenPrices"`
}

func (r *ExecutionReportingPlugin) Query(ctx context.Context, timestamp types.ReportTimestamp) (types.Query, error) {
	return types.Query{}, nil
}

func (r *ExecutionReportingPlugin) Observation(ctx context.Context, timestamp types.ReportTimestamp, query types.Query) (types.Observation, error) {
	lggr := r.lggr.Named("ExecutionObservation")
	if isCommitStoreDownNow(lggr, r.config.commitStore) {
		return nil, ErrCommitStoreIsDown
	}
	// Expire any inflight reports.
	r.expireInflight(lggr)

	// Read and make a copy for the builder.
	r.inFlightMu.RLock()
	inFlight := make([]InflightInternalExecutionReport, len(r.inFlight))
	copy(inFlight[:], r.inFlight[:])
	r.inFlightMu.RUnlock()

	batchBuilderStart := time.Now()
	// IMPORTANT: We build executable set based on the leaders token prices, ensuring consistency across followers.
	executableSequenceNumbers, err := r.getExecutableSeqNrs(inFlight)
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

func (r *ExecutionReportingPlugin) getExecutableSeqNrs(inflight []InflightInternalExecutionReport) ([]uint64, error) {
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
	bucket, err := r.config.offRamp.CalculateCurrentTokenBucketState(nil)
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
		dst, err2 := r.config.offRamp.GetDestinationToken(nil, sourceToken)
		if err2 != nil {
			return nil, err2
		}
		srcToDst[sourceToken] = dst
	}

	supportedDestTokenAmounts := make([]common.Address, 0, len(srcToDst))
	for _, destTokenAmounts := range srcToDst {
		supportedDestTokenAmounts = append(supportedDestTokenAmounts, destTokenAmounts)
	}

	destTokenPrices, err := r.config.offRamp.GetPricesForTokens(nil, supportedDestTokenAmounts)
	if err != nil {
		return nil, err
	}

	pricePerDestToken := make(map[common.Address]*big.Int)
	for i, destToken := range supportedDestTokenAmounts {
		pricePerDestToken[destToken] = destTokenPrices[i]
	}

	for _, unexpiredReport := range unexpiredReports {
		snoozeUntil, haveSnoozed := r.snoozedRoots[unexpiredReport.MerkleRoot]
		if haveSnoozed && time.Now().Before(snoozeUntil) {
			incSkippedRequests(reasonSnoozed)
			continue
		}
		blessed, err := r.config.commitStore.IsBlessed(nil, unexpiredReport.MerkleRoot)
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

		batch, allMessagesExecuted := r.config.builder.BuildBatch(srcToDst, srcLogs, executedMp, inflight, allowedTokenAmount, pricePerDestToken)
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

func (r *ExecutionReportingPlugin) parseSeqNr(log gethtypes.Log) (uint64, error) {
	s, err := r.config.onRamp.ParseCCIPSendRequested(log)
	if err != nil {
		return 0, err
	}
	return s.Message.SequenceNumber, nil
}

// Assumes non-empty report. Messages to execute can span more than one report, but are assumed to be in order of increasing
// sequence number.
func (r *ExecutionReportingPlugin) buildReport(lggr logger.Logger, finalSeqNums []uint64) ([]byte, error) {
	me, err := buildExecution(
		lggr,
		r.config.source,
		r.config.dest,
		r.config.onRamp.Address(),
		finalSeqNums,
		r.config.commitStore,
		int(r.offchainConfig.SourceIncomingConfirmations),
		r.config.eventSignatures,
		r.parseSeqNr,
		r.config.leafHasher.HashLeaf,
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
	if isCommitStoreDownNow(lggr, r.config.commitStore) {
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

	report, err := r.buildReport(lggr, finalSequenceNumbers)
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

func (r *ExecutionReportingPlugin) expireInflight(lggr logger.Logger) {
	r.inFlightMu.Lock()
	defer r.inFlightMu.Unlock()
	// Reap old inflight txs and check if any messages in the report are inflight.
	var stillInFlight []InflightInternalExecutionReport
	for _, report := range r.inFlight {
		if time.Since(report.createdAt) > r.config.inflightCacheExpiry {
			// Happy path: inflight report was successfully transmitted onchain, we remove it from inflight and onchain state reflects inflight.
			// Sad path: inflight report reverts onchain, we remove it from inflight, onchain state does not reflect the change so we retry.
			lggr.Infow("Inflight report expired", "seqNums", report.seqNrs)
		} else {
			stillInFlight = append(stillInFlight, report)
		}
	}
	r.inFlight = stillInFlight
}

func (r *ExecutionReportingPlugin) addToInflight(lggr logger.Logger, seqNrs []uint64, encMsgs [][]byte) error {
	r.inFlightMu.Lock()
	defer r.inFlightMu.Unlock()
	for _, report := range r.inFlight {
		// TODO: Think about if this fails in reorgs
		if (len(report.seqNrs) > 0 && len(seqNrs) > 0) && (report.seqNrs[0] == seqNrs[0]) {
			return errors.Errorf("report is already in flight")
		}
	}
	// Otherwise not already in flight, add it.
	lggr.Infow("Added report to inflight",
		"seqNums", seqNrs)
	r.inFlight = append(r.inFlight, InflightInternalExecutionReport{
		createdAt:   time.Now(),
		seqNrs:      seqNrs,
		encMessages: encMsgs,
	})
	return nil
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
	if err = r.addToInflight(lggr, seqNrs, encMsgs); err != nil {
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
